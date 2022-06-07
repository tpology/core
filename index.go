package core

import (
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
		var res v1.Resource
		var tmp v1.Template
		if err = yaml.Unmarshal(yamlBytes, &res); err != nil {
			if err = yaml.Unmarshal(yamlBytes, &tmp); err != nil {
				errs = append(errs, err)
				return nil
			}
			i.template[tmp.Template.Name] = &tmp
		} else {
			if _, ok := i.resourceByKind[res.Resource.Kind]; !ok {
				i.resourceByKind[res.Resource.Kind] = map[string]*v1.Resource{}
			}
			i.resourceByKind[res.Resource.Kind][res.Resource.Name] = &res
		}
		return nil
	})
	return errs
}
