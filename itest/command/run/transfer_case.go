package run

import (
	"github.com/iost-official/go-iost/itest"
	"github.com/urfave/cli"
)

// TransferCaseCommand is the command of transfer test case
var TransferCaseCommand = cli.Command{
	Name:      "transfer_case",
	ShortName: "t_case",
	Usage:     "run transfer test case",
	Flags:     TransferCaseFlags,
	Action:    TransferCaseAction,
}

// TransferCaseFlags is the flags of transfer test case
var TransferCaseFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "number, n",
		Value: 1000,
		Usage: "number of transaction",
	},
	cli.StringFlag{
		Name:  "output, o",
		Value: "accounts.json",
		Usage: "output of account information",
	},
}

// TransferCaseAction is the action of transfer test case
var TransferCaseAction = func(c *cli.Context) error {
	afile := c.GlobalString("account")
	output := c.String("output")
	tnum := c.Int("number")
	keysfile := c.GlobalString("keys")
	configfile := c.GlobalString("config")
	memoSize := c.Int("memo")

	it, err := itest.Load(keysfile, configfile)
	if err != nil {
		return err
	}

	accounts, err := itest.LoadAccounts(afile)
	if err != nil {
		return err
	}

	if _, err := it.TransferN(tnum, accounts, memoSize, true); err != nil {
		return err
	}

	if err := it.CheckAccounts(accounts); err != nil {
		return err
	}

	if err := itest.DumpAccounts(accounts, output); err != nil {
		return err
	}

	return nil
}
