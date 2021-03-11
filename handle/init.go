package handle

import (
	"known01/brain"
	"sync"
)

var onceControl sync.Once
var baCenter brain.Center

func init() {
	onceControl.Do(func() {
		baCenter = brain.CreateCenter()
	})
}
