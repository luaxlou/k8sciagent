/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new app version",
	Long: `Usage:
		k8sciagent deploy -a appName -i imageName -v version -e env`,
	Run: func(cmd *cobra.Command, args []string) {

		version, _ := cmd.Flags().GetString("version")
		image, _ := cmd.Flags().GetString("image")
		app, _ := cmd.Flags().GetString("app")
		env, _ := cmd.Flags().GetString("env")
		path, _ := cmd.Flags().GetString("path")

		Deploy(app, image, version, env, path)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("version", "v", "", "new app version")
	deployCmd.Flags().StringP("image", "i", "", "docker image name")
	deployCmd.Flags().StringP("app", "a", "", "app name")
	deployCmd.Flags().StringP("env", "e", "", "env")
	deployCmd.Flags().StringP("path", "p", ".", "docker context path")
}

func Deploy(app, image, version, env, path string) {

	iv := image + ":" + version
	il := image + ":latest"

	ctx := context.Background()

	k8sDeployName := "deployment.v1.apps/" + app

	runCmd(ctx, "docker", "build", "-t", iv, "-t", il, path)
	runCmd(ctx, "docker", "push", iv)
	runCmd(ctx, "docker", "push", il)
	runCmd(ctx, "kubectl", "--kubeconfig", "/etc/kubectl/"+env+".config", "--record", k8sDeployName, "set", "image", k8sDeployName, app+"="+iv)

}

func runCmd(ctx context.Context, cmdName string, arg ...string) {

	cmd := exec.CommandContext(ctx, cmdName, arg...)
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("\n%s\n", string(out))
}
