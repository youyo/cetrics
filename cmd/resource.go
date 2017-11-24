package cmd

import (
	"github.com/spf13/cobra"
)

var resourceCmd = &cobra.Command{
	Use: "resource",
	//Short: "",
	//Long: ``,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.AddCommand(resourceCmd)
}
