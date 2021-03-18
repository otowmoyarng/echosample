package model

import (
	repo "echosample/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exp_userlist = []repo.User{
	repo.User{Name: "taro", Birthday: "2010/01/01"},
	repo.User{Name: "ken", Birthday: "2015/03/11"},
	repo.User{Name: "nana", Birthday: "2005/01/31"},
	repo.User{Name: "tomo", Birthday: "2010/11/21"},
}

func TestGetUsers(t *testing.T) {

	result_userlist := GetUsers()

	// 件数比較
	assert.Equal(t, len(exp_userlist), len(result_userlist))

	// 内容比較
	for index := 0; index < len(exp_userlist); index++ {
		assert.Equal(t, exp_userlist[index].Name, result_userlist[index].Name)
		assert.Equal(t, exp_userlist[index].Birthday, result_userlist[index].Birthday)
	}
}

func TestGetUser_found(t *testing.T) {

	condition := exp_userlist[2]
	result, isok := GetUser(condition)
	assert.Equal(t, condition.Name, result.Name)
	assert.Equal(t, condition.Birthday, result.Birthday)
	assert.Nil(t, isok)
}

func TestGetUser_notfound(t *testing.T) {

	condition := exp_userlist[2]
	condition.Name = "maki"
	result, isok := GetUser(condition)
	assert.NotEqual(t, condition.Name, result.Name)
	assert.NotEqual(t, condition.Birthday, result.Birthday)
	assert.NotNil(t, isok)
}

func TestInsertUser(t *testing.T) {

	insertuser := new(repo.User)
	insertuser.Name = "koji"
	insertuser.Birthday = "1990/12/31"
	exp_userlist = append(exp_userlist, *insertuser)

	InsertUser(*insertuser)

	result_userlist := GetUsers()

	// 内容比較
	for index := 0; index < len(exp_userlist); index++ {
		assert.Equal(t, exp_userlist[index].Name, result_userlist[index].Name)
		assert.Equal(t, exp_userlist[index].Birthday, result_userlist[index].Birthday)
	}
}

func TestUpdateUser(t *testing.T) {

	exp_userlist[4].Birthday = "1995/10/13"
	updateuser := exp_userlist[4]

	UpdateUser(updateuser)

	result_userlist := GetUsers()

	// 内容比較
	for index := 0; index < len(exp_userlist); index++ {
		assert.Equal(t, exp_userlist[index].Name, result_userlist[index].Name)
		assert.Equal(t, exp_userlist[index].Birthday, result_userlist[index].Birthday)
	}
}

func TestDeleteUser(t *testing.T) {

	deleteuser := exp_userlist[4]
	DeleteUser(deleteuser)

	exp_userlist = exp_userlist[:3]
	result_userlist := GetUsers()

	// 内容比較
	for index := 0; index < len(exp_userlist); index++ {
		assert.Equal(t, exp_userlist[index].Name, result_userlist[index].Name)
		assert.Equal(t, exp_userlist[index].Birthday, result_userlist[index].Birthday)
	}
}
