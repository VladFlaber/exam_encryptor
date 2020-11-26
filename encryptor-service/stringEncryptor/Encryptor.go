package stringEncryptor

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Sha256Encryptor struct{
	workersCount int
}
type Job struct {
	ID int
	Str string
}
//Encrypt one string to sha256
func (e Sha256Encryptor) doJob(workerId int,Jobs <-chan Job, results chan <-string)  {
	for j:=range Jobs{
		fmt.Println("Worker # " , workerId , "received string # " , j.ID , " with value : " , j.Str)
		hash:= sha256.New()
		hash.Write([]byte(j.Str))
		res:=base64.URLEncoding.EncodeToString(hash.Sum(nil))
		fmt.Println("Worker # " , workerId, "encrypted string # " , j.ID , " to : " , res)
		results <-res
	}
}
//SetWorkersCount
//set the environment variable ,named "encr_wrkcnt", for worker pool count. If count <2 the function sets 4 by default
func (e Sha256Encryptor) SetWorkersCount(count int)  {
		if count<2{count =4}
		os.Setenv(workersCountEVName,strconv.Itoa(count))
		e.workersCount=count
}
//GetWorkersCount
// Returns count of workers , red from the environment variable ,named "encr_wrkcnt".
// If variable is empty  , it will  create the EV with default value 4
func (e Sha256Encryptor) GetWorkersCount() int{
	val,ok:=os.LookupEnv(workersCountEVName)
	if !ok {
		e.SetWorkersCount(4)
		return e.GetWorkersCount()
	} else{
		count,_:=strconv.Atoi(val)
		e.workersCount=count
		return e.workersCount
	}
}
//creates workers pool which will encrypt list of string
func (e Sha256Encryptor) EncryptStrings(listRandomStrings []string ) (*[]string,*string) {
	jobsCount := len (listRandomStrings)
//	workersCount :=e.GetWorkersCount()
	jobs := make(chan Job, jobsCount)
	results := make(chan string, jobsCount)
	defer close(results)
	for w := 0; w < 20; w++ {
		go e.doJob(w, jobs, results)
	}
	for j := 0; j < jobsCount; j++ {
		jobs <- Job{ID: j,Str: listRandomStrings[j]}
	}
	close(jobs)
	var arr []string
	for i:=0 ; i<jobsCount; i++ {
		arr=append(arr,<-results)//receiving encrypted strings from the results channel
	}
	fmt.Println(len(arr), " strings encrypted")
	// misunderstanding with task
	oneStr:=strings.Join(arr,"")
	return  &arr,&oneStr
}

