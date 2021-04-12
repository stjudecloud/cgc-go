package cmd

import (
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "CGC projects endpoints",
	Long: `NOT IMPLEMENTED: Support for projects API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/projects`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
