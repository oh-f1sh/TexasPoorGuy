package conf

import (
	"io"
	"os"

	"github.com/oh-f1sh/TexasPoorGuy/client"
	"gopkg.in/yaml.v3"
)

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")

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
	client.DEFAULTUSERNAME = m["username"]
	client.DEFAULTPWD = m["password"]
	client.SERVER_ADDR = m["server"]
	client.PORT = m["port"]
}
