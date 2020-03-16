// 测试读取 duration

package goconfig

import "testing"

func TestDuration(t *testing.T) {
	var timeout dr = "5s"
	var td dr = "5d"
	var th dr = "5h"
	var tm dr = "5M"
	var te dr = "5e"
	t.Log(timeout.Duration())
	t.Log(td.Duration())
	t.Log(th.Duration())
	t.Log(tm.Duration())
	m, err := te.Duration()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(m)
}
