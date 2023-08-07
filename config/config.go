package config

type Annotation struct {
	Name string `yaml:"name"`
	Val  string `yaml:"value"`
}

type ContainerEnv struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Config interface {
	GetName() string
}

type CloudRunConfig interface {
	GetProjectID() string
}

type CloudRunServiceConfig struct {
	Annotations  []Annotation   `yaml:"annotations"`
	ContainerEnv []ContainerEnv `yaml:"containerEnv"`
	ProjectId    string         `yaml:"projectId"`
	Image        string         `yaml:"image"`
	Kpt          bool           `yaml:"kpt"`
	Name         string         `yaml:"name"`
}

func (c CloudRunServiceConfig) GetName() string {
	return c.Name
}

func (c CloudRunServiceConfig) GetProjectID() string {
	return c.ProjectId
}

type KubernetesConfig interface {
	GetEnvironment() string
}

type KubernetesDeploymentConfig struct {
	Annotations  []Annotation   `yaml:"annotations"`
	ContainerEnv []ContainerEnv `yaml:"containerEnv"`
	Deployment   bool           `yaml:"deployment"`
	Environment  string         `yaml:"env"`
	Image        string         `yaml:"image"`
	Kpt          bool           `yaml:"kpt"`
	Name         string         `yaml:"name"`
}

type KubernetesJobConfig struct {
	Annotations  []Annotation   `yaml:"annotations"`
	ContainerEnv []ContainerEnv `yaml:"containerEnv"`
	Environment  string         `yaml:"env"`
	Image        string         `yaml:"image"`
	JobName      string         `yaml:"jobName"`
	Kpt          bool           `yaml:"kpt"`
	Name         string         `yaml:"name"`
}

func (k KubernetesDeploymentConfig) GetName() string {
	return k.Name
}

func (k KubernetesDeploymentConfig) GetEnvironment() string {
	return k.Environment
}

func (k KubernetesJobConfig) GetEnvironment() string {
	return k.Environment
}
