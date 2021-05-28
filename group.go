package goconfig

func (f *config) newGroup(module string, notes ...string) {
	f.Groups = append(f.Groups, &groupLine{
		note: notes,  // 组注释
		name: module, // 组名
	})
}

func (c *config) newGroupKeyValue(module string, k string, v string, notes ...string) {
	// 新组的key
	tmp := node{
		key:   k,
		value: v,
		note:  notes,
	}
	g := make([]node, 0)
	g = append(g, tmp)
	c.Groups = append(c.Groups, &groupLine{
		group: g,                 // 组的行
		note:  make([]string, 0), // 组注释
		name:  module,            // 组名
	})
	// c.KeyValue[module+"."+k] = v
}

func (c *config) addGroupKeyValue(i int, k string, v string, notes ...string) {
	// 新组的key
	tmp := node{
		key:   k,
		value: v,
		note:  notes,
	}
	c.Groups[i].group = append(c.Groups[i].group, tmp)
	newkey := string(c.Groups[i].name) + "." + k
	if k[0:1] == "$" {
		c.env[newkey] = v
	}
	c.KeyValue[newkey] = v

}
