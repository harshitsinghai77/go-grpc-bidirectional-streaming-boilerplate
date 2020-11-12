package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	pb "github.com/harshitsinghai/grpc/protofiles"
)

func main() {
	//Employees details
	employees := &pb.Employees{
		Employees: []*pb.Employee{
			{
				Id:     "1234",
				Name:   "Harshit Singhai",
				Salary: 40000,
			},
			{
				Id:     "1235",
				Name:   "Neil Patrick",
				Salary: 80000,
			},
			{
				Id:     "1236",
				Name:   "Robert Andrew James",
				Salary: 100000,
			},
		},
	}

	p1 := pb.Employee{}

	data, err := proto.Marshal(employees)
	fmt.Println(string(data))
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println("Raw data", data)

	err = proto.Unmarshal(data, &p1)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("Original struct loaded from proto file:", employees)

	for _, person := range employees.Employees {
		fmt.Println("==============")
		fmt.Println(person.GetId())
		fmt.Println(person.GetName())
		fmt.Println(person.GetSalary())
	}
}
