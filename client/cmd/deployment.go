/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"shellclient/pkg/cli"

	"github.com/spf13/cobra"
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Deployment resource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Parent().Use == "create" {
			fileDescriptorString, _ := cmd.Flags().GetString("file")
			fileDescriptor, err := os.ReadFile(fileDescriptorString)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			cli.CreateDeployment(fileDescriptor)
		} else if cmd.Parent().Use == "get" {
			specificId, _ := cmd.Flags().GetInt64("id")
			if specificId != 0 {
				cli.GetDeploymentById(specificId)
			} else {
				cli.GetDeployment()
			}
		} else if cmd.Parent().Use == "update" {
			specificId, _ := cmd.Flags().GetInt64("id")
			fileDescriptorString, _ := cmd.Flags().GetString("file")
			fileDescriptor, err := os.ReadFile(fileDescriptorString)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			res := cli.UpdateDeployment(specificId, fileDescriptor)
			fmt.Println(res)
		} else if cmd.Parent().Use == "delete" {
			specificId, _ := cmd.Flags().GetInt64("id")
			res := cli.DeleteDeployment(specificId)
			fmt.Println(res)
		}
	},
}

func init() {

	var createDeploymentCmd = *deploymentCmd
	var getDeploymentCmd = *deploymentCmd
	var updateDeploymentCmd = *deploymentCmd
	var deleteDeploymentCmd = *deploymentCmd
	createCmd.AddCommand(&createDeploymentCmd)
	deleteCmd.AddCommand(&deleteDeploymentCmd)
	updateCmd.AddCommand(&updateDeploymentCmd)
	getCmd.AddCommand(&getDeploymentCmd)

	createDeploymentCmd.PersistentFlags().StringP("file", "", "", "App descriptor file")
	updateDeploymentCmd.PersistentFlags().StringP("file", "", "", "App descriptor file")
	getDeploymentCmd.PersistentFlags().Int64P("id", "", 0, "ID of the deployment")
	updateDeploymentCmd.PersistentFlags().Int64P("id", "", 0, "ID of the deployment")
	deleteDeploymentCmd.PersistentFlags().Int64P("id", "", 0, "ID of the deployment")

}
