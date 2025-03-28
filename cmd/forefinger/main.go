package main

import (
	"fmt"
	"github.com/s4bb4t/forefinger/internal/models"
	"reflect"
	"unsafe"
)

func main() {
	t := *models.NewBlock()

	fmt.Println(reflect.TypeOf(t).Size())
	fmt.Println(unsafe.Sizeof(t))

	b := *models.NewLightWeightBlock()

	fmt.Println(reflect.TypeOf(b).Size())
	fmt.Println(unsafe.Sizeof(b))
}
