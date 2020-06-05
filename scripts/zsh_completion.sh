#compdef _findy-cli findy-cli


function _findy-cli {
  local -a commands

  _arguments -C \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "agency:Parent command for starting and pinging agency"
      "completion:Generates bash completion scripts"
      "help:Help about any command"
      "pool:Parent command for pool commands"
      "service:Parent command for service client"
      "user:Parent command for user client"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  agency)
    _findy-cli_agency
    ;;
  completion)
    _findy-cli_completion
    ;;
  help)
    _findy-cli_help
    ;;
  pool)
    _findy-cli_pool
    ;;
  service)
    _findy-cli_service
    ;;
  user)
    _findy-cli_user
    ;;
  esac
}


function _findy-cli_agency {
  local -a commands

  _arguments -C \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "ping:Command for pinging agency"
      "start:Command for starting agency"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  ping)
    _findy-cli_agency_ping
    ;;
  start)
    _findy-cli_agency_start
    ;;
  esac
}

function _findy-cli_agency_ping {
  _arguments \
    '--base-addr[base address of agency]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]'
}

function _findy-cli_agency_start {
  _arguments \
    '--a2a[URL path for A2A protocols]:' \
    '--did[steward DID]:' \
    '--genesis[pool genesis file]:' \
    '--hostaddr[host address]:' \
    '--hostport[host port]:' \
    '--pool[pool name]:' \
    '--protocol[pool protocol]:' \
    '--psmdb[state machine db'\''s filename]:' \
    '--register[handshake registry'\''s filename]:' \
    '--reset[reset register]' \
    '--seed[steward seed]:' \
    '--serverport[server port]:' \
    '--service-name[service name]:' \
    '--steward-key[steward key]:' \
    '--steward-name[steward name]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]'
}

function _findy-cli_completion {
  _arguments \
    '(-h --help)'{-h,--help}'[help for completion]' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]'
}

function _findy-cli_help {
  _arguments \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]'
}


function _findy-cli_pool {
  local -a commands

  _arguments -C \
    '--poolname[name of the pool]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "create:Command for creating creating pool"
      "ping:Command for pinging pool"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  create)
    _findy-cli_pool_create
    ;;
  ping)
    _findy-cli_pool_ping
    ;;
  esac
}

function _findy-cli_pool_create {
  _arguments \
    '--pool-genesis[pool genesis file]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--poolname[name of the pool]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]'
}

function _findy-cli_pool_ping {
  _arguments \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--poolname[name of the pool]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]'
}


function _findy-cli_service {
  local -a commands

  _arguments -C \
    '--url[endpoint base address]:' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "connect:Command for creating a2a connection between 2 agents"
      "createkey:Command for creating valid wallet keys"
      "creddef:Parent command for operating with Credential definations"
      "export:Command for exporting wallet"
      "invitation:Command for creating invitation message for agent"
      "onboard:Command for onboarding agent"
      "ping:Command for pinging services and agents"
      "schema:Parent command for operating with schemas"
      "send:Command for sending basic message to another agent"
      "steward:Command for creating steward wallet"
      "trustping:Command for making trustping to another agent"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  connect)
    _findy-cli_service_connect
    ;;
  createkey)
    _findy-cli_service_createkey
    ;;
  creddef)
    _findy-cli_service_creddef
    ;;
  export)
    _findy-cli_service_export
    ;;
  invitation)
    _findy-cli_service_invitation
    ;;
  onboard)
    _findy-cli_service_onboard
    ;;
  ping)
    _findy-cli_service_ping
    ;;
  schema)
    _findy-cli_service_schema
    ;;
  send)
    _findy-cli_service_send
    ;;
  steward)
    _findy-cli_service_steward
    ;;
  trustping)
    _findy-cli_service_trustping
    ;;
  esac
}

function _findy-cli_service_connect {
  _arguments \
    '--pwendp[pairwise endpoint]:' \
    '--pwkey[pairwise endpoint key]:' \
    '--pwname[name of the pairwise connection]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_createkey {
  _arguments \
    '--seed[Seed for wallet key creation]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}


function _findy-cli_service_creddef {
  local -a commands

  _arguments -C \
    '--schema-id[schema ID]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "create:Command for creating new credential definition"
      "read:Command for getting credential definition by id"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  create)
    _findy-cli_service_creddef_create
    ;;
  read)
    _findy-cli_service_creddef_read
    ;;
  esac
}

