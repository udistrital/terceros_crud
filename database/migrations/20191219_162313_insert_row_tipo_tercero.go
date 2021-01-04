package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertRowTipoTercero_20191219_162313 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertRowTipoTercero_20191219_162313{}
	m.Created = "20191219_162313"

	migration.Register("InsertRowTipoTercero_20191219_162313", m)
}

// Run the migrations
func (m *InsertRowTipoTercero_20191219_162313) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (14, 'COLEGIO_VERIFICAR', 'Registro de colegio cuya información está pendiente de verificación.', 'COLEGIO_VERIFICAR', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertRowTipoTercero_20191219_162313) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_tercero WHERE id = 14;")
}
