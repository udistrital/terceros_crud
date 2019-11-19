package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterTableTercerosAddUsuarioWso2_20191119_144700 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterTableTercerosAddUsuarioWso2_20191119_144700{}
	m.Created = "20191119_144700"

	migration.Register("AlterTableTercerosAddUsuarioWso2_20191119_144700", m)
}

// Run the migrations
func (m *AlterTableTercerosAddUsuarioWso2_20191119_144700) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE terceros.tercero ADD usuario_wso2 character varying(20);")
	m.SQL("COMMENT ON COLUMN terceros.tercero.usuario_wso2 IS 'Campo opcional para guardar el usuario de WSO2 de una persona_natural.';")
}

// Reverse the migrations
func (m *AlterTableTercerosAddUsuarioWso2_20191119_144700) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE terceros.tercero DROP COLUMN usuario_wso2;")
}
