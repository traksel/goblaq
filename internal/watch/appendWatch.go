package watch

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type data struct {
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
	Schema string `yaml:"schema"`
	Path   string `yaml:"path"`
}

var home = os.Getenv("HOME")

func (d *data) getActualData(name string, target string, schema string, path string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/.goblaq/%s/", home, name)); os.IsNotExist(err) {
		os.MkdirAll(fmt.Sprintf("%s/.goblaq/%s", home, name), os.ModePerm)
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, name))
	yaml.Unmarshal(file, d)
	return err
}

func (d *data) fillData(name string, target string, schema string, path string) *data {
	d.Name = name
	d.Target = target
	d.Schema = schema
	d.Path = path
	return d
}

func AppendWatch(name string, target string, schema string, path string) error {
	var d data
	err := d.getActualData(name, target, schema, path)
	if err != nil {
		file := d.fillData(name, target, schema, path)
		writeData, _ := yaml.Marshal(file)
		os.WriteFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, name), writeData, os.ModePerm)
	}
	fmt.Printf(
		"SERVICE ADDED:\nName: %s\nTarget: %s\nSchema: %s\nTargets: %v",
		name,
		target,
		schema,
		path,
	)
	return nil
}
