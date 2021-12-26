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

type options struct {
	language  string
	inputText string
}

type result struct {
	content string
}

func toResult(content string) *result {
	return &result{
		content: content,
	}
}

type cmdSetting struct {
	isCopyToClipboard   bool
	isOutputStdout      bool
	useClipboardToInput bool
}
type outputSetting struct {
	isCopyToClipboard bool
	isOutputStdout    bool
	inputText         string
}

func (c *cmdSetting) buildOutputSetting(args []string, opt *options) (*outputSetting, error) {
	res := &outputSetting{
		isCopyToClipboard: true,
		inputText:         "",
		isOutputStdout:    c.isOutputStdout,
	}
	scanner := bufio.NewScanner(os.Stdin)

	if c.isCopyToClipboard {
		fmt.Printf("Copy snippet to clipboard? (Y/n) (default: Y) > ")
		scanner.Scan()
		in := scanner.Text()
		if in == "n" {
			res.isCopyToClipboard = false
		}
	}
	if opt.inputText != "" {
		res.inputText = opt.inputText
	} else if c.useClipboardToInput && opt.inputText == "" {
		input, err := clipboard.ReadAll()
		if err != nil {
			os.Exit(1)
		}
		res.inputText = input
	}
	return res, nil
}

func (o *outputSetting) exec(res *result) {
	if o.isCopyToClipboard {
		clipboard.WriteAll(res.content)
	}
	if o.isOutputStdout {
		fmt.Println("```")
		fmt.Printf(res.content)
		fmt.Println("```")
	}
}

func Cmd() *cobra.Command {
	rootCmd.AddCommand(
		getSnippetCmd(&cmdSetting{
			isCopyToClipboard:   true,
			isOutputStdout:      true,
			useClipboardToInput: false,
		}),
		getReplaceCmd(&cmdSetting{
			isCopyToClipboard:   true,
			isOutputStdout:      true,
			useClipboardToInput: true,
		}),
	)
	return rootCmd
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
