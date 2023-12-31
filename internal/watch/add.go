package watch

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Data struct {
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
	Schema string `yaml:"schema"`
	Path   string `yaml:"path"`
	Chain  string `yaml:"chain"`
}

func NewData() *Data {
	d := Data{}
	d.Chain = ""
	return &d
}

var home = os.Getenv("HOME")

func (d *Data) Get(name string) (*Data, error) {
	if _, err := os.Stat(fmt.Sprintf("%s/.goblaq/%s/", home, name)); os.IsNotExist(err) {
		os.MkdirAll(fmt.Sprintf("%s/.goblaq/%s", home, name), os.ModePerm)
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, name))
	yaml.Unmarshal(file, d)
	return d, err
}

func (d *Data) Fill(name string, target string, schema string, path string, chain string) *Data {
	d.Name = name
	d.Target = target
	d.Schema = schema
	d.Path = path
	if chain != "" {
		d.Chain = chain
	}
	return d
}

func Add(name string, target string, schema string, path string, chain string) error {
	d := NewData()
	_, err := d.Get(name)
	if err != nil {
		file := d.Fill(name, target, schema, path, chain)
		writeData, _ := yaml.Marshal(file)
		err := os.WriteFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, name), writeData, os.ModePerm)
		if err == nil {
			fmt.Printf(
				"SERVICE ADDED:\nName: %s\nTarget: %s\nSchema: %s\nTargets: %v\nChain: %s",
				name,
				target,
				schema,
				path,
				chain,
			)
		}
	}
	return nil
}
