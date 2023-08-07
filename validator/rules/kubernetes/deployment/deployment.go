package deployment

import (
	"fmt"
	"validator-refactor/config"
)

type KubernetesDeploymentValidator struct{}

func (v *KubernetesDeploymentValidator) Validate(config config.KubernetesDeploymentConfig) error {
	if err := ValidateDeployment(config.Deployment); err != nil {
		return err
	}
	return nil
}

func ValidateDeployment(deployment bool) error {
	if deployment == true {
		fmt.Printf("%v is a valid deployment\n", deployment)
		return nil
	} else {
		fmt.Println("Deployment: ", deployment)
		return fmt.Errorf("Deployment must be set to true")
	}
}
