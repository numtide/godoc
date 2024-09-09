package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomarkdoc",
	Short: "Go Doc generation in Markdown format",
	// todo Long description
	// todo add Long description
	RunE: func(cmd *cobra.Command, args []string) error {
		println("hello")
		return nil
	},
}

func init() {

}
