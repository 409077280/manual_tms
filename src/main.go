package main

import (
	"rwfile"
	"loadcfg"
	"rwfile"
	"tmsctl"
	"datactl"
)
type Manual interface {
	ReadData() []string
	WrtieLost() []string
	WritePass() []string
	WriteNoPass() []string
}
type  Do struct {
	GetData []string
	LostData []string
	PassData []string
	NoPassData []string
}

func (d *Do)ReadData(filename string)  {
	opfile := new(rwfile.FileControl)
	opfile.Readname = filename
	d.GetData = datactl.Deletesame(opfile.ReadFileToMap())
}

func (d *Do)WrtieLost()  {
	datactl.CheckLost(d.GetData,)
}

func (d Do)WritePass()  {

}

func (d Do)WriteNoPass()  {

}

func main() {
	var manual Manual
	conf := new(loadcfg.IniConfig)
	conf.ReadConfig("./config.ini")


	do := new(Do)
	do.ReadData("./1.txt")

}