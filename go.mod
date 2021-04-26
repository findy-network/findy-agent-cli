module github.com/findy-network/findy-agent-cli

go 1.16

require (
	github.com/findy-network/findy-agent-auth v0.1.4-0.20210421160857-fa97baa3b52a
	github.com/findy-network/findy-common-go v0.1.2-0.20210421160228-49a5213c3ab5
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/uuid v1.2.0
	github.com/lainio/err2 v0.6.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.36.0
)

replace github.com/findy-network/findy-common-go => ../findy-common-go
