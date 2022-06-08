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

// Test_Valid_InvalidResourceField tests that an invalid resource field is not valid
func Test_Validate_InvalidResourceField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/validate/012-invalid-resource-field")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "invalid resource field `invalid`" {
		t.Errorf("expected 'invalid resource field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidResourceSpecField tests that an invalid resource spec field is not valid
func Test_Validate_InvalidResourceSpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/validate/013-invalid-resource-spec-field")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "invalid resource spec field `invalid`" {
		t.Errorf("expected 'invalid resource spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidTemplateField tests that an invalid template field is not valid
func Test_Validate_InvalidTemplateField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/validate/014-invalid-template-field")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "invalid template field `invalid`" {
		t.Errorf("expected 'invalid template field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidTemplateSpecField tests that an invalid template spec field is not valid
func Test_Validate_InvalidTemplateSpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/validate/015-invalid-template-spec-field")
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "invalid template spec field `invalid`" {
		t.Errorf("expected 'invalid template spec field `invalid`', got '%s'", errs[0].Error())
	}
}
