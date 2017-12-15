
PROJECT_ROOT?=$(shell pwd)/../../
PROJECT_PKG:=marathon
TARGET_IMAGE=marathonserver:0.1

all: image

binary:
	CGO_ENABLED=0 go build -o /go/src/$(PROJECT_PKG)/dist/cargo-assistant $(PROJECT_PKG)

target:
	mkdir -p $(PROJECT_ROOT)/dist
	docker run --rm -ti -v $(PROJECT_ROOT):/go/src/$(PROJECT_PKG) \
	  -w /go/src/$(PROJECT_PKG)/image/marathonserver golang:1.9.2 \
	  make binary

image: target
	temp=`mktemp -d` && \
	cp -af $(PROJECT_ROOT)/dist/cargo-assistant $$temp && \
	cp -af Dockerfile $$temp && \
	docker build -t $(TARGET_IMAGE) $$temp && \
	rm -rf $$temp

clean:

.PHONY: image all clean target
