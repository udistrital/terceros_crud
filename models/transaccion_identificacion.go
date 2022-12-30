package models

import "github.com/astaxie/beego/orm"

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

func GetAllDatosIdentificacionVinculacion(search string, tiposVinculacion string, terceros *[]DatosIdentificacionTercero_) (err error) {

	o := orm.NewOrm()

	search_ := "%" + search + "%"
	script := scriptVinculaciones(tiposVinculacion)
	if _, err = o.Raw(script, search_, search_).QueryRows(terceros); err != nil {
		return err
	}

	return

}

func scriptVinculaciones(tipos string) (script string) {

	script = `
	WITH con_doc AS (
		SELECT DISTINCT ON (1)
			t.id,
			t.nombre_completo,
			di.numero,
			di.digito_verificacion,
			td.codigo_abreviacion,
			td.id tipo
		FROM
			terceros.tercero t,
			terceros.datos_identificacion di,
			terceros.tipo_documento td,
			terceros.vinculacion v,
			CONCAT(
				UPPER(td.codigo_abreviacion), ' ',
				UPPER(di.numero), ' ',
				UPPER(t.nombre_completo)
				) compuesto
		WHERE
			v.tercero_principal_id = t.id
			`

	if len(tipos) > 0 {
		script += "AND v.tipo_vinculacion_id IN (" + tipos + ")"
	}

	script += `
			AND compuesto LIKE UPPER(?)
			AND di.tercero_id = t.id
			AND td.id = di.tipo_documento_id
			AND v.activo = TRUE
			AND t.activo = TRUE
			AND di.activo = TRUE
		ORDER BY 1, 6 DESC
	), sin_doc AS (
		SELECT
			t.id,
			t.nombre_completo
		FROM
			terceros.tercero t,
			terceros.vinculacion v
		WHERE
			UPPER(t.nombre_completo) LIKE UPPER(?)
			`

	if len(tipos) > 0 {
		script += "AND v.tipo_vinculacion_id IN (" + tipos + ")"
	}

	script += `
			AND v.tercero_principal_id = t.id
			AND v.activo = TRUE
			AND t.activo = TRUE
			AND NOT EXISTS (
				SELECT 1
					FROM terceros.datos_identificacion di
					WHERE di.tercero_id = t.id
			)
	)

	SELECT
		*,
		NULL AS numero,
		NULL AS digito_verificacion,
		NULL AS codigo_abreviacion
	FROM sin_doc
	UNION
	SELECT
		id,
		nombre_completo,
		numero,
		digito_verificacion,
		codigo_abreviacion
	FROM con_doc
	ORDER BY nombre_completo;
	`

	return
}
