package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var enableVerboseLogging bool
var Logger *logrus.Logger

var rootCmd = &cobra.Command{
	Use:   "kubeception",
	Short: "kubeception",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		Logger = logrus.New()

		Logger.SetOutput(os.Stdout)
		Logger.SetLevel(logrus.InfoLevel)

		if enableVerboseLogging {
			Logger.SetLevel(logrus.DebugLevel)
		}
	},
	Run: nil,
}

func main() {
	rootCmd.PersistentFlags().BoolVarP(&enableVerboseLogging,
		"verbose", "v", false, "log more information (debug level)")

	commands := []*cobra.Command{
		createClusterCommands(rootCmd),
	}

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
