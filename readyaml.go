package goconfig

import "encoding/json"

func (c *config) readYaml() error {

	err := json.Unmarshal(c.Read, &c.sjson)
	if err != nil {
		return err
	}
	if err := c.parseJson("", c.sjson); err != nil {
		return err
	}

	return nil
}
