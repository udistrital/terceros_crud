package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaTipoDocumento_20191107_105856 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTipoDocumento_20191107_105856{}
	m.Created = "20191107_105856"

	migration.Register("InsertTablaTipoDocumento_20191107_105856", m)
}

// Run the migrations
func (m *InsertTablaTipoDocumento_20191107_105856) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('REGISTRO CIVIL DE NACIMIENTO', 'REGISTRO CIVIL DE NACIMIENTO', 'RC', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('TARJETA DE IDENTIDAD', 'TARJETA DE IDENTIDAD', 'TI', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CÉDULA DE CIUDADANÍA', 'CÉDULA DE CIUDADANÍA', 'CC', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CERTIFICADO REGISTRADURÍA SIN IDENTIFICACIÓN', 'CERTIFICADO REGISTRADURÍA SIN IDENTIFICACIÓN', 'CRSI', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('TARJETA DE EXTRANJERÍA', 'TARJETA DE EXTRANJERÍA', 'TE', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CÉDULA DE EXTRANJERÍA', 'CÉDULA DE EXTRANJERÍA', 'CE', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('NIT', 'NIT', 'NIT', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('IDENTIFICACIÓN DE EXTRANJEROS DIFERENTE AL NIT ASIGNADO DIAN', 'IDENTIFICACIÓN DE EXTRANJEROS DIFERENTE AL NIT ASIGNADO DIAN', 'IEDD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('PASAPORTE', 'PASAPORTE', 'PAS', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO', 'DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO', 'DIE', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('SIN IDENTIFICACIÓN DEL EXTERIOR O PARA USO DEFINIDO POR DIAN', 'SIN IDENTIFICACIÓN DEL EXTERIOR O PARA USO DEFINIDO POR DIAN', 'SIED', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO PERSONA JURÍDICA', 'DOCUMENTO DE IDENTIFICACIÓN EXTRANJERO PERSONA JURÍDICA', 'DIEPJ', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.tipo_documento (nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES ('CARNÉ DIPLOMÁTICO', 'CARNÉ DIPLOMÁTICO', 'CD', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	
}

// Reverse the migrations
func (m *InsertTablaTipoDocumento_20191107_105856) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.tipo_documento;")
}
