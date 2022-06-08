package core

import "testing"

// Test_Validate_MissingResourceKind tests that a resource without a kind is not valid
func Test_Validate_MissingResourceKind(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/validate/009-missing-resource-kind")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "resource kind is required" {
		t.Errorf("expected 'resource kind is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingResourceName tests that a resource without a name is not valid
func Test_Validate_MissingResourceName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/validate/010-missing-resource-name")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "resource name is required" {
		t.Errorf("expected 'resource name is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingTemplateName tests that a template without a name is not valid
func Test_Validate_MissingTemplateName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/validate/011-missing-template-name")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "template name is required" {
		t.Errorf("expected 'template name is required', got '%s'", errs[0].Error())
	}
}
