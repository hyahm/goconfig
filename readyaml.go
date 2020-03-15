package goconfig

import (
	"sigs.k8s.io/yaml"
)

func (c *config) readYaml() error {

	j, err := yaml.YAMLToJSON(c.Read)
	if err != nil {
		return err
	}
	c.Read = j
	if err := c.readJson(); err != nil {
		return err
	}

	return nil
}
