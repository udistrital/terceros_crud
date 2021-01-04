package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTipoTercero_20191217_095251 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoTercero_20191217_095251{}
	m.Created = "20191217_095251"

	migration.Register("InsertTipoTercero_20191217_095251", m)
}

// Run the migrations
func (m *InsertTipoTercero_20191217_095251) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (13, 'DOCENTE HORA CÁTEDRA SALARIOS', 'DOCENTE HORA CÁTEDRA SALARIOS.', 'DHC_SALARIOS', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (14, 'DOCENTE HORA CÁTEDRA HONORARIOS', 'DOCENTE HORA CÁTEDRA HONORARIOS.', 'DHC_HONORARIOS', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (15, 'DOCENTE DE PLANTA', 'DOCENTE DE PLANTA.', 'D_PLANTA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (16, 'FUNCIONARIO DE PLANTA', 'FUNCIONARIO DE PLANTA.', 'FUNCIONARIO_PLANTA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (17, 'CONTRATISTA', 'CONTRATISTA.', 'CONTRATISTA', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_tercero (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (18, 'PENSIONADO', 'PENSIONADO.', 'PENSIONADO', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertTipoTercero_20191217_095251) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_tercero WHERE id BETWEEN 13 AND 18;")
}
