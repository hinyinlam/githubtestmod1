package package1

import (
	"fmt"
	"time"
)

func PublicFunction() {
	fmt.Println("This is public func")
}

var ExportMe = time.Now()
