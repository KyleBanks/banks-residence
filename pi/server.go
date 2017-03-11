package main

import (
	"log"
	"net/http"

	//cv "github.com/lazywei/go-opencv/opencv"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
)

const (
	httpPort = ":8080"
)

func main() {
	go updateIp()

	camera := opencv.NewCameraDriver(0)
	work := func() {
		camera.On(camera.Event("frame"), func(data interface{}) {
			log.Printf("FRAME: %v", data)
		})
	}

	robot := gobot.NewRobot("cameraBot",
		[]gobot.Device{camera},
		work,
	)
	robot.Start()

	log.Printf("Listening on Port %v", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
