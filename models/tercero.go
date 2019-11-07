package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Tercero struct {
	Id                  int                `orm:"column(id);pk;auto"`
	NombreCompleto      string             `orm:"column(nombre_completo)"`
	PrimerNombre        string             `orm:"column(primer_nombre);null"`
	SegundoNombre       string             `orm:"column(segundo_nombre);null"`
	PrimerApellido      string             `orm:"column(primer_apellido);null"`
	SegundoApellido     string             `orm:"column(segundo_apellido);null"`
	LugarOrigen         int                `orm:"column(lugar_origen);null"`
	FechaNacimiento     time.Time          `orm:"column(fecha_nacimiento);type(timestamp without time zone);null"`
	Activo              bool               `orm:"column(activo)"`
	TipoContribuyenteId *TipoContribuyente `orm:"column(tipo_contribuyente_id);rel(fk)"`
	FechaCreacion       time.Time          `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion   time.Time          `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *Tercero) TableName() string {
	return "tercero"
}

func init() {
	orm.RegisterModel(new(Tercero))
}

// AddTercero insert a new Tercero into database and returns
// last inserted Id on success.
func AddTercero(m *Tercero) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTerceroById retrieves Tercero by Id. Returns error if
// Id doesn't exist
func GetTerceroById(id int) (v *Tercero, err error) {
	o := orm.NewOrm()
	v = &Tercero{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTercero retrieves all Tercero matches certain condition. Returns empty list if
// no records exist
func GetAllTercero(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tercero)).RelatedSel()
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

	var l []Tercero
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

// UpdateTercero updates Tercero by Id and returns error if
// the record to be updated doesn't exist
func UpdateTerceroById(m *Tercero) (err error) {
	o := orm.NewOrm()
	v := Tercero{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTercero deletes Tercero by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTercero(id int) (err error) {
	o := orm.NewOrm()
	v := Tercero{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Tercero{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
