package goconfig


func (f *config) newKeyValue(k string, v []byte) {
	// 同一个组里面添加注释
	f.Lines = make([]*node, 0)
	n := &node{
		key: k,
		value :v,
		note: make([][]byte, 0),
	}
	f.Lines = append(f.Lines, n)
	f.KeyValue[k] = v
}

func (f *config) newNote(note []byte) {
	// 同一个组里面添加注释
	nt := make([][]byte, 0)
	nt =append(nt, note)
	f.Lines = make([]*node, 0)
	n := &node{
		note: nt,
	}
	f.Lines = append(f.Lines, n)
}

func (f *config) addKeyValue(k string, v []byte) {
	// 同一个组里面添加注释
	ll := len(f.Lines)
	f.Lines[ll-1].key = k
	f.Lines[ll-1].value = v
	f.KeyValue[k] = v
}

func (f *config) addNote(note []byte) {
	// 里面添加注释
	ll := len(f.Lines)
	f.Lines[ll-1].note = append(f.Lines[ll-1].note, note)
}




