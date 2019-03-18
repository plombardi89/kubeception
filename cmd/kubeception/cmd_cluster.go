package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func createClusterCommands(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Manage clusters",
		Run:   nil,
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a cluster",
		Run:   createCluster,
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a cluster",
		Run:   deleteCluster,
	}

	cmd.AddCommand(createCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

func createCluster(cmd *cobra.Command, args []string) {
	fmt.Println("create cluster")
}

func deleteCluster(cmd *cobra.Command, args []string) {
	fmt.Println("delete cluster")
}