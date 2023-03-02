package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var list []ToDo

type ToDo struct {
	gorm.Model
	ID       int64  `json:"id"`
	Task     string `gorm:"" json:"task"`
	Due      string `json:"due"`
	Complete bool   `json:"complete"`
}

func init() {
	//config.Connect()
	//db = config.GetDB()
	//db.AutoMigrate(&ToDo{})
}

//1:43:03
func (t *ToDo) CreateToDo() *ToDo {
	//db.NewRecord(t)
	//db.Create(&t)
	list = append(list, *t)
	return t
}

func GetAllToDo() []ToDo {
	var ToDos []ToDo
	//db.Find(&ToDos)
	ToDos = list
	return ToDos
}

func GetToDoByID(Id int64) (*ToDo, *gorm.DB) {
	var getToDo ToDo
	//db := db.Where("ID=?", Id).Find(&getToDo)
	for _, t := range list {
		if t.ID == Id {
			getToDo = t
		}
	}
	return &getToDo, db
}

func GetToDoByIDList(Id int64) (*ToDo, *[]ToDo) {
	var getToDo ToDo
	for _, t := range list {
		if t.ID == Id {
			getToDo = t
		}
	}
	return &getToDo, &list
}

func DeleteToDo(Id int64) ToDo {
	var todo ToDo
	//db.Where("ID=?", Id).Delete(todo)
	for i, t := range list {
		if t.ID == Id {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	return todo
}

func GetList() []ToDo {
	return list
}
