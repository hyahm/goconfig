package goconfig

func (f *config) newKeyValue(k string, v string, notes ...string) {
	// 同一个组里面添加注释
	if len(f.Lines) == 0 {
		f.Lines = make([]*node, 0)
	}
	n := &node{
		key:   k,
		value: v,
		note:  notes,
	}
	f.Lines = append(f.Lines, n)
	f.KeyValue[k] = v
}
