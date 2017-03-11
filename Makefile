# Builds and deploys the b-light application to the Raspberry PI.
b-light:
	cd b-light && GOARM=7 GOARCH=arm GOOS=linux go build github.com/KyleBanks/banks-residence/b-light
	
	scp b-light/b-light pi@192.168.0.200:/home/pi/
	rm b-light/b-light
.PHONY: b-light
