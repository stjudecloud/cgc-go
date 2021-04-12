package cmd

import (
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "CGC files endpoints",
	Long: `NOT IMPLEMENTED: Support for files API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/files`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
