#!/bin/bash

rm -rf output/
rm -rf release/
mkdir -p output/
mkdir -p release/
P_NAME=ws_forward

VerStr=""
VersionStrInit() {
    VERSION_FIRST=$(cat version/version.go | grep "VersionFirst" |grep -v "fmt"|awk '{print $3}')
    VERSION_SECOND=$(cat version/version.go | grep "VersionSecond"|grep -v "fmt"|awk '{print $3}')
    VERSION_THIRD=$(cat version/version.go | grep "VersionThird" | grep -v "fmt"|awk '{print $3}')
    VerStr=${VERSION_FIRST}.${VERSION_SECOND}.${VERSION_THIRD}
    echo "version init to:${VerStr}"
}

Build() {
    OS=${1}
    ARCH=${2}

    rm -rf ./output/*

    GOOS="${OS}" GOARCH="${ARCH}" go build -o output/${P_NAME}
    if [ $? -ne 0 ]; then
		    echo "build fail"
		    exit 1
		fi
    cd ./output/ || return

    ArchType=nuknown
    Bits=32
    if [ "${ARCH}" == "amd64" ]; then
		    ArchType="x86"
		    Bits=64
	  fi

	  if [ "${ARCH}" == "unknown" ]; then
	      echo "not support this arch"
	      exit 1
	  fi
    tar -czf ../release/dm_"${ArchType}"_"${Bits}"_"${VerStr}".tgz ./*
    cd ..
}

VersionStrInit

Build linux amd64
