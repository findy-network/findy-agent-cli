module github.com/findy-network/findy-agent-cli

go 1.16

require (
	github.com/findy-network/findy-agent v0.0.0-20210302063538-627f8f3c8758
	github.com/findy-network/findy-agent-api v0.0.0-20210203142917-ee7d471ffd4b
	github.com/findy-network/findy-agent-auth v0.1.1-0.20210318145233-c65e941d9b2b
	github.com/findy-network/findy-common-go v0.1.2-0.20210314110550-1711fcae935e
	github.com/findy-network/findy-wrapper-go v0.0.0-20210302063517-bb98c7f07ea4
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/uuid v1.2.0 // indirect
	github.com/lainio/err2 v0.6.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.36.0
)

replace github.com/findy-network/findy-agent-api => ../findy-agent-api
