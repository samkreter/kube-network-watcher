# Image URL to use all building/pushing image targets
IMG ?= pskreter/kube-network-watcher:v0.0.1-alpha

# Run tests
test: fmt vet
	go test ./... -coverprofile=cover.out -covermode=atomic

# Generate release artifects for latest tag
release-latest:
	cd config/controller && kustomize edit set image ${IMG}
	kustomize build config/controller > deploy/deploy.yaml

# Deploy latest build image to k8s cluster
.PHONY: deploy
deploy:
	cd config/controller && kustomize edit set image ${IMG}
	kustomize build config/controller | kubectl apply -f -

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push: docker-build
	docker push ${IMG}
