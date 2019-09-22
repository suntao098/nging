// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
)

// Config 配置
type Config struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*Config
	namer   func(string) string
	connID  int
	
	Key        	string  	`db:"key,pk" bson:"key" comment:"键" json:"key" xml:"key"`
	Label      	string  	`db:"label" bson:"label" comment:"选项名称" json:"label" xml:"label"`
	Description	string  	`db:"description" bson:"description" comment:"简介" json:"description" xml:"description"`
	Value      	string  	`db:"value" bson:"value" comment:"值" json:"value" xml:"value"`
	Group      	string  	`db:"group,pk" bson:"group" comment:"组" json:"group" xml:"group"`
	Type       	string  	`db:"type" bson:"type" comment:"值类型(list-以半角逗号分隔的值列表)" json:"type" xml:"type"`
	Sort       	int     	`db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Disabled   	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Encrypted  	string  	`db:"encrypted" bson:"encrypted" comment:"是否加密" json:"encrypted" xml:"encrypted"`
}

func (this *Config) Trans() *factory.Transaction {
	return this.trans
}

func (this *Config) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *Config) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *Config) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *Config) Objects() []*Config {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *Config) NewObjects() *[]*Config {
	this.objects = []*Config{}
	return &this.objects
}

func (this *Config) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *Config) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *Config) Short_() string {
	return "config"
}

func (this *Config) Struct_() string {
	return "Config"
}

func (this *Config) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *Config) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *Config) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *Config) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *Config) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Config) GroupBy(keyField string, inputRows ...[]*Config) map[string][]*Config {
	var rows []*Config
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*Config{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*Config{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *Config) KeyBy(keyField string, inputRows ...[]*Config) map[string]*Config {
	var rows []*Config
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*Config{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *Config) AsKV(keyField string, valueField string, inputRows ...[]*Config) map[string]interface{} {
	var rows []*Config
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

func (this *Config) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Config) Add() (pk interface{}, err error) {
	
	if len(this.Type) == 0 { this.Type = "text" }
	if len(this.Encrypted) == 0 { this.Encrypted = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	
	return
}

func (this *Config) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.Type) == 0 { this.Type = "text" }
	if len(this.Encrypted) == 0 { this.Encrypted = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *Config) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *Config) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *Config) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["type"] = "text" } }
	if val, ok := kvset["encrypted"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["encrypted"] = "N" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *Config) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.Type) == 0 { this.Type = "text" }
	if len(this.Encrypted) == 0 { this.Encrypted = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	},func(){
		
	if len(this.Type) == 0 { this.Type = "text" }
	if len(this.Encrypted) == 0 { this.Encrypted = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	})
	
	return 
}

func (this *Config) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *Config) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *Config) Reset() *Config {
	this.Key = ``
	this.Label = ``
	this.Description = ``
	this.Value = ``
	this.Group = ``
	this.Type = ``
	this.Sort = 0
	this.Disabled = ``
	this.Encrypted = ``
	return this
}

func (this *Config) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Key"] = this.Key
	r["Label"] = this.Label
	r["Description"] = this.Description
	r["Value"] = this.Value
	r["Group"] = this.Group
	r["Type"] = this.Type
	r["Sort"] = this.Sort
	r["Disabled"] = this.Disabled
	r["Encrypted"] = this.Encrypted
	return r
}

func (this *Config) Set(key interface{}, value ...interface{}) {
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
				case "Key": this.Key = param.AsString(vv)
				case "Label": this.Label = param.AsString(vv)
				case "Description": this.Description = param.AsString(vv)
				case "Value": this.Value = param.AsString(vv)
				case "Group": this.Group = param.AsString(vv)
				case "Type": this.Type = param.AsString(vv)
				case "Sort": this.Sort = param.AsInt(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "Encrypted": this.Encrypted = param.AsString(vv)
			}
	}
}

func (this *Config) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["key"] = this.Key
	r["label"] = this.Label
	r["description"] = this.Description
	r["value"] = this.Value
	r["group"] = this.Group
	r["type"] = this.Type
	r["sort"] = this.Sort
	r["disabled"] = this.Disabled
	r["encrypted"] = this.Encrypted
	return r
}

func (this *Config) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *Config) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

