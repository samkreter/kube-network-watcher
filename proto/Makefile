SHELL := /bin/bash

IMAGE_NAME     = protobuild
IMAGE_VERSION  = 1.0
PROTOFILES     = $(shell find ./ -iname "*.proto" | grep -v vendor)
DOCKER_RUN_CMD = docker run --volume "$(shell pwd):/workspace" --workdir /workspace ${IMAGE_NAME}:${IMAGE_VERSION}
PROTOC_CMD     = ${DOCKER_RUN_CMD} protoc
INSERT_TAG_CMD = ${DOCKER_RUN_CMD} protoc-go-inject-tag
BUF_LINT_CMD   = ${DOCKER_RUN_CMD} buf check lint

all: docker-checks proto-compile
docker-checks: check_docker_dependency build_docker_image

.PHONY: proto-compile
proto-compile: proto_lint
	for file in ${PROTOFILES} ; do \
		${PROTOC_CMD} --proto_path=./ --go-grpc_out=paths=source_relative,requireUnimplementedServers=false:./ --go_out=paths=source_relative:./ $${file}; \
	done

.PHONY: check_docker_dependency
check_docker_dependency:
	@bash -c "source ./make-dependency-checks.sh && verify_docker_present"

.PHONY: build_docker_image
build_docker_image:
	@bash -c "source ./make-dependency-checks.sh && verify_docker_image_build ${IMAGE_NAME} ${IMAGE_VERSION}"

.PHONY: proto_lint
proto_lint:
	${BUF_LINT_CMD}
