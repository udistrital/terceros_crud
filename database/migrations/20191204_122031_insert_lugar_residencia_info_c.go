package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertLugarResidenciaInfoC_20191204_122031 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertLugarResidenciaInfoC_20191204_122031{}
	m.Created = "20191204_122031"

	migration.Register("InsertLugarResidenciaInfoC_20191204_122031", m)
}

// Run the migrations
func (m *InsertLugarResidenciaInfoC_20191204_122031) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('LUGAR_RESIDENCIA', 'LUGAR_RESIDENCIA', TRUE, 10, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertLugarResidenciaInfoC_20191204_122031) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE codigo_abreviacion = 'LUGAR_RESIDENCIA';")
}
