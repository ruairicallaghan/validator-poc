package cmd

import (
	"fmt"
	"log"
	"validator-refactor/config"
	"validator-refactor/utility"
	"validator-refactor/validator/rules/common"
	"validator-refactor/validator/rules/kubernetes"
	"validator-refactor/validator/rules/kubernetes/deployment"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Declaring a variable for command's Run function to override at test time
var runDeploymentFunction = runDeployment

func NewDeploymentCommand() *cobra.Command {

	var deploymentCmd = &cobra.Command{
		PreRun: func(cmd *cobra.Command, args []string) {},
		Use:    "deployment",
		Short:  "Generates a Kubernetes deployment.",
		Run:    runDeploymentFunction,
	}

	return deploymentCmd
}

func runDeployment(_ *cobra.Command, _ []string) {
	utility.InitViper()
	config := config.KubernetesDeploymentConfig{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to unmarshal to ChartConfig struct: %s", err)
	}

	v := kubernetes.KubernetesValidator{}
	d := deployment.KubernetesDeploymentValidator{}
	c := common.CommonValidator{}

	if err := v.Validate(config); err != nil {
		log.Fatalf("Unable to validate config: %s", err)
	}

	if err := d.Validate(config); err != nil {
		log.Fatalf("Unable to validate config: %s", err)
	}

	if err := c.Validate(config); err != nil {
		log.Fatalf("Unable to validate config: %s", err)
	}

	fmt.Println("Validated OK!")
}
