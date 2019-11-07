package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaInfoComplementaria_20191025_120313 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaInfoComplementaria_20191025_120313{}
	m.Created = "20191025_120313"

	migration.Register("CrearTablaInfoComplementaria_20191025_120313", m)
}

// Run the migrations
func (m *CrearTablaInfoComplementaria_20191025_120313) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.info_complementaria( id serial NOT NULL, nombre character varying(100) NOT NULL, codigo_abreviacion character varying(20), activo boolean NOT NULL, tipo_de_dato character varying(20), --Evaluar uso de UNIQUE KEY grupo_info_complementaria_id integer NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_info_complementaria PRIMARY KEY (id), CONSTRAINT fk_grupo_info_complementaria_info_complementaria FOREIGN KEY (grupo_info_complementaria_id) REFERENCES terceros.grupo_info_complementaria(id) );")
	m.SQL("ALTER TABLE terceros.info_complementaria OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE terceros.info_complementaria IS 'Tabla que parametriza los tipos concretos de grupos sanguineo, grupos etnicos, discapacidades, estados civiles, generos y cualquier informacion complementaria relacionada a un tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.id IS 'Identificador unico de la tabla info_complementaria.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.tipo_de_dato IS 'Valor que indica el tipo de dato almacenado en la tabla info_complementaria_tercero en la columna dato.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.grupo_info_complementaria_id IS 'Identificador unico de la tabla tipo_informacion_basica.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaInfoComplementaria_20191025_120313) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.info_complementaria")
}
