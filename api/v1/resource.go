package v1

type OutputSpec struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
	Path       string `yaml:"path"`
	Template   string `yaml:"template"`
}

// ResourceSpec is the specification of a resource.
type ResourceSpec struct {
	Name        string                 `yaml:"name"`
	Kind        string                 `yaml:"kind"`
	Labels      map[string]string      `yaml:"labels"`
	Annotations map[string]string      `yaml:"annotations"`
	Data        map[string]interface{} `yaml:"data"`
	Outputs     []OutputSpec           `yaml:"outputs"`
}

// ValidResourceSpecFields is the list of valid fields in a ResourceSpec.
var ValidResourceSpecFields = []string{"name", "kind", "labels", "annotations", "data", "outputs"}

// Resource represents a Tpology resource
type Resource struct {
	APIVersion string       `yaml:"apiVersion"`
	Resource   ResourceSpec `yaml:"resource"`
}

// ValidResourceFields is the list of valid fields in a Resource.
var ValidResourceFields = []string{"apiVersion", "resource"}
