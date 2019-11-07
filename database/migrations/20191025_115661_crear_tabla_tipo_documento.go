package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoDocumento_20191025_115661 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoDocumento_20191025_115661{}
	m.Created = "20191025_115661"

	migration.Register("CrearTablaTipoDocumento_20191025_115661", m)
}

// Run the migrations
func (m *CrearTablaTipoDocumento_20191025_115661) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.tipo_documento( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_documento PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE terceros.tipo_documento OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE terceros.tipo_documento IS 'Tabla que parametriza los diferentes tipos de documento que puede tener una persona juridica o natural.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.id IS 'Identificador unico de la tabla tipo_documento.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo_documento.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_documento.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTipoDocumento_20191025_115661) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.tipo_documento")
}
