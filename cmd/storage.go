package cmd

import (
	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "CGC storage endpoint",
	Long: `NOT IMPLEMENTED: Support for storage API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/storage`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
