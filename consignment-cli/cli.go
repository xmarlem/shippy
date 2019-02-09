package main

import (
	"log"
	pb "github.com/xmarlem/shippy/consignment-service/proto/consignment"
	"io/ioutil"
	"encoding/json"
	"os"
	"context"
	"google.golang.org/grpc"
)

const (
	address = "192.168.99.100:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error){
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err!=nil{
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}


func main(){

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	file := defaultFilename
	if len(os.Args)>1{
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err!=nil{
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err!=nil{
		log.Fatalf("Could not green: %v", err)
	}

	log.Printf("Created: %t", r.Created)


	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err!=nil{
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments{
		log.Println(v)
	}
}
