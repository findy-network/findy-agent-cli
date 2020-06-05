package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/findy-network/findy-agent/agent/ssi"
	"github.com/lainio/err2"
)

const (
	stewardTmpWalletName1 = "test_steward_wallet1"
	stewardTmpWalletKey1  = "6cih1cVgRH8yHD54nEYyPKLmdv67o8QbufxaTHot3Qxp"

	walletName1 = "test_wallet1"
	walletName2 = "test_wallet2"
	walletKey   = "6cih1cVgRH8yHD54nEYyPKLmdv67o8QbufxaTHot3Qxp"
	email1      = "test_email1"
	email2      = "test_email2"
	testGenesis = "../configs/test/genesis_tranactions"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func tearDown() {
	currentUser, err := user.Current()
	err2.Check(err)
	home := currentUser.HomeDir

	removeFiles(home, "/.indy_client/worker/test_wallet*")
	removeFiles(home, "/.indy_client/worker/test_email*")
	removeFiles(home, "/.indy_client/wallet/test_*")
	removeFiles(home, "/.indy_client/wallet/test_email*")
	removeFiles(home, "/test_export_wallets/*")
	removeFile(testGenesis)
	ssi.ClosePool()
}

func removeFiles(home, nameFilter string) {
	filter := filepath.Join(home, nameFilter)
	files, _ := filepath.Glob(filter)
	for _, f := range files {
		if err := os.RemoveAll(f); err != nil {
			panic(err)
		}
	}
}
func removeFile(filename string) {
	if err := os.Remove(filename); err != nil {
		panic(err)
	}

}

func setUp() {
	defer err2.CatchTrace(func(err error) {
		fmt.Println("error on setup", err)
	})
	err2.Try(createTestWallets())
	f, e := os.Create(testGenesis)
	err2.Check(e)
	defer f.Close()
}

func createTestWallets() (err error) {
	wallet1 := ssi.NewRawWalletCfg(walletName1, walletKey)
	exist := wallet1.Create()
	if exist {
		return errors.New("test wallet exist already")
	}
	return nil
}

func TestExecute(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Define tests
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "user create key",
			args: []string{"cmd",
				"user", "createkey", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--seed", "00000000000000000000thisisa_test",
			},
		},

		{
			name: "service create key",
			args: []string{"cmd",
				"service", "createkey", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--seed", "00000000000000000000thisisa_test",
			},
		},
		{
			name: "user ping",
			args: []string{"cmd",
				"user", "ping", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
			},
		},
		{
			name: "service ping",
			args: []string{"cmd",
				"service", "ping", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
			},
		},
		{
			name: "user send basic msg",
			args: []string{"cmd",
				"user", "send", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--msg", "test message",
				"--from", "senderName",
				"--con-id", "connectionName",
			},
		},
		{
			name: "service send basic msg",
			args: []string{"cmd",
				"service", "send", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--msg", "test message",
				"--from", "senderName",
				"--con-id", "connectionName",
			},
		},
		{
			name: "user trustping",
			args: []string{"cmd",
				"user", "trustping", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--con-id", "my_connection",
			},
		},
		{
			name: "service trustping",
			args: []string{"cmd",
				"service", "trustping", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--con-id", "my_connection",
			},
		},
		{
			name: "user onboard",
			args: []string{"cmd",
				"user", "onboard", "--dry-run",
				"--email", email2,
				"--walletname", walletName2,
				"--walletkey", walletKey,
			},
		},
		{
			name: "service onboard",
			args: []string{"cmd",
				"service", "onboard", "--dry-run",
				"--email", email2,
				"--walletname", walletName2,
				"--walletkey", walletKey,
			},
		},
		{
			name: "create schema (config file)",
			args: []string{"cmd",
				"service", "schema", "create", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--config", "../configs/test/createSchema.yaml",
			},
		},
		{
			name: "service read schema (config file)",
			args: []string{"cmd",
				"service", "schema", "read", "--dry-run",
				"--config", "../configs/test/readSchema.yaml",
			},
		},
		{
			name: "user read schema (config file)",
			args: []string{"cmd",
				"user", "schema", "read", "--dry-run",
				"--config", "../configs/test/readSchema.yaml",
			},
		},
		{
			name: "create creddef (config file)",
			args: []string{"cmd",
				"service", "creddef", "create", "--dry-run",
				"--config", "../configs/test/createCreddef.yaml",
			},
		},
		{
			name: "service read creddef (config file)",
			args: []string{"cmd",
				"service", "creddef", "read", "--dry-run",
				"--config", "../configs/test/readCreddef.yaml",
			},
		},
		{
			name: "user read creddef (config file)",
			args: []string{"cmd",
				"user", "creddef", "read", "--dry-run",
				"--config", "../configs/test/readCreddef.yaml",
			},
		},
		{
			name: "create steward (config file)",
			args: []string{"cmd",
				"service", "steward", "--dry-run",
				"--walletname", stewardTmpWalletName1,
				"--walletkey", stewardTmpWalletKey1,
				"--config", "../configs/test/createSteward.yaml",
			},
		},
		{
			name: "create pool",
			args: []string{"cmd",
				"pool", "create", "--dry-run",
				"--poolname", "findy-pool",
				"--pool-genesis", testGenesis,
			},
		},
		{
			name: "ping pool",
			args: []string{"cmd",
				"pool", "ping", "--dry-run",
				"--poolname", "findy-pool",
			},
		},
		{
			name: "start agency (config file)",
			args: []string{"cmd",
				"agency", "start", "--dry-run",
				"--config", "../configs/test/startAgency.yaml",
			},
		},
		{
			name: "ping agency",
			args: []string{"cmd",
				"agency", "ping", "--dry-run",
				"--base-addr", "my_agency_base_address.com",
			},
		},
		{
			name: "user create invitation",
			args: []string{"cmd",
				"user", "invitation", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--label", "connection_name",
			},
		},
		{
			name: "service create invitation",
			args: []string{"cmd",
				"service", "invitation", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--label", "connection_name",
			},
		},
		{
			name: "service connect (config file & no invitation)",
			args: []string{"cmd",
				"service", "connect", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--config", "../configs/test/connect.yaml",
			},
		},
		{
			name: "user connect (config file & no invitation)",
			args: []string{"cmd",
				"user", "connect", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--config", "../configs/test/connect.yaml",
			},
		},
		{
			name: "service export",
			args: []string{"cmd",
				"service", "export", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--export-file", "../configs/test/my-export-wallet",
			},
		},
		{
			name: "user export",
			args: []string{"cmd",
				"user", "export", "--dry-run",
				"--walletname", walletName1,
				"--walletkey", walletKey,
				"--export-file", "../configs/test/my-export-wallet",
			},
		},
		{
			name: "service connect invitation (config file)",
			args: []string{"cmd",
				"service", "connect", "--dry-run",
				"--config", "../configs/test/connectInvitation.yaml",
				"../configs/test/test_invitation",
			},
		},
		{
			name: "user connect invitation(config file)",
			args: []string{"cmd",
				"user", "connect", "--dry-run",
				"--config", "../configs/test/connectInvitation.yaml",
				"../configs/test/test_invitation",
			},
		},
	}

	// Iterate tests
	for _, test := range tests {
		os.Args = test.args
		rootCmd.SilenceUsage = true
		rootCmd.SilenceErrors = true

		t.Run(test.name, func(t *testing.T) {
			if err := rootCmd.Execute(); err != nil {
				t.Errorf("Test error = %v", err)
			}
		})
	}
}
