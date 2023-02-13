package main

import (
	"os"

	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func Start() error {
	svc := s3.New(session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           os.Getenv("AWS_PROFILE"),
		SharedConfigState: session.SharedConfigEnable,
	})))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		input := &s3.ListBucketsInput{}
		result, err := svc.ListBuckets(input)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, result.GoString())
	})
	return e.Start("0.0.0.0:1323")
}

func main() {
	logrus.Fatal(Start())
}
