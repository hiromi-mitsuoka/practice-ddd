package main

import (
	"fmt"

	models "github.com/hiromi-mitsuoka/practice-ddd/models/values"
)

func main() {
	nameA, _ := models.NewFullName("hoge", "huga")
	nameB, _ := models.NewFullName("hogehoge", "hugahuga")
	nameC, _ := models.NewFullName("hogehoge", "hugahuga")

	fmt.Printf("reflect.DeepEqual(nameA, nameB) : 等価 : %t\n", nameA.Equals(nameB))
	fmt.Printf("reflect.DeepEqual(nameB, nameC) : 等価 : %t\n", nameB.Equals(nameC))
}
