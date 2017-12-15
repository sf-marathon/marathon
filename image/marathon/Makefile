
PROJECT_ROOT?=$(shell pwd)/../../
TARGET_IMAGE=marathon:latest

all: image

image:
	temp=`mktemp -d` && \
	cp -af $(PROJECT_ROOT)/webapp/dist $$temp && \
	cp -af nginx.conf $$temp && \
	cp -af Dockerfile $$temp && \
	docker build -t $(TARGET_IMAGE) $$temp && \
	rm -rf $$temp

clean:

.PHONY: image all clean
