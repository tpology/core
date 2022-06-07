package v1

import "testing"

// NewTestResource returns a new Resource for testing
func NewTestResource(name, kind string) *Resource {
	return &Resource{
		APIVersion: "v1",
		Resource: ResourceSpec{
			Name: name,
			Kind: kind,
		},
	}
}

func Test_Resource_Equal(t *testing.T) {
	r1 := NewTestResource("foo", "bar")
	r2 := NewTestResource("foo", "bar")
	if r1 == r2 {
		t.Error("Expected r1 != r2")
	}
}
