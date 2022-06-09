package core

import "testing"

// Test_Validate_MissingResourceKind tests that a resource without a kind is not valid
func Test_Validate_MissingResourceKind(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/009-missing-resource-kind", nil)
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
	errs := i.Load("testdata/010-missing-resource-name", nil)
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
	errs := i.Load("testdata/011-missing-template-name", nil)
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
	errs := i.Load("testdata/012-invalid-resource-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/012-invalid-resource-field/resource-1.yaml: invalid resource field `invalid`" {
		t.Errorf("expected 'testdata/012-invalid-resource-field/resource-1.yaml: invalid resource field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidResourceSpecField tests that an invalid resource spec field is not valid
func Test_Validate_InvalidResourceSpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/013-invalid-resource-spec-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/013-invalid-resource-spec-field/resource-1.yaml: invalid resource spec field `invalid`" {
		t.Errorf("expected 'testdata/013-invalid-resource-spec-field/resource-1.yaml: invalid resource spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidTemplateField tests that an invalid template field is not valid
func Test_Validate_InvalidTemplateField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/014-invalid-template-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/014-invalid-template-field/template-1.yaml: invalid template field `invalid`" {
		t.Errorf("expected 'testdata/014-invalid-template-field/template-1.yaml: invalid template field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidTemplateSpecField tests that an invalid template spec field is not valid
func Test_Validate_InvalidTemplateSpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/015-invalid-template-spec-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/015-invalid-template-spec-field/template-1.yaml: invalid template spec field `invalid`" {
		t.Errorf("expected 'testdata/015-invalid-template-spec-field/template-1.yaml: invalid template spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidRepositoryField tests that an invalid repository field is not valid
func Test_Validate_InvalidRepositoryField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/021-invalid-repository-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/021-invalid-repository-field/repository-1.yaml: invalid repository field `invalid`" {
		t.Errorf("expected 'testdata/021-invalid-repository-field/repository-1.yaml: invalid repository field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidRepositorySpecField tests that an invalid repository spec field is not valid
func Test_Validate_InvalidRepositorySpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/022-invalid-repository-spec-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/022-invalid-repository-spec-field/repository-1.yaml: invalid repository spec field `invalid`" {
		t.Errorf("expected 'testdata/022-invalid-repository-spec-field/repository-1.yaml: invalid repository spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingRepositoryName tests that a repository without a name is not valid
func Test_Validate_MissingRepositoryName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/018-missing-repository-name", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "repository name is required" {
		t.Errorf("expected 'repository name is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingRepository tests that a repository without a repo is not valid
func Test_Validate_MissingRepository(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/019-missing-repository", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "repository is required" {
		t.Errorf("expected 'repository is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingRepositoryBranch tests that a repository without a branch is not valid
func Test_Validate_MissingRepositoryBranch(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/020-missing-repository-branch", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "repository branch is required" {
		t.Errorf("expected 'repository branch is required', got '%s'", errs[0].Error())
	}
}
