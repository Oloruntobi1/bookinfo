package main

import (
	pb "github.com/Oloruntobi1/bookinfo/book_merce/bookmerce"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	
	defer conn.Close()
	c := pb.NewBookMerceClient(conn)
	name := "Chronicles of a Happy People"
	author:= "Wole Soyinka"
	language:= "English"
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddBook(ctx,
		&pb.Book{Name: name, Author: author, Language: language})
	if err != nil {
		log.Fatalf("Could not add book: %v", err)
	}
	
	log.Printf("Book ID: %s added successfully", r.Value)
	book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get book: %v", err)
	}
	log.Printf("Book: ", book.String())
}
