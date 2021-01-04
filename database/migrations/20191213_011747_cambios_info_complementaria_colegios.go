package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambiosInfoComplementariaColegios_20191213_011747 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambiosInfoComplementariaColegios_20191213_011747{}
	m.Created = "20191213_011747"

	migration.Register("CambiosInfoComplementariaColegios_20191213_011747", m)
}

// Run the migrations
func (m *CambiosInfoComplementariaColegios_20191213_011747) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 13;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 13;")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (13, 'Calendario Colegios', 'Calendario Colegios.', 'Grupo_13', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CALENDARIO A', 'CALENDARIO A', TRUE, 13, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CALENDARIO A-B', 'CALENDARIO A-B', TRUE, 13, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CALENDARIO B', 'CALENDARIO B', TRUE, 13, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CALENDARIO B-OTRO', 'CALENDARIO B-OTRO', TRUE, 13, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('CALENDARIO OTRO', 'CALENDARIO OTRO', TRUE, 13, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('LOCALIDAD', 'LOCALIDAD', TRUE, 10, LOCALTIMESTAMP, LOCALTIMESTAMP);")

}

// Reverse the migrations
func (m *CambiosInfoComplementariaColegios_20191213_011747) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 13;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 13;")
}
