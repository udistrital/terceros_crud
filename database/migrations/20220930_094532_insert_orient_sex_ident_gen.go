package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertOrientSexIdentGen_20220930_094532 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertOrientSexIdentGen_20220930_094532{}
	m.Created = "20220930_094532"

	migration.Register("InsertOrientSexIdentGen_20220930_094532", m)
}

// Run the migrations
func (m *InsertOrientSexIdentGen_20220930_094532) Up() {
	file, err := ioutil.ReadFile("../scripts/20220930_094532_insert_orient_sex_ident_gen.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")
	requests = []string{}
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *InsertOrientSexIdentGen_20220930_094532) Down() {
	file, err := ioutil.ReadFile("../scripts/20220930_094532_insert_orient_sex_ident_gen.down.sql")

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
