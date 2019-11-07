package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
