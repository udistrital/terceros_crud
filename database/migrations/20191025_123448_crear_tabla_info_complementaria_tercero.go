package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaInfoComplementariaTercero_20191025_123448 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaInfoComplementariaTercero_20191025_123448{}
	m.Created = "20191025_123448"

	migration.Register("CrearTablaInfoComplementariaTercero_20191025_123448", m)
}

// Run the migrations
func (m *CrearTablaInfoComplementariaTercero_20191025_123448) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.info_complementaria_tercero( id serial NOT NULL, tercero_id integer NOT NULL, info_complementaria_id integer NOT NULL, dato json, activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_info_complementaria_tercero PRIMARY KEY (id), CONSTRAINT fk_tercero_info_complementaria_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id), CONSTRAINT fk_info_complementaria_info_complementaria_tercero FOREIGN KEY (info_complementaria_id) REFERENCES terceros.info_complementaria(id) );")
	m.SQL("COMMENT ON TABLE terceros.info_complementaria_tercero IS 'Tabla que relaciona a un Tercero con su información complementaria, el valor concreto con el que se relaciona puede ser encontrado en la columna dato o en la tabla info_complementaria dependiendo de si es parametrisable o no.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria_tercero.id IS 'Identificador unico de la tabla info_complementaria_tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria_tercero.tercero_id IS 'Identificador unico de la tabla tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria_tercero.info_complementaria_id IS 'Identificador unico de la tabla info_complementaria';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria_tercero.dato IS 'Campo para guardar informacion que no se pueda parametrizar, el tipo de dato que se guarda en esta estructura se relacionara en la tabla info_complementaria.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.info_complementaria_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaInfoComplementariaTercero_20191025_123448) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.info_complementaria_tercero")
}
