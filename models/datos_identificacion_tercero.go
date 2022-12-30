package models

import (
	"github.com/astaxie/beego/orm"
)

type DatosIdentificacionTercero struct {
	TerceroId int    `json:"tercero_id"`
	Compuesto string `json:"compuesto"`
}

func GetAllDatosIdentificacionTercero(search string, terceros *[]DatosIdentificacionTercero_) (err error) {
	o := orm.NewOrm()

	query := `
	WITH sin_doc AS (
		SELECT
			t.id,
			t.nombre_completo
		FROM
			terceros.tercero t
		WHERE
				UPPER(t.nombre_completo) LIKE UPPER(?)
			AND t.activo = TRUE
			AND NOT EXISTS (
				SELECT 1
					FROM terceros.datos_identificacion di
					WHERE di.tercero_id = t.id
			)
	), con_doc AS (
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
			CONCAT(
				UPPER(td.codigo_abreviacion), ' ',
				UPPER(di.numero), ' ',
				UPPER(t.nombre_completo)
				) compuesto
		WHERE
				compuesto LIKE UPPER(?)
			AND di.tercero_id = t.id
			AND td.id = di.tipo_documento_id
			AND t.activo = TRUE
			AND di.activo = TRUE
		ORDER BY 1, 6 DESC
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
	ORDER BY 2 ASC;
	`

	search_ := "%" + search + "%"
	if _, err := o.Raw(query, search_, search_).QueryRows(terceros); err != nil {
		return err
	}

	return
}

// GetTrIdentificacionTercero Consulta el nombre del tercero y el documento de identificación (uno solo) sí existe
func GetTrIdentificacionTercero(terceroId int, data *DatosIdentificacionTercero_) (err error) {
	o := orm.NewOrm()

	query := `
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
			terceros.tipo_documento td
		WHERE
				t.id = ?
			AND di.tercero_id = t.id
			AND td.id = di.tipo_documento_id
			AND di.activo = TRUE
		ORDER BY 1, 6 DESC
	), sin_doc AS (
		SELECT
			t.id,
			t.nombre_completo
		FROM
			terceros.tercero t
		WHERE
				t.id = ?
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
	FROM con_doc;
	`

	var data_ []DatosIdentificacionTercero_
	if _, err = o.Raw(query, terceroId, terceroId).QueryRows(&data_); err != nil {
		return
	} else if len(data_) == 1 {
		*data = data_[0]
	}

	return
}
