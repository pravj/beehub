package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Name     string `orm:"size(128)"`
	UserName string `orm:"size(128)"`
	Email    string `orm:"size(128)"`
	Token    string `orm:"size(128)"`
        Avatar   string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// CreateUser checks for an existing user's email and
// create a new user instance, if no result found.
func CreateUser(m *User) {
  o := orm.NewOrm()
  user := User{Email: m.Email}

  err := o.Read(&user, "Email")
  if err == orm.ErrNoRows {
    beego.Info("no result found")
    o.Insert(m)
  } else if err == orm.ErrMissPK {
    beego.Info("no primary key found")
  } else {
    beego.Info(user.Name)
  }
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
