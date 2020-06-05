package main

import (
	"io/ioutil"

	// "gopkg.in/yaml.v2"
	m_pack "core"
	// _ "core"
)

func main() {
	// file, err := ioutil.ReadFile("./conf.yml")

	// if err != nil {

	// }

	// var c interface{}

	// err = yaml.Unmarshal(file, &c)

	conf, err := ioutil.ReadFile("./conf.json")

	if err != nil {

	}
	dd := m_pack.NewPackCore(string(conf), m_pack.CONF_FORMAT_JSON)

	dd.Pack(nil)

	if dd != nil {

	}

}
