/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	// "fmt"
	"google.golang.org/grpc"
	pb "login/LoginTest"
	"log"
	"net"
)

const (
	port = ":9090"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// counter counts how many post messages it has received in the current session
var counter int

// SayHello implements helloworld.GreeterServer
func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	username := in.GetUsername()
	password := in.GetPassword()
	var message string
	counter++
	if Login(username, password) == true {
		message = "Access granted, welcome " + username
	} else {
		message = "Acces denied"
	}
	return &pb.LoginReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
