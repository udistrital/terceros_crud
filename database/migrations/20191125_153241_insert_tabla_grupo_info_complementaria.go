package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaGrupoInfoComplementaria_20191125_153241 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaGrupoInfoComplementaria_20191125_153241{}
	m.Created = "20191125_153241"

	migration.Register("InsertTablaGrupoInfoComplementaria_20191125_153241", m)
}

// Run the migrations
func (m *InsertTablaGrupoInfoComplementaria_20191125_153241) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (6, 'Genero', 'Generos.', 'Grupo_6', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (7, 'Grupo sanguineo', 'Clasificacion de los grupos sanguineos establecida por Karl Landsteiner.', 'Grupo_7', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (8, 'Factor RH', 'El factor RH es una proteína de los glóbulos rojos y permite detectar el tipo de sangre sea + o -.', 'Grupo_8', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (9, 'Información Socioeconómica', 'Información Socioeconómica de una persona natural.', 'Grupo_9', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (10, 'Información Contacto', 'Información de Contacto de una persona natural o juridica.', 'Grupo_10', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertTablaGrupoInfoComplementaria_20191125_153241) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id BETWEEN 6 AND 10;")
}
