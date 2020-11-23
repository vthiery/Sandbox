#!/usr/bin/env bash

execution() {
    echo "Creating cassandra keyspace & table"
    cqlsh -f $(dirname "${BASH_SOURCE[0]}")/schema.cql
}

until execution; do
    echo "Cassandra is still initializing, will retry later.."
    sleep 10
done &

exec /docker-entrypoint.sh "$@"