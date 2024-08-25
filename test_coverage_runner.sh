#!/bin/sh

# Environment variables
UNIT_COV=tmp/unit_coverage.out
INTEGRATION_COV=tmp/integration_coverage.out
OUTPUT_COV=tmp/coverage.out

# echo_info echoes messages in blue
echo_info() {
    echo -e "\\033[1;34m$1\\033[0m"
}

# Cleaning test cache
go clean -testcache && \

echo_info "### Running unit tests" && \

# Running unit tests and collecting coverage
ENV=test go test -v -short -coverprofile=$UNIT_COV ./... && \

echo_info "### Running integration tests" && \

# Running integration tests and collecting coverage
ENV=test go test -v -p 1 -coverprofile=$INTEGRATION_COV ./... -run Integration && \

echo_info "### Building $OUTPUT_COV"

# Building output coverage file
head -n 1 $UNIT_COV > $OUTPUT_COV && \
tail -n +2 $UNIT_COV >> $OUTPUT_COV && \
tail -n +2 $INTEGRATION_COV >> $OUTPUT_COV && \

go tool cover -func=$OUTPUT_COV && \

echo_info "Done"