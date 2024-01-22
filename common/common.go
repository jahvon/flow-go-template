package common

import (
	"fmt"
	"reflect"
)

func StopStructureLogging() {
	fmt.Println("---break")
}

func StartStructureLogging() {
	fmt.Println("---endbreak")
}

// HasValue returns true if the object has a value, false if it is nil
func HasValue[T any](obj *T) (any, bool) {
	if obj == nil {
		return nil, true
	}

	return &obj, reflect.ValueOf(obj).IsZero()
}
