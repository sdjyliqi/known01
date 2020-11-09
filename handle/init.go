package handle

import (
	"fmt"
	"known01/brain"
	"sync"
)

var onceControl sync.Once
var baCenter brain.Center

func init() {
	onceControl.Do(func() {
		fmt.Println("Init handle")
		baCenter = brain.CreateCenter()
	})
}
