package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaTipoContribuyente_20191125_142715 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTipoContribuyente_20191125_142715{}
	m.Created = "20191125_142715"

	migration.Register("InsertTablaTipoContribuyente_20191125_142715", m)
}

// Run the migrations
func (m *InsertTablaTipoContribuyente_20191125_142715) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_contribuyente (id, nombre, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1, 'PERSONA NATURAL', 'P_NATURAL', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_contribuyente (id, nombre, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2, 'PERSONA JURIDICA', 'P_JURIDICA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
}

// Reverse the migrations
func (m *InsertTablaTipoContribuyente_20191125_142715) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_contribuyente WHERE id BETWEEN 1 AND 2;")
}
