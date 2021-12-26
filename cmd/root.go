/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pit",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

type result struct {
	content string
}

func toResult(content string) *result {
	return &result{
		content: content,
	}
}

type inputSetting struct {
	isCopyToClipboard bool
}
type inputSettingResult struct {
	isCopyToClipboard bool
}

func (i *inputSetting) Exec() *inputSettingResult {
	res := &inputSettingResult{
		isCopyToClipboard: true,
	}
	scanner := bufio.NewScanner(os.Stdin)

	if i.isCopyToClipboard {
		fmt.Printf("Copy snippet to clipboard? (Y/n) (default: Y) > ")
		scanner.Scan()
		in := scanner.Text()
		if in == "n" {
			res.isCopyToClipboard = false
		}
	}
	return res
}

func (r *inputSettingResult) Exec(res *result) {
	if r.isCopyToClipboard {
		clipboard.WriteAll(res.content)
	}
}

func Cmd() *cobra.Command {
	rootCmd.AddCommand(
		getSnippetCmd(&inputSetting{
			isCopyToClipboard: true,
		}),
	)
	return rootCmd
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
