#!/bin/bash

# colors holds variables responsible for stdout colors
declare -A colors=(
  ["warning"]="\e[33m\e[1m" # for warning
  ["info"]="\e[36m\e[1m" # for info
  ["normal"]="\e[0m"        # for normal (for making text normal again)
)

export POSTGRES_HOST="localhost"
export POSTGRES_PORT="5432"
export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="password"
export POSTGRES_DB_NAME="postgres"

export GOOGLE_AUTH_REDIRECT_URL="http://localhost:8080/api/auth/google/callback"


# reading ./secrets.config file, that includes secrets formatted as the following
# SECRET_ENV_1=something;
# SECRET_ENV_2=something_more;
# SECRET_ENV_3=something_more_more;

# reading the contents of file
SECRETS=$(<dev-keys.secret)
# replacing `\n` from SECRETS string
SECRETS=$(echo $SECRETS|tr --delete '\n')
# adding a ` ` at the end of the string, so later can be splitted at `; `
SECRETS=$(printf "${SECRETS} ")

# splitting teh string from `; `
IFS='; ' read -ra EACH <<< "${SECRETS}"
# for each chunk
for i in "${EACH[@]}"; do
    # again splitting at `=`
    IFS='=' read -ra PART <<< "$i"
    # eventually exporting the environment variable
    export "${PART[0]}"="${PART[1]}"
done


# secret environment variables that are expected to be declared
declare -a secrets=("GOOGLE_CLIENT_ID" "GOOGLE_CLIENT_SECRET")

# checking if secret config variables exists or not
for i in "${secrets[@]}"
do
  if [[ ! -v "$i" ]]; then
    # warning if some secret variable is not present in the environment
    echo -e "${colors[warning]}WARN${colors[normal]} environment variable not found - \"${i}\""
  else
    :
  fi
done


# starting the go server
echo -e "${colors[info]}INFO${colors[normal]} starting the server..."
go run main.go
