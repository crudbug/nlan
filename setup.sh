#!/bin/bash

HOMEDIR=$GOPATH/src/github.com/araobp/nlan 

echo "Compiling tega driver..."
go install github.com/araobp/tega/driver

echo "Compiling NLAN model..."
MODELDIR=$HOMEDIR/model/nlan
protoc -I $MODELDIR $MODELDIR/nlan.proto --go_out=plugins=grpc:$MODELDIR

echo "Building NLAN master..."
cd master
go build -o master
cd $HOMEDIR

echo "Building containers with NLAN agent embedded..."
cd docker
go build docker_mng.go
./restart.sh
cd $HOMEDIR

echo "Done!"
echo ""
