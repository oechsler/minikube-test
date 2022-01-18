.PHONY: build
build: ##Builds the key-value services Docker image
	docker build -t key-value:latest .
	minikube image load key-value:latest

.PHONY: deploy
deploy: ##Deploys the key-value service to the minikube cluster
	make build

	kubectl apply -f .k8s/0-namespace.yaml
	kubectl apply -f .k8s/1-service.yaml
	kubectl apply -f .k8s/2-deployment.yaml

.PHONY: stop
stop: ##Stops the key-value services deployment on the minikube cluster
	kubectl scale deploy -n key-value key-value-deployment --replicas=0

.PHONY: start
start: ##Starts the key-value services deployment on the minikube cluster
	kubectl apply -f .k8s/2-deployment.yaml

.PHONY: dev
dev: ## Runs the key-value services dev environment on the minikube cluster
	okteto up -f .okteto/0-devenv.yaml

	make prod

.PHONY: prod
prod: ##Removes the key-value services dev environment from the minikube cluster
	okteto down -f .okteto/0-devenv.yaml

.PHONY: update
update: ##Updates the key-value service on the minikube cluster
	make prod
	make build
	make stop
	make start

.PHONY: clean
clean: ##Removes the key-value service from the minikube cluster
	make prod

	kubectl delete -f .k8s/2-deployment.yaml
	kubectl delete -f .k8s/1-service.yaml
	kubectl delete -f .k8s/0-namespace.yaml

.PHONY: help
help: ##Show the help information
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'