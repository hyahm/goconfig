package goconfig

func (f *config) newGroup(module string, notes ...string) {
	// 新组的key
	if len(f.Groups) == 0 {
		f.Groups = make([]*groupLine, 0)
	}

	f.Groups = append(f.Groups, &groupLine{
		note: notes,  // 组注释
		name: module, // 组名
	})
}

func (f *config) newGroupKeyValue(module string, k string, v string, notes ...string) {
	// 新组的key
	if len(f.Groups) == 0 {
		f.Groups = make([]*groupLine, 0)
	}
	tmp := node{
		key:   k,
		value: v,
		note:  notes,
	}
	g := make([]node, 0)
	g = append(g, tmp)
	f.Groups = append(f.Groups, &groupLine{
		group: g,                 // 组的行
		note:  make([]string, 0), // 组注释
		name:  module,            // 组名
	})
}

func (f *config) addGroupKeyValue(i int, k string, v string, notes ...string) {
	// 新组的key
	tmp := node{
		key:   k,
		value: v,
		note:  notes,
	}
	f.Groups[i].group = append(f.Groups[i].group, tmp)
	f.KeyValue[string(f.Groups[i].name)+"."+k] = v
}
