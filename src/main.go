package main

import (
	"fmt"
	"io/ioutil"

	// "gopkg.in/yaml.v2"
	m_pack "core"
	// _ "core"
)

func main() {

	conf, err := ioutil.ReadFile("./conf.json")

	if err != nil {

	}
	dd := m_pack.NewPackCore(string(conf), m_pack.CONF_FORMAT_JSON)

	// dd.Pack(nil)
	// datas := dd.Pack([]byte("test"))

	datas := []byte{
		0,
		0,
		118,
		49,
		0,
		0,
		0,
		12,
		116,
		101,
		115,
		116,
	}

	if dd != nil {
		dd1, err1 := dd.Unpack(datas)
		if err1 != nil {
			fmt.Println(dd1)
		}
	}

}
