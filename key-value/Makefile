.PHONY: build
build: ##Builds the key-value services Docker image
	minikube image build -t key-value:latest .

.PHONY: deploy
deploy: ##Deploys the key-value service to the minikube cluster
	make build

	kubectl apply -f .minikube/0-namespace.yaml
	kubectl apply -f .minikube/1-service.yaml
	kubectl apply -f .minikube/2-deployment.yaml

.PHONY: clean
clean: ##Removes the key-value service from the minikube cluster
	kubectl delete -f .minikube/2-deployment.yaml
	kubectl delete -f .minikube/1-service.yaml
	kubectl delete -f .minikube/0-namespace.yaml

.PHONY: update
update: ##Updates the key-value service on the minikube cluster
	make build
	make stop
	make start

.PHONY: stop
stop: ##Stops the key-value services deployment on the minikube cluster
	kubectl scale deploy -n key-value key-value-deployment --replicas=0

.PHONY: start
start: ##Starts the key-value services deployment on the minikube cluster
	kubectl apply -f .minikube/2-deployment.yaml

.PHONY: help
help: ##Show the help information
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'