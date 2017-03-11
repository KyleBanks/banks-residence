# Builds and deploys the pi application to the Raspberry PI.
pi:
	cd pi && GOARM=7 GOARCH=arm GOOS=linux go build github.com/KyleBanks/banks-residence/pi
	
	scp pi/pi pi@192.168.0.200:/home/pi/
	rm pi/pi
.PHONY: pi
