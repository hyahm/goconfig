# goconfig
read go config from key value file file

### 安装
```
go get github.com/hyahm/goconfig
```
### 注意
- 一行是一个key ，value
- 键值 用等号分割
- 隐形支持json， 存入json []byte 
- 注释使用 #开头
- 建议使用Write* 方法写入， 格式更整齐
- 已经存在的键值不会重复写入

# 使用
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

# 辅助方法， 方便调试
PrintLines() // 打印读取的配置文件
PrintKeyValue()  // 打印kv数据

