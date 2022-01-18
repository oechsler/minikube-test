.PHONY: build
build: ##Builds the reader services Docker image
	docker build -t reader:latest .
	minikube image load reader:latest

.PHONY: deploy
deploy: ##Deploys the reader service to the minikube cluster
	make build

	kubectl apply -f .k8s/0-namespace.yaml
	kubectl apply -f .k8s/1-service.yaml
	kubectl apply -f .k8s/2-deployment.yaml

.PHONY: stop
stop: ##Stops the reader services deployment on the minikube cluster
	kubectl scale deploy -n reader reader-deployment --replicas=0

.PHONY: start
start: ##Starts the reader services deployment on the minikube cluster
	kubectl apply -f .k8s/2-deployment.yaml

.PHONY: dev
dev: ## Runs the reader services dev environment on the minikube cluster
	okteto up -f .okteto/0-devenv.yaml
	
	make prod

.PHONY: prod
prod: ##Removes the reader services dev environment from the minikube cluster
	okteto down -f .okteto/0-devenv.yaml

.PHONY: update
update: ##Updates the reader service on the minikube cluster
	make prod
	make build
	make stop
	make start

.PHONY: clean
clean: ##Removes the reader service from the minikube cluster
	make prod

	kubectl delete -f .k8s/2-deployment.yaml
	kubectl delete -f .k8s/1-service.yaml
	kubectl delete -f .k8s/0-namespace.yaml

.PHONY: help
help: ##Show the help information
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'