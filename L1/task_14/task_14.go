package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := complex(1, 1)

	println(switchTypeof(x))
	println(reflectTypeof(x))
	println(reflectKindTypeof(x))
	println(formatTypeof(x))
}

// не все типы можно распознать (например chan) и меньшк гибкости из-за switch
func switchTypeof(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

// определяем любой тип
func reflectTypeof(x interface{}) string {
	return reflect.TypeOf(x).String()
}

// Kind возвращает представление типа (что-то типа enum)
func reflectKindTypeof(x interface{}) string {
	return reflect.TypeOf(x).Kind().String()
}

func formatTypeof(x interface{}) string {
	return fmt.Sprintf("%T", x)
}
