package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaTipoParentesco_20191202_123000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTipoParentesco_20191202_123000{}
	m.Created = "20191202_123000"

	migration.Register("InsertTablaTipoParentesco_20191202_123000", m)
}

// Run the migrations
func (m *InsertTablaTipoParentesco_20191202_123000) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_parentesco (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('MADRE', 'MADRE.', 'MD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_parentesco (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('PADRE', 'PADRE.', 'PD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_parentesco (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('HERMANO', 'HERMANO.', 'HMNO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertTablaTipoParentesco_20191202_123000) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_parentesco;")
}
