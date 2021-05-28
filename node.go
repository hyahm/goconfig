package goconfig

func (c *config) newKeyValue(k string, v string, notes ...string) {
	// 同一个组里面添加注释
	n := &node{
		key:   k,
		value: v,
		note:  notes,
	}
	c.Lines = append(c.Lines, n)
	newkey := k
	if k[0:1] == "$" {
		newkey = k[1:]
		c.env[newkey] = v
	}
	c.KeyValue[newkey] = v
}
