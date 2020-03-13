package goconfig

import "fmt"

func PrintLines() {
	for _, v := range Config.Lines {
		// 打印注释
		for _, n := range v.note {
			fmt.Println("# ", string(n))
		}
		// 打印kv
		fmt.Println(v.key, ":", string(v.value))
	}
	for _, v := range Config.Groups {
		// 打印组注释
		for _, n := range v.note {
			fmt.Println(string(n))
		}
		// 打印组
		fmt.Println("[", string(v.name), "]")
		for _, gn := range v.group {
			// 组key 注释
			for _, nn := range gn.note {
				fmt.Println("# ", string(nn))
			}
			// 打印kv
			fmt.Println(gn.key, ":", string(gn.value))
		}
	}
}

func PrintKeyValue() {
	for k, v := range Config.KeyValue {
		fmt.Println(k, ":", string(v))
	}
}
