package main

import (
	"database/sql"
	"fmt"
	"github.com/abrorbeksoft/auth-service/config"
	"github.com/abrorbeksoft/auth-service/genproto/auth"
	"github.com/abrorbeksoft/auth-service/service"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
)

//type AuthServer struct {
//
//}
//
//func (a AuthServer) HelloUser(ctx context.Context, message *auth.Message) (*auth.Message, error) {
//	return &auth.Message{
//		Message: fmt.Sprintf("Hello user %s", message.Name),
//		Name: "Ahmadboy",
//	}, nil
//}

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age int `json:"age"`
}

func main()  {
	cfg:=config.Load()

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",cfg.DBUsername,cfg.DBDatbase,cfg.DBPassword,cfg.DBHost)
	db,err:=sql.Open("postgres",connStr)

	if err!=nil {
		fmt.Println(err)
	}
	defer db.Close()

	authService:=service.NewAuthService(db)

	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",cfg.AuthServiceHost,cfg.AuthServicePort))
	if err!=nil {
		fmt.Println(err)
	}

	grpcServer:=grpc.NewServer();
	auth.RegisterAuthSerivceServer(grpcServer,authService)


	if err:=grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %s",err)
	}

}