package main

import (
 "fmt"
 pr "github.com/VladFlaber/exam_encryptor/encryptor-service/proto"
 "github.com/VladFlaber/exam_encryptor/randomizer-service/client"
 "github.com/VladFlaber/exam_encryptor/randomizer-service/stringGenerator"
 "log"
 "os"
 "strconv"
)
//environment variable for workers count. Used here just for comfortability
const workersCountEVName string="ENCR_WRKCNT"
func main() {
 //ENCRYPTOR-SERVICE REQUIRED BEING STARTED
 // args[1] = count of wanted strings , args[2] = charset .
 // if none of args declared application starts simple console ui to get this data from stdin
 if len(os.Args)==1{
   InteractiveMenu()
 } else if len(os.Args)==2 {
  count,err:=strconv.Atoi(os.Args[1])
  if err != nil {
    log.Println(err)
  }
  //charset:=os.Args[2]
  charset:=""
  createAndSend(count,charset,stringGenerator.StringGenerator{})
 }
}
//
func createAndSend(stringsCount int ,charset string ,generator stringGenerator.IStringGenerator)   {
 var randomStrings []string
 randomStrings = *generator.CreateListOfStrings(stringsCount,charset)
 fmt.Println(stringsCount, " strings created ")
 for index, value := range randomStrings {
  fmt.Println(index+1, ": ", value)
 }
 if randomStrings != nil && len(randomStrings) > 0 {

  message:= pr.ListOfStrings{
   List: randomStrings,
  }
  cl := client.Client{}
  list,onestr:=cl.Response(message)
  fmt.Println("\n\n\n encrypted strings received")
  for index, value := range list {
   fmt.Println(index+1, ": ", value)
  }
  fmt.Println("\n\n\n encrypted string ")
 fmt.Println(onestr)
 }
}

// simple ui to interact with user
func InteractiveMenu() {
 const menuOutput string = "1. Create and send random strings  to encryptor-service\n" +
     "2. Set workers count for encryptor-service \n" +
     "3. Exit \n"
 var choice string = ""
 for choice != "3" {
  fmt.Println(menuOutput)
  fmt.Println("enter your choice")
  fmt.Scanln(&choice)
  switch choice {
  case "1":
   {
    fmt.Println("Enter amount of strings you want to create")
    var stringsCount int
    var charset string=""
    fmt.Scanln(&stringsCount)
    fmt.Println("Enter charset u want to use in generator")
    fmt.Scanln(&charset)
    createAndSend(stringsCount,charset,stringGenerator.StringGenerator{})
    }
    //Just for comfortability for task demonstration
  case "2":
   {
    fmt.Println("Current amount of workers = ", os.Getenv(workersCountEVName))
    fmt.Println("Enter new count of workers")
    var count int
    fmt.Scanln(&count)
    os.Setenv(workersCountEVName, strconv.Itoa(count))
    break
   }
  case "3":
   {
    os.Exit(0)
    break
   }
  default:
   fmt.Println("incorrect input")
  }
 }
}
