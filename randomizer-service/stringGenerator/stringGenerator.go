package stringGenerator

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultCharset= "qwertyuiopasdfghjklzxcvbnm"+
	"QWERTYUIOPASDFGHJKLZXCVBNM"+
	"0123456789"
//generate random string using passed charset
type StringGenerator struct {}

//CreateListOfStrings
//Create slice of random strings with size of stringsCount.
//If charset passed as "" (empty string) , the generator will use default charset (upper + lower case letters + 0-9 numbers)
func (s StringGenerator)CreateListOfStrings(stringsCount int,charset string) *[]string {
	if charset==""||charset=="\n"{charset=defaultCharset}
	var strings []string
	rand.Seed(time.Now().UnixNano())
	max,min:=20,1
	fmt.Println(charset)
	c :=make(chan string , stringsCount)
	defer close(c)
	for i:=0;i<stringsCount;i++{
		//creating single random string
		go func() {
			length:=rand.Intn(max - min + 1) + min
			res :=make([]byte ,length)
			for  i :=0 ; i < length ; i++ {
				res[i]=charset[rand.Intn(len(charset))]
			}
			c<-string(res)
		}()
		strings=append(strings,<-c)
	}
	return  &strings
}

