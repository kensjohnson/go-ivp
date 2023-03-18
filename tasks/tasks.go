package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

		//skip the csv header line (line 1)
		//TODO: should probably do a bit of sanity checking to make sure there is a header row
		if lineCount == 1 {
			fmt.Println("skipping header row==>", rec)
			continue
		}

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

	// print counts for comparison - since 1st line is headers they are off by 1
	fmt.Printf("CSV Lines read: %d  taskList Length: %d \n ", lineCount, len(taskList))

	// dump map
	// fmt.Println(taskList)

	//sortList is used to sort the keys for a subset of tasks to order them  by taskPosition
	sortlist := make([]task, len(taskList))

	//walk taskList and identify top level tasks(parent = null). add taskPositon and taskID to map and taskPosition to slice
	//for subsequent sorting. this assumes that the positions are unique (ie 2 tasks cant have the same position in the list)
	//note only need values since the value also contains the taskID
	for _, v := range taskList {
		if v.taskParent == "null" {
			//fmt.Println(k, v.taskParent, v.taskTitle)
			//manually deref files from task pointer in master list to a new task for the slice
			//TODO: decide whether taskList should be a pointer or a instance since this is gettin crazy
			t := task{
				taskLineNo:   v.taskLineNo,
				taskParent:   v.taskParent,
				taskPosition: v.taskPosition,
				taskID:       v.taskID,
				taskTitle:    v.taskTitle,
			}
			sortlist = append(sortlist, t)

		}
	}

	//sort the slice by taskposition
	sort.SliceStable(sortlist, func(i, j int) bool {
		return sortlist[i].taskPosition < sortlist[j].taskPosition
	})

	fmt.Println("Sorted slice")
	for _, u := range sortlist {
		fmt.Printf("SEQ: %s Title: %s \n", u.taskPosition, u.taskTitle)
	}

	//sort the slice to get in order by taskPosition

	fmt.Println("end of program")

}
