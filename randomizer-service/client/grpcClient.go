package client

import (
	"context"
	pr "exam_encryptor/encryptor-service/proto"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
}
func (c* Client) Response(message pr.ListOfStrings) ([]string ,string){

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error   ", err)
	}
	defer conn.Close()
	client := pr.NewStringsServiceClient(conn)
	//message := pr.ListOfStrings{List: randomStrings}
	response, err := client.GetSha256Strings(context.Background(), &message)
	if err != nil {
		log.Fatal("error occurred while was trying to get response : ", err)
	}
	return response.List,response.OneStr
}