package core

import (
	"fmt"

	v1 "github.com/tpology/core/api/v1"
)

// validate validates the Index
func (i *Index) validate() []error {
	errs := []error{}
	// validate resources
	for _, resources := range i.resourceByKind {
		for _, r := range resources {
			errs = append(errs, validateResource(r)...)
		}
	}
	// validate templates
	for _, t := range i.template {
		errs = append(errs, validateTemplate(t)...)
	}
	// validate repositories
	for _, r := range i.repository {
		errs = append(errs, validateRepository(r)...)
	}
	return errs
}

// validateResource validates the resource
func validateResource(r *v1.Resource) []error {
	errs := []error{}
	// validate kind
	if r.Resource.Kind == "" {
		errs = append(errs, fmt.Errorf("resource kind is required"))
	}
	// validate name
	if r.Resource.Name == "" {
		errs = append(errs, fmt.Errorf("resource name is required"))
	}
	return errs
}

// validateTemplate validates the template
func validateTemplate(t *v1.Template) []error {
	errs := []error{}
	// validate name
	if t.Template.Name == "" {
		errs = append(errs, fmt.Errorf("template name is required"))
	}
	return errs
}

// validateRepository validates the repository
func validateRepository(r *v1.Repository) []error {
	errs := []error{}
	// validate name
	if r.Repository.Name == "" {
		errs = append(errs, fmt.Errorf("repository name is required"))
	}
	// validate repository
	if r.Repository.Repository == "" {
		errs = append(errs, fmt.Errorf("repository is required"))
	}
	// validate branch
	if r.Repository.Branch == "" {
		errs = append(errs, fmt.Errorf("repository branch is required"))
	}
	return errs
}

// validateFields validates the fields against a list of valid fields.
func validateFields(kind string, r map[string]interface{}, validFields []string) []error {
FIELD:
	for k := range r {
		for _, f := range validFields {
			if k == f {
				continue FIELD
			}
		}
		return []error{fmt.Errorf("invalid %s field `%s`", kind, k)}
	}
	return nil
}

// validateSpecFields validates the fields against a list of valid fields.
func validateSpecFields(kind string, r map[interface{}]interface{}, validFields []string) []error {
FIELD:
	for k := range r {
		for _, f := range validFields {
			if k == f {
				continue FIELD
			}
		}
		return []error{fmt.Errorf("invalid %s spec field `%s`", kind, k)}
	}
	return nil
}

// validateResourceFields validates the fields in a Resource.
func validateResourceFields(r map[string]interface{}) []error {
	return validateFields("resource", r, v1.ValidResourceFields)
}

// validateResourceSpecFields validates the fields in a ResourceSpec.
func validateResourceSpecFields(r map[interface{}]interface{}) []error {
	return validateSpecFields("resource", r, v1.ValidResourceSpecFields)
}

// validateOutputSpecFields validates the fields in a OutputSpec.
func validateOutputSpecFields(r map[interface{}]interface{}) []error {
	return validateSpecFields("output", r, v1.ValidOutputSpecFields)
}

// validateTemplateFields validates the fields in a Template.
func validateTemplateFields(r map[string]interface{}) []error {
	return validateFields("template", r, v1.ValidTemplateFields)
}

// validateTemplateSpecFields validates the fields in a TemplateSpec.
func validateTemplateSpecFields(r map[interface{}]interface{}) []error {
	return validateSpecFields("template", r, v1.ValidTemplateSpecFields)
}

// validateRepositoryFields validates the fields in a Repository.
func validateRepositoryFields(r map[string]interface{}) []error {
	return validateFields("repository", r, v1.ValidRepositoryFields)
}

// validateRepositorySpecFields validates the fields in a RepositorySpec.
func validateRepositorySpecFields(r map[interface{}]interface{}) []error {
	return validateSpecFields("repository", r, v1.ValidRepositorySpecFields)
}
