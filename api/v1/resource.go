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

// Resource represents a Tpology resource
type Resource struct {
	APIVersion string       `yaml:"apiVersion"`
	Resource   ResourceSpec `yaml:"resource"`
}

// TemplateSpec is the specification of a template.
type TemplateSpec struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
}

// Template represents a Tpology template
type Template struct {
	APIVersion string       `yaml:"apiVersion"`
	Template   TemplateSpec `yaml:"template"`
}
