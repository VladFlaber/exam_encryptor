package main

import (
	encryptorServer "github.com/VladFlaber/exam_encryptor/encryptor-service/Server"
	pr "github.com/VladFlaber/exam_encryptor/encryptor-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main()  {
	address,ok:=os.LookupEnv("")
	if !ok{
		address = ":9000"
	}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server started listening : ",address)
	ser:=encryptorServer.Server{}
	grpcServer := grpc.NewServer()
	pr.RegisterStringsServiceServer(grpcServer, &ser)
	grpcServer.Serve(lis)
 }
