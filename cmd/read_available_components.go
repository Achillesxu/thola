package cmd

import (
	"github.com/inexio/thola/core/request"
	"github.com/spf13/cobra"
)

func init() {
	readCMD.AddCommand(readAvailableComponentsCMD)
}

var readAvailableComponentsCMD = &cobra.Command{
	Use:   "available-components",
	Short: "Returns the available components for the device",
	Long:  "Returns the available components for the device.",
	Run: func(cmd *cobra.Command, args []string) {
		request := request.ReadAvailableComponentsRequest{
			ReadRequest: getReadRequest(),
		}
		handleRequest(&request)
	},
}
