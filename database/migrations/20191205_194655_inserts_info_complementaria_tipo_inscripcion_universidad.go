package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsInfoComplementariaTipoInscripcionUniversidad_20191205_194655 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsInfoComplementariaTipoInscripcionUniversidad_20191205_194655{}
	m.Created = "20191205_194655"

	migration.Register("InsertsInfoComplementariaTipoInscripcionUniversidad_20191205_194655", m)
}

// Run the migrations
func (m *InsertsInfoComplementariaTipoInscripcionUniversidad_20191205_194655) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (16, 'Tipo de inscripción a la universidad', 'Tipo de inscripción a la universidad', 'Grupo_16', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Normal', 'Normal', TRUE, 16, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Indígenas', 'Indígenas', TRUE, 16, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Minorías Étnicas y Culturales', 'Minorías Étnicas', TRUE, 16, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Desplazados', 'Desplazados', TRUE, 16, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Mejores Bachilleres Col Distrital Oficial', 'Mejores Bachilleres', TRUE, 16, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Beneficiarios Ley 1084 de 2006', 'Ben. Ley 1084-2006', TRUE, 16, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertsInfoComplementariaTipoInscripcionUniversidad_20191205_194655) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 16;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 16;")
}
