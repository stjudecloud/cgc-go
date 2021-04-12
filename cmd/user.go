package cmd

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "CGC user endpoint",
	Long: `NOT IMPLEMENTED: Support for user API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/user`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
