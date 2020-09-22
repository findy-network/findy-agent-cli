#!/bin/bash
# e2e-test.sh

CLI=$GOPATH/bin/findy-agent-cli

CURRENT_DIR=$(dirname "$BASH_SOURCE")

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[1;94m'
NC='\033[0m'

set -e

clean() {
  echo -e "${BLUE}*** dev - clean ***${NC}"
  echo -e "${RED}WARNING: erasing all local data stored by indy!${NC}"
  rm -rf ~/.indy_client/
  echo "{}" >findy.json
  set +e
  rm findy.bolt
  docker stop findy-pool
  docker rm findy-pool
  docker volume rm sandbox
  set -e
}

dev_ledger() {
  echo -e "${BLUE}*** dev - start dev ledger ***${NC}"
  docker run -itd -p 9701-9708:9701-9708 \
    -p 9000:9000 \
    -v sandbox:/var/lib/indy/sandbox/ \
    --name findy-pool \
    optechlab/indy-pool-browser:latest
}

set_envs() {
    export FCLI_POOL_NAME="findy"
    export FCLI_POOL_GENESIS_TXN_FILE="${CURRENT_DIR}/genesis_transactions"

    export FCLI_STEWARD_POOL_NAME="findy"
    export FCLI_STEWARD_SEED="000000000000000000000000Steward1"
    export FCLI_STEWARD_WALLET_NAME="sovrin_steward_wallet"
    export FCLI_STEWARD_WALLET_KEY="9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY"

    export FCLI_AGENCY_POOL_NAME="findy"
    export FCLI_AGENCY_STEWARD_WALLET_NAME="sovrin_steward_wallet"
    export FCLI_AGENCY_STEWARD_WALLET_KEY="9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY"
    export FCLI_AGENCY_STEWARD_DID="Th7MpTaRZVRYnPiabds81Y"
    export FCLI_AGENCY_STEWARD_SEED="000000000000000000000000Steward1"
    export FCLI_AGENCY_SALT="my_test_salt"

    export FCLI_SERVICE_WALLET_NAME="test_wallet1"
    export FCLI_SERVICE_WALLET_KEY="9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY"
    export FCLI_SERVICE_AGENCY_URL="http://localhost:8080"
    export FCLI_ONBOARD_EMAIL="test_email1"
    export FCLI_ONBOARD_EXPORT_FILE="${CURRENT_DIR}/test_wallet1.export"
    export FCLI_ONBOARD_EXPORT_KEY="9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY"
    export FCLI_ONBOARD_SALT="my_test_salt"

    export FCLI_SCHEMA_NAME="my_schema1"
    export FCLI_SCHEMA_VERSION="2.0"
    export FCLI_SCHEMA_ATTRIBUTES="[\"field1\", \"field2\", \"field3\"]"
}

agency_env() {
  # set env variables
  set_envs
  # remove and reset all stored data
  clean

  # install latest version of findy-agency
  make install

  # launch and create pool
  dev_ledger
  echo -e "${BLUE}*** env - create pool ***${NC}"
  $CLI ledger pool create
  echo -e "${BLUE}*** env - create steward ***${NC}"
  $CLI ledger steward create

  # run agency
  echo -e "${BLUE}*** env - run agency ***${NC}"
  docker start findy-pool
  $CLI agency start
}

cmds_env() {
  set_envs

  # onboard
  echo -e "${BLUE}*** env - onboard ***${NC}"
  set +e
  rm ${CURRENT_DIR}/test_wallet1.export
  rm -rf ~/.indy_client/wallet/test_wallet1
  rm -rf ~/.indy_client/wallet/test_email1
  set -e
  $CLI service onboard

  # create schema
  echo -e "${BLUE}*** env - create schema ***${NC}"
  $CLI service schema create
}

agency_conf() {
  # remove and reset all stored data
  clean

  # install latest version of findy-agency
  make install

  # launch and create pool
  dev_ledger
  echo -e "${BLUE}*** conf - create pool ***${NC}"
  $CLI ledger pool create \
    --config=${CURRENT_DIR}/configs/createPool.yaml \
    --genesis-txn-file=${CURRENT_DIR}/genesis_transactions
  echo -e "${BLUE}*** conf - create steward ***${NC}"
  $CLI ledger steward create \
    --config=${CURRENT_DIR}/configs/createSteward.yaml

  # run agency
  echo -e "${BLUE}*** conf - run agency ***${NC}"
  docker start findy-pool
  $CLI agency start --config=${CURRENT_DIR}/configs/startAgency.yaml
}

cmds_conf() {
  # onboard
  echo -e "${BLUE}*** conf - onboard ***${NC}"
  set +e
  rm ${CURRENT_DIR}/test_wallet2.export
  rm -rf ~/.indy_client/wallet/test_wallet2
  rm -rf ~/.indy_client/wallet/test_email2
  set -e
  $CLI service onboard \
    --config=${CURRENT_DIR}/configs/onboard.yaml \
    --export-file=${CURRENT_DIR}/test_wallet2.export

  # create schema
  echo -e "${BLUE}*** conf - create schema ***${NC}"
  $CLI service schema create \
    --config=${CURRENT_DIR}/configs/createSchema.yaml
}

agency_flag() {
  # remove and reset all stored data
  clean

  # install latest version of findy-agency
  make install

  # launch and create pool
  dev_ledger
  echo -e "${BLUE}*** flag - create pool ***${NC}"
  $CLI ledger pool create \
    --name=findy \
    --genesis-txn-file=${CURRENT_DIR}/genesis_transactions
  echo -e "${BLUE}*** flag - create steward ***${NC}"
  $CLI ledger steward create \
    --pool-name=findy \
    --seed=000000000000000000000000Steward1 \
    --wallet-name=sovrin_steward_wallet \
    --wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY

  # run agency
  echo -e "${BLUE}*** flag - run agency ***${NC}"
  docker start findy-pool
  $CLI agency start \
    --pool-name=findy \
    --steward-wallet-name=sovrin_steward_wallet \
    --steward-wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
    --steward-did=Th7MpTaRZVRYnPiabds81Y \
    --steward-seed=000000000000000000000000Steward1 \
    --salt=my_test_salt
}

cmds_flag() {
  # onboard
  echo -e "${BLUE}*** flag - onboard ***${NC}"
  set +e
  rm ${CURRENT_DIR}/test_wallet3.export
  rm -rf ~/.indy_client/wallet/test_wallet3
  rm -rf ~/.indy_client/wallet/test_email3
  set -e
  $CLI service onboard \
    --export-file=${CURRENT_DIR}/test_wallet3.export \
    --export-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
    --agency-url=http://localhost:8080 \
    --wallet-name=test_wallet3 \
    --wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
    --email=test_email3 \
    --salt=my_test_salt

  # create schema
  echo -e "${BLUE}*** flag - create schema ***${NC}"
  $CLI service schema create \
    --wallet-name=test_wallet3 \
    --agency-url=http://localhost:8080 \
    --wallet-key=9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
    --name=my_schema3 \
    --version="2.0" \
    --attributes=["field1", "field2", "field3"]
}

cmds() {
  cmds_flag
  cmds_conf
  cmds_env
}
"$@"
