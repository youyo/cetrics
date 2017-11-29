package cmd

import (
	"github.com/spf13/cobra"
)

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "machine",
	//Long: ``,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.AddCommand(machineCmd)
}
