module github.com/findy-network/findy-agent-cli

go 1.16

require (
	github.com/findy-network/findy-agent-auth v0.1.22
	github.com/findy-network/findy-common-go v0.1.26-0.20211129161717-174e67c10dcf
	github.com/ghodss/yaml v1.0.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/uuid v1.2.0
	github.com/lainio/err2 v0.7.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	google.golang.org/grpc v1.40.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/findy-network/findy-common-go => ../findy-common-go
