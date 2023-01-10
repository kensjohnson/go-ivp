package main

import "fmt"

func main() {
	//a table of fake defects so I can play with branches etc
	var fakeDefects [5]string
	fakeDefects[0] = "i need to be fixed"
	fakeDefects[1] = "i need to be fixed"
	fakeDefects[2] = "i need to be fixed"
	fakeDefects[3] = "i need to be fixed"
	fakeDefects[4] = "i need to be fixed"
	fmt.Println("Entering program")

	//dump the defects
	fmt.Println("defect table")
	for index, element := range fakeDefects {
		fmt.Println("Defect:", index, "=>", element)
	}
	fmt.Println("<<<< ======== >>>>")

	fmt.Println("end of program")

}
