NAME = horodata.www
DOCKERID = 392789183703.dkr.ecr.us-east-1.amazonaws.com

all: bindata
	go build

clean:
	rm -fr horo

test: bindata
	ginkgo -r \
	./api \
	./config \
	./helpers \
	./middlewares \
	./models \
	./services \
	./www

bindata: mails templates

mails:
	go-bindata -o=./services/mail/messages/bin.go \
	-pkg=messages \
	-prefix="services/mail/messages" \
	./services/mail/messages/*

templates:
	go-bindata -o=./html/bin.go \
	-pkg=html \
	-prefix="html" \
	./html/**/*.html ./html/*.html

container: clean bindata
	GOOS=linux GOARCH=amd64 go build
	docker build -t $(DOCKERID)/$(NAME) .

push:
	docker push $(DOCKERID)/$(NAME)

.PHONY: build push
