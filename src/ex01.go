package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) {
	val := reflect.ValueOf(plant)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag
		key := field.Name
		value := val.Field(i)

		// Проверяем наличие тегов и изменяем ключ, если необходимо
		if tagValue, ok := tag.Lookup("color_scheme"); ok {
			key += fmt.Sprintf("(%s)", tagValue)
		} else if tagValue, ok := tag.Lookup("unit"); ok {
			key += fmt.Sprintf("(%s)", tagValue)
		}

		fmt.Printf("%s:%v\n", key, value)
	}
}

func main() {
	plant1 := UnknownPlant{
		FlowerType: "Orchid",
		LeafType:   "Broad",
		Color:      123456,
	}

	plant2 := AnotherUnknownPlant{
		FlowerColor: 654321,
		LeafType:    "Needle",
		Height:      30,
	}

	describePlant(plant1)
	describePlant(plant2)
}
