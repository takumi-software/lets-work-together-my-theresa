#!/bin/bash

export PROTO_PATH=/go/src/github.com/takumi-software/lets-work-together-my-theresa

echo "Going to the repository directory"
cd ${PROTO_PATH};
echo "Formatting the proto files inside ${PROTO_PATH}"
prototool format -w .
echo "Generating the proto files inside ${PROTO_PATH}"
prototool generate
echo "Generating the Gateway proto files"
protoc -I${PROTO_PATH} -I${PROTO_PATH}/third_party/googleapis  --grpc-gateway_out=logtostderr=true:protos/go ${PROTO_PATH}/mytheresa/api/products/products.proto