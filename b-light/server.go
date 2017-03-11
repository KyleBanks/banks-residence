package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	gpioPinNumber = 4
	httpPort      = ":8080"
)

func main() {
	go updateIp()

	l, err := newLight(gpioPinNumber)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		l.toggleState()

		out := "ON"
		if l.State != On {
			out = "OFF"
		}
		fmt.Fprintf(w, out)
	})

	log.Printf("Listening on Port %v", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}

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

// package main

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/ec2"
// )

// func main() {
// 	sess, err := session.NewSession()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create an EC2 service object in the "us-west-2" region
// 	// Note that you can also configure your region globally by
// 	// exporting the AWS_REGION environment variable
// 	svc := ec2.New(sess, &aws.Config{Region: aws.String("us-west-2")})

// 	// Call the DescribeInstances Operation
// 	resp, err := svc.DescribeInstances(nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// resp has all of the response data, pull out instance IDs:
// 	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
// 	for idx, res := range resp.Reservations {
// 		fmt.Println("  > Number of instances: ", len(res.Instances))
// 		for _, inst := range resp.Reservations[idx].Instances {
// 			fmt.Println("    - Instance ID: ", *inst.InstanceId)
// 		}
// 	}
// }
