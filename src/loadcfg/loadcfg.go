package loadcfg

import(
	"os"
	"io"
	"bufio"
	"strings"
	"fmt"
)

var Config = map[string]string{}

const defaultpath = "./config.ini"

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func init() {
	b, e := pathExists(defaultpath)
	if b != true || e != nil {
		panic("Please check your config file : './config.ini'")
	}
	file, err := os.Open(defaultpath)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	r := bufio.NewReader(file)

	var strcet string
	for {
		tmp, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//trim the line null byte
		s := strings.TrimSpace(string(tmp))
		//去掉注释行，#出现的位置为0
		if strings.Index(s, "#") == 0 {
			continue
		}
		//去掉注释
		if strings.Index(s, "//") == 0 {
			continue
		}

		tag1index := strings.Index(s, "[")
		//最后 ] 出现的位置，读取配置名
		tag2index := strings.LastIndex(s, "]")
		//优化为 tag1index > -1 && tag2index > tag1index+1
		//if tag1index > -1 && tag2index > -1 && tag2index > tag1index+1 {
		//如果出现[]且符合结构，返回索引名称
		if tag1index > -1 && tag2index > tag1index+1 {
			strcet = strings.TrimSpace(s[tag1index+1 : tag2index])
			continue  //??
		}
		//如果没有索引，视为无效配置
		if len(strcet) == 0 {
			continue
		}

		index := strings.Index(s, "=")

		// = 位置最少应该在第2位
		//if index < 0 {
		if index < 1 {
			continue
		}
		//获得 配置的key
		name := strings.TrimSpace(s[:index])
		// = 前的长度为0
		if len(name) == 0 {
			continue
		}
		// = 后的值
		value := strings.TrimSpace(s[index+1:])
		// = 到 # 注释之间的值
		pos := strings.Index(value, "#")
		if pos > -1 {

			value = strings.TrimSpace(value[0:pos])
		}
		pos = strings.Index(value, "//")
		if pos > -1 {
			value = strings.TrimSpace(value[0:pos])
		}

		Config[name] = value
	}
}


func List()  {
	if len(Config) == 0 {
		panic("Config not set.")
	}
	keytb := 30
	valtb := 40
	var space string
	for i := 0; i < keytb + valtb; i++ {
		space += "-"
	}
	fmt.Printf("%s\n",space)
	for k, v := range Config {
		k = "  " + k
		v = "  " + v
		spacelen := keytb - len(k) -2
		for i := 0; i < spacelen; i++ {
			k = k + " "
		}
		spacelen = valtb - len(v) -1
		for i := 0; i < spacelen; i++ {
			v = v + " "
		}
		fmt.Printf("|%s|%s|\n", k, v)
		fmt.Printf("%s\n",space)
	}
}