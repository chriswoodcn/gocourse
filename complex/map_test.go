package complex

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	mapInstance := map[string]string{
		"language": "golang",
		"arch":     "x86_64",
		"os":       "windows",
	}
	fmt.Println(mapInstance)
	// make space
	makeMap := make(map[string]interface{}, 2)
	fmt.Println(makeMap)
	makeMap["aaa"] = 10
	makeMap["bbb"] = "ig390gs0oab0alb"
	makeMap["ccc"] = 15.33
	fmt.Println(makeMap)
	// iterate map
	for k, v := range makeMap {
		fmt.Printf("key: %s value-type: %T value: %v \n", k, v, v)
		fmt.Println(reflect.TypeOf(v).Name())
	}
}