function _findy-cli_service_creddef_create {
  _arguments \
    '--tag[cred def tag]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--schema-id[schema ID]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_creddef_read {
  _arguments \
    '--creddef-id[cred def id]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--schema-id[schema ID]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_export {
  _arguments \
    '--export-file[filename for wallet export with whole path]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_invitation {
  _arguments \
    '--label[invitation label]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_onboard {
  _arguments \
    '--email[onboarding email]:' \
    '--export-file[filename for wallet export with path]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_ping {
  _arguments \
    '(-s --sa)'{-s,--sa}'[ping CA and connected SA (me) as well]' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}


function _findy-cli_service_schema {
  local -a commands

  _arguments -C \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "create:Command for creating new schema"
      "read:Command for getting schema by id"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  create)
    _findy-cli_service_schema_create
    ;;
  read)
    _findy-cli_service_schema_read
    ;;
  esac
}

function _findy-cli_service_schema_create {
  _arguments \
    '*--schema-attrs[schema attributes]:' \
    '--schema-name[schema name]:' \
    '--schema-v[schema version]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_schema_read {
  _arguments \
    '--schema-id[schema id]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_send {
  _arguments \
    '--con-id[connection id]:' \
    '--from[name of the msg sender]:' \
    '--msg[message to be send]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_steward {
  _arguments \
    '--poolname[Pool name]:' \
    '--steward-seed[Steward seed]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_service_trustping {
  _arguments \
    '--con-id[connection id]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}


function _findy-cli_user {
  local -a commands

  _arguments -C \
    '--url[endpoint base address]:' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "connect:Command for creating a2a connection between 2 agents"
      "createkey:Command for creating valid wallet keys"
      "creddef:Parent command for operating with Credential definations"
      "export:Command for exporting wallet"
      "invitation:Command for creating invitation message for agent"
      "onboard:Command for onboarding agent"
      "ping:Command for pinging services and agents"
      "schema:Parent command for operating with schemas"
      "send:Command for sending basic message to another agent"
      "trustping:Command for making trustping to another agent"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  connect)
    _findy-cli_user_connect
    ;;
  createkey)
    _findy-cli_user_createkey
    ;;
  creddef)
    _findy-cli_user_creddef
    ;;
  export)
    _findy-cli_user_export
    ;;
  invitation)
    _findy-cli_user_invitation
    ;;
  onboard)
    _findy-cli_user_onboard
    ;;
  ping)
    _findy-cli_user_ping
    ;;
  schema)
    _findy-cli_user_schema
    ;;
  send)
    _findy-cli_user_send
    ;;
  trustping)
    _findy-cli_user_trustping
    ;;
  esac
}

function _findy-cli_user_connect {
  _arguments \
    '--pwendp[pairwise endpoint]:' \
    '--pwkey[pairwise endpoint key]:' \
    '--pwname[name of the pairwise connection]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_user_createkey {
  _arguments \
    '--seed[Seed for wallet key creation]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}


function _findy-cli_user_creddef {
  local -a commands

  _arguments -C \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "read:Command for getting credential definition by id"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  read)
    _findy-cli_user_creddef_read
    ;;
  esac
}

function _findy-cli_user_creddef_read {
  _arguments \
    '--creddef-id[cred def id]:' \
    '--schema-id[schema ID]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_user_export {
  _arguments \
    '--export-file[filename for wallet export with whole path]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_user_invitation {
  _arguments \
    '--label[invitation label]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_user_onboard {
  _arguments \
    '--email[onboarding email]:' \
    '--export-file[filename for wallet export with path]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_user_ping {
  _arguments \
    '(-s --sa)'{-s,--sa}'[ping CA and connected SA (me) as well]' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}


function _findy-cli_user_schema {
  local -a commands

  _arguments -C \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "read:Command for getting schema by id"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  read)
    _findy-cli_user_schema_read
    ;;
  esac
}

function _findy-cli_user_schema_read {
  _arguments \
    '--schema-id[schema id]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_user_send {
  _arguments \
    '--con-id[connection id]:' \
    '--from[name of the msg sender]:' \
    '--msg[message to be send]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

function _findy-cli_user_trustping {
  _arguments \
    '--con-id[connection id]:' \
    '--apiurl[api base address]:' \
    '--config[config file]:' \
    '--data[path for data files]:' \
    '(-n --dry-run)'{-n,--dry-run}'[perform a trial run with no changes made]' \
    '--logging[logging startup arguments]:' \
    '--salt[salt]:' \
    '--url[endpoint base address]:' \
    '(-v --verbose)'{-v,--verbose}'[verbose]' \
    '--walletkey[wallet key]:' \
    '--walletname[wallet name]:'
}

