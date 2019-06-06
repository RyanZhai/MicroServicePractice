package main

import (
	"log"
	"os"

	"Ethan/MicroServicePractice/common"
	"Ethan/MicroServicePractice/config"
	pb "Ethan/MicroServicePractice/interface-center/out/user"

	"context"
)

const service = "user"

var (
	serviceName string
)

func init() {
	serviceName = config.GetServiceName(service)
}

func main() {
	srv := common.GetMicroClient(service)

	// 创建 user-service 微服务的客户端
	client := pb.NewUserServiceClient(serviceName, srv.Client())

	name := "Ethan"
	email := "test@gmail.com"
	password := "test123"
	company := "test company"

	resp, err := client.Create(context.Background(), &pb.User{
		Id:       "5cf7c408afe05ce7e837f2a4",
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("call Create error: %v", err)
	}
	log.Println("created: ", resp.User.Id)

	allResp, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("call GetAll error: %v", err)
	}
	for i, u := range allResp.Users {
		log.Printf("user_%d: %v\n", i, u)
	}

	authResp, err := client.Auth(context.Background(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("auth failed: %v", err)
	}
	log.Println("token: ", authResp.Token)

	os.Exit(0)
}
