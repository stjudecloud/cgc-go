package cmd

import (
	"github.com/spf13/cobra"
)

var markersCmd = &cobra.Command{
	Use:   "markers",
	Short: "CGC markerss endpoints",
	Long: `NOT IMPLEMENTED: Support for markers API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/markers`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
