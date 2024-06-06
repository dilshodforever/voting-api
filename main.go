package main

import (
	"fmt"
	"log"
	"root/api"
	"root/api/handler"
	pb "root/genprotos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":50051"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()
	ProductConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":50052"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer ProductConn.Close()

	us := pb.NewPublicServiceClient(UserConn)
	ps := pb.NewPartyServiceClient(ProductConn)

	h := handler.NewHandler(us, ps)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run()
	if err!=nil{
		log.Fatal("Error while Run: ", err.Error())
	}
}
