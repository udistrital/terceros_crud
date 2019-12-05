package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsInfoComplementariaLocalidadesBogota_20191204_204509 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsInfoComplementariaLocalidadesBogota_20191204_204509{}
	m.Created = "20191204_204509"

	migration.Register("InsertsInfoComplementariaLocalidadesBogota_20191204_204509", m)
}

// Run the migrations
func (m *InsertsInfoComplementariaLocalidadesBogota_20191204_204509) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (12, 'Medio por el que se entero de la Univsersidad', 'Medio por el cual se entero de la Universidad Distrital', 'Grupo_12', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Amigo', 'Amigo', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Afiche', 'Afiche', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Prensa', 'Prensa', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Feria Social', 'Feria Social', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Internet', 'Internet', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Radio', 'Radio', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Visita al Colegio', 'Visita al Colegio', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Televisión', 'Televisión', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Otro Medio', 'Otro Medio', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertsInfoComplementariaLocalidadesBogota_20191204_204509) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 12;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 12;")
}
