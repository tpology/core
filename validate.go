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
