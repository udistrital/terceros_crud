package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InfoComplementaria_20191128_152420 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InfoComplementaria_20191128_152420{}
	m.Created = "20191128_152420"

	migration.Register("InfoComplementaria_20191128_152420", m)
}

// Run the migrations
func (m *InfoComplementaria_20191128_152420) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('ESTRATO_RESPONSABLE', 'ESTRATO_RESPONSABLE', TRUE, 9, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InfoComplementaria_20191128_152420) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE codigo_abreviacion = 'ESTRATO_RESPONSABLE' ;")
}
