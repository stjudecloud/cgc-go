package cmd

import (
	"github.com/spf13/cobra"
)

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "CGC actions endpoints",
	Long: `NOT IMPLEMENTED: Support for actions API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/actions`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
