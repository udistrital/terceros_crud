package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertRowTipoTerceroUniversidades_20191219_163324 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertRowTipoTerceroUniversidades_20191219_163324{}
	m.Created = "20191219_163324"

	migration.Register("InsertRowTipoTerceroUniversidades_20191219_163324", m)
}

// Run the migrations
func (m *InsertRowTipoTerceroUniversidades_20191219_163324) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (15, 'UNIVERSIDAD_OFICIAL', 'UNIVERSIDAD_OFICIAL.', 'UNIVERSIDAD_OFICIAL', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (16, 'UNIVERSIDAD_PRIVADA', 'UNIVERSIDAD_PRIVADA.', 'UNIVERSIDAD_PRIVADA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertRowTipoTerceroUniversidades_20191219_163324) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_tercero WHERE id BETWEEN 15 AND 16;")
}
