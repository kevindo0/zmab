package ab

import (
	"encoding/json"
	"fmt"
	"os"
)

// 将配置文件内容加载到内存中
func Load(fileName string) ABOnes {
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := ABOnes{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf)
	return conf
}
