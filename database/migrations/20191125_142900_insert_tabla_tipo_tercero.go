package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaTipoTercero_20191125_142900 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTipoTercero_20191125_142900{}
	m.Created = "20191125_142900"

	migration.Register("InsertTablaTipoTercero_20191125_142900", m)
}

// Run the migrations
func (m *InsertTablaTipoTercero_20191125_142900) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1, 'BANCO', 'Entidad Bancaria', 'BANCO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2, 'E_ASEGURADORA', 'Entidad Aseguradora', 'E_ASEGURADORA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (3, 'EPS', 'Entidad Prestadora de Salud ', 'EPS', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (4, 'ARL', 'Administradora de Riesgos Laborales', 'ARL', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (5, 'F_PENSION', 'Fondo de Pensión', 'F_PENSION', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (6, 'CAJA_COMPENSACIÓN', 'Caja de Compensación', 'CAJA_COMPENSACIÓN', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (7, 'COLEGIO', null, 'COLEGIO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (8, 'ENTIDAD_PUBLICA', null, 'ENTIDAD_PUBLICA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (9, 'ENTIDAD_PRIVADA', null, 'ENTIDAD_PRIVADA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (10, 'ENTIDAD_MIXTA', 'Eentidad de caracter tanto publica como privada', 'ENTIDAD_MIXTA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (11, 'FAMILIAR', 'Persona natural que tiene un parentesco con otra persona natural de la organización', 'FAMILIAR', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
}

// Reverse the migrations
func (m *InsertTablaTipoTercero_20191125_142900) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_tercero WHERE id BETWEEN 1 AND 11;")
}
