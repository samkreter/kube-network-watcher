package main

import (

	"github.com/spf13/cobra"
)



// startCmd represents the start command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long:  `A longer description of your command`,
	Run:   startHelper,
}

func init() {
	rootCmd.AddCommand(agentCmd)

	agentCmd.Flags().StringVar(&vars, "vars", "", "the vars")
}

func startHelper(cmd *cobra.Command, args []string) {
	/*
		?StartNetScan
		?EndNetscan
		Healthz
		Scrap

	*/
}

