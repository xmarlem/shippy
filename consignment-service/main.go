package main

import (
	pb "github.com/xmarlem/shippy/consignment-service/proto/consignment"
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)


type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct{
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error){
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

type service struct {
	repo IRepository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error){
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created:true, Consignment:consignment}, nil
}

func main(){
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	pb.RegisterShippingServiceServer(srv, &service{repo})

	reflection.Register(srv)
	if err := srv.Serve(lis); err!=nil{
		log.Fatalf("Failed to serve: %v", err)
	}

}
