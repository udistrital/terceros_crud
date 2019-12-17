package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertGrupoBasicaSecundaria_20191217_141600 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertGrupoBasicaSecundaria_20191217_141600{}
	m.Created = "20191217_141600"

	migration.Register("InsertGrupoBasicaSecundaria_20191217_141600", m)
}

// Run the migrations
func (m *InsertGrupoBasicaSecundaria_20191217_141600) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (17, 'Basica Secundaria', 'Grupo informacion de la basica secundaria del tercero.', 'Grupo_17', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('INFORMACIO COLEGIO', 'INFORMACION DEL COLEGIO DEL TERCERO', TRUE, 17, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertGrupoBasicaSecundaria_20191217_141600) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 17;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 17;")
}
