package utils

import (
	"testing"
	"time"
)

func Test_TimeFormat(t *testing.T) {
	curtime := time.Now()
	t.Log(curtime.Format(FullTime))
}
