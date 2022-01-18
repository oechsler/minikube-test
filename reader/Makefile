.PHONY: build
build: ##Builds the reader services Docker image
	minikube image build -t reader:latest .

.PHONY: deploy
deploy: ##Deploys the reader service to the minikube cluster
	make build

	kubectl apply -f .minikube/0-namespace.yaml
	kubectl apply -f .minikube/1-service.yaml
	kubectl apply -f .minikube/2-deployment.yaml

.PHONY: clean
clean: ##Removes the reader service from the minikube cluster
	kubectl delete -f .minikube/2-deployment.yaml
	kubectl delete -f .minikube/1-service.yaml
	kubectl delete -f .minikube/0-namespace.yaml

.PHONY: update
update: ##Updates the reader service on the minikube cluster
	make build
	make stop
	make start

.PHONY: stop
stop: ##Stops the reader services deployment on the minikube cluster
	kubectl scale deploy -n reader reader-deployment --replicas=0

.PHONY: start
start: ##Starts the reader services deployment on the minikube cluster
	kubectl apply -f .minikube/2-deployment.yaml

.PHONY: help
help: ##Show the help information
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'