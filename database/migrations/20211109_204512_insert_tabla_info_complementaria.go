package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaInfoComplementaria_20211109_204512 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaInfoComplementaria_20211109_204512{}
	m.Created = "20211109_204512"

	migration.Register("InsertTablaInfoComplementaria_20211109_204512", m)
}

// Run the migrations
func (m *InsertTablaInfoComplementaria_20211109_204512) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20211109_204512_insert_tabla_info_complementaria_up.sql")
	
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
func (m *InsertTablaInfoComplementaria_20211109_204512) Down() {
	// use m.SQL("CREATE TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20211109_204512_insert_tabla_info_complementaria_down.sql")
	
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
