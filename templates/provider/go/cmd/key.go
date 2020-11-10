package main

import (
	"fmt"
	"io/ioutil"

	gp "github.com/howeyc/gopass"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/spf13/cobra"
)

var (
	KeysCmd = &cobra.Command{
		Use:   "keys",
		Short: " Key management commands",
	}
)

// KeysAddCmd implements the keys add command
func KeysAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name] [config-file]",
		Short: "Generate a new key",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase := readPassphrase(2)

			configPath := ""

			if len(args) == 1 {
				configPath = common.ConfigPath
			} else {
				configPath = args[1]
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), passphrase)

			addr, mnemonic, err := serviceClient.AddKey(args[0], serviceClient.Password)
			if err != nil {
				return err
			}

			fmt.Printf("key generated successfully: \n\nname: %s\naddress: %s\nmnemonic: %s\n\n", args[0], addr, mnemonic)

			return nil
		},
	}

	return cmd
}

// KeysShowCmd implements the keys show command
func KeysShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [name] [config-file]",
		Short: "Show the key information by name",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase := readPassphrase(1)

			configPath := ""

			if len(args) == 1 {
				configPath = common.ConfigPath
			} else {
				configPath = args[1]
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), passphrase)

			addr, err := serviceClient.ShowKey(args[0], serviceClient.Password)
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", addr)

			return nil
		},
	}

	return cmd
}

// KeysImportCmd implements the keys import command
func KeysImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import [name] [key-file] [config-file]",
		Short: "Import a key from the private key armor file",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase := readPassphrase(1)

			configPath := ""

			if len(args) == 2 {
				configPath = common.ConfigPath
			} else {
				configPath = args[2]
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			keyArmor, err := ioutil.ReadFile(args[1])
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), passphrase)

			addr, err := serviceClient.ImportKey(args[0], serviceClient.Password, string(keyArmor))
			if err != nil {
				return err
			}

			fmt.Printf("key imported successfully: %s\n", addr)

			return nil
		},
	}

	return cmd
}

func init() {
	KeysCmd.AddCommand(
		KeysAddCmd(),
		KeysShowCmd(),
		KeysImportCmd(),
	)
}

func readPassphrase(times int) string {
	// Get user's password
	fmt.Print("Please enter your password: ")
	pwd0, _ := gp.GetPasswd()
	if times == 2 {
		fmt.Print("Please confirm your password: ")
		pwd1, _ := gp.GetPasswd()
		if string(pwd0) == string(pwd1) {
			return string(pwd0)
		}
		panic("The two passwords do not match.")
	}
	return string(pwd0)
}
