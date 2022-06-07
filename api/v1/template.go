package v1

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
