package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Oloruntobi1/bookinfo/book_merce/bookmerce"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const port = ":50051"

type server struct {
	pb.UnimplementedBookMerceServer
	bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {

	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Book ID", err)
	}
	in.Id = out.String()
	
	if s.bookMap == nil {
		s.bookMap = make(map[string]*pb.Book)
	}
	s.bookMap[in.Id] = in
	
	return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()

}

// GetBook implements bookmerce.GetBook
func (s *server) GetBook(ctx context.Context, in *pb.BookID) (*pb.Book, error) {

	value, exists := s.bookMap[in.Value]
		if exists {
			return value, status.New(codes.OK, "").Err()
		}
		return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
	log.Fatalf("failed to listen: %v", err)
	}
		s:= grpc.NewServer()
		pb.RegisterBookMerceServer(s, &server{})
	
	log.Printf("Starting gRPC listener on port " + port)
		if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
}


}