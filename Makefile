NAME = horodata.www
DOCKERID = 392789183703.dkr.ecr.us-east-1.amazonaws.com

all: templates
	go build

clean:
	rm -fr horo

test: templates
	ginkgo -r \
	./api \
	./config \
	./helpers \
	./middlewares \
	./models \
	./services \
	./www

templates:
	go-bindata -o=./html/bin.go \
	-pkg=html \
	-prefix="html" \
	./html/**/**/*.html ./html/**/*.html ./html/*.html

container: clean templates
	GOOS=linux GOARCH=amd64 go build
	docker build -t $(DOCKERID)/$(NAME) .

push:
	docker push $(DOCKERID)/$(NAME)

.PHONY: build container push templates
