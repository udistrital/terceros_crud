package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoTercero_20191025_120726 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoTercero_20191025_120726{}
	m.Created = "20191025_120726"

	migration.Register("CrearTablaTipoTercero_20191025_120726", m)
}

// Run the migrations
func (m *CrearTablaTipoTercero_20191025_120726) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.tipo_tercero( id serial NOT NULL DEFAULT, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_tercero PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE terceros.tipo_tercero OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE terceros.tipo_tercero IS 'Tabla que parametriza a los tipos de terceros, sean: Personas/Empresas/Eps/Caja Compensacion/Arl, Entidades de caracter Publico/Privado/Mixto entre otros.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_tercero.id IS 'Identificador unico de la tabla tipo_tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_tercero.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_tercero.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_tercero.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaTipoTercero_20191025_120726) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.tipo_tercero")
}
