package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type ModifyTableInfoComplementariaTercero_20210118_112024 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModifyTableInfoComplementariaTercero_20210118_112024{}
	m.Created = "20210118_112024"

	migration.Register("ModifyTableInfoComplementariaTercero_20210118_112024", m)
}

// Run the migrations
func (m *ModifyTableInfoComplementariaTercero_20210118_112024) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210105_152820_agregar_campos_info_complementaria_tercero_padre.up.sql")

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
func (m *ModifyTableInfoComplementariaTercero_20210118_112024) Down() {
	file, err := ioutil.ReadFile("../scripts/20210105_152820_agregar_campos_info_complementaria_tercero_padre.down.sql")

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
	// use m.SQL("DROP TABLE ...") to reverse schema update


}
