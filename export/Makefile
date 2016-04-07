NAME = horodata.export
DOCKERID = 392789183703.dkr.ecr.us-east-1.amazonaws.com

all: build push

build:
	docker build -t $(DOCKERID)/$(NAME) .

push:
	docker push $(DOCKERID)/$(NAME)

.PHONY: all build push
