package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type DeleteRowsTipoTercero_20191219_161441 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DeleteRowsTipoTercero_20191219_161441{}
	m.Created = "20191219_161441"

	migration.Register("DeleteRowsTipoTercero_20191219_161441", m)
}

// Run the migrations
func (m *DeleteRowsTipoTercero_20191219_161441) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *DeleteRowsTipoTercero_20191219_161441) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
