package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func configureShell() {
	fmt.Println(rice.MustFindBox("fish").MustString("yam.fish"))
}

func listKeys() {
	cmd := exec.Command("pass", "list", "aws-keys")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", stdoutStderr)
}

func getKey(profile string) {
	passPath := fmt.Sprintf("aws-keys/%s", profile)
	cmd := exec.Command("pass", passPath)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", stdoutStderr)
}

func main() {
	var rootCmd = &cobra.Command{
		Use: "yam",
	}

	var listCmd = &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			listKeys()
		},
	}

	var getCmd = &cobra.Command{
		Use: "__get PROFILE",
		Run: func(cmd *cobra.Command, args []string) {
			profile := args[0]
			getKey(profile)
		},
		Args:   cobra.ExactArgs(1),
		Hidden: true,
	}

	var configureCmd = &cobra.Command{
		Use: "__configure",
		Run: func(cmd *cobra.Command, args []string) {
			configureShell()
		},
		Hidden: true,
	}

	var activateCmd = &cobra.Command{
		Use: "activate",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Do not call this on the helper. Use yam activate instead.")
		},
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(activateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
