package main

import (
	"github.com/spf13/cobra"
)

var (
	tenantID     string
	clientID     string
)

// startCmd represents the start command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "A brief description of your command",
	Long:  `A longer description of your command`,
	Run:   startHelper,
}

func init() {
	rootCmd.AddCommand(controllerCmd)


}

func startHelper(cmd *cobra.Command, args []string) {
	/*


	* Register
		* Agent -> PodName/deployment/NS: (Pod Name is ID)
	* Deregister

	Server

	*/

	// Register

}

