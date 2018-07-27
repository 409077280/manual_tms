package main

import (
	"loadcfg"
	"rwfile"
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
//	datactl.CheckLost(d.GetData,)
}

func (d Do)WritePass()  {

}

func (d Do)WriteNoPass()  {

}


func main() {
	base := rwfile.FileControl{"base.txt","Checklost"}
	bs := base.ReadFileToMap()
	basenew := rwfile.FileControl{"basenew.txt","No same"}
	bn := basenew.ReadFileToMap()
	base.WriteListToFile(datactl.CheckLost(bs,bn))
	return
	rw := rwfile.FileControl{"in.txt","No same"}
	list := datactl.Deletesame(rw.ReadFileToMap())
	rw.WriteListToFile(list)
	return
	/*
	config := loadcfg.IniConfig{}
	config.ReadConfig("config.ini")

	pg := &tmsctl.Pgconfig{User:config.Data["username"],
	Pass:config.Data["password"],
	Host:config.Data["hostname"],
	Port:config.Data["port"],
	Dbname:config.Data["database"],
	Sslmodel:config.Data["sslmodel"],
	}
	tmsList := []tmsctl.DataModel{}
	pg.GetPackageIdList(tmsList, list)

	fmt.Println(tmsList)
*/
/*	var manual Manual
	conf := new(loadcfg.IniConfig)
	conf.ReadConfig("./config.ini")


	do := new(Do)
	do.ReadData("./1.txt")
*/
}