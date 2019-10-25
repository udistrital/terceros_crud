package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTerceroTipoTercero_20191025_122927 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTerceroTipoTercero_20191025_122927{}
	m.Created = "20191025_122927"

	migration.Register("CrearTablaTerceroTipoTercero_20191025_122927", m)
}

// Run the migrations
func (m *CrearTablaTerceroTipoTercero_20191025_122927) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.tercero_tipo_tercero( id serial NOT NULL, tercero_id integer NOT NULL, tipo_tercero_id integer NOT NULL, activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tercero_tipo_tercero PRIMARY KEY (id), CONSTRAINT fk_tercero_tercero_tipo_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id), CONSTRAINT fk_tipo_tercero_tercero_tipo_tercero FOREIGN KEY (tipo_tercero_id) REFERENCES terceros.tipo_tercero(id) );")
	m.SQL("ALTER TABLE terceros.tercero_tipo_tercero OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE terceros.tercero_tipo_tercero IS 'Tabla que relaciona un tercero con uno o mas tipos de tercero Ej: Un colegio es de tipo colegio / organizacion  y a su vez es de carácter publico-privado-mixto.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_tipo_tercero.id IS 'Identificador unico de la tabla tercero_tipo_tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_tipo_tercero.tercero_id IS 'Identificador unico de la tabla tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_tipo_tercero.tipo_tercero_id IS 'Identificador unico de la tabla tipo_tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_tipo_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_tipo_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.tercero_tipo_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTerceroTipoTercero_20191025_122927) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.tercero_tipo_tercero")
}
