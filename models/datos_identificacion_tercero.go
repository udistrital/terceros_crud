package models

import (
	"github.com/astaxie/beego/orm"
)

type DatosIdentificacionTercero struct {
	TerceroId int    `json:"tercero_id"`
	Compuesto string `json:"compuesto"`
}

func GetAllDatosIdentificacionTercero(search string, terceros *[]*DatosIdentificacionTercero) (err error) {
	o := orm.NewOrm()

	query := `
	WITH sin_doc AS (
		SELECT t.id tercero_id,
			t.nombre_completo compuesto
		FROM terceros.tercero t
		WHERE UPPER(t.nombre_completo) LIKE UPPER(?)
			AND t.activo = TRUE
			AND NOT EXISTS (
				(SELECT 1
					FROM terceros.datos_identificacion di
					WHERE di.tercero_id = t.id
				)
			)
		ORDER BY t.nombre_completo ASC
	), con_doc AS (
		SELECT t.id tercero_id,
			CONCAT(codigo_abreviacion, ' ', numero, ' ', nombre_completo) compuesto
		FROM terceros.tercero t,
			terceros.datos_identificacion di,
			terceros.tipo_documento td,
			CONCAT(
				UPPER(td.codigo_abreviacion), ' ',
				UPPER(di.numero), ' ',
				UPPER(t.nombre_completo)
				) compuesto_upper
		WHERE compuesto_upper LIKE UPPER(?)
			AND t.activo = TRUE
			AND di.activo = TRUE
			AND di.tercero_id = t.id
			AND td.id = di.tipo_documento_id
		ORDER BY t.nombre_completo ASC
	)

	SELECT *
	FROM sin_doc
	UNION ALL
	SELECT *
	FROM con_doc;
	`

	search_ := "%" + search + "%"
	if _, err := o.Raw(query, search_, search_).QueryRows(terceros); err != nil {
		return err
	}

	return
}

// GetTrIdentificacionTercero Consulta el nombre del tercero y el documento de identificación (uno solo) sí existe
func GetTrIdentificacionTercero(terceroId int, data *DatosIdentificacionTercero) (err error) {
	o := orm.NewOrm()

	query := `
	WITH sin_doc AS (
		SELECT t.id tercero_id,
			t.nombre_completo compuesto
		FROM terceros.tercero t
		WHERE t.id = ?
			AND NOT EXISTS (
				(SELECT 1
					FROM terceros.datos_identificacion di
					WHERE di.tercero_id = t.id
				)
			)
	), con_doc AS (
		SELECT di.tercero_id,
			CONCAT(codigo_abreviacion, ' ', numero, ' ', nombre_completo) compuesto
		FROM terceros.tercero t,
			terceros.datos_identificacion di,
			terceros.tipo_documento td
		WHERE t.id = ?
			AND di.activo = TRUE
			AND di.tercero_id = t.id
			AND td.id = di.tipo_documento_id
		ORDER BY td.id DESC
		LIMIT 1
	)

	SELECT *
	FROM sin_doc
	UNION ALL
	SELECT *
	FROM con_doc;
	`

	var data_ []*DatosIdentificacionTercero
	if _, err := o.Raw(query, terceroId, terceroId).QueryRows(&data_); err != nil {
		return err
	} else if len(data_) == 1 {
		*data = *data_[0]
	}

	return
}
