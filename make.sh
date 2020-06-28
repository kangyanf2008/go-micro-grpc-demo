#!/bin/bash -x

if [ ! -f "$0" ]; then
    echo 'make must be run within its container folder' 1>&2
    exit 1
fi


if [ "$1" == "clean" ]; then
        rm bin/* pkg/* utility/pkg/* -rf
        exit
fi

CURDIR=`pwd`
export GOPATH="$CURDIR"
export GOBIN=${GOPATH}/bin
author=${USER}
date=`date "+%Y-%m-%d_%H:%I:%S"`

goversion=`go version`

#客户端
serverLdflags="-X server/logic._AUTHOR_=$author -X server/logic._COMPILETIME_=\"$date\""
echo clientLdflags ${clientLdflags}
#服务端
clientLdflags="-X client/logic._AUTHOR_=$author -X client/logic._COMPILETIME_=\"$date\""
echo clientLdflags ${clientLdflags}

echo "formating code..."
gofmt -w src/

# 以下命令可以使用go get golang.org/x/tools/cmd/goimports获取
#${GOROOT}/bin/goimports -w=true src/

cd ./src/
#${GOROOT}/bin/go install -v -clientLdflags "$clientLdflags"  main/main.go
go install -v -ldflags "$clientLdflags"  main/client.go
if [ $? == 0 ]; then
	echo "build client success"
else
	echo "build client error"
fi
go install -v -ldflags "$serverLdflags"  main/server.go
if [ $? == 0 ]; then
	echo "build server success"
else
	echo "build client error"
fi

