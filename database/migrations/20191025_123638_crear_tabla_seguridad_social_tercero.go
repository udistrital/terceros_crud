package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaSeguridadSocialTercero_20191025_123638 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaSeguridadSocialTercero_20191025_123638{}
	m.Created = "20191025_123638"

	migration.Register("CrearTablaSeguridadSocialTercero_20191025_123638", m)
}

// Run the migrations
func (m *CrearTablaSeguridadSocialTercero_20191025_123638) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.seguridad_social_tercero( id serial NOT NULL, tercero_id integer NOT NULL, tercero_entidad_id integer NOT NULL, activo boolean NOT NULL, fecha_inicio_vinculacion TIMESTAMP NOT NULL, fecha_fin_vinculacion TIMESTAMP, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_seguridad_social_tercero PRIMARY KEY (id), CONSTRAINT fk_tercero_seguridad_social_tercero FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id), CONSTRAINT fk_tercero_entidad_seguridad_social_tercero FOREIGN KEY (tercero_entidad_id) REFERENCES terceros.tercero(id) );")
	m.SQL("COMMENT ON TABLE terceros.seguridad_social_tercero IS 'Tabla que relaciona a un Tercero con su:  Eps / Caja Compensacion / Arl / fondo de Pension entre otros.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.id IS 'Identificador unico de la tabla info_complementaria_tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.tercero_id IS 'Identificador unico de la tabla tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.tercero_entidad_id IS 'Identificador unico de la tabla tercero que corresponde a la entidad.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_inicio_vinculacion IS 'Fecha en la que el tercero se vincula a la entidad.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_fin_vinculacion IS 'Fecha en la que el tercero se retira/desvincula de la entidad, cambio de activo a FALSE.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.seguridad_social_tercero.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaSeguridadSocialTercero_20191025_123638) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.seguridad_social_tercero")
}
