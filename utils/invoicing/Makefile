NAME = horodata.invoicing
DOCKERID = 392789183703.dkr.ecr.us-east-1.amazonaws.com

all:
	go build

clean:
	rm -fr invoicing

container: clean bindata
	GOOS=linux GOARCH=amd64 go build
	docker build -t $(DOCKERID)/$(NAME) .

push:
	docker push $(DOCKERID)/$(NAME)

.PHONY: build container push
