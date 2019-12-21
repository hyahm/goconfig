配置文件包含2部分
注释优先

普通kv（类型map）
- 注释
- kv

type line struct {
    key string
    value []byte
    note []byte
}

组kv （类型： 数组）
- 组注释
- 组名
- kv注释
- kv

type groupLine struct {
    group map[string]*line  // 组的行
    note [][]byte  // 组注释
    name []byte   // 组名
}

type config struct {
    Groups map[string]*groupLine  // 组
    Lines map[string]*line    // 单key
    KeyValue map[string][]byte   // 键值缓存， key的值  key or group.key
}
