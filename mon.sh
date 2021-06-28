#!/bin/bash

help() {
    echo -e "Usage: go <command>"
    echo -e
    echo -e "help                Print this help"
    echo -e
    echo -e "Supported commands:"
    echo -e "    run             Run Endpoint Monitor"
    # echo -e "    test            Run unit tests"
    echo -e "    stop            Stop Endpoint Monitor"
    echo -e "    rebuild         Delete previous Docker containers and images related to Endpoint Monitor and rebuild from scratch"
    echo ""
    exit 0
}

run() {
    type docker-compose > /dev/null 2>&1 || { echo >&2 "docker-compose must be installed"; exit 1; }

    if [[ -z "$1" ]]; then
      export GRAFANA_PORT=3000
    else
      export MY_SCRIPT_VARIABLE="$1"
    fi

    if [[ -z "$2" ]]; then
      export PROMETHEUS_PORT=9090
    else
      export PROMETHEUS_PORT="$2"
    fi

    if [[ -z "$3" ]]; then
      export METRICS_PORT=2112
    else
      export METRICS_PORT="$3"
    fi

    # Alter environment variables for prometheus.yml
    sed -i '' -e "s/PROMETHEUS_PORT/$PROMETHEUS_PORT/" -e "s/METRICS_PORT/$METRICS_PORT/" prometheus-config/prometheus.yml

    #cat gocd/kubernetes/varnish-api-cache.yaml | envsubst > tmp/manifest.yaml
    docker-compose -f docker-compose.yml up

}

stop() {
    type docker-compose > /dev/null 2>&1 || { echo >&2 "docker-compose must be installed"; exit 1; }

    docker-compose -f docker-compose.yml down
}

rebuild() {
    type docker >/dev/null 2>&1 || { echo >&2 "docker must be installed"; exit 1; }

    docker container rm -f $(docker container ls -aq) || true
    docker image rm -f $(docker image ls 'go-web-monitor_go-mon' -q) || true
}


if [[ $1 =~ ^(help|run|stop|rebuild)$ ]]; then
  COMMAND=$1
  shift
  $COMMAND "$@"
else
  help
  exit 1
fi
