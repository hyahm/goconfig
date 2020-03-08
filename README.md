# goconfig
read go config from key value file file

### 安装
```
go get github.com/hyahm/goconfig
```
### 注意
- key不能有.
- 支持读取任何单行配置文件
- 自定义符号
- 支持写入配置文件

# 使用

### 自定义符号, 加载配置文件之前定义
```ini
SEP         = "="  // key 和 value 分隔符
NOTE        = "#;" // #开头的为注释
MODEL_START = "["  // 模块开头符号
MODEL_END   = "]"  // 模块结尾符号
WELL        = "#"  // 写入的注释， 
```


### 读取配置
> 先初始化 test.conf
```vim
[mysql]
host=192.168.80.2
port=name
```
```
第二个参数表示， 如果不存在key， 或者读取的值不符合读取类型的默认值
goconfig.InitConf("test.conf")
goconfig.ReadString("mysql.host", "127.0.0.1") // return 192.168.80.2
goconfig.ReadInt("mysql.port", 3306)    // return 3306
goconfig.ReadString("mysql.db", "name")    // return name
```


> 从[]byte 读取
```vim
```go
x := `[mysql]
host=192.168.80.2
port=name
`
当然也可以是传过来的数据进行解析
goconfig.InitFromBytes([]byte(x))
goconfig.ReadString("mysql.host", "127.0.0.1") // return 192.168.80.2
goconfig.ReadInt("mysql.port", 3306)    // return 3306
goconfig.ReadString("mysql.db", "name")    // return name
```
### 写入配置文件
```
goconfig.InitWriteConf("write.conf")  // 与InitConf的区别是， 这个会清空里面原有数据
goconfig.InitConf("test.conf")  // 原有配置文件添加
goconfig.WriteString("mysql.host", "127.0.0.1","mysql数据库host") 
goconfig.WriteInt("mysql.port", 3306)   
goconfig.WriteString("mysql.db", "name")   
goconfig.FlushWrite()  // 缓存一次写入文件， 写入别忘了这行
```
> test.conf
```
[mysql]
# mysql数据库host
host = 127.0.0.1
port = 3306
db = name
```

### 配置文件软加载方法
err = Reload(),  仅限InitConf(file)  从文件读取的软加载， 自己写接口调用此方法， 配置文件会刷新， 如果配置文件有错则不会更新


# 辅助方法， 方便调试
PrintLines() // 打印读取的配置文件  
PrintKeyValue()  // 打印kv数据  

