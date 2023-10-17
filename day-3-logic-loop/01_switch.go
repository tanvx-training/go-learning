package main

import (
	"fmt"
	"log"
	"reflect"
)

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	ts := []interface{}{"A", true, 1, 1.5, 'a'}
	for _, t := range ts {
		eval(t)
	}
}

func eval(t interface{}) {
	switch typedValue := t.(type) {
	default:
		log.Fatalf("%v is %v", typedValue, reflect.TypeOf(typedValue))
	case bool:
		fmt.Println("bool:", typedValue, "is", reflect.TypeOf(typedValue))
	case int:
		fmt.Println("int:", typedValue, "is", reflect.TypeOf(typedValue))
	case float64:
		fmt.Println("float64:", typedValue, "is", reflect.TypeOf(typedValue))
	case string:
		fmt.Println("string:", typedValue, "is", reflect.TypeOf(typedValue))
	}
}
