package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"practice/IMDB/cms/handler"
	"practice/IMDB/utility"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	// "google.golang.org/grpc"
)

//go:embed assets
var assetFiles embed.FS

// //go:embed migrations
// var migrationFiles embed.FS

var sessionManager *scs.SessionManager

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	decoder := form.NewDecoder()
	postGresStorage, err := utility.NewPostgresStorage(config)
	if err != nil {
		log.Fatal(err)
	}

	// goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}
	// if err := goose.Up(postGresStorage.DB.DB, "migrations"); err != nil {
	// 	log.Fatalln(err)
	// }
	lt := config.GetDuration("session.lifetime")
	it := config.GetDuration("session.idletime")
	sessionManager = scs.New()
	sessionManager.Lifetime = lt * time.Hour
	sessionManager.IdleTimeout = it * time.Minute
	sessionManager.Cookie.Name = "student_session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true
	sessionManager.Store = utility.NewSqlxStore(postGresStorage.DB)

	var assetFS = fs.FS(assetFiles)
	staticFiles, err := fs.Sub(assetFS, "assets/src")
	if err != nil {
		log.Fatal(err)
	}

	templateFiles, err := fs.Sub(assetFS, "assets/templates")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	urmgmUrl := config.GetString("usermgm.url")
	usermgmConn, err := grpc.Dial(urmgmUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%#v", err)
	}

	chi := handler.NewHandler(sessionManager, decoder, usermgmConn, staticFiles, templateFiles)
	nosurfHandler := nosurf.New(chi)
	p := config.GetString("server.port")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", p))
	if err != nil {
		log.Fatalf("%#v", err)
	}

	fmt.Println("cms server running at:", lis.Addr())
	if err := http.Serve(lis, nosurfHandler); err != nil {
		log.Fatalf("%#v", err)
	}
}
