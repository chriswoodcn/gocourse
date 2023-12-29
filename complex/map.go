package main

import (
	"fmt"
	"reflect"
)

func main() {
	mapInstance := map[string]string{
		"language": "golang",
		"arch":     "x86_64",
		"os":       "windows",
	}
	fmt.Println(mapInstance)
	makeMap := make(map[string]interface{}, 2)
	fmt.Println(makeMap)
	makeMap["aaa"] = 10
	makeMap["bbb"] = "ig390gs0oab0alb"
	makeMap["ccc"] = 15.33
	fmt.Println(makeMap)
	for k, v := range makeMap {
		fmt.Printf("key: %s value-type: %T value: %v \n", k, v, v)
		fmt.Println(reflect.TypeOf(v).Name())
	}
}
