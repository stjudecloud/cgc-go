package cmd

import (
	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "CGC users endpoints",
	Long: `NOT IMPLEMENTED: Support for users API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/users`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
