package goconfig

func (f *config) newGroupKeyValue(modele []byte, k string, v []byte) {
	// 新组的key
	tmp := &node{
		key: k,
		value: v,
		note: make([][]byte, 0),
	}
	g := make([]*node,0)
	g = append(g, tmp)
	f.Groups = append(f.Groups, &groupLine{
		group: g,  // 组的行
		note: make([][]byte, 0),  // 组注释
		name: modele,   // 组名
	})
	f.KeyValue[string(modele) + "." +k] = v
}

func (f *config) newGroupNote(modele []byte,value []byte) {
	// 新组的注释
	nt := make([][]byte, 0)
	nt = append(nt, value)
	tmp := &node{
		note: nt,
	}
	g := make([]*node,0)
	g = append(g, tmp)
	f.Groups = append(f.Groups, &groupLine{
		group: g,  // 组的行
		note: make([][]byte, 0),  // 组注释
		name: modele,   // 组名
	})
}

func (f *config) addGroupNote(note []byte) {
	// 同一个组里面添加注释
	gl := len(f.Groups)
	f.Groups[gl-1].note = append(f.Groups[gl-1].note, note)
}

func (f *config) addGroupKeyValue(k string, v []byte) {
	// 同一个组里面增加key
	tmp := &node{
		key: k,
		value: v,
		note: make([][]byte, 0),
	}
	gl := len(f.Groups)
	f.Groups[gl-1].group = append(f.Groups[gl-1].group, tmp)
	f.KeyValue[string(f.Groups[gl-1].name) + "." +k] = v
}