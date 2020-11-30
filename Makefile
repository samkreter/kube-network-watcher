# Image URL to use all building/pushing image targets
IMG ?= pskreter/kube-network-watcher:v0.0.1-alpha

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile=cover.out -covermode=atomic

# Generate release artifects for latest tag
release-latest: manifests
	kustomize build config/crd > deploy/latest/crd.yaml
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default > deploy/latest/deploy.yaml

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
docker-push:
	docker push ${IMG}
