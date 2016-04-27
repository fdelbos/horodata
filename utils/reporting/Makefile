NAME = horodata.reporting
DOCKERID = 392789183703.dkr.ecr.us-east-1.amazonaws.com

all:
	go build

container:
	GOOS=linux GOARCH=amd64 go build
	docker build -t $(DOCKERID)/$(NAME) .

push:
	docker push $(DOCKERID)/$(NAME)

.PHONY: build container push
