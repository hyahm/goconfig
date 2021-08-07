package goconfig

import (
	"sync"
	"testing"
)

func TestRead(t *testing.T) {
	x := `
me = true
aaa = 888



$bbb=999
[server]
$domain=http://192.168.50.250
listen = :9090
always = true

[u5.redis_download]
on = false
key = u5_download${server.domain}

`
	// 这个是ini的config
	kvconfig = &config{
		Groups:   make([]*groupLine, 0),
		Lines:    make([]*node, 0),
		Read:     []byte(x),
		KeyValue: make(map[string]string),
		env:      make(map[string]string),
		mu:       &sync.RWMutex{},
	}
	kvconfig.readIni()
	PrintKeyValue()
}
