package handler

import (
	"log"
	"net/http"
	userpb "practice/IMDB/gunk/v1/user"

	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"
)

type RegisterUser struct {
	FirstName string
	LastName  string
	Email     string
	UserName  string
	Password  string
}

func (u RegisterUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&u.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&u.UserName,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&u.Email,
			validation.Required.Error("The email field is required."),
			is.Email.Error("The email field must be a valid email."),
		),
		validation.Field(&u.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}

type UserRegistrationForm struct {
	User      RegisterUser
	FormError map[string]error
	CSRFToken string
}

func (h Handler) UserRegistration(w http.ResponseWriter, r *http.Request) {
	h.parseRegisterTemplate(w, UserRegistrationForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) UserRegistrationPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	form := UserRegistrationForm{}
	ru := RegisterUser{}
	if err := h.decoder.Decode(&ru, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	form.User = ru
	if err := ru.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			for key, val := range vErr {
				form.FormError[strings.Title(key)] = val
			}
		}
		h.parseRegisterTemplate(w, form)
		return
	}

	_, err := h.usermgmSvc.Register(r.Context(), &userpb.RegisterRequest{
		FirstName: ru.FirstName,
		LastName:  ru.LastName,
		UserName:  ru.UserName,
		Email:     ru.Email,
		Password:  ru.Password,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h Handler) parseRegisterTemplate(w http.ResponseWriter, form UserRegistrationForm) {
	t := h.Templates.Lookup("user-registration.html")
	if t == nil {
		log.Println("unable to lookup register template")
		h.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, form); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
