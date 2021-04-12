package cmd

import (
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "CGC tasks endpoint",
	Long: `NOT IMPLEMENTED: Support for tasks API endpoint
	Ref: https://docs.cancergenomicscloud.org/docs/tasks`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
