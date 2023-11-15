API_URL = http://localhost

create_cluster:
		@echo  "\e[1;34mCreating cluster\e[0m"
		kind create cluster --config kube/kind.yaml
		kind get kubeconfig > ~/.kube/kind-config
		KUBECONFIG=~/.kube/kind-config kubectl apply -f kube/ingress.yaml
		KUBECONFIG=~/.kube/kind-config kubectl wait --namespace ingress-nginx \
		--for=condition=ready pod \
		--selector=app.kubernetes.io/component=controller \
		--timeout=90s
		@echo  "\e[1;32mCluster created\e[0m"
		@echo "To use kubectl on the new cluster, please run: export KUBECONFIG=~/.kube/kind-config"

delete_cluster:
		kind delete cluster
		@echo "To use again your default kube config file run: export KUBECONFIG=~/.kube/config"
		@echo  "\e[1;31mCluster delete\e[0m"

deploy:
		kind load docker-image config-api:v0.1
		KUBECONFIG=~/.kube/kind-config kubectl apply -f kube/app.yaml
		@echo  "\e[1;32mApplication deployed\e[0m"

list:
		@curl -X GET -s '$(API_URL)/configs'
		@echo
create:
		curl  -i -X POST -d @examples/create.json $(API_URL)/configs
		@echo
get:
		@curl -X GET $(API_URL)/configs/$(PARAM)
		@echo

updatePUT:
		curl -i -X PUT -d @examples/updatePUT.json $(API_URL)/configs/$(PARAM)
		@echo

updatePATCH:
		curl -i  -X PATCH -d @examples/updatePATCH.json $(API_URL)/configs/$(PARAM)
		@echo

delete:
		curl -i -X DELETE $(API_URL)/configs/$(PARAM)
		@echo

query:
		curl -X GET '$(API_URL)/search?metada.calories=230'
		@echo
		@echo
		curl -X GET '$(API_URL)/search?metada.monitoring.enabled=true'
		@echo

test:
		cd go && go test -v 
