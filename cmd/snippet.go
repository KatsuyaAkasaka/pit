/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/KatsuyaAkasaka/pit/pkg/components/snippet"
	"github.com/spf13/cobra"
)

func (o *options) toSnippetPkgVars(subCommand string) *snippet.Variable {
	return &snippet.Variable{
		Subcommand: subCommand,
		Option: snippet.Options{
			Language: o.language,
		},
	}
}

// sniCmd represents the sni command
func getSnippetCmd(cs *cmdSetting) *cobra.Command {
	opt := &options{}

	var snippetCmd = &cobra.Command{
		Use:   "sni",
		Short: "my snippets",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("please specify subcommand")
				os.Exit(1)
			}

			outputSetting, err := cs.buildOutputSetting(args, opt)
			if err != nil {
				fmt.Println("error occured", err)
				os.Exit(1)
			}

			subCommand := args[0]
			err, res := snippet.Exec(opt.toSnippetPkgVars(subCommand))
			if err != nil {
				fmt.Println("error occured", err)
				os.Exit(1)
			}

			outputSetting.exec(toResult(res))
		},
	}
	snippetCmd.Flags().StringVarP(&opt.language, "language", "l", "", "snippet language")
	return snippetCmd
}
