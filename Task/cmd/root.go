package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Task",
	Short: "Task is CLI task manager",
	// Do Stuff Here
}

// func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }
