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
	WITH terceros AS (
		SELECT DISTINCT ON (1)
			t.id,
			t.nombre_completo,
			di.numero,
			di.digito_verificacion,
			td.codigo_abreviacion,
			td.id tipo
		FROM
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
			and t.activo = true
			and (di.id is null or di.activo = true)
			and (td.id is null or td.codigo_abreviacion != 'CODE')
		ORDER BY 1, 6 DESC
	)

	SELECT *
	FROM terceros
	ORDER BY 2;
	`

	_, err = o.Raw(query, "%"+search+"%").QueryRows(terceros)
	return
}

// GetTrIdentificacionTercero Consulta el nombre del tercero y el documento de identificación (uno solo) sí existe
func GetTrIdentificacionTercero(terceroId int, data *DatosIdentificacionTercero_) (err error) {
	o := orm.NewOrm()

	query := `
	SELECT DISTINCT ON (1)
		t.id,
		t.nombre_completo,
		di.numero,
		di.digito_verificacion,
		td.codigo_abreviacion,
		td.id tipo
	FROM
		terceros.tercero t
		LEFT JOIN terceros.datos_identificacion di
			ON di.tercero_id = t.id
		LEFT JOIN terceros.tipo_documento td
			ON td.id = di.tipo_documento_id
	WHERE	t.id = ?
		and (di.id is null or di.activo = true)
		and (td.id is null or td.codigo_abreviacion != 'CODE');
	`

	var data_ []DatosIdentificacionTercero_
	_, err = o.Raw(query, terceroId).QueryRows(&data_)
	if err == nil && len(data_) == 1 {
		*data = data_[0]
	}

	return
}
