#!/bin/bash

BLUE='\033[0;34m'
GREEN='\033[0;32m'
ORANGE='\033[0;33m'
RED='\033[0;31m'
NO_COLOR='\033[0m'

function log {
    msg="$2"
    component="$3"
    if [[ $1 == "INFO" ]]; then
        printf "${BLUE}INFO:${NO_COLOR} %s %s \n" "$component" "$msg"
    elif [[ $1 == "DONE" ]]; then
        printf "${GREEN}SUCCESS:${NO_COLOR} %s %s \n" "$component" "$msg"
    elif [[ $1 == "WARN" ]]; then
        printf "${ORANGE}WARNING:${NO_COLOR} %s %s \n" "$component" "$msg"
    else
        printf "${RED}FAILED:${NO_COLOR} %s %s \n" "$component" "$msg"
    fi
}

# build the shell-client (TO-DO)
# cd client && ./build.sh

cd client/ || exit

# Add controller to lighthouse (TO-DO)
# ICOS_AUTH_TOKEN=$ICOS_AUTH_TOKEN go run main.go --config=config_client.yml add controller -a "192.168.1.1" -n "test2"

# Get controllers from lighthouse
unset CONTROLLERS
CONTROLLERS=$(go run main.go --config=config_client.yml get controller 2> /dev/null) 
if [ "$CONTROLLERS" != ""  ]; then
    log "DONE" "Controllers retrieved successfully" [lighthouse]
else
    log "FAIL" "Error while retrieving controllers" [lighthouse]
fi

# healthcheck shell-backend from controller
unset CONTROLLER
CONTROLLER="localhost:8080" go run main.go --config=config_client.yml 2> /dev/null
if [ $? -eq 0 ]; then
    log "DONE" "Healthcheck to the shell-backend was successful" [shell-backend]
else
    log "FAIL" "Error when trying healthcheck to the shell-backend" [shell-backend]
fi

# Login to keycloak
unset ICOS_AUTH_TOKEN
eval `go run main.go --config=config_client.yml auth login 2> /dev/null` 
if [ "$ICOS_AUTH_TOKEN" != "" ]; then
  log "DONE" "Token returned successfully" [keycloak]
else
  log "FAIL" "failed to retrieve the token" [keycloak]
fi

# Create deployment (TO-DO)
# ICOS_AUTH_TOKEN=$ICOS_AUTH_TOKEN go run main.go --config=config_client.yml create deployment --file any_file.yml
