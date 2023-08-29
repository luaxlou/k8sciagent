/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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

		Deploy(app, image, version, env)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	rootCmd.Flags().StringP("version", "v", "", "new app version")
	rootCmd.Flags().StringP("image", "i", "", "docker image name")
	rootCmd.Flags().StringP("app", "a", "", "app name")
	rootCmd.Flags().StringP("env", "e", "", "env")
}

func Deploy(app, image, version, env string) {

	iv := image + ":" + version

	k8sDeployName := "deployment.v1.apps/" + app

	runCmd("docker", "build", "-t", iv, ".")
	runCmd("docker", "build", "-t", image+":latest", ".")
	runCmd("docker", "push", iv)
	runCmd("docker", "push", "latest")
	runCmd("kubectl", "--kubeconfig", env+".config", "--record", k8sDeployName, "set", "image", k8sDeployName, app+"="+iv)

}

func runCmd(cmdName string, arg ...string) {
	cmd := exec.Command(cmdName, arg...)
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("\n%s\n", string(out))
}
