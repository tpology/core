package v1

type OutputSpec struct {
	// Name is the name of the output artifact being generated.
	Name string `yaml:"name"`
	// Repository is the name of the repository the output artifact will be committed to.
	Repository string `yaml:"repository"`
	// File is the full path to the output artifact in the repository.
	File string `yaml:"file"`
	// Template is the name of the template that produces the output artifact.
	Template string `yaml:"template"`
	// Context specifies the context used to render the template.
	Context string `yaml:"context"`
	// PostProcessor is the name of a post-processor to run on the template output
	// to produce the output artifact.
	PostProcessor string `yaml:"postProcessor"`
}

// ValidOutputSpecFields is the list of valid fields in a OutputSpec.
var ValidOutputSpecFields = []string{"name", "repository", "file", "template", "context", "postProcessor"}

// ResourceSpec is the specification of a resource.
type ResourceSpec struct {
	Name        string            `yaml:"name"`
	Kind        string            `yaml:"kind"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations"`
	Data        interface{}       `yaml:"data"`
	Outputs     []OutputSpec      `yaml:"outputs"`
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
