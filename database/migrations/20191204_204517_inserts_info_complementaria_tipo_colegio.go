package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsInfoComplementariaTipoColegio_20191204_204517 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsInfoComplementariaTipoColegio_20191204_204517{}
	m.Created = "20191204_204517"

	migration.Register("InsertsInfoComplementariaTipoColegio_20191204_204517", m)
}

// Run the migrations
func (m *InsertsInfoComplementariaTipoColegio_20191204_204517) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO terceros.grupo_info_complementaria (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (13, 'Tipo Colegio', 'Tipos de colegio', 'Grupo_13', TRUE, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Oficial', 'Colegios oficiales', TRUE, 13, LOCALTIMESTAMP, LOCALTIMESTAMP);")
	m.SQL("INSERT INTO terceros.info_complementaria (nombre, codigo_abreviacion, activo, grupo_info_complementaria_id, fecha_creacion, fecha_modificacion) VALUES ('Privado', 'Colegios privados', TRUE, 13, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}
// Reverse the migrations
func (m *InsertsInfoComplementariaTipoColegio_20191204_204517) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM terceros.info_complementaria WHERE grupo_info_complementaria_id = 13;")
	m.SQL("DELETE FROM terceros.grupo_info_complementaria WHERE id = 13;")
}
