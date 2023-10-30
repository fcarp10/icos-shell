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

# Login to keycloak
unset ICOS_AUTH_TOKEN
COMPONENTS="[shell-backend --> keycloak]"
eval `CONTROLLER="localhost:8080" go run main.go --config=config_client.yml auth login 2> /dev/null` 
if [ "$ICOS_AUTH_TOKEN" != "" ]; then
  log "DONE" "Token returned successfully" "$COMPONENTS"
else
  log "FAIL" "failed to retrieve the token" "$COMPONENTS"
fi

# Add controller to lighthouse
COMPONENTS="[lighthouse --> keycloak]"
RESPONSE=$(ICOS_AUTH_TOKEN=$ICOS_AUTH_TOKEN go run main.go --config=config_client.yml add controller -a 127.0.0.1 -n helloworld_controller 2> /dev/null)
if [[ $RESPONSE == "201" ]]; then
    log "DONE" "Controller added to the lighthouse successfully" "$COMPONENTS"
elif [[ $RESPONSE == "202" ]]; then
    log "INFO" "Controller already exists in the lighthouse (reset timeout)" "$COMPONENTS"
else
    log "FAIL" "Error while trying to add a controller to the lighthouse" "$COMPONENTS"
fi

# Get controllers from lighthouse
unset CONTROLLERS
COMPONENTS="[lighthouse]"
CONTROLLERS=$(go run main.go --config=config_client.yml get controller 2> /dev/null) 
if [ "$CONTROLLERS" != ""  ]; then
    log "DONE" "Controllers retrieved successfully" "$COMPONENTS"
else
    log "FAIL" "Error while retrieving controllers" "$COMPONENTS"
fi

# healthcheck shell-backend from controller
unset CONTROLLER
COMPONENTS="[shell-backend]"
CONTROLLER="localhost:8080" go run main.go --config=config_client.yml 2> /dev/null
if [ $? -eq 0 ]; then
    log "DONE" "Healthcheck to the shell-backend was successful" "$COMPONENTS"
else
    log "FAIL" "Error while trying healthcheck to the shell-backend" "$COMPONENTS"
fi

# Create deployment
COMPONENTS="[shell-backend --> job-manager]"
RESPONSE=$(ICOS_AUTH_TOKEN=$ICOS_AUTH_TOKEN CONTROLLER="localhost:8080" go run main.go --config=config_client.yml create deployment --file ../openapi.yaml 2> /dev/null)
if [[ $RESPONSE ]]; then
    log "DONE" "Deployment added to the controller successfully" "$COMPONENTS"
else
    log "FAIL" "Error while trying to add a deployment to the controller" "$COMPONENTS"
fi

# Get deployment
COMPONENTS="[shell-backend --> job-manager]"
RESPONSE=$(ICOS_AUTH_TOKEN=$ICOS_AUTH_TOKEN CONTROLLER="localhost:8080" go run main.go --config=config_client.yml get deployment 2> /dev/null)
if [[ $RESPONSE ]]; then
    log "DONE" "Deployments retrieved successfully" "$COMPONENTS"
else
    log "FAIL" "Error while retrieving deployments" "$COMPONENTS"
fi