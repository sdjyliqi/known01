package handle

import (
	"github.com/sdjyliqi/known01/brain"
	"sync"
)

var onceControl sync.Once
var baCenter brain.Center

func InitBrain() {
	onceControl.Do(func() {
		baCenter = brain.CreateCenter()
	})
}
