package main

import (
	"ARRAYSTRUCT/view"
	"fmt"
)

func main() {
	view.Insert()
	view.Insert()
	view.Views()
	fmt.Println()

	//// test update
	// view.Update()
	// view.Views()
	// fmt.Println()

	// test delete
	view.Delete()
	view.Views()

}
