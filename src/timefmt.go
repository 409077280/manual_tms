package ttime

import (
	"fmt"
	"time"
	"datactl"
)
/*
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700"
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
*/
func test() {
	a := []int{1,2,3,4}
	b := []int{2,3,4}
	if a != b {
		fmt.Println("err")
	}
	a1 := time.Now()
	fmt.Println(a1)
	a2 := a1.Format("2006-01-02T15:04:05+00:00")
	fmt.Println(a2)
	/*
	args := os.Args
	if len(args) != 3 {
		os.Exit(1)
	}
	file1 := args[1]
	file2 := args[2]
	var lostList []string
	arr1 := deletesame(readFileToMap(file1))
	arr2 := deletesame(readFileToMap(file2))
	fmt.Println("distinc:", file1, len(arr1))
	fmt.Println("distinc:", file2, len(arr2))
	lostList = checkLost(arr1, arr2)
	WriteListToFile(lostList)
	fmt.Println("Excute success!")
	*/

}
