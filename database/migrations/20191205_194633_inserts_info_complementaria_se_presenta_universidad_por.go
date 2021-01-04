package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsInfoComplementariaSePresentaUniversidadPor_20191205_194633 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsInfoComplementariaSePresentaUniversidadPor_20191205_194633{}
	m.Created = "20191205_194633"

	migration.Register("InsertsInfoComplementariaSePresentaUniversidadPor_20191205_194633", m)
}

// Run the migrations
func (m *InsertsInfoComplementariaSePresentaUniversidadPor_20191205_194633) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (15, 'Se presenta a la universidad por', 'Se presenta a la universidad por', 'Grupo_15', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Primera vez', 'Primera vez', TRUE, 15, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Segunda vez', 'Segunda vez', TRUE, 15, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Tercera o más veces', 'Tercera o más veces', TRUE, 15, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertsInfoComplementariaSePresentaUniversidadPor_20191205_194633) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 15;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 15;")
}
