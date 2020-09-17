#!/bin/bash
# dev.sh

CLI=$GOPATH/bin/findy-agent-cli

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

set -e

clean() {
  echo -e "${GREEN}*** dev - clean ***${NC}"
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

run() {
  # run agency
  echo -e "${GREEN}*** dev - run agency ***${NC}"
  docker start findy-pool
  $CLI agency start \
    --pool-name findy \
    --steward-wallet-name sovrin_steward_wallet \
    --steward-wallet-key 9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY \
    --steward-did Th7MpTaRZVRYnPiabds81Y
}

scratch() {
  CURRENT_DIR=$(dirname "$BASH_SOURCE")

  # remove and reset all stored data
  clean

  # install latest version of findy-agency
  make install

  # launch and create pool
  echo -e "${GREEN}*** dev - start dev ledger ***${NC}"
  docker run -itd -p 9701-9708:9701-9708 \
    -p 9000:9000 \
    -v sandbox:/var/lib/indy/sandbox/ \
    --name findy-pool \
    optechlab/indy-pool-browser:latest
  echo -e "${GREEN}*** dev - create pool ***${NC}"
  $CLI ledger pool create \
    --name findy \
    --genesis-txn-file $CURRENT_DIR/genesis_transactions
  echo -e "${GREEN}*** dev - create steward ***${NC}"
  $CLI ledger steward create \
    --pool-name findy \
    --seed 000000000000000000000000Steward1 \
    --wallet-name sovrin_steward_wallet \
    --wallet-key 9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY

  run
}

install_run() {
  make install

  run
}

onboard() {
  make install

  echo -e "${GREEN}*** dev - onboard ***${NC}"

  # example
  # ./tools/dev.sh onboard myName 9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY .
  EXPORT_NAME=$1
  EXPORT_KEY=$2
  EXPORT_DIR=$3
  echo "name: $EXPORT_NAME, key: $EXPORT_KEY, dir: $EXPORT_DIR"
  set +e
  rm $EXPORT_DIR/${EXPORT_NAME}.export
  rm -rf ~/.indy_client/wallet/${EXPORT_NAME}_client
  rm -rf ~/.indy_client/wallet/${EXPORT_NAME}_server
  set -e
  $AGENT client handshakeAndExport \
    -wallet ${EXPORT_NAME}_client \
    -email ${EXPORT_NAME}_server \
    -pwd ${EXPORT_KEY} \
    -url http://localhost:8080 \
    -exportpath ${EXPORT_DIR}/${EXPORT_NAME}.export
}

"$@"
