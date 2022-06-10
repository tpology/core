package v1

// DefaultContext is the context passed to a template if some other context
// is not specified.
type DefaultContext struct {
	// Self is the resource spec of the resource generating the output.
	Self *ResourceSpec `yaml:"self"`
	// Output is the current output of the resource.
	Output *OutputSpec `yaml:"output"`
	// Resources is a map of all the resource specs, keyed by kind and then
	// name.
	Resources map[string]map[string]*ResourceSpec `yaml:"resources"`
	// Templates is a map of all the template specs, keyed by name.
	Templates map[string]*TemplateSpec `yaml:"templates"`
	// Repositories is a map of all the repository specs, keyed by name.
	Repositories map[string]*RepositorySpec `yaml:"repositories"`
}
