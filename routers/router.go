// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/terceros_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_parentesco",
			beego.NSInclude(
				&controllers.TipoParentescoController{},
			),
		),

		beego.NSNamespace("/tipo_tercero",
			beego.NSInclude(
				&controllers.TipoTerceroController{},
			),
		),

		beego.NSNamespace("/datos_identificacion",
			beego.NSInclude(
				&controllers.DatosIdentificacionController{},
			),
		),

		beego.NSNamespace("/tercero",
			beego.NSInclude(
				&controllers.TerceroController{},
			),
		),

		beego.NSNamespace("/info_complementaria_tercero",
			beego.NSInclude(
				&controllers.InfoComplementariaTerceroController{},
			),
		),

		beego.NSNamespace("/seguridad_social_tercero",
			beego.NSInclude(
				&controllers.SeguridadSocialTerceroController{},
			),
		),

		beego.NSNamespace("/tercero_familiar",
			beego.NSInclude(
				&controllers.TerceroFamiliarController{},
			),
		),

		beego.NSNamespace("/tercero_tipo_tercero",
			beego.NSInclude(
				&controllers.TerceroTipoTerceroController{},
			),
		),

		beego.NSNamespace("/tipo_contribuyente",
			beego.NSInclude(
				&controllers.TipoContribuyenteController{},
			),
		),

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/grupo_info_complementaria",
			beego.NSInclude(
				&controllers.GrupoInfoComplementariaController{},
			),
		),

		beego.NSNamespace("/info_complementaria",
			beego.NSInclude(
				&controllers.InfoComplementariaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
