package core

import v1 "github.com/tpology/core/api/v1"

// CheckLoadPolicy checks the given resource and filename against the load
// policies. The effect of the first policy that matches will be returned.
func CheckLoadPolicy(resource interface{}, filename string, policies []*v1.LoadPolicy) v1.Effect {
	return ""
}
