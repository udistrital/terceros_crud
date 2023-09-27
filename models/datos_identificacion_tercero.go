package models

import (
	"strings"

	"github.com/astaxie/beego/orm"
)

type DatosIdentificacionTercero struct {
	TerceroId int    `json:"tercero_id"`
	Compuesto string `json:"compuesto"`
}

func GetAllDatosIdentificacionTercero(documento, search string, terceros *[]DatosIdentificacionTercero_) (err error) {
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
				ON td.id = di.tipo_documento_id AND di.activo = TRUE
		WHERE COND_TERCERO
			and t.activo = true
			and (td.id is null or td.codigo_abreviacion != 'CODE')
		ORDER BY 1 ASC, 6 DESC
	)

	SELECT *
	FROM terceros
	ORDER BY 2;
	`

	parameter := ""
	filter := ""
	if documento != "" {
		filter = "di.numero = ?"
		parameter = documento
	} else {
		filter = `CONCAT(
				UPPER(td.codigo_abreviacion), ' ',
				UPPER(di.numero), ' ',
				UPPER(t.nombre_completo)) LIKE UPPER(?)`
		parameter = "%" + search + "%"
	}

	query = strings.ReplaceAll(query, "COND_TERCERO", filter)
	_, err = o.Raw(query, parameter).QueryRows(terceros)
	return
}

// GetTrIdentificacionTercero Consulta el nombre del tercero y el documento de identificación (uno solo) si existe
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
