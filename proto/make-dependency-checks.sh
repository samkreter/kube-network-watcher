#!/usr/bin/env bash

verify_docker_present() {
  if [[ -z "$(command -v docker)" ]]; then
    echo "Can't find 'docker' in PATH, please fix and retry."
    return 2
  fi
}

verify_docker_image_build() {
  proto_build_image=$(docker image list | grep $1 | awk '{print $1}')
  proto_build_version=$(docker image list | grep $1 | awk '{print $2}')
  if [[ ${proto_build_image} != $1 ]] || [[ ${proto_build_version} != $2 ]]; then
    echo "Building Dockerfile..."
    docker build --tag $1:$2 ./images/
  fi
}
