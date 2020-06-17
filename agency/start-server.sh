#!/bin/bash

if [ -f "/tmp/env" ]; then
  echo "Adding environment variables from configuration file."
  eval $(cat /tmp/env)
  cp /tmp/genesis_transactions /genesis_transactions
  cp /tmp/steward.exported /steward.exported
  cp /tmp/aps.p12 /aps.p12
fi

FOLDER=~/.indy_client/wallet/$WALLET_NAME/
if [ -d "$FOLDER" ]; then
  echo "$FOLDER exists"
else
  echo "$FOLDER does not exist, importing wallet"
  echo "pool create findy gen_txn_file=/genesis_transactions" >$1/indy-cli.cmd
  echo "wallet import $WALLET_NAME key=$WALLET_KEY export_path=/steward.exported export_key=$IMPORTED_WALLET_KEY" >>$1/indy-cli.cmd
  echo "exit" >>$1/indy-cli.cmd
  indy-cli $1/indy-cli.cmd &>/dev/null
fi

cd $1
# other settings defined in Docker image as env variables
./findy-agent-cli agency start \
  --steward-name $WALLET_NAME \
  --steward-key $WALLET_KEY \
  --steward-did $STEWARD_DID \
  --host-port $HOST_PORT \
  --host-address $HOST_ADDR \
  --register-file $REGISTRY_PATH \
  --psm-database-file $PSMDB_PATH \
  --salt $FINDY_AGENT_SALT
