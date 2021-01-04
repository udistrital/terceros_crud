package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertGrupoExperienciaLaboralGrupo19_20201227_142847 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertGrupoExperienciaLaboralGrupo19_20201227_142847{}
	m.Created = "20201227_142847"

	migration.Register("InsertGrupoExperienciaLaboralGrupo19_20201227_142847", m)
}

// Run the migrations
func (m *InsertGrupoExperienciaLaboralGrupo19_20201227_142847) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

	file, err := ioutil.ReadFile("../scripts/20201227_142847_insert_grupo_experiencia_laboral_grupo_19.up.sql")
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
func (m *InsertGrupoExperienciaLaboralGrupo19_20201227_142847) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201227_142847_insert_grupo_experiencia_laboral_grupo_19.down.sql")
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
