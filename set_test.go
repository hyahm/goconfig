package goconfig

import (
	"strings"
	"testing"
)

var a = `# 是否使用了代码，为了获取ip，可能不起作用
httpproxy = true
# 自动初始化数据库（为了减少启动时间，第一次初始化完成就关掉
initdb = true
# 监听地址
listenaddr = :10001
# 存放图片的目录
imgdir = /data/bugimg/
# 图片显示的地址(用接口的地址)
showbaseurl = http://127.0.0.1:10001/showimg
# 盐值，建议修改，然后用curl http://ip:10001/admin/reset?password=123 来修改root密码
salt = hjkkaksjdhfryuooweqzmbvc

# ssl, 使用ssl
[ssl]
on = false
cert = 
key = 

[log]
# 日志目录, 不设置就控制台输出
path = 
# 日志大小备份一次， 0为不切割大小
size = 0
# 每天备份一次 大小也存在的话，此项优先 ，false为不每天备份一次
everyday = false

[mysql]
user = root
pwd = 
host = 127.0.0.1
port = 3306
db = project
`

func TestSet(t *testing.T) {
	t.Log("start test write profile to ini file")
	InitConf("test.ini", INI)
	WriteBool("httpproxy", true, "是否使用了代码，为了获取ip，可能不起作用")
	WriteBool("initdb", true, "自动初始化数据库（为了减少启动时间，第一次初始化完成就关掉")
	WriteString("listenaddr", ":10001", "监听地址")
	WriteString("imgdir", "/data/bugimg/", "存放图片的目录")
	WriteString("showbaseurl", "http://127.0.0.1:10001/showimg", "图片显示的地址(用接口的地址)")
	WriteString("salt", "hjkkaksjdhfryuooweqzmbvc", "盐值，建议修改，然后用curl http://ip:10001/admin/reset?password=123 来修改root密码")
	WriteInt64("httpproxy", 120, "token 过期时间")
	WriteString("httpproxy", "/share/", "共享文件夹根目录")
	WriteBool("httpproxy", true, "排除记录这些ip日志")
	WriteBool("httpproxy", true, "是否初始化api帮助文档")
	WriteString("httpproxy", "help", "api项目名")

	WriteNotesForModule("ssl", "ssl, 使用ssl")
	WriteBool("ssl.on", false)
	WriteString("ssl.cert", "")
	WriteString("ssl.key", "")

	WriteString("log.path", "", "日志目录, 不设置就控制台输出")
	WriteInt64("log.size", 0, "日志大小备份一次， 0为不切割大小")
	WriteBool("log.everyday", false, "每天备份一次 大小也存在的话，此项优先 ，false为不每天备份一次")

	WriteString("mysql.user", "root")
	WriteString("mysql.pwd", "")
	WriteString("mysql.host", "127.0.0.1")
	WriteInt("mysql.port", 3306)
	WriteString("mysql.db", "project")

	content := GetBytesAndClear()
	if strings.Trim(a, " ") != string(content) {
		t.Error("error")
	}
}
