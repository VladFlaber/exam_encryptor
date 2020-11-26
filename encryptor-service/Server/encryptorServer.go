package encryptorServer

import (
	"context"
	pr "exam_encryptor/encryptor-service/proto"
	"exam_encryptor/encryptor-service/stringEncryptor"
	"log"
)
type Server struct {
	pr.UnimplementedStringsServiceServer
}
func (s *Server) GetSha256Strings(ctx context.Context,message *pr.ListOfStrings) (*pr.Response,error)  {
	log.Println("server received ",len(message.List) ," strings")
	encryptor:=stringEncryptor.Sha256Encryptor{}
	encryptedStrings,str :=s.processData(message.List,encryptor)
	return &pr.Response{List: encryptedStrings,OneStr: str},nil
}

func (s Server) processData(strings []string ,encryptor stringEncryptor.EncryptorWorker) ([]string,string)  {
	List,oneStr:=encryptor.EncryptStrings(strings)
	return *List,*oneStr
}






