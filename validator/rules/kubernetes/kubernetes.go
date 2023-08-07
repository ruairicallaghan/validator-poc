package kubernetes

import (
	"fmt"
	"validator-refactor/config"
)

type KubernetesValidator struct{}

func (v *KubernetesValidator) Validate(config config.KubernetesDeploymentConfig) error {
	if err := ValidateName(config.Name); err != nil {
		return err
	}

	return nil
}

func ValidateName(name string) error {
	if name == "kubernetes" {
		fmt.Printf("%s is a valid name\n", name)
		return nil
	} else {
		fmt.Println("Name: ", name)
		return fmt.Errorf("Name must be kubernetes")
	}
}
