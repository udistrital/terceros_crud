package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaGrupoInfoComplementaria_20191107_094657 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaGrupoInfoComplementaria_20191107_094657{}
	m.Created = "20191107_094657"

	migration.Register("InsertTablaGrupoInfoComplementaria_20191107_094657", m)
}

// Run the migrations
func (m *InsertTablaGrupoInfoComplementaria_20191107_094657) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1, 'Tipo Discapacidad', 'Tipos de Discapacidades que puede tener un tercero / persona natural.', 'Grupo_1', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2, 'Estado Civil', 'Tipos de Estado Civil que puede tener un tercero / persona_natural.', 'Grupo_2', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (3, 'Tipo Etnia', 'Tipos de Etnias a la que puede pertenecer un tercero / persona_natural.', 'Grupo_3', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (4, 'Estados', 'Tipos de Estados.', 'Grupo_4', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (5, 'Tipo Perfil', 'Tipos de Perfiles.', 'Grupo_5', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertTablaGrupoInfoComplementaria_20191107_094657) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id BETWEEN 1 AND 5;")
}
