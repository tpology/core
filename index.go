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
	repository     map[string]*v1.Repository
}

// NewIndex returns a new Index
func NewIndex() *Index {
	return &Index{
		resourceByKind: map[string]map[string]*v1.Resource{},
		template:       map[string]*v1.Template{},
		repository:     map[string]*v1.Repository{},
	}
}

// AddResource adds a resource to the index
func (i *Index) AddResource(r *v1.Resource) error {
	if _, ok := i.resourceByKind[r.Resource.Kind]; !ok {
		i.resourceByKind[r.Resource.Kind] = map[string]*v1.Resource{}
	}
	if _, ok := i.resourceByKind[r.Resource.Kind][r.Resource.Name]; ok {
		return fmt.Errorf("resource %s of kind %s already exists", r.Resource.Name, r.Resource.Kind)
	}
	i.resourceByKind[r.Resource.Kind][r.Resource.Name] = r
	return nil
}

// RemoveResource removes a resource from the index
func (i *Index) RemoveResource(r *v1.Resource) error {
	if _, ok := i.resourceByKind[r.Resource.Kind]; ok {
		delete(i.resourceByKind[r.Resource.Kind], r.Resource.Name)
		if len(i.resourceByKind[r.Resource.Kind]) == 0 {
			delete(i.resourceByKind, r.Resource.Kind)
		}
		return nil
	}
	return fmt.Errorf("resource %s of kind %s does not exist", r.Resource.Name, r.Resource.Kind)
}

// AddTemplate adds a template to the index
func (i *Index) AddTemplate(t *v1.Template) error {
	if _, ok := i.template[t.Template.Name]; ok {
		return fmt.Errorf("template %s already exists", t.Template.Name)
	}
	i.template[t.Template.Name] = t
	return nil
}

// RemoveTemplate removes a template from the index
func (i *Index) RemoveTemplate(t *v1.Template) error {
	if _, ok := i.template[t.Template.Name]; ok {
		delete(i.template, t.Template.Name)
		return nil
	}
	return fmt.Errorf("template %s does not exist", t.Template.Name)
}

// AddRepository adds a repository to the index
func (i *Index) AddRepository(r *v1.Repository) error {
	if _, ok := i.repository[r.Repository.Name]; ok {
		return fmt.Errorf("repository %s already exists", r.Repository.Name)
	}
	i.repository[r.Repository.Name] = r
	return nil
}

// RemoveRepository removes a repository from the index
func (i *Index) RemoveRepository(r *v1.Repository) error {
	if _, ok := i.repository[r.Repository.Name]; ok {
		delete(i.repository, r.Repository.Name)
		return nil
	}
	return fmt.Errorf("repository %s does not exist", r.Repository.Name)
}

func (i *Index) Load(dir string) []error {
	errs := []error{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
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
			// Format the error to prepend the resource path
			errs = append(errs, fmt.Errorf("%s: %s", path, err))
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
		if _, ok := doc["resource"]; ok {
			// If there is a resource key, unmarshal as Resource
			verrs := validateResourceFields(doc)
			if len(verrs) > 0 {
				// Format the errors to prepend the resource path
				for _, err := range verrs {
					errs = append(errs, fmt.Errorf("%s: %s", path, err))
				}
				return nil
			}
			verrs = validateResourceSpecFields(doc["resource"].(map[interface{}]interface{}))
			if len(verrs) > 0 {
				// Format the errors to prepend the resource path
				for _, err := range verrs {
					errs = append(errs, fmt.Errorf("%s: %s", path, err))
				}
				return nil
			}
			output := doc["resource"].(map[interface{}]interface{})["output"]
			if output != nil {
				verrs = validateOutputSpecFields(doc["resource"].(map[interface{}]interface{})["output"].(map[interface{}]interface{}))
				if len(verrs) > 0 {
					// Format the errors to prepend the resource path
					for _, err := range verrs {
						errs = append(errs, fmt.Errorf("%s: %s", path, err))
					}
					return nil
				}
			}
			var resource v1.Resource
			err = yaml.Unmarshal(yamlBytes, &resource)
			if err != nil {
				// Format the error to prepend the resource path
				errs = append(errs, fmt.Errorf("%s: %s", path, err))
				return nil
			}
			err = i.AddResource(&resource)
			if err != nil {
				// Format the error to prepend the resource path
				errs = append(errs, fmt.Errorf("%s: %s", path, err))
				return nil
			}
		} else if _, ok := doc["template"]; ok {
			// If there is a template key, unmarshal as Template
			verrs := validateTemplateFields(doc)
			if len(verrs) > 0 {
				// Format the errors to prepend the resource path
				for _, err := range verrs {
					errs = append(errs, fmt.Errorf("%s: %s", path, err))
				}
				return nil
			}
			verrs = validateTemplateSpecFields(doc["template"].(map[interface{}]interface{}))
			if len(verrs) > 0 {
				// Format the errors to prepend the resource path
				for _, err := range verrs {
					errs = append(errs, fmt.Errorf("%s: %s", path, err))
				}
				return nil
			}
			var template v1.Template
			err = yaml.Unmarshal(yamlBytes, &template)
			if err != nil {
				// Format the error to prepend the resource path
				errs = append(errs, fmt.Errorf("%s: %s", path, err))
				return nil
			}
			err = i.AddTemplate(&template)
			if err != nil {
				// Format the error to prepend the resource path
				errs = append(errs, fmt.Errorf("%s: %s", path, err))
				return nil
			}
		} else if _, ok := doc["repository"]; ok {
			// If there is a repository key, unmarshal as Repository
			verrs := validateRepositoryFields(doc)
			if len(verrs) > 0 {
				// Format the errors to prepend the resource path
				for _, err := range verrs {
					errs = append(errs, fmt.Errorf("%s: %s", path, err))
				}
				return nil
			}
			verrs = validateRepositorySpecFields(doc["repository"].(map[interface{}]interface{}))
			if len(verrs) > 0 {
				// Format the errors to prepend the resource path
				for _, err := range verrs {
					errs = append(errs, fmt.Errorf("%s: %s", path, err))
				}
				return nil
			}
			var repository v1.Repository
			err = yaml.Unmarshal(yamlBytes, &repository)
			if err != nil {
				// Format the error to prepend the resource path
				errs = append(errs, fmt.Errorf("%s: %s", path, err))
				return nil
			}
			err = i.AddRepository(&repository)
			if err != nil {
				// Format the error to prepend the resource path
				errs = append(errs, fmt.Errorf("%s: %s", path, err))
				return nil
			}
		} else {
			errs = append(errs, fmt.Errorf("%s: no resource or template", path))
		}
		return nil
	})
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs
	}
	return i.validate()
}
