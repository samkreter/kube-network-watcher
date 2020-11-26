package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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

	controllerCmd.Flags().StringVar(&tenantID, "tenant-id", "", "Tenant ID")
	controllerCmd.Flags().StringVar(&clientID, "aad-client-id", "", "Service Principal ID")
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

