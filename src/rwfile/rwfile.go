package rwfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)
type FileControl struct {
	Readname string
	Writename string
}

func (f FileControl)ReadFileToList() []string {
	file, err := os.Open(f.Readname)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	var list []string
	for i := 0; ; {
		tmp, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(tmp))
		list = append(list, s)
		i++
	}
	fmt.Println(f.Readname, "lines", len(list))
	return list
}

func (f FileControl)WriteListToFile(list []string) {
	name := fmt.Sprintf("%s-%v",f.Writename, time.Now().Format("2006-01-02")) + ".csv"
	wfile, error := os.Create(name)
	defer wfile.Close()
	if error != nil {
		fmt.Println(error)
	}
	for i := 0; i < len(list); i++ {
		_, err := wfile.Write([]byte(list[i] + "\n"))
		if err != nil {
			os.Exit(1)
		}
	}
}