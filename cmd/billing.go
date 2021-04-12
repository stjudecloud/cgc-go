package cmd

import (
	"github.com/spf13/cobra"
)

var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "CGC billing endpoints",
	Long: `NOT IMPLEMENTED: Support for billing API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/billing`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
