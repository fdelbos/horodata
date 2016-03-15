NAME = horodata.static
DOCKERID = 392789183703.dkr.ecr.us-east-1.amazonaws.com

all: compress build push

compress:
	uglifyjs --compress --mangle -- app/horodata.js > app/horodata.min.js

build:
	docker build -t $(DOCKERID)/$(NAME) .

push:
	docker push $(DOCKERID)/$(NAME)

.PHONY: all build push
