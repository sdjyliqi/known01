package handle

import (
	"fmt"
	"github.com/sdjyliqi/known01/conf"
	"github.com/sdjyliqi/known01/control"
	"sync"
)

var onceControl sync.Once


func init() {
	onceControl.Do(func() {
		fmt.Println("Init handle")
	})
}
