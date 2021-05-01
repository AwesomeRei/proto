#!/usr/bin/env sh

# ROOT=$(pwd)
# if [ -n "$1" ]
# then
#     ROOT="$1"
# fi

# cd "${ROOT}" 

GO_OUT="$(pwd)/go"
# PROTO_PATH="$(pwd)/proto"
# PROTO_VERSIONS_PATHS=$(find $PROTO_PATH -type d -regex ".*/v[0-9]+")

if [ -z $GOPATH ]
then
  GOPATH=$(go env GOPATH)
fi

if [ -z $GRPC_GW_PROTO_PATH ]
then
  GRPC_GW_PROTO_PATH="${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"
fi

mkdir -p $GO_OUT

echo 'Clean up old go files.'
# rm -rf ${GO_OUT:?}/*

echo 'Generating Go code from protobuf...'
# for v in ${PROTO_VERSIONS_PATHS}
# do
    echo "Process proto files in: $v"
    # protos="$v/*.proto"
protoc -I . -I${GRPC_GW_PROTO_PATH} --go_out=go $1     
protoc -I . -I${GRPC_GW_PROTO_PATH} --go-grpc_out=go $1
protoc -I . -I${GRPC_GW_PROTO_PATH} --grpc-gateway_out go $1

    # protoc $protos -I${PROTO_PATH} -I. -I${GRPC_GW_PROTO_PATH} --go_out=plugins=grpc:${GO_OUT} || exit 1
    # protoc $protos -I${PROTO_PATH} -I. -I${GRPC_GW_PROTO_PATH} --grpc-gateway_out=${GO_OUT} || exit 1
    # TODO: add swagger
# done
echo 'Done.'
