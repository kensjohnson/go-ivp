package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("Entering program")

	type task struct {
		taskLineNo   int
		taskParent   string
		taskPosition string
		taskID       string
		taskTitle    string
	}

	taskList := make(map[string]*task)

	// open file
	f, err := os.Open("data/gtasks_today_dump.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
	lineCount := 0

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// increment line counter
		lineCount++

		// print the entire line
		//fmt.Printf("%+v\n", rec)

		//this would let me parse thru the csv but since fields 0 (datetime)  and 1 (loglevel) are throwaway wil skip but leave
		//for i, v := range rec {
		//	fmt.Printf("Index: %d   Value: %s\n", i, v)
		//}

		//test splitting the 3rd csv field(which is really multiple csv fields that gone munged on export into one quote dleimitted csv filed)
		// "task.parent, task.position, task.id, task.title"
		// s := strings.Split(rec[2], ",")
		// t := (strings.Split(rec[2], ",")[3])    //field 3 is the task description
		// fmt.Println(s)
		// fmt.Println((t))
		//TODO strip leading and trailing quotes from everything
		t := task{
			taskLineNo:   lineCount,
			taskParent:   (strings.Split(rec[2], ",")[0]),
			taskPosition: (strings.Split(rec[2], ",")[1]),
			taskID:       (strings.Split(rec[2], ",")[2]),
			taskTitle:    (strings.Split(rec[2], ",")[3]),
		}
		fmt.Println((t))
		taskList[t.taskID] = &t
	}
	// dump map
	// fmt.Println(taskList)

	//walk list and display titles just to work out dererencing

	for k, v := range taskList {
		fmt.Println(k, v.taskTitle)
	}

	fmt.Println("end of program")

}
