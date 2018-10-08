package logrecord

import (
	"log"
	"os"
)

func main(){
	fileName:= "log_debug.log"
	//logFile,err := os.Create(fileName)
	logFile,err := os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error!")
	}
	debugLog := log.New(logFile,"[Debug]",log.Llongfile)
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A debug message here")
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here")
	//debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")

}