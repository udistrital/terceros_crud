package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaGrupoInfoComplementaria_20191025_115951 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaGrupoInfoComplementaria_20191025_115951{}
	m.Created = "20191025_115951"

	migration.Register("CrearTablaGrupoInfoComplementaria_20191025_115951", m)
}

// Run the migrations
func (m *CrearTablaGrupoInfoComplementaria_20191025_115951) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.grupo_info_complementaria( id serial NOT NULL DEFAULT, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_grupo_info_complementaria PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE terceros.grupo_info_complementaria OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE terceros.grupo_info_complementaria IS 'Tabla que parametriza la información complementaria de un tercero que pueda ser relacionada a un mismo grupo o tipo como pueden ser: discapacidades, estados civiles, grupos sanguíneo, grupos étnicos, géneros entre otros. También se requiere un grupo para valores cuyos datos internos no van a ser parametrizados.';")
	m.SQL("COMMENT ON COLUMN terceros.grupo_info_complementaria.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN terceros.grupo_info_complementaria.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del grupo o tipo.';")
	m.SQL("COMMENT ON COLUMN terceros.grupo_info_complementaria.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN terceros.grupo_info_complementaria.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.grupo_info_complementaria.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.grupo_info_complementaria.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaGrupoInfoComplementaria_20191025_115951) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.grupo_info_complementaria")
}
