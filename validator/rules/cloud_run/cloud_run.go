package cloud_run

import (
	"fmt"
	"validator-refactor/config"
)

type CloudRunValidator struct{}

func (v *CloudRunValidator) Validate(config config.CloudRunServiceConfig) error {
	return ValidateName(config.Name)
}

// Rule
func ValidateName(name string) error {
	if name == "cloud_run" {
		fmt.Printf("%s is a valid name\n", name)
		return nil
	} else {
		fmt.Println("Name: ", name)
		return fmt.Errorf("name must be cloud_run")
	}
}
