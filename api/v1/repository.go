package v1

// RepositorySpec is the specification of a repository.
type RepositorySpec struct {
	Name        string            `yaml:"name"`
	Repository  string            `yaml:"repository"`
	Branch      string            `yaml:"branch"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations"`
}

// ValidRepositorySpecFields is the list of valid fields in a RepositorySpec.
var ValidRepositorySpecFields = []string{"name", "repository", "branch", "labels", "annotations"}

// Repository represents a Tpology repository
type Repository struct {
	APIVersion string         `yaml:"apiVersion"`
	Repository RepositorySpec `yaml:"repository"`
}

// ValidRepositoryFields is the list of valid fields in a Repository.
var ValidRepositoryFields = []string{"apiVersion", "repository"}
