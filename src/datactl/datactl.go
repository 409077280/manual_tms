package datactl

import (
	"log"
)

func CheckLost(long, short []string) []string {
	if len(long) < len(short) {
		long, short = short, long
	}
	lost := make([]string, 0, len(long) - len(short))
	for i := 0; i < len(long); i++ {
		check := true
		for j := 0; j < len(short); j++ {
			if long[i] == short[j] {
				check = false
				break
			}
		}
		//means lost[i] is alone
		if check == true {
			lost = append(lost, long[i])
		}
	}
	return lost
}

func Deletesame(list []string) []string {
	if len(list) == 0 {
		log.Fatal("The data is nil")
	}
	temp := make([]string, 0, len(list))
	for i := 0; i < len(list); i++ {
		check := true
		for j := i + 1; j < len(list); j++ {
			if list[i] == list[j] {
				check = false
				break
			}
		}
		if check == true {
			temp = append(temp, list[i])
		}
	}
	return temp
}