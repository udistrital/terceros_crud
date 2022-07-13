package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDatosIdentificacion_20191025_123156 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDatosIdentificacion_20191025_123156{}
	m.Created = "20191025_123156"

	migration.Register("CrearTablaDatosIdentificacion_20191025_123156", m)
}

// Run the migrations
func (m *CrearTablaDatosIdentificacion_20191025_123156) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS terceros.datos_identificacion( id serial NOT NULL, tipo_documento_id integer NOT NULL, tercero_id integer NOT NULL, numero character varying(50) NOT NULL, digito_verificacion numeric(2,0), ciudad_expedicion integer, fecha_expedicion TIMESTAMP, activo boolean NOT NULL, documento_soporte integer, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_datos_identificacion PRIMARY KEY (id), CONSTRAINT fk_tipo_documento_datos_identificacion FOREIGN KEY (tipo_documento_id) REFERENCES terceros.tipo_documento(id), CONSTRAINT fk_tercero_datos_identificacion FOREIGN KEY (tercero_id) REFERENCES terceros.tercero(id) );")
	m.SQL("COMMENT ON TABLE terceros.datos_identificacion IS 'Tabla que almacena el numero de identificación de un tercero y lo relaciona con su tipo pudiendo ser este CC, CE, NIT entre otros.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.id IS 'Identificador unico de la tabla datos_identificacion.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.tipo_documento_id IS 'Identificador unico de la tabla tipo_documento';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.tercero_id IS 'Identificador unico de la tabla tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.numero IS 'Numero de Identificacion del Tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.digito_verificacion IS 'Dígito de verificación de la DIAN. Toda persona natural o jurídica que tenga un RUT debe tener un DV (NOT NULL para personas_juridicas).';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.ciudad_expedicion IS 'Valor que almacena el ID de la ciudad en que fue expedido el documento.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.fecha_expedicion IS 'Fecha en la que fue expedido el documento de identificación.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.documento_soporte IS 'Identificador asociado al soporte del documento, correspondiente para el tercero.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN terceros.datos_identificacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaDatosIdentificacion_20191025_123156) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS terceros.datos_identificacion")
}
