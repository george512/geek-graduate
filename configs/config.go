package configs

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
)

// 程序配置
type Config struct {
	HttpAddress string `json:"http_address"`
	GrpcAddress string `json:"grpc_address"`
	EndPoint    string `json:"end_point"`
}

// 配置文件名
const (
	CONFIG_FILENAME = "config.json"
)

type Conf struct {
	*Config
}

// 加载配置
func (c *Conf) Load() (err error) {
	var (
		content []byte
		conf    Config
	)

	p := getCurrentAbPathByCaller()
	p = filepath.Join(p, CONFIG_FILENAME)

	// 1. 把配置文件读进来
	if content, err = ioutil.ReadFile(p); err != nil {
		return
	}

	// 2. json做反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3. 复制单例
	c.Config = &conf
	return
}

func NewConf() *Conf {
	return &Conf{}
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
