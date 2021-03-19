package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarTablaVinculacion_20210318_185036 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarTablaVinculacion_20210318_185036{}
	m.Created = "20210318_185036"

	migration.Register("ModificarTablaVinculacion_20210318_185036", m)
}

// Run the migrations
func (m *ModificarTablaVinculacion_20210318_185036) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210318_185036_modificar_tabla_vinculacion_up.sql")

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
func (m *ModificarTablaVinculacion_20210318_185036) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210318_185036_modificar_tabla_vinculacion_down.sql")

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
