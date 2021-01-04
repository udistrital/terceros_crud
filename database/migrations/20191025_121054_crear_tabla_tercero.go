package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTercero_20191025_121054 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTercero_20191025_121054{}
	m.Created = "20191025_121054"

	migration.Register("CrearTablaTercero_20191025_121054", m)
}

// Run the migrations
func (m *CrearTablaTercero_20191025_121054) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.tercero( id serial NOT NULL, nombre_completo character varying(255) NOT NULL, primer_nombre character varying(100), segundo_nombre character varying(100), primer_apellido character varying(100), segundo_apellido character varying(100), lugar_origen integer, fecha_nacimiento TIMESTAMP, activo boolean NOT NULL, tipo_contribuyente_id integer NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tercero PRIMARY KEY (id), CONSTRAINT fk_tipo_contribuyente_tercero FOREIGN KEY (tipo_contribuyente_id) REFERENCES terceros.tipo_contribuyente(id) );")
	m.SQL("ALTER TABLE terceros.tercero OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE terceros.tercero IS 'Tabla que parametriza los Tipos de contribuyentes: PersonaNatural - Persona Juridica';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.id IS 'Identificador unico de la tabla tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.nombre_completo IS 'Campo obligatorio de la tabla que indica el nombre de la persona natural o juridica';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.primer_nombre IS 'Primer nombre persona natural.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.segundo_nombre IS 'Segundo nombre persona natural.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.primer_apellido IS 'Primer apellido persona natural.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.segundo_apellido IS 'Segundo apellido persona natural.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.lugar_origen IS 'Id del pais de origen de la persona / empresa / entidad.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.fecha_nacimiento IS 'Campo que indica la fecha de nacimiento de la persona natural / indica la fecha de creacion de la empresa.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.tipo_contribuyente_id IS 'Identificador de la tabla tipo_contribuyente, diferencia entre persona_natural, persona_juridica.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaTercero_20191025_121054) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.tercero")
}
