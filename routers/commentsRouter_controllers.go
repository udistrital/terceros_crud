package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:DatosIdentificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:GrupoInfoComplementariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:InfoComplementariaTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:SeguridadSocialTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "PostInformacionFamiliar",
            Router: `/informacion_familiar`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroFamiliarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TerceroTipoTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoContribuyenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoParentescoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/terceros_crud/controllers:TipoTerceroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
