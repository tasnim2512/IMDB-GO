package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	adminpb "practice/IMDB/gunk/v1/admin"
	userpb "practice/IMDB/gunk/v1/user"
	ag "practice/IMDB/usermgm/core/admin"
	ca "practice/IMDB/usermgm/core/user"
	"practice/IMDB/usermgm/service/user"
	"practice/IMDB/usermgm/service/admin"
	"practice/IMDB/usermgm/storage/postgres"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:embed migrations
var migrationFiles embed.FS

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

	port := config.GetString("server.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("%v", err)
	}
	postGresStorage, err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatal(err)
	}

	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}
	if err := goose.Up(postGresStorage.DB.DB, "migrations"); err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()

	userCore := ca.NewCoreUser(postGresStorage)
	userSvc := user.NewUserSvc(userCore)
	userpb.RegisterUserServiceServer(grpcServer, userSvc)

	adminCore := ag.NewCoreAdmin (postGresStorage)
	adminSvc := admin.NewAdminSvc(adminCore)
	adminpb.RegisterAdminServiceServer(grpcServer, adminSvc)

	reflection.Register(grpcServer)
	fmt.Println("usermgm server running at ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve port %v", err)
	}
}
