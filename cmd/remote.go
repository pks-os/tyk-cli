package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tyk-cli/commands/remote"
	"github.com/TykTechnologies/tyk-cli/utils"
)

var verbose bool
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Select a remote",
	Run: func(cmd *cobra.Command, args []string) {
		conf := utils.ParseJSONFile("example.conf.json")["remotes"].([]interface{})
		remote.List(os.Stdout, conf, verbose)
	},
}

func init() {
	RootCmd.AddCommand(remoteCmd)
	remoteCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "List available remotes and URLs")
}
