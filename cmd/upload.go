package cmd

import (
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "uploads",
	Short: "CGC uploads endpoints",
	Long: `NOT IMPLEMENTED: Support for file uploads API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/upload-files`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
