// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

// Command 指令集
type Command struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*Command
	namer   func(string) string
	connID  int
	
	Id            	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name          	string  	`db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Description   	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Command       	string  	`db:"command" bson:"command" comment:"命令" json:"command" xml:"command"`
	WorkDirectory 	string  	`db:"work_directory" bson:"work_directory" comment:"工作目录" json:"work_directory" xml:"work_directory"`
	Env           	string  	`db:"env" bson:"env" comment:"环境变量(一行一个，格式为：var1=val1)" json:"env" xml:"env"`
	Created       	uint    	`db:"created" bson:"created" comment:"添加时间" json:"created" xml:"created"`
	Updated       	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled      	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Remote        	string  	`db:"remote" bson:"remote" comment:"是否执行远程SSH命令" json:"remote" xml:"remote"`
	SshAccountId  	uint    	`db:"ssh_account_id" bson:"ssh_account_id" comment:"SSH账号ID" json:"ssh_account_id" xml:"ssh_account_id"`
}

func (this *Command) Trans() *factory.Transaction {
	return this.trans
}

func (this *Command) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *Command) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *Command) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *Command) Objects() []*Command {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *Command) NewObjects() *[]*Command {
	this.objects = []*Command{}
	return &this.objects
}

func (this *Command) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *Command) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *Command) Short_() string {
	return "command"
}

func (this *Command) Struct_() string {
	return "Command"
}

func (this *Command) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *Command) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *Command) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *Command) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *Command) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Command) GroupBy(keyField string, inputRows ...[]*Command) map[string][]*Command {
	var rows []*Command
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*Command{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*Command{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *Command) KeyBy(keyField string, inputRows ...[]*Command) map[string]*Command {
	var rows []*Command
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*Command{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *Command) AsKV(keyField string, valueField string, inputRows ...[]*Command) map[string]interface{} {
	var rows []*Command
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

func (this *Command) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Command) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Remote) == 0 { this.Remote = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *Command) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Remote) == 0 { this.Remote = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *Command) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *Command) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *Command) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["remote"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["remote"] = "N" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *Command) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Remote) == 0 { this.Remote = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Remote) == 0 { this.Remote = "N" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *Command) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *Command) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *Command) Reset() *Command {
	this.Id = 0
	this.Name = ``
	this.Description = ``
	this.Command = ``
	this.WorkDirectory = ``
	this.Env = ``
	this.Created = 0
	this.Updated = 0
	this.Disabled = ``
	this.Remote = ``
	this.SshAccountId = 0
	return this
}

func (this *Command) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["Description"] = this.Description
	r["Command"] = this.Command
	r["WorkDirectory"] = this.WorkDirectory
	r["Env"] = this.Env
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["Disabled"] = this.Disabled
	r["Remote"] = this.Remote
	r["SshAccountId"] = this.SshAccountId
	return r
}

func (this *Command) Set(key interface{}, value ...interface{}) {
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
				case "Description": this.Description = param.AsString(vv)
				case "Command": this.Command = param.AsString(vv)
				case "WorkDirectory": this.WorkDirectory = param.AsString(vv)
				case "Env": this.Env = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "Remote": this.Remote = param.AsString(vv)
				case "SshAccountId": this.SshAccountId = param.AsUint(vv)
			}
	}
}

func (this *Command) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["name"] = this.Name
	r["description"] = this.Description
	r["command"] = this.Command
	r["work_directory"] = this.WorkDirectory
	r["env"] = this.Env
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["disabled"] = this.Disabled
	r["remote"] = this.Remote
	r["ssh_account_id"] = this.SshAccountId
	return r
}

func (this *Command) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *Command) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

