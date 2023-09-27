package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/terceros_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

// TerceroController operations for Tercero
type TerceroController struct {
	beego.Controller
}

// URLMapping ...
func (c *TerceroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllTrTerceroIdentificacion", c.GetAllTrTerceroIdentificacion)
	c.Mapping("GetTrIdentificacionTercero", c.GetTrIdentificacionTercero)
}

// Post ...
// @Title Post
// @Description create Tercero
// @Param	body		body 	models.Tercero	true		"body for Tercero content"
// @Success 201 {int} models.Tercero
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TerceroController) Post() {
	var v models.Tercero
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaCreacion = time_bogota.TiempoBogotaFormato()
		v.FechaModificacion = time_bogota.TiempoBogotaFormato()
		if _, err := models.AddTercero(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Tercero by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tercero
// @Failure 404 not found resource
// @router /:id [get]
func (c *TerceroController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTerceroById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Tercero
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Tercero
// @Failure 404 not found resource
// @router / [get]
func (c *TerceroController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllTercero(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Tercero
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Tercero	true		"body for Tercero content"
// @Success 200 {object} models.Tercero
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TerceroController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Tercero{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		infoAd, _ := models.GetTerceroById(id)
		if infoAd != nil {
			v.FechaCreacion = time_bogota.TiempoCorreccionFormato(infoAd.FechaCreacion)
			v.FechaModificacion = time_bogota.TiempoBogotaFormato()
		}
		if err := models.UpdateTerceroById(&v); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Tercero
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *TerceroController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTercero(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetAllTrTerceroIdentificacion ...
// @Title GetAllTrTerceroIdentificacion
// @Description Consulta en las tablas tercero y datos_identificacion el string suministrado
// @Param	query		query	string	true	"tipo documento + numero documento + nombre completo"
// @Param	documento	query	string	true	"numero documento"
// @Success 200 {object} []models.DatosIdentificacionTercero
// @Failure 404 not found resource
// @router /identificacion [get]
func (c *TerceroController) GetAllTrTerceroIdentificacion() {

	query := c.GetString("query")
	documento := c.GetString("documento")
	var terceros = make([]models.DatosIdentificacionTercero_, 0)
	if query != "" || documento != "" {
		if err := models.GetAllDatosIdentificacionTercero(documento, query, &terceros); err != nil {
			c.Abort("500")
		}
	}

	c.Data["json"] = terceros
	c.ServeJSON()
}

// GetTrIdentificacionTercero ...
// @Title GetAllTrTerceroIdentificacion
// @Description Consulta en las tablas tercero y datos_identificacion el tercero indicado
// @Param	id	path	string	true	"tercero_id que se consulta"
// @Success 200 {object} models.DatosIdentificacionTercero
// @Failure 404 not found resource
// @router /identificacion/:id [get]
func (c *TerceroController) GetTrIdentificacionTercero() {

	var id int
	if v, err := c.GetInt(":id"); err != nil || v <= 0 {
		c.Abort("400")
	} else {
		id = v
	}

	var tercero models.DatosIdentificacionTercero_
	if err := models.GetTrIdentificacionTercero(id, &tercero); err != nil {
		c.Abort("500")
	}

	c.Data["json"] = tercero
	c.ServeJSON()
}
