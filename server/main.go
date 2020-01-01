package main

import (
	"context"
	"google.golang.org/grpc/reflection"
	"math"
	"google.golang.org/grpc"
	"golang-grpc-boilerplate/proto"
	"net"
)

type server struct {}


func (s *server) Add(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := a+b
	return &proto.Response{Ans:result}, nil
}

func (s *server) Subtract(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := a-b
	return &proto.Response{Ans:result}, nil
}

func (s *server) Multiply(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := a*b
	return &proto.Response{Ans:result}, nil
}

func (s *server) Divide(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := a/b
	return &proto.Response{Ans:result}, nil
}

func (s *server) Power(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := math.Pow(a,b)
	return &proto.Response{Ans:result}, nil
}

func main() {

	listner, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	svr := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(svr, &server{})
	reflection.Register(svr)
	if e := svr.Serve(listner); e!=nil {
		panic(err)
	}
}
