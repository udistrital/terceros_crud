package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type DatosIdentificacionTercero struct {
	TerceroId int    `json:"tercero_id"`
	Compuesto string `json:"compuesto"`
}

func GetAllDatosIdentificacionTercero(search string, terceros *[]*DatosIdentificacionTercero) (err error) {
	o := orm.NewOrm()

	var query string
	query = `
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

	for _, tr := range *terceros {
		fmt.Println(tr.TerceroId, tr.Compuesto)
	}

	fmt.Println("Número", len(*terceros))

	return
}
