// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
	
	"time"
)

type Slice_Vhost []*Vhost

func (s Slice_Vhost) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_Vhost) RangeRaw(fn func(m *Vhost) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// Vhost 虚拟主机
type Vhost struct {
	base    factory.Base
	objects []*Vhost
	
	Id      	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name    	string  	`db:"name" bson:"name" comment:"网站名称" json:"name" xml:"name"`
	GroupId 	uint    	`db:"group_id" bson:"group_id" comment:"组" json:"group_id" xml:"group_id"`
	Domain  	string  	`db:"domain" bson:"domain" comment:"域名" json:"domain" xml:"domain"`
	Root    	string  	`db:"root" bson:"root" comment:"网站物理路径" json:"root" xml:"root"`
	Created 	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated 	uint    	`db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
	Setting 	string  	`db:"setting" bson:"setting" comment:"设置" json:"setting" xml:"setting"`
	Disabled	string  	`db:"disabled" bson:"disabled" comment:"是否停用" json:"disabled" xml:"disabled"`
}

// - base function

func (this *Vhost) Trans() *factory.Transaction {
	return this.base.Trans()
}

func (this *Vhost) Use(trans *factory.Transaction) factory.Model {
	this.base.Use(trans)
	return this
}

func (this *Vhost) SetContext(ctx echo.Context) factory.Model {
	this.base.SetContext(ctx)
	return this
}

func (this *Vhost) Context() echo.Context {
	return this.base.Context()
}

func (this *Vhost) SetConnID(connID int) factory.Model {
	this.base.SetConnID(connID)
	return this
}

func (this *Vhost) SetNamer(namer func (string) string) factory.Model {
	this.base.SetNamer(namer)
	return this
}

func (this *Vhost) Namer() func(string) string {
	return this.base.Namer()
}

func (this *Vhost) SetParam(param *factory.Param) factory.Model {
	this.base.SetParam(param)
	return this
}

func (this *Vhost) Param() *factory.Param {
	if this.base.Param() == nil {
		return this.NewParam()
	}
	return this.base.Param()
}

// - current function

func (this *Vhost) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.base.ConnID()).Use(this.trans)
}

func (this *Vhost) Objects() []*Vhost {
	if this.bjects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *Vhost) NewObjects() factory.Ranger {
	return &Slice_Vhost{}
}

func (this *Vhost) InitObjects() *[]*Vhost {
	this.objects = []*Vhost{}
	return &this.objects
}

func (this *Vhost) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *Vhost) Short_() string {
	return "vhost"
}

func (this *Vhost) Struct_() string {
	return "Vhost"
}

func (this *Vhost) Name_() string {
	if this.base.Namer() != nil {
		return WithPrefix(this.base.Namer()(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *Vhost) CPAFrom(source factory.Model) factory.Model {
	this.SetContext(source.Context())
	this.Use(source.Trans())
	this.SetNamer(source.Namer())
	return this
}

func (this *Vhost) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := this.base
	err := this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
	this.base = base
	return err
}

func (this *Vhost) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Vhost) GroupBy(keyField string, inputRows ...[]*Vhost) map[string][]*Vhost {
	var rows []*Vhost
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*Vhost{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*Vhost{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *Vhost) KeyBy(keyField string, inputRows ...[]*Vhost) map[string]*Vhost {
	var rows []*Vhost
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*Vhost{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *Vhost) AsKV(keyField string, valueField string, inputRows ...[]*Vhost) map[string]interface{} {
	var rows []*Vhost
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *Vhost) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Vhost) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	err = DBI.Fire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", this, nil)
	}
	return
}

func (this *Vhost) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if err = DBI.Fire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", this, mw, args...)
}

func (this *Vhost) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *Vhost) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *Vhost) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	m := *this
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (this *Vhost) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
		return DBI.Fire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
		return DBI.Fire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", this, mw, args...)
		} else {
			err = DBI.Fire("created", this, nil)
		}
	} 
	return 
}

func (this *Vhost) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.Fire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", this, mw, args...)
}

func (this *Vhost) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *Vhost) Reset() *Vhost {
	this.Id = 0
	this.Name = ``
	this.GroupId = 0
	this.Domain = ``
	this.Root = ``
	this.Created = 0
	this.Updated = 0
	this.Setting = ``
	this.Disabled = ``
	return this
}

func (this *Vhost) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["GroupId"] = this.GroupId
	r["Domain"] = this.Domain
	r["Root"] = this.Root
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["Setting"] = this.Setting
	r["Disabled"] = this.Disabled
	return r
}

func (this *Vhost) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
			case "id": this.Id = param.AsUint(value)
			case "name": this.Name = param.AsString(value)
			case "group_id": this.GroupId = param.AsUint(value)
			case "domain": this.Domain = param.AsString(value)
			case "root": this.Root = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
			case "updated": this.Updated = param.AsUint(value)
			case "setting": this.Setting = param.AsString(value)
			case "disabled": this.Disabled = param.AsString(value)
		}
	}
}

func (this *Vhost) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "GroupId": this.GroupId = param.AsUint(vv)
				case "Domain": this.Domain = param.AsString(vv)
				case "Root": this.Root = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "Setting": this.Setting = param.AsString(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
			}
	}
}

func (this *Vhost) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["name"] = this.Name
	r["group_id"] = this.GroupId
	r["domain"] = this.Domain
	r["root"] = this.Root
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["setting"] = this.Setting
	r["disabled"] = this.Disabled
	return r
}

func (this *Vhost) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *Vhost) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

