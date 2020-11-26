ARG go_version=1.13.14

FROM golang:$go_version AS build

ARG protoc_version=3.12.4
ARG grpc_version=1.29.1
ARG protoc_gen_go_version=1.25.0
ARG protoc_gen_go_grpc_version=0.0.0-20200806181047-b830b5f36107

RUN apt update && apt install -y \
    wget \
    unzip

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v${protoc_version}/protoc-${protoc_version}-linux-x86_64.zip \
    && unzip protoc-${protoc_version}-linux-x86_64.zip \
    && cp ./bin/protoc /usr/local/bin/protoc \
    && cp -r ./include/* /usr/local/include/ \
    && chmod +x /usr/local/bin/protoc

ENV GO111MODULE=on
RUN git clone https://github.com/googleapis/googleapis
RUN go get -u google.golang.org/grpc@v${grpc_version}
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v${protoc_gen_go_version}
RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${protoc_gen_go_grpc_version}
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v${protoc_gen_go_gateway_version}
RUN go get -u github.com/favadi/protoc-go-inject-tag@v${protoc_go_inject_tag_version}
RUN BIN="/usr/local/bin" && \
    VERSION="0.20.5" && \
    BINARY_NAME="buf" && \
      curl -sSL \
        "https://github.com/bufbuild/buf/releases/download/v${VERSION}/${BINARY_NAME}-$(uname -s)-$(uname -m)" \
        -o "${BIN}/${BINARY_NAME}" && \
      chmod +x "${BIN}/${BINARY_NAME}"
