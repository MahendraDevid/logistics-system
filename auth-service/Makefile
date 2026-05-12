build:
	docker build -t imamaziz/auth-service:v1 .

push:
	docker push imamaziz/auth-service:v1

deploy:
	kubectl apply -f deployments/kubernetes/

all: build push deploy