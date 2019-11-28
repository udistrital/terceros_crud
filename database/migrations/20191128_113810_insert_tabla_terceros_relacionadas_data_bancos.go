package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaTercerosRelacionadasDataBancos_20191128_113810 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTercerosRelacionadasDataBancos_20191128_113810{}
	m.Created = "20191128_113810"

	migration.Register("InsertTablaTercerosRelacionadasDataBancos_20191128_113810", m)
}

// Run the migrations
func (m *InsertTablaTercerosRelacionadasDataBancos_20191128_113810) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTablaTercerosRelacionadasDataBancos_20191128_113810) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
