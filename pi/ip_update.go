package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// updateIp periodically updates the server's remote IP in a Dynamo table.
func updateIp() {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	firstRun := true
	for {
		if !firstRun {
			time.Sleep(time.Minute * 3)
		}
		firstRun = false

		// Get the IP address
		log.Println("Updating IP...")
		resp, err := http.Get("http://ipv4bot.whatismyipaddress.com/")
		if err != nil {
			log.Println("Error fetching IP:", err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading body:", err)
			continue
		}
		log.Println("Found IP:", string(body))

		// Upload it as a file to S3
		log.Println("Uploading...")
		s3res, err := svc.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String("banks-residence"),
			Key:         aws.String("ip.txt"),
			Body:        bytes.NewReader(body),
			ContentType: aws.String("text/plain"),
		})
		if err != nil {
			log.Println("Error uploading IP:", err)
			continue
		}
		log.Println("Upload complete:", s3res)
	}
}
