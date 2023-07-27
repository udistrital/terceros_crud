package models

import (
	"strings"

	"github.com/astaxie/beego/orm"
)

type DatosIdentificacionTercero_ struct {
	Tercero struct {
		Id             int
		NombreCompleto string
	}
	Identificacion struct {
		Numero             string
		DigitoVerificacion float64
		TipoDocumentoId    struct {
			CodigoAbreviacion string
		}
	}
}

func GetAllDatosIdentificacionVinculacion(search string, tiposVinculacion, cargos, dependencias string, terceros *[]DatosIdentificacionTercero_) (err error) {

	o := orm.NewOrm()

	script := scriptVinculaciones(tiposVinculacion, cargos, dependencias)
	_, err = o.Raw(script, "%"+search+"%").QueryRows(terceros)
	return

}

func scriptVinculaciones(tiposVinculacion, cargos, dependencia string) (script string) {

	script = `
	WITH vinculaciones AS (
		SELECT DISTINCT ON (1)
			t.id,
			t.nombre_completo,
			di.numero,
			di.digito_verificacion,
			td.codigo_abreviacion,
			td.id tipo
		FROM
			terceros.vinculacion v,
			terceros.tercero t
			LEFT JOIN terceros.datos_identificacion di
				ON di.tercero_id = t.id
			LEFT JOIN terceros.tipo_documento td
				ON td.id = di.tipo_documento_id,
			CONCAT(
				UPPER(td.codigo_abreviacion), ' ',
				UPPER(di.numero), ' ',
				UPPER(t.nombre_completo)
				) compuesto
		WHERE
				compuesto LIKE UPPER(?)
			AND t.activo = true
			AND v.activo = true
			-- AND (di.id is null or di.activo = true)
			AND (td.id is null or td.codigo_abreviacion != 'CODE')
			AND v.tercero_principal_id = t.id
			COND_TIPO_VINCULACION_ID
			COND_CARGO_ID
			COND_DEPENDENCIA_ID
		ORDER BY 1, 6 DESC
	)

	SELECT *
	FROM vinculaciones
	ORDER BY 2;
	`

	condVinculacion := ""
	if len(tiposVinculacion) > 0 {
		condVinculacion = "AND v.tipo_vinculacion_id IN (" + tiposVinculacion + ")"

	}

	condCargos := ""
	if len(cargos) > 0 {
		condCargos = "AND v.cargo_id IN (" + cargos + ")"
	}

	condDependencia := ""
	if len(cargos) > 0 {
		condCargos = "AND v.dependencia_id IN (" + dependencia + ")"
	}

	script = strings.ReplaceAll(script, "COND_TIPO_VINCULACION_ID", condVinculacion)
	script = strings.ReplaceAll(script, "COND_CARGO_ID", condCargos)
	script = strings.ReplaceAll(script, "COND_DEPENDENCIA_ID", condDependencia)

	return
}
