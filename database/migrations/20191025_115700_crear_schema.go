package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchema_20191025_115700 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchema_20191025_115700{}
	m.Created = "20191025_115700"

	migration.Register("CrearSchema_20191025_115700", m)
}

// Run the migrations
func (m *CrearSchema_20191025_115700) Up() {
	m.SQL("CREATE SCHEMA IF NOT EXISTS terceros;")
	m.SQL("ALTER SCHEMA terceros OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,terceros;")
}

// Reverse the migrations
func (m *CrearSchema_20191025_115700) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS terceros");
}
