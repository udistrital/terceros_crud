package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertGrupoInfoComplementariaDerechosPecunarios_20210912_173323 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertGrupoInfoComplementariaDerechosPecunarios_20210912_173323{}
	m.Created = "20210912_173323"

	migration.Register("InsertGrupoInfoComplementariaDerechosPecunarios_20210912_173323", m)
}

// Run the migrations
func (m *InsertGrupoInfoComplementariaDerechosPecunarios_20210912_173323) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210912_173323_insert_grupo_info_complementaria_derechos_pecunarios_up.sql")
	
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
func (m *InsertGrupoInfoComplementariaDerechosPecunarios_20210912_173323) Down() {
	// use m.SQL("CREATE TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210912_173323_insert_grupo_info_complementaria_derechos_pecunarios_down.sql")
	
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
