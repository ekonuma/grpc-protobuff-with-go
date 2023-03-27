package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ekonuma/grpc-protobuff-with-go/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Eike",
		Email:       "konumaeike5618@gmail.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Eike",
		},
		Birthday: &pb.Date{
			Year:  2001,
			Month: 7,
			Day:   12,
		},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize, err")
	}

	if err := ioutil.WriteFile("test.bin", binData, 0127); err != nil {
		log.Fatalln("Can't write", err)
	}

	in, err := ioutil.ReadFile("test.bin")

	if err != nil {
		log.Fatalf("Can't read file")
	}

	readEmployee := &pb.Employee{}

	err = proto.Unmarshal(in, readEmployee)

	if err != nil {
		log.Fatalln("Can't deserialize, err")
	}

	fmt.Println(readEmployee)
}
