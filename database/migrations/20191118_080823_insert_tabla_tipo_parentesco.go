package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaTipoParentesco_20191118_080823 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTipoParentesco_20191118_080823{}
	m.Created = "20191118_080823"

	migration.Register("InsertTablaTipoParentesco_20191118_080823", m)
}

// Run the migrations
func (m *InsertTablaTipoParentesco_20191118_080823) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_parentesco (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('MADRE', 'Madre', 'MD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_parentesco (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('PADRE', 'Padre', 'PD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_parentesco (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('HERMANO', 'Hermano', 'HMNO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertTablaTipoParentesco_20191118_080823) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_parentesco;")
}
