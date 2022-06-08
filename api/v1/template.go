package v1

// TemplateSpec is the specification of a template.
type TemplateSpec struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
}

// ValidTemplateSpecFields is the list of valid fields in a TemplateSpec.
var ValidTemplateSpecFields = []string{"name", "content"}

// Template represents a Tpology template
type Template struct {
	APIVersion string       `yaml:"apiVersion"`
	Template   TemplateSpec `yaml:"template"`
}

// ValidTemplateFields is the list of valid fields in a Template.
var ValidTemplateFields = []string{"apiVersion", "template"}
