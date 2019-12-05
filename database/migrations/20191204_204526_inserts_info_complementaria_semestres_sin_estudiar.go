package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsInfoComplementariaSemestresSinEstudiar_20191204_204526 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsInfoComplementariaSemestresSinEstudiar_20191204_204526{}
	m.Created = "20191204_204526"

	migration.Register("InsertsInfoComplementariaSemestresSinEstudiar_20191204_204526", m)
}

// Run the migrations
func (m *InsertsInfoComplementariaSemestresSinEstudiar_20191204_204526) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (14, 'Semestres transcurridos sin estudiar', 'Semestres transcurridos sin estudiar', 'Grupo_14', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Recién graduado', 'Recién graduado', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('1 Semestre', '1 Semestre', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('2 Semestres', '2 Semestres', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('3 Semestres', '3 Semestres', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('4 Semestres', '4 Semestres', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('5 Semestres', '5 Semestres', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('6 Semestres', '6 Semestres', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Más de tres años', 'Más de tres años', TRUE, 14, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertsInfoComplementariaSemestresSinEstudiar_20191204_204526) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 14;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 14;")
}
