package main

import (
	"fmt"
	"log"
	"net"
	// adminpb "practice/IMDB/gunk/v1/admin"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

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
	if err !=nil{
		log.Fatalf("%v", err)
	}

	grpcServer := grpc.NewServer()
	// adminpb.RegisterAdminServiceServer(grpcServer,)
	fmt.Println("usermgm server running at ", lis.Addr())
	if err := grpcServer.Serve(lis) ; err!=nil{
		log.Fatalf("unable to serve port %v", err)
	}
}
