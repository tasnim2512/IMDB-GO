package handler

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	userpb "practice/IMDB/gunk/v1/user"
	"strconv"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
	"google.golang.org/grpc"
)

type usermgmService struct {
	userpb.UserServiceClient
}

type Handler struct {
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	usermgmSvc     usermgmService
	Templates      *template.Template
	staticFiles    fs.FS
	templateFiles  fs.FS
}

type ErrorPage struct {
	Code    int
	Message string
}

func (h Handler) Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	ep := ErrorPage{
		Code:    code,
		Message: error,
	}

	tf := "default"
	switch code {
	case 400, 401, 402, 403, 404:
		tf = "4xx"
	case 500, 501, 503:
		tf = "5xx"
	}

	tpl := fmt.Sprintf("templates/errors/%s.html", tf)
	t, err := template.ParseFiles(tpl)
	if err != nil {
		log.Fatalln(err)
	}

	if err := t.Execute(w, ep); err != nil {
		log.Fatalln(err)
	}
}

func NewHandler(sm *scs.SessionManager, formDecoder *form.Decoder, usermgmConn *grpc.ClientConn, staticFiles, templateFiles fs.FS) *chi.Mux {
	h := &Handler{
		sessionManager: sm,
		decoder:        formDecoder,
		usermgmSvc:     usermgmService{userpb.NewUserServiceClient(usermgmConn)},
		staticFiles:    staticFiles,
		templateFiles:  templateFiles,
	}

	h.ParseTemplates()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(Method)

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Get("/", h.Home)
		r.Get("/registration", h.UserRegistration)
		r.Post("/registration", h.UserRegistrationPost)
		r.Get("/login", h.Login)
		r.Post("/login", h.LoginPostHandler)
	})

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(h.staticFiles))))

	r.Route("/students", func(r chi.Router) {
		r.Use(h.Authenticator)

		// 	r.Get("/", h.ListOfStudent)

		// 	r.Get("/{id:[0-9]+}/delete", h.DeleteStudent)

		// 	r.Get("/{id:[0-9]+}/add/marks/{classId:[0-9]+}", h.GetMarks)

		// 	r.Post("/{id:[0-9]+}/store/marks/{classId:[0-9]+}", h.StoreMarks)

		// 	r.Get("/{id:[0-9]+}/edit/marks/{classId:[0-9]+}", h.EditMarks)

		// 	r.Post("/{id:[0-9]+}/update/marks/{classId:[0-9]+}", h.UpdateMarks)

		// 	r.Get("/{id:[0-9]+}/detail", h.DetailStudent)

		// 	r.Get("/{id:[0-9]+}/result", h.ShowResult)
	})

	// r.Route("/admin", func(r chi.Router) {
	// 	r.Use(h.Authenticator)

	// 	r.Get("/options", h.OPTIONS)

	// 	r.Get("/create/class", h.CreateClass)

	// 	r.Get("/classList", h.ListOfClass)

	// 	r.Post("/store/class", h.StoreClass)

	// 	r.Get("/create/subject", h.CreateSubject)

	// 	r.Post("/store/subject", h.StoreSubject)

	// 	r.Get("/create/student", h.CreateStudent)

	// 	r.Post("/store/student", h.StoreStudent)

	// 	r.Get("/{id:[0-9]+}/edit/student", h.EditStudent)

	// 	r.Put("/{id:[0-9]+}/update/student", h.UpdateStudent)

	// 	r.Get("/show/class/{id:[0-9]+}/result", h.GetResultSheet)

	// })

	r.Group(func(r chi.Router) {
		r.Use(h.Authenticator)
		r.Get("/logout", h.Logout)
	})
	return r
}

func Method(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch strings.ToLower(r.PostFormValue("_method")) {
			case "put":
				r.Method = http.MethodPut
			case "patch":
				r.Method = http.MethodPatch
			case "delete":
				r.Method = http.MethodDelete
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (h Handler) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := h.sessionManager.GetString(r.Context(), "userId")
		uid, err := strconv.Atoi(userId)

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if uid <= 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ParseTemplates() error {
	templates := template.New("web-templates").Funcs(template.FuncMap{
		"calculatePreviousPage": func(currentPageNumber int) int {
			if currentPageNumber == 1 {
				return 0
			}

			return currentPageNumber - 1
		},

		"calculateNextPage": func(currentPageNumber, totalPage int) int {
			if currentPageNumber == totalPage {
				return 0
			}

			return currentPageNumber + 1
		},
	}).Funcs(sprig.FuncMap())

	tmpl := template.Must(templates.ParseFS(h.templateFiles, "*/*/*.html", "*/*.html", "*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
