using namespace System.Management.Automation
using namespace System.Management.Automation.Language
Register-ArgumentCompleter -Native -CommandName 'findy-cli' -ScriptBlock {
    param($wordToComplete, $commandAst, $cursorPosition)
    $commandElements = $commandAst.CommandElements
    $command = @(
        'findy-cli'
        for ($i = 1; $i -lt $commandElements.Count; $i++) {
            $element = $commandElements[$i]
            if ($element -isnot [StringConstantExpressionAst] -or
                $element.StringConstantType -ne [StringConstantType]::BareWord -or
                $element.Value.StartsWith('-')) {
                break
            }
            $element.Value
        }
    ) -join ';'
    $completions = @(switch ($command) {
        'findy-cli' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('agency', 'agency', [CompletionResultType]::ParameterValue, 'Parent command for starting and pinging agency')
            [CompletionResult]::new('completion', 'completion', [CompletionResultType]::ParameterValue, 'Generates bash completion scripts')
            [CompletionResult]::new('help', 'help', [CompletionResultType]::ParameterValue, 'Help about any command')
            [CompletionResult]::new('pool', 'pool', [CompletionResultType]::ParameterValue, 'Parent command for pool commands')
            [CompletionResult]::new('service', 'service', [CompletionResultType]::ParameterValue, 'Parent command for service client')
            [CompletionResult]::new('user', 'user', [CompletionResultType]::ParameterValue, 'Parent command for user client')
            break
        }
        'findy-cli;agency' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('ping', 'ping', [CompletionResultType]::ParameterValue, 'Command for pinging agency')
            [CompletionResult]::new('start', 'start', [CompletionResultType]::ParameterValue, 'Command for starting agency')
            break
        }
        'findy-cli;agency;ping' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--base-addr', 'base-addr', [CompletionResultType]::ParameterName, 'base address of agency')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            break
        }
        'findy-cli;agency;start' {
            [CompletionResult]::new('--a2a', 'a2a', [CompletionResultType]::ParameterName, 'URL path for A2A protocols')
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('--did', 'did', [CompletionResultType]::ParameterName, 'steward DID')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--genesis', 'genesis', [CompletionResultType]::ParameterName, 'pool genesis file')
            [CompletionResult]::new('--hostaddr', 'hostaddr', [CompletionResultType]::ParameterName, 'host address')
            [CompletionResult]::new('--hostport', 'hostport', [CompletionResultType]::ParameterName, 'host port')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--pool', 'pool', [CompletionResultType]::ParameterName, 'pool name')
            [CompletionResult]::new('--protocol', 'protocol', [CompletionResultType]::ParameterName, 'pool protocol')
            [CompletionResult]::new('--psmdb', 'psmdb', [CompletionResultType]::ParameterName, 'state machine db''s filename')
            [CompletionResult]::new('--register', 'register', [CompletionResultType]::ParameterName, 'handshake registry''s filename')
            [CompletionResult]::new('--reset', 'reset', [CompletionResultType]::ParameterName, 'reset register')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--seed', 'seed', [CompletionResultType]::ParameterName, 'steward seed')
            [CompletionResult]::new('--serverport', 'serverport', [CompletionResultType]::ParameterName, 'server port')
            [CompletionResult]::new('--service-name', 'service-name', [CompletionResultType]::ParameterName, 'service name')
            [CompletionResult]::new('--steward-key', 'steward-key', [CompletionResultType]::ParameterName, 'steward key')
            [CompletionResult]::new('--steward-name', 'steward-name', [CompletionResultType]::ParameterName, 'steward name')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            break
        }
        'findy-cli;completion' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('-h', 'h', [CompletionResultType]::ParameterName, 'help for completion')
            [CompletionResult]::new('--help', 'help', [CompletionResultType]::ParameterName, 'help for completion')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            break
        }
        'findy-cli;help' {
            break
        }
        'findy-cli;pool' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--poolname', 'poolname', [CompletionResultType]::ParameterName, 'name of the pool')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('create', 'create', [CompletionResultType]::ParameterValue, 'Command for creating creating pool')
            [CompletionResult]::new('ping', 'ping', [CompletionResultType]::ParameterValue, 'Command for pinging pool')
            break
        }
        'findy-cli;pool;create' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--pool-genesis', 'pool-genesis', [CompletionResultType]::ParameterName, 'pool genesis file')
            [CompletionResult]::new('--poolname', 'poolname', [CompletionResultType]::ParameterName, 'name of the pool')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            break
        }
        'findy-cli;pool;ping' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--poolname', 'poolname', [CompletionResultType]::ParameterName, 'name of the pool')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            break
        }
        'findy-cli;service' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            [CompletionResult]::new('connect', 'connect', [CompletionResultType]::ParameterValue, 'Command for creating a2a connection between 2 agents')
            [CompletionResult]::new('createkey', 'createkey', [CompletionResultType]::ParameterValue, 'Command for creating valid wallet keys')
            [CompletionResult]::new('creddef', 'creddef', [CompletionResultType]::ParameterValue, 'Parent command for operating with Credential definations')
            [CompletionResult]::new('export', 'export', [CompletionResultType]::ParameterValue, 'Command for exporting wallet')
            [CompletionResult]::new('invitation', 'invitation', [CompletionResultType]::ParameterValue, 'Command for creating invitation message for agent')
            [CompletionResult]::new('onboard', 'onboard', [CompletionResultType]::ParameterValue, 'Command for onboarding agent')
            [CompletionResult]::new('ping', 'ping', [CompletionResultType]::ParameterValue, 'Command for pinging services and agents')
            [CompletionResult]::new('schema', 'schema', [CompletionResultType]::ParameterValue, 'Parent command for operating with schemas')
            [CompletionResult]::new('send', 'send', [CompletionResultType]::ParameterValue, 'Command for sending basic message to another agent')
            [CompletionResult]::new('steward', 'steward', [CompletionResultType]::ParameterValue, 'Command for creating steward wallet')
            [CompletionResult]::new('trustping', 'trustping', [CompletionResultType]::ParameterValue, 'Command for making trustping to another agent')
            break
        }
        'findy-cli;service;connect' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--pwendp', 'pwendp', [CompletionResultType]::ParameterName, 'pairwise endpoint')
            [CompletionResult]::new('--pwkey', 'pwkey', [CompletionResultType]::ParameterName, 'pairwise endpoint key')
            [CompletionResult]::new('--pwname', 'pwname', [CompletionResultType]::ParameterName, 'name of the pairwise connection')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;createkey' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--seed', 'seed', [CompletionResultType]::ParameterName, 'Seed for wallet key creation')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;creddef' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--schema-id', 'schema-id', [CompletionResultType]::ParameterName, 'schema ID')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            [CompletionResult]::new('create', 'create', [CompletionResultType]::ParameterValue, 'Command for creating new credential definition')
            [CompletionResult]::new('read', 'read', [CompletionResultType]::ParameterValue, 'Command for getting credential definition by id')
            break
        }
        'findy-cli;service;creddef;create' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--schema-id', 'schema-id', [CompletionResultType]::ParameterName, 'schema ID')
            [CompletionResult]::new('--tag', 'tag', [CompletionResultType]::ParameterName, 'cred def tag')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;creddef;read' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--creddef-id', 'creddef-id', [CompletionResultType]::ParameterName, 'cred def id')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--schema-id', 'schema-id', [CompletionResultType]::ParameterName, 'schema ID')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;export' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--export-file', 'export-file', [CompletionResultType]::ParameterName, 'filename for wallet export with whole path')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;invitation' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--label', 'label', [CompletionResultType]::ParameterName, 'invitation label')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;onboard' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--email', 'email', [CompletionResultType]::ParameterName, 'onboarding email')
            [CompletionResult]::new('--export-file', 'export-file', [CompletionResultType]::ParameterName, 'filename for wallet export with path')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;ping' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('-s', 's', [CompletionResultType]::ParameterName, 'ping CA and connected SA (me) as well')
            [CompletionResult]::new('--sa', 'sa', [CompletionResultType]::ParameterName, 'ping CA and connected SA (me) as well')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;schema' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            [CompletionResult]::new('create', 'create', [CompletionResultType]::ParameterValue, 'Command for creating new schema')
            [CompletionResult]::new('read', 'read', [CompletionResultType]::ParameterValue, 'Command for getting schema by id')
            break
        }
        'findy-cli;service;schema;create' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--schema-attrs', 'schema-attrs', [CompletionResultType]::ParameterName, 'schema attributes')
            [CompletionResult]::new('--schema-name', 'schema-name', [CompletionResultType]::ParameterName, 'schema name')
            [CompletionResult]::new('--schema-v', 'schema-v', [CompletionResultType]::ParameterName, 'schema version')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;schema;read' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--schema-id', 'schema-id', [CompletionResultType]::ParameterName, 'schema id')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;send' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--con-id', 'con-id', [CompletionResultType]::ParameterName, 'connection id')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--from', 'from', [CompletionResultType]::ParameterName, 'name of the msg sender')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--msg', 'msg', [CompletionResultType]::ParameterName, 'message to be send')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;steward' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--poolname', 'poolname', [CompletionResultType]::ParameterName, 'Pool name')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--steward-seed', 'steward-seed', [CompletionResultType]::ParameterName, 'Steward seed')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;service;trustping' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--con-id', 'con-id', [CompletionResultType]::ParameterName, 'connection id')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            [CompletionResult]::new('connect', 'connect', [CompletionResultType]::ParameterValue, 'Command for creating a2a connection between 2 agents')
            [CompletionResult]::new('createkey', 'createkey', [CompletionResultType]::ParameterValue, 'Command for creating valid wallet keys')
            [CompletionResult]::new('creddef', 'creddef', [CompletionResultType]::ParameterValue, 'Parent command for operating with Credential definations')
            [CompletionResult]::new('export', 'export', [CompletionResultType]::ParameterValue, 'Command for exporting wallet')
            [CompletionResult]::new('invitation', 'invitation', [CompletionResultType]::ParameterValue, 'Command for creating invitation message for agent')
            [CompletionResult]::new('onboard', 'onboard', [CompletionResultType]::ParameterValue, 'Command for onboarding agent')
            [CompletionResult]::new('ping', 'ping', [CompletionResultType]::ParameterValue, 'Command for pinging services and agents')
            [CompletionResult]::new('schema', 'schema', [CompletionResultType]::ParameterValue, 'Parent command for operating with schemas')
            [CompletionResult]::new('send', 'send', [CompletionResultType]::ParameterValue, 'Command for sending basic message to another agent')
            [CompletionResult]::new('trustping', 'trustping', [CompletionResultType]::ParameterValue, 'Command for making trustping to another agent')
            break
        }
        'findy-cli;user;connect' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--pwendp', 'pwendp', [CompletionResultType]::ParameterName, 'pairwise endpoint')
            [CompletionResult]::new('--pwkey', 'pwkey', [CompletionResultType]::ParameterName, 'pairwise endpoint key')
            [CompletionResult]::new('--pwname', 'pwname', [CompletionResultType]::ParameterName, 'name of the pairwise connection')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;createkey' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--seed', 'seed', [CompletionResultType]::ParameterName, 'Seed for wallet key creation')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;creddef' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            [CompletionResult]::new('read', 'read', [CompletionResultType]::ParameterValue, 'Command for getting credential definition by id')
            break
        }
        'findy-cli;user;creddef;read' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--creddef-id', 'creddef-id', [CompletionResultType]::ParameterName, 'cred def id')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--schema-id', 'schema-id', [CompletionResultType]::ParameterName, 'schema ID')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;export' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--export-file', 'export-file', [CompletionResultType]::ParameterName, 'filename for wallet export with whole path')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;invitation' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--label', 'label', [CompletionResultType]::ParameterName, 'invitation label')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;onboard' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--email', 'email', [CompletionResultType]::ParameterName, 'onboarding email')
            [CompletionResult]::new('--export-file', 'export-file', [CompletionResultType]::ParameterName, 'filename for wallet export with path')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;ping' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('-s', 's', [CompletionResultType]::ParameterName, 'ping CA and connected SA (me) as well')
            [CompletionResult]::new('--sa', 'sa', [CompletionResultType]::ParameterName, 'ping CA and connected SA (me) as well')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;schema' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            [CompletionResult]::new('read', 'read', [CompletionResultType]::ParameterValue, 'Command for getting schema by id')
            break
        }
        'findy-cli;user;schema;read' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--schema-id', 'schema-id', [CompletionResultType]::ParameterName, 'schema id')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;send' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--con-id', 'con-id', [CompletionResultType]::ParameterName, 'connection id')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--from', 'from', [CompletionResultType]::ParameterName, 'name of the msg sender')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--msg', 'msg', [CompletionResultType]::ParameterName, 'message to be send')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
        'findy-cli;user;trustping' {
            [CompletionResult]::new('--apiurl', 'apiurl', [CompletionResultType]::ParameterName, 'api base address')
            [CompletionResult]::new('--con-id', 'con-id', [CompletionResultType]::ParameterName, 'connection id')
            [CompletionResult]::new('--config', 'config', [CompletionResultType]::ParameterName, 'config file')
            [CompletionResult]::new('--data', 'data', [CompletionResultType]::ParameterName, 'path for data files')
            [CompletionResult]::new('-n', 'n', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--dry-run', 'dry-run', [CompletionResultType]::ParameterName, 'perform a trial run with no changes made')
            [CompletionResult]::new('--logging', 'logging', [CompletionResultType]::ParameterName, 'logging startup arguments')
            [CompletionResult]::new('--salt', 'salt', [CompletionResultType]::ParameterName, 'salt')
            [CompletionResult]::new('--url', 'url', [CompletionResultType]::ParameterName, 'endpoint base address')
            [CompletionResult]::new('-v', 'v', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--verbose', 'verbose', [CompletionResultType]::ParameterName, 'verbose')
            [CompletionResult]::new('--walletkey', 'walletkey', [CompletionResultType]::ParameterName, 'wallet key')
            [CompletionResult]::new('--walletname', 'walletname', [CompletionResultType]::ParameterName, 'wallet name')
            break
        }
    })
    $completions.Where{ $_.CompletionText -like "$wordToComplete*" } |
        Sort-Object -Property ListItemText
}