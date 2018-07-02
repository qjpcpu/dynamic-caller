/*
the codes runs on ropsten network
*/
package main

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/ethereum/contracts"
	"github.com/qjpcpu/ethereum/key"
	"github.com/urfave/cli"
	"math/big"
	"os"
	"strings"
)

const (
	AnyContractAddress  = `0x2f44fc640F9708FD969620466F9eddD21859e8E9`
	DynamicCallerAddres = `0x5ec567cf2137da526945f4820d0c0621ddcd02ce`
)

var (
	conn   *ethclient.Client
	dc     *DynamicCaller
	anyABI abi.ABI
	pk     *ecdsa.PrivateKey
)

func init() {
	url := os.Getenv("ETH_ROPSTEN")
	var err error
	conn, err = ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	anyABI, _ = contracts.ParseABI(AnyContractABI)
	dc, err = NewDynamicCaller(common.HexToAddress(DynamicCallerAddres), conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter private key: ")
	text, _ := reader.ReadString('\n')
	pkstr := strings.TrimSuffix(text, "\n")
	pk, err = key.PrivateKeyFromHex(pkstr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name: "add",
			Action: func(c *cli.Context) error {
				data, err := contracts.PackArguments(anyABI, "add", big.NewInt(1), big.NewInt(2))
				if err != nil {
					return err
				}
				builder := contracts.NewTxOptsBuilderFromPK(pk)
				tx, err := dc.DynCall(builder.Get(), common.HexToAddress(AnyContractAddress), data)
				if err != nil {
					return err
				}
				fmt.Println("add tx:", tx.Hash().Hex())
				return nil
			},
		},
		{
			Name: "write",
			Action: func(c *cli.Context) error {
				data, err := contracts.PackArguments(anyABI, "write", "hello world")
				if err != nil {
					return err
				}
				builder := contracts.NewTxOptsBuilderFromPK(pk)
				tx, err := dc.DynCall(builder.Get(), common.HexToAddress(AnyContractAddress), data)
				if err != nil {
					return err
				}
				fmt.Println("write tx:", tx.Hash().Hex())
				return nil
			},
		},
		{
			Name: "batchWrite",
			Action: func(c *cli.Context) error {
				data, err := contracts.PackArguments(anyABI, "batchWrite", big.NewInt(10), big.NewInt(20), "Hello World")
				if err != nil {
					return err
				}
				builder := contracts.NewTxOptsBuilderFromPK(pk)
				builder.BuildValue(big.NewInt(3))
				tx, err := dc.DynCall(builder.Get(), common.HexToAddress(AnyContractAddress), data)
				if err != nil {
					return err
				}
				fmt.Println("batchWrite tx:", tx.Hash().Hex())
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
