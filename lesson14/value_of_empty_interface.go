package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = nil
	v := reflect.ValueOf(x)
	fmt.Printf("value %v of type %T\n", x, x)
	fmt.Printf("value %v of type %T\n", v, v)
}
