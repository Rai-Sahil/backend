package utils

import (
	"fmt"
	"runtime"
)

func ErrorWithFileInfo(err error) string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return err.Error()
	}
	return fmt.Sprintf("[%s:%d] %s", file, line, err.Error())
}
