package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaInfoComplementaria_20210105_095124 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaInfoComplementaria_20210105_095124{}
	m.Created = "20210105_095124"

	migration.Register("InsertTablaInfoComplementaria_20210105_095124", m)
}

// Run the migrations
func (m *InsertTablaInfoComplementaria_20210105_095124) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('AREA_CONOCIMIENTO', 'AREA_CONOCIMIENTO', TRUE, 20, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('FORMACION_ACADEMICA', 'FORMACION_ACADEMICA', TRUE, 20, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('INSTITUCION', 'INSTITUCION', TRUE, 20, LOCALTIMESTAMP, LOCALTIMESTAMP);")

}

// Reverse the migrations
func (m *InsertTablaInfoComplementaria_20210105_095124) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 20 ;")
}
