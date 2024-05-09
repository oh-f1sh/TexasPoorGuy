package conf

import (
	"io"
	"os"

	"github.com/oh-f1sh/TexasPoorGuy/common"
	"gopkg.in/yaml.v3"
)

func init() {
	f, err := os.Open("config.yaml")
	if err != nil {
		return
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return
	}
	m := make(map[string]string)
	err = yaml.Unmarshal(b, m)
	if err != nil {
		return
	}
	common.DEFAULTUSERNAME = m["username"]
	common.DEFAULTPWD = m["password"]
	common.SERVER_ADDR = m["server"]
	common.PORT = m["port"]

	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
}
