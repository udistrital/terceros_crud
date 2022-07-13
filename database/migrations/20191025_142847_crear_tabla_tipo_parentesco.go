package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoParentesco_20191025_142847 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoParentesco_20191025_142847{}
	m.Created = "20191025_142847"

	migration.Register("CrearTablaTipoParentesco_20191025_142847", m)
}

// Run the migrations
func (m *CrearTablaTipoParentesco_20191025_142847) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.tipo_parentesco( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_parentesco PRIMARY KEY (id) );")
	m.SQL("COMMENT ON TABLE terceros.tipo_parentesco IS 'Tabla que parametriza a los tipos de parentescos que puede tener un tercero en relacion a otro, sean estos: Madre, Padre, Conyuge, Hijos, entre otros.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_parentesco.id IS 'Identificador unico de la tabla tipo_parentesco.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_parentesco.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_parentesco.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del tipo de parentesco.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_parentesco.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_parentesco.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_parentesco.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.tipo_parentesco.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTipoParentesco_20191025_142847) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.tipo_parentesco")
}
