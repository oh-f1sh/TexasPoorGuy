package conf

import (
	"os"
)

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
}
