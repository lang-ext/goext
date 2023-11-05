package log

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	t.Log("xxx")
	InitLog(LogCfg{ToFile: "./l.log", Level: "3"})
	defer Flush() // must be called
	L().Info("x", "xx", "d")
	L().Info("x", "xx", `
	1
	2
	3
	`)
	L().Error(fmt.Errorf("exx"), "asd")
	L().V(1).Info("1", "xx", "d")
	L().V(2).Info("2", "xx", "d")
	L().V(3).Info("3", "xx", "d")
	L().V(4).Info("4", "xx", "d")
	L().V(5).Info("5", "xx", "d")

}
