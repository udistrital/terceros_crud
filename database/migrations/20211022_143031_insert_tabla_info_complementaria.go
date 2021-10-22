package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaInfoComplementaria_20211022_143031 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaInfoComplementaria_20211022_143031{}
	m.Created = "20211022_143031"

	migration.Register("InsertTablaInfoComplementaria_20211022_143031", m)
}

// Run the migrations
func (m *InsertTablaInfoComplementaria_20211022_143031) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20211022_143031_insert_tabla_info_complementaria_up.sql")
	
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
func (m *InsertTablaInfoComplementaria_20211022_143031) Down() {
	// use m.SQL("CREATE TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20211022_143031_insert_tabla_info_complementaria_down.sql")
	
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
