package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type TerceroFamiliar struct {
	Id                int             `orm:"column(id);pk;auto"`
	TerceroId         *Tercero        `orm:"column(tercero_id);rel(fk)"`
	TerceroFamiliarId *Tercero        `orm:"column(tercero_familiar_id);rel(fk)"`
	TipoParentescoId  *TipoParentesco `orm:"column(tipo_parentesco_id);rel(fk)"`
	CodigoAbreviacion string          `orm:"column(codigo_abreviacion);null"`
	Activo            bool            `orm:"column(activo)"`
	FechaCreacion     string       `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string       `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

type TerceroFamiliarConInfoComplementaria struct {
	Familiar            *TerceroFamiliar
	InformacionContacto *[]InfoComplementariaTercero
}

type TrPostInformacionFamiliar struct {
	// Tercero_Familiar *Tercero
	Familiares  		 *[]TerceroFamiliarConInfoComplementaria
}

func (t *TerceroFamiliar) TableName() string {
	return "tercero_familiar"
}

func init() {
	orm.RegisterModel(new(TerceroFamiliar))
}

// AddTerceroFamiliar insert a new TerceroFamiliar into database and returns
// last inserted Id on success.
func AddTerceroFamiliar(m *TerceroFamiliar) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// TrPostInformacionFamiliar transaction to insert all information
// last inserted Id on success.
func AddInformacionFamiliar(m *TrPostInformacionFamiliar) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()

	date := time_bogota.TiempoBogotaFormato()
	for _, v := range *m.Familiares {

		familiar := v.Familiar

		// familiar.TerceroId.Id = int(idTercero)

		familiar.TerceroFamiliarId.Activo = true
		familiar.TerceroFamiliarId.FechaCreacion = date
		familiar.TerceroFamiliarId.FechaModificacion = date

		if idFamiliar, errTr := o.Insert(familiar.TerceroFamiliarId); errTr == nil {
			fmt.Println("Familiar registrado", idFamiliar)
			familiar.TerceroFamiliarId.Id = int(idFamiliar)
		} else {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
			return
		}

		familiar.Activo = true
		familiar.FechaCreacion = date
		familiar.FechaModificacion = date

		if _, errTr := o.Insert(familiar); errTr == nil {
			fmt.Println("Relación Tercero-Familiar registrado")
		} else {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
			return
		}

		for _, dato := range *v.InformacionContacto {

			dato.TerceroId.Id = familiar.TerceroFamiliarId.Id
			dato.Activo = true
			dato.FechaCreacion = date
			dato.FechaModificacion = date

			if _, errTr := o.Insert(&dato); errTr == nil {
				fmt.Println("dato de contacto registrado", dato.Dato)
			} else {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

	}

	_ = o.Commit()
	return
}

// GetTerceroFamiliarById retrieves TerceroFamiliar by Id. Returns error if
// Id doesn't exist
func GetTerceroFamiliarById(id int) (v *TerceroFamiliar, err error) {
	o := orm.NewOrm()
	v = &TerceroFamiliar{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTerceroFamiliar retrieves all TerceroFamiliar matches certain condition. Returns empty list if
// no records exist
func GetAllTerceroFamiliar(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TerceroFamiliar)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TerceroFamiliar
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTerceroFamiliar updates TerceroFamiliar by Id and returns error if
// the record to be updated doesn't exist
func UpdateTerceroFamiliarById(m *TerceroFamiliar) (err error) {
	o := orm.NewOrm()
	v := TerceroFamiliar{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTerceroFamiliar deletes TerceroFamiliar by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTerceroFamiliar(id int) (err error) {
	o := orm.NewOrm()
	v := TerceroFamiliar{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TerceroFamiliar{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
