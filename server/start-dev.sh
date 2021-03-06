#!/bin/bash

# colors holds variables responsible for stdout colors
declare -A colors=(
  ["warning"]="\e[33m\e[1m" # for warning
  ["info"]="\e[36m\e[1m" # for info
  ["normal"]="\e[0m"        # for normal (for making text normal again)
)

export DEV_MODE="true"
export MONGO_URI="mongodb://localhost:27017"
export DB_NAME="houseofbosons-dev"
export TEST_DB_NAME="houseofbosons-test"
export GOOGLE_AUTH_REDIRECT_URL="http://localhost:8080/api/auth/google/callback"
export ADMIN_ORIGIN="http://localhost:3000"

# reading ./secrets.config file, that includes secrets formatted as the following
# SECRET_ENV_1=something;
# SECRET_ENV_2=something_more;
# SECRET_ENV_3=something_more_more;

# reading the contents of file
SECRETS=$(<dev-keys.secret)
# replacing `\n` from SECRETS string
# SECRETS=$(echo $SECRETS|tr --delete '\n')
SECRETS=$(echo $SECRETS|tr -d '\n')
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
declare -a secrets=("GOOGLE_CLIENT_ID" "GOOGLE_CLIENT_SECRET" "AUTHORIZED_EMAIL" "SESSION_KEY" "MONGO_CLUSTER_URI" "MONGO_CLUSTER_DB_NAME" "MONGO_CLUSTER_USER" "MONGO_CLUSTER_PASSWORD") 

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
