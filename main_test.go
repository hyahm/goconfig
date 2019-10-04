package goconfig

import (
	"testing"
)

func Test_main(t *testing.T) {
	InitConf("showgo.conf")
	Print()
	//gaconfig.ReadString("")
}



