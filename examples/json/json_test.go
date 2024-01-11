package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Order struct { //tag
	Id       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Price    int    `json:"price"`
}

func marshal() {
	o := Order{
		Id:       "1234",
		Quantity: 10,
		Price:    199,
	}
	fmt.Printf("%+v\n", o)
	of := reflect.TypeOf(o)
	fmt.Println(of.NumField())
	for i := 0; i < of.NumField(); i++ {
		fmt.Println(of.Field(i))
	}
	fmt.Println(of.NumMethod())
	for i := 0; i < of.NumMethod(); i++ {
		fmt.Println(of.Method(i))
	}
	fmt.Println(of)
	m, err := json.Marshal(o)
	if err != nil {
		panic("marshal error")
	}
	fmt.Printf("%s\n", m)
}

type ColorGroup struct {
	Id     int      `json:"id"`
	Name   string   `json:"name,omitempty"`
	Colors []string `json:"colors"`
}

func unMarshal() {
	str := "{\"ID\":1,\"Name\":\"Reds\",\"Colors\":[\"Crimson\",\"Red\",\"Ruby\",\"Maroon\"]}"
	var group ColorGroup
	json.Unmarshal([]byte(str), &group)
	fmt.Printf("%+v \n", group)
}

func TestJson(t *testing.T) {
	unMarshal()
}

func TestParseNlp(t *testing.T) {
	res := `{
"data": [
    {
        "synonym":"",
        "weight":"0.6",
        "word": "真丝",
        "tag":"材质"
    },
    {
        "synonym":"",
        "weight":"0.8",
        "word": "韩都衣舍",
        "tag":"品牌"
    },
    {
        "synonym":"连身裙;联衣裙",
        "weight":"1.0",
        "word": "连衣裙",
        "tag":"品类"
    }
]
}`
	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Weight  string `json:"weight"`
			Word    string `json:"word"`
			Tag     string `json:"tag"`
		} `json:"data"`
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", m.Data[2])
}
