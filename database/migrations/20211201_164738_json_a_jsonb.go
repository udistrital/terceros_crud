package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type JsonAJsonb_20211201_164738 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &JsonAJsonb_20211201_164738{}
	m.Created = "20211201_164738"

	migration.Register("JsonAJsonb_20211201_164738", m)
}

// Run the migrations
func (m *JsonAJsonb_20211201_164738) Up() {
	file, err := ioutil.ReadFile("../scripts/20211201_164738_json_a_jsonb.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *JsonAJsonb_20211201_164738) Down() {
	file, err := ioutil.ReadFile("../scripts/20211201_164738_json_a_jsonb.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
