
PROJECT_ROOT?=$(shell pwd)/../../
PROJECT_PKG:=marathon/groupmanager
TARGET_IMAGE=groupmanager:0.1

all: image

binary:
	CGO_ENABLED=0 go build -o /go/src/marathon/dist/groupmanager $(PROJECT_PKG)

target:
	mkdir -p $(PROJECT_ROOT)/dist
	docker run --rm -ti -v $(PROJECT_ROOT):/go/src/marathon \
	  -w /go/src/marathon/image/groupmanager golang:1.9.2 \
	  make binary

image: target
	temp=`mktemp -d` && \
	cp -af $(PROJECT_ROOT)/dist/groupmanager $$temp && \
	cp -af Dockerfile $$temp && \
	docker build -t $(TARGET_IMAGE) $$temp && \
	rm -rf $$temp

clean:

.PHONY: image all clean target
