/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/KatsuyaAkasaka/pit/pkg/components/replace"
	"github.com/spf13/cobra"
)

func (o *options) toReplacePkgVars(subCommand string, input string) *replace.Variable {
	return &replace.Variable{
		Subcommand: subCommand,
		Option: replace.Options{
			InputText: input,
		},
	}
}

// sniCmd represents the sni command
func getReplaceCmd(cs *cmdSetting) *cobra.Command {
	opt := &options{}

	var replaceCmd = &cobra.Command{
		Use:   "rep",
		Short: "my replacement",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("please specify subcommand -h")
				os.Exit(1)
			}

			outputSetting, err := cs.buildOutputSetting(args, opt)
			if err != nil {
				fmt.Println("error occured", err)
				os.Exit(1)
			}

			subCommand := args[0]
			res, err := replace.Exec(opt.toReplacePkgVars(subCommand, outputSetting.inputText))
			if err != nil {
				fmt.Println("error occured", err)
				os.Exit(1)
			}

			outputSetting.exec(toResult(res))
		},
	}
	replaceCmd.Flags().StringVarP(&opt.inputText, "input", "i", "", "input text")
	return replaceCmd
}
