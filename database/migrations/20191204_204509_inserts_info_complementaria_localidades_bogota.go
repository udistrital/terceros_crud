package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsInfoComplementariaLocalidadesBogota_20191204_204509 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsInfoComplementariaLocalidadesBogota_20191204_204509{}
	m.Created = "20191204_204509"

	migration.Register("InsertsInfoComplementariaLocalidadesBogota_20191204_204509", m)
}

// Run the migrations
func (m *InsertsInfoComplementariaLocalidadesBogota_20191204_204509) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (12, 'Localidades Bogotá', 'Localidades de Bogotá', 'Grupo_12', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Fuera de Bogotá', 'Fuera de Bogotá', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Usaquen', 'Usaquen', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Chapinero', 'Chapinero', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Santa Fe', 'Santa Fe', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('San Cristóbal', 'San Cristóbal', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Usem', 'Usem', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Tunjuelito', 'Tunjuelito', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Bosa', 'Bosa', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Kennedy', 'Kennedy', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Fontibón', 'Fontibón', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Engativa', 'Engativa', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Suba', 'Suba', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Barrios Unidos', 'Barrios Unidos', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Teusaquillo', 'Teusaquillo', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Los Mártires', 'Loos Mártires', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Antonio Nariño', 'Antonio Nariño', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Puente Aranda', 'Puente Aranda', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('La Candelaria', 'La Candelaria', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Rafael Uribe Uribe', 'Rafael Uribe Uribe', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Ciudad Bolivar', 'Ciudad Bolivar', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Sumapaz', 'Sumapaz', TRUE, 12, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertsInfoComplementariaLocalidadesBogota_20191204_204509) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 12;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 12;")
}
