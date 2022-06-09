package v1

// Effect is the effect of a policy rule.
type Effect string

// Deny is the effect of a policy rule that denies access.
const Deny Effect = "deny"

// Allow is the effect of a policy rule that allows access.
const Allow Effect = "allow"

// LoadPolicy represents a load policy which allows or denies a resource to
// be loaded from the directory structure based on its location and contents.
type LoadPolicy struct {
	// Name is the name of the load policy.
	Name string `yaml:"name"`
	// Effect is the effect of the load policy.
	Effect Effect `yaml:"effect"`
	// Paths is a list of path patterns of the resources covered by the load policy.
	Paths []string `yaml:"path"`
	// Resource is the spec of a resource covered by the policy. All non-empty
	// fields of this resource must match the actual resource being loaded for
	// the policy to match.
	Resource *ResourceSpec `yaml:"resource"`
	// Template is the spec of a template covered by the policy. All non-empty
	// fields of this template must match the actual template being loaded for
	// the policy to match.
	Template *TemplateSpec `yaml:"template"`
	// Repository is the spec of a repository covered by the policy. All non-empty
	// fields of this repository must match the actual repository being loaded for
	// the policy to match.
	Repository *RepositorySpec `yaml:"repository"`
}

// ValidLoadPolicyFields is the list of valid fields in a LoadPolicy.
var ValidLoadPolicyFields = []string{"name", "effect", "path", "resource", "template", "repository"}
