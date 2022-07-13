package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTerceroFamiliar_20191025_143040 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTerceroFamiliar_20191025_143040{}
	m.Created = "20191025_143040"

	migration.Register("CrearTablaTerceroFamiliar_20191025_143040", m)
}

// Run the migrations
func (m *CrearTablaTerceroFamiliar_20191025_143040) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.tercero_familiar( id serial NOT NULL, tercero_id integer NOT NULL, tercero_familiar_id integer NOT NULL, tipo_parentesco_id integer NOT NULL, codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tercero_familiar PRIMARY KEY (id), CONSTRAINT fk_tercero_tercero_familiar FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id), CONSTRAINT fk_tercero_familiar_tercero_familiar FOREIGN KEY (tercero_familiar_id) REFERENCES terceros.tercero(id), CONSTRAINT fk_tipo_parentesco_tercero_familiar FOREIGN KEY (tipo_parentesco_id) REFERENCES terceros.tipo_parentesco(id) );")
	m.SQL("COMMENT ON TABLE terceros.tercero_familiar IS 'Tabla que parametriza a los tipos de parentescos que puede tener un tercero en relacion a otro, sean estos: Madre, Padre, Conyuge, Hijos, entre otros.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.id IS 'Identificador unico de la tabla tercero_familiar.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.tercero_id IS 'Identificador unico de la tabla tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.tercero_familiar_id IS 'Identificador unico de la tabla tercero correspondiente al familiar.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.tipo_parentesco_id IS 'Identificador unico de la relacion de parentesco que existe entre tercero_id y tercero_familiar_id.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_familiar.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTerceroFamiliar_20191025_143040) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.tercero_familiar")
}
