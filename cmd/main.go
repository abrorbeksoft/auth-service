package main

import (
	"context"
	"fmt"
	"github.com/abrorbeksoft/auth-service/config"
	"github.com/abrorbeksoft/auth-service/genproto/auth"
	"google.golang.org/grpc"
	"net"
)

type AuthServer struct {

}

func (a AuthServer) HelloUser(ctx context.Context, message *auth.Message) (*auth.Message, error) {
	return &auth.Message{
		Message: fmt.Sprintf("Hello user %s", message.Name),
		Name: "Ahmadboy",
	}, nil
}

func main()  {
	cfg:=config.Load()
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",cfg.AuthServiceHost,cfg.AuthServicePort))
	if err!=nil {
		fmt.Println(err)
	}
	s:=AuthServer{}
	grpcServer:=grpc.NewServer();
	auth.RegisterRouteGuideServer(grpcServer,&s)
	if err:=grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %s",err)
	}
}