package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaTipoContribuyente_20191118_083332 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTipoContribuyente_20191118_083332{}
	m.Created = "20191118_083332"

	migration.Register("InsertTablaTipoContribuyente_20191118_083332", m)
}

// Run the migrations
func (m *InsertTablaTipoContribuyente_20191118_083332) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_contribuyente (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('No contribuyenye', 'No contribuyenye', 'NCONT', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertTablaTipoContribuyente_20191118_083332) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_contribuyente;")
}
