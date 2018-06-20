#!/bin/bash

. ./scripts/common.sh

ARGS_NUMBER="$#"
COMMAND="$1"

usage_message="Useage: $0 start | restart | clean | status | stop | cleanpeer"

function verifyArg() {
    if [ $ARGS_NUMBER -ne 1 ]; then
        echo $usage_message
        exit 1;
    fi
}

function pullDockerImages(){

  for IMAGES in ca peer orderer ccenv tools; do
      echo "==> FABRIC IMAGE: $IMAGES"
      echo
      docker pull hyperledger/fabric-$IMAGES:$FABRIC_VERSION
      docker tag hyperledger/fabric-$IMAGES:$FABRIC_VERSION hyperledger/fabric-$IMAGES
  done

  docker pull hyperledger/fabric-couchdb:$COUCHDB_VERSION
  docker tag hyperledger/fabric-couchdb:$COUCHDB_VERSION hyperledger/fabric-couchdb
}

function startNetwork() {

    echo
    echo "----------------------------"
    echo "--- Starting the network ---"
    echo "----------------------------"
    docker-compose -f ./fabric-network.yaml up -d
}

function cleanNetwork(){
    docker rm -f $(docker ps -aq)
    docker rmi -f $(docker images | grep hyperledger | tr -s ' ' | cut -d ' ' -f 3)
    docker rmi -f $(docker images | grep aladdin | tr -s ' ' | cut -d ' ' -f 3)
    docker rmi -f $(docker images | grep chainhero | tr -s ' ' | cut -d ' ' -f 3)

}

function networkStatus(){
    docker ps --format "table {{.Names}}\t{{.Status}}" --filter "name=genesis.cn"
}

function networkRestart(){
    docker rm -f $(docker ps -aq)
    startNetwork
}

function stopNetwork(){
    docker-compose -f ./fabric-network.yaml stop	
}

function cleanPeer(){
    docker rm -f $(docker ps -aq)
    docker rmi -f $(docker images | grep hyperledger/fabric-peer | tr -s ' ' | cut -d ' ' -f 3)	
}

verifyArg
case $COMMAND in
    "start")
        pullDockerImages
        startNetwork
        ;;
    "restart")
        networkRestart
        ;;
    "status")
        networkStatus
        ;;
    "clean")
        cleanNetwork
        ;;
    "stop")
	stopNetwork
	;;
    "cleanpeer")
	cleanPeer
	;;	  				
    *)
        echo $usage_message
        exit 1
esac
