#!/usr/bin/env bash

#WORKDIR=$(docker-compose --file docker/docker-compose.yml run --rm -u utente php pwd)
WORKDIR=/home/utente/matventure
PROJECT_NAME=$(basename $(pwd) | tr  '[:upper:]' '[:lower:]')
COMPOSE_OVERRIDE=

if [[ -f "docker/docker-compose.override.yml" ]]; then
    COMPOSE_OVERRIDE="--file docker/docker-compose.override.yml"
fi

if [[ "$1" = "up" ]]; then

    # shift 1

    # mkdir -p docker/php/git/
    # cp ~/.gitconfig docker/php/git/

    # if [[ ! -d docker/php/ssh ]]; then
    #     mkdir -p docker/php/ssh
    #     cp ~/.ssh/id_rsa docker/php/ssh/
    #     cp ~/.ssh/id_rsa.pub docker/php/ssh/
    # fi

    docker-compose \
        --file docker/docker-compose.yml \
        ${COMPOSE_OVERRIDE} \
        -p ${PROJECT_NAME} \
        up $@

elif [[ "$1" = "enter-root" ]]; then

    docker-compose \
        --file docker/docker-compose.yml \
        ${COMPOSE_OVERRIDE} \
        -p ${PROJECT_NAME} \
        run \
        --name ${PROJECT_NAME} \
        -u root \
        golang /bin/zsh

elif [[ "$1" = "enter" ]]; then

    docker-compose \
        --file docker/docker-compose.yml \
        ${COMPOSE_OVERRIDE} \
        -p ${PROJECT_NAME} \
        run \
        --name ${PROJECT_NAME} \
        -u utente \
        -w ${WORKDIR} \
        golang /bin/zsh

elif [[ "$1" = "down" ]]; then

    shift 1
    docker-compose \
	    --file docker/docker-compose.yml \
	    ${COMPOSE_OVERRIDE} \
	    -p ${PROJECT_NAME} \
		down $@

elif [[ "$1" = "purge" ]]; then

    docker-compose \
	    --file docker/docker-compose.yml \
	    ${COMPOSE_OVERRIDE} \
	    -p ${PROJECT_NAME} \
		down \
        --rmi=all \
        --volumes \
        --remove-orphans

elif [[ "$1" = "log" ]]; then

    docker-compose \
        --file docker/docker-compose.yml \
        ${COMPOSE_OVERRIDE} \
        -p ${PROJECT_NAME} \
        logs -f

elif [[ $# -gt 0 ]]; then
    docker-compose \
        --file docker/docker-compose.yml \
        ${COMPOSE_OVERRIDE} \
        -p ${PROJECT_NAME} \
        "$@"

else
    docker-compose \
        --file docker/docker-compose.yml \
        ${COMPOSE_OVERRIDE} \
        -p ${PROJECT_NAME} \
        ps
fi
