package v1

// DefaultContext is the context passed to a template if some other context
// is not specified.
type DefaultContext struct {
	// SelfType is the type of the Self object.
	SelfType string `yaml:"selfType"`
	// Self is the resource spec of the resource generating the output.
	Self interface{} `yaml:"self"`
	// Resources is a map of all the resource specs, keyed by kind and then
	// name.
	Resources map[string]map[string]*ResourceSpec `yaml:"resources"`
	// Templates is a map of all the template specs, keyed by name.
	Templates map[string]*TemplateSpec `yaml:"templates"`
	// Repositories is a map of all the repository specs, keyed by name.
	Repositories map[string]*RepositorySpec `yaml:"repositories"`
}
