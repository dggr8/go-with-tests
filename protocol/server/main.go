package main

import (
	"net/http"

	pb "github.com/dggr8/go-with-tests/protocol/server/userpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	u := &pb.User{
		Name:  "Inigo Montoya",
		Id:    1234,
		Email: "inigo@montoya.example.com",
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/x-protobuf")
	res.Write(body)
}
