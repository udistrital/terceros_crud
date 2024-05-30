package main

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"

	_ "github.com/udistrital/terceros_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
	"github.com/udistrital/utils_oas/xray"
)

func test() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func init() {
	test()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}

func main() {
	//Prueba CI - 2
	orm.RegisterDataBase("default", "postgres",
		"postgres://"+beego.AppConfig.String("PGuser")+
			":"+url.QueryEscape(beego.AppConfig.String("PGpass"))+
			"@"+beego.AppConfig.String("PGurls")+
			// TODO: Descomentar una vez exista TERCEROS_CRUD_PGPORT en el entorno
			// ":"+beego.AppConfig.String("PGport")+
			"/"+beego.AppConfig.String("PGdb")+
			"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
	xray.InitXRay()
	apistatus.Init()
	beego.ErrorController(&customerror.CustomErrorController{})
	beego.Run()
}
