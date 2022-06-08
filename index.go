package core

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	v1 "github.com/tpology/core/api/v1"
	"gopkg.in/yaml.v2"
)

// Index is the index of all resources
type Index struct {
	resourceByKind map[string]map[string]*v1.Resource
	template       map[string]*v1.Template
}

// NewIndex returns a new Index
func NewIndex() *Index {
	return &Index{
		resourceByKind: map[string]map[string]*v1.Resource{},
		template:       map[string]*v1.Template{},
	}
}

// AddResource adds a resource to the index
func (i *Index) AddResource(r *v1.Resource) {
	if _, ok := i.resourceByKind[r.Resource.Kind]; !ok {
		i.resourceByKind[r.Resource.Kind] = map[string]*v1.Resource{}
	}
	i.resourceByKind[r.Resource.Kind][r.Resource.Name] = r
}

// RemoveResource removes a resource from the index
func (i *Index) RemoveResource(r *v1.Resource) {
	if _, ok := i.resourceByKind[r.Resource.Kind]; ok {
		delete(i.resourceByKind[r.Resource.Kind], r.Resource.Name)
		if len(i.resourceByKind[r.Resource.Kind]) == 0 {
			delete(i.resourceByKind, r.Resource.Kind)
		}
	}
}

// AddTemplate adds a template to the index
func (i *Index) AddTemplate(t *v1.Template) {
	i.template[t.Template.Name] = t
}

// RemoveTemplate removes a template from the index
func (i *Index) RemoveTemplate(t *v1.Template) {
	delete(i.template, t.Template.Name)
}

func (i *Index) Load(dir string) []error {
	errs := []error{}
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			errs = append(errs, err)
			return nil
		}
		if d.IsDir() || (filepath.Ext(d.Name()) != ".yaml" && filepath.Ext(d.Name()) != ".yml") {
			return nil
		}
		yamlBytes, err := ioutil.ReadFile(path)
		if err != nil {
			errs = append(errs, err)
			return nil
		}
		var doc map[string]interface{}
		err = yaml.Unmarshal(yamlBytes, &doc)
		if err != nil {
			errs = append(errs, err)
			return nil
		}
		// There must be APIVersion = v1
		apiVersion, ok := doc["apiVersion"].(string)
		if !ok {
			errs = append(errs, fmt.Errorf("%s: no apiVersion", path))
			return nil
		}
		if apiVersion != "v1" {
			errs = append(errs, fmt.Errorf("%s: invalid apiVersion", path))
			return nil
		}
		// If there is a resource key, unmarshal as Resource
		if _, ok := doc["resource"]; ok {
			verrs := validateResourceFields(doc)
			if len(verrs) > 0 {
				errs = append(errs, verrs...)
				return nil
			}
			verrs = validateResourceSpecFields(doc["resource"].(map[interface{}]interface{}))
			if len(verrs) > 0 {
				errs = append(errs, verrs...)
				return nil
			}
			var resource v1.Resource
			err = yaml.Unmarshal(yamlBytes, &resource)
			if err != nil {
				errs = append(errs, err)
				return nil
			}
			i.AddResource(&resource)
			// If there is a template key, unmarshal as Template
		} else if _, ok := doc["template"]; ok {
			verrs := validateTemplateFields(doc)
			if len(verrs) > 0 {
				errs = append(errs, verrs...)
				return nil
			}
			verrs = validateTemplateSpecFields(doc["template"].(map[interface{}]interface{}))
			if len(verrs) > 0 {
				errs = append(errs, verrs...)
				return nil
			}
			var template v1.Template
			err = yaml.Unmarshal(yamlBytes, &template)
			if err != nil {
				errs = append(errs, err)
				return nil
			}
			i.AddTemplate(&template)
		} else {
			errs = append(errs, fmt.Errorf("%s: no resource or template", path))
		}
		return nil
	})
	if len(errs) > 0 {
		return errs
	}
	return i.validate()
}
