package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Start() error {
	svc := s3.New(session.New())
    input := &s3.ListBucketsInput{}
    result, _ := svc.ListBuckets(input)
    fmt.Println(result)

}


func main() {
	Start()
}