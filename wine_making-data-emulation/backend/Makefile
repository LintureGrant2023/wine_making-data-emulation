.PHONY: all docker push

all: 
	go build -o backend main.go
docker:
	docker build . -t xiwen2001/backend:v2.6
push: 
	docker push xiwen2001/backend:v2.6
