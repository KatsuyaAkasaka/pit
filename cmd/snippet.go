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

type options struct {
	language string
}

func (o *options) toPkgVars(subCommand string) *snippet.Variable {
	return &snippet.Variable{
		Subcommand: subCommand,
		Option: snippet.Options{
			Language: o.language,
		},
	}
}

// sniCmd represents the sni command
func getSnippetCmd(inputSetting *inputSetting) *cobra.Command {
	opt := &options{}

	var snippetCmd = &cobra.Command{
		Use:   "sni",
		Short: "my snippets",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("please specify subcommand -h")
				os.Exit(1)
			}

			inputRes := inputSetting.Exec()

			err, res := snippet.Exec(opt.toPkgVars(args[0]))
			if err != nil {
				fmt.Println("error occured", err)
				os.Exit(1)
			}

			inputRes.Exec(toResult(res))
		},
	}
	snippetCmd.Flags().StringVarP(&opt.language, "language", "l", "", "snippet language")
	return snippetCmd
}
