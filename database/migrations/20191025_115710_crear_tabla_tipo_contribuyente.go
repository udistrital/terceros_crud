package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoContribuyente_20191025_115710 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoContribuyente_20191025_115710{}
	m.Created = "20191025_115710"

	migration.Register("CrearTablaTipoContribuyente_20191025_115710", m)
}

// Run the migrations
func (m *CrearTablaTipoContribuyente_20191025_115710) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.tipo_contribuyente( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_contribuyente PRIMARY KEY (id) );")
	m.SQL("COMMENT ON TABLE terceros.tipo_contribuyente IS 'Tabla que parametriza los Tipos de contribuyentes: PersonaNatural - Persona Juridica';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_contribuyente.id IS 'Identificador unico de la tabla tipo_contribuyente.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_contribuyente.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_contribuyente.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_contribuyente.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_contribuyente.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_contribuyente.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_contribuyente.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_contribuyente.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTipoContribuyente_20191025_115710) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.tipo_contribuyente")
}
