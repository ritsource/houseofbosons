#!/bin/bash

export POSTGRES_HOST="localhost"
export POSTGRES_PORT="5432"
export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="password"
export POSTGRES_DB_NAME="postgres"

go run main.go