package utils

import (
	"fmt"
	"strings"

	"github.com/lainio/err2/try"
)

func ParseLoggingArgs(s string) (use bool, lvl int) {
	if s == "" {
		return false, 0
	}

	use = strings.Contains(s, "logtostderr")
	if use {
		s = strings.Replace(s, "-v ", "-v=", -1)
		ss := strings.Split(s, "-v=")
		if len(ss) > 1 {
			try.Out1(fmt.Sscanf("-v="+ss[1], "-v=%d", &lvl)).Catch(0)
		}
	}
	return
}
