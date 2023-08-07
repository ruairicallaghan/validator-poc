package cmd

import (
	"fmt"
	"log"
	"validator-refactor/config"
	"validator-refactor/utility"
	"validator-refactor/validator/rules/cloud_run"
	"validator-refactor/validator/rules/common"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCloudRunCommand() *cobra.Command {
	var cloudRunCmd = &cobra.Command{
		Use:   "cloud_run",
		Short: "Generates a Cloud Run ",
		Run:   runCloudRun,
	}

	return cloudRunCmd
}

func runCloudRun(_ *cobra.Command, _ []string) {
	utility.InitViper()
	config := config.CloudRunServiceConfig{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to unmarshal to ChartConfig struct: %s", err)
	}

	v := cloud_run.CloudRunValidator{}
	c := common.CommonValidator{}

	if err := v.Validate(config); err != nil {
		log.Fatalf("Unable to validate config: %s", err)
	}

	if err := c.Validate(config); err != nil {
		log.Fatalf("Unable to validate config: %s", err)
	}

	fmt.Println("Validated OK!")

}
