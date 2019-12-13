package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambioTipoTercero_20191213_011554 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambioTipoTercero_20191213_011554{}
	m.Created = "20191213_011554"

	migration.Register("CambioTipoTercero_20191213_011554", m)
}

// Run the migrations
func (m *CambioTipoTercero_20191213_011554) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("UPDATE terceros.tipo_tercero SET nombre = 'COLEGIO_OFICIAL', codigo_abreviacion = 'COLEGIO_OFICIAL' WHERE codigo_abreviacion = 'COLEGIO';")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (12, 'COLEGIO_NO_OFICIAL', null, 'COLEGIO_NO_OFICIAL', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *CambioTipoTercero_20191213_011554) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_tercero WHERE id = 12;")
}
