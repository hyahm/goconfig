package goconfig

func (f *config) newKeyValue(k string, v []byte, notes ...[]byte) {
	// 同一个组里面添加注释
	if len(f.Lines) == 0 {
		f.Lines = make([]*node, 0)
	}
	n := &node{
		key: k,
		value :v,
		note: notes,
	}
	f.Lines = append(f.Lines, n)
	f.KeyValue[k] = v
}



////        --------------------------------------------------

//func (f *config) writeNewKeyValue(note []byte) {
//	// 里面添加注释
//	ll := len(f.Lines)
//	f.Lines[ll-1].note = append(f.Lines[ll-1].note, note)
//}
//
//func (f *config) writeNote(note []byte) {
//	// 里面添加注释
//	ll := len(f.Lines)
//	f.Lines[ll-1].note = append(f.Lines[ll-1].note, note)
//}
//
//func (f *config) writeKeyValue(note []byte) {
//	// 里面添加注释
//	ll := len(f.Lines)
//	f.Lines[ll-1].note = append(f.Lines[ll-1].note, note)
//}
//
//func (f *config) writeGroupNote(note []byte) {
//	// 里面添加注释
//	ll := len(f.Lines)
//	f.Lines[ll-1].note = append(f.Lines[ll-1].note, note)
//}
//





