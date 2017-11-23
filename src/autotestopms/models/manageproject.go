package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//一个结构作为mysql映射的对象
//beego本身会有一个对应的字段名对应关系，不过一般不要用默认的，容易出错，可以显示的说明对应关系
type Projects struct{
	Id int64 `orm:"column(ID)"`
	Projectname string `orm:"column(PROJECT_NAME)"`
	Projectdirector string `orm:"column(PROJECT_DIRECTOR)"`
	Projectmembers string `orm:"column(PROJECT_MEMBERS)"`
	Projectbriefing string `orm:"column(PROJECT_BRIEFING)"`
	Projecttime time.Time `orm:"column(PROJECT_TIME)"`
	Projectcomment string `orm:"column(PROJECT_COMMENT)"`
	Projectstate string `orm:"column(PROJECT_STATE)"`
	Createtime time.Time `orm:"column(CREATE_TIME)"`
	Edittime time.Time `orm:"column(EDIT_TIME)"`
}
//对应数据表的名字，系统自动调用这个函数
func (this *Projects) TableName() string{
	return "TB_MANAGE_PROJECT"
}

//注册模型
func init(){
	orm.RegisterModel(new(Projects))
}

//增加数据表的一行,测试一下数据库
func AddProject(project Projects) error{
	o := orm.NewOrm()
	pro := new(Projects)

	pro.Id = project.Id
	pro.Projectname = project.Projectname
	pro.Projectdirector = project.Projectdirector
	pro.Projectmembers = project.Projectmembers
	pro.Projectbriefing = project.Projectbriefing
	pro.Projecttime = project.Projecttime
	pro.Projectcomment = project.Projectcomment
	pro.Projectstate = project.Projectstate
	//这两个时间是插入数据库的时候需要修改的
	pro.Createtime = time.Now()
	pro.Edittime = time.Now()
	_, err := o.Insert(pro)
	return err
}


