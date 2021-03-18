package model

import (
	repo "echosample/repository"
	"fmt"
)

var (
	userlist = []repo.User{
		repo.User{Name: "taro", Birthday: "2010/01/01"},
		repo.User{Name: "ken", Birthday: "2015/03/11"},
		repo.User{Name: "nana", Birthday: "2005/01/31"},
		repo.User{Name: "tomo", Birthday: "2010/11/21"},
	}
)

func GetUsers() []repo.User {
	return userlist
}

func GetUser(condition repo.User) (repo.User, error) {

	user := new(repo.User)

	for _, value := range userlist {
		if isSameUser(condition, value) {
			return value, nil
		}
	}
	return *user, fmt.Errorf("Couldn't find")
}

func InsertUser(insertuser repo.User) {
	userlist = append(userlist, insertuser)
}

func UpdateUser(updateuser repo.User) {

	for index, value := range userlist {
		if isSameUser(updateuser, value) {
			userlist[index] = updateuser
		}
	}
}

func DeleteUser(deleteuser repo.User) {

	newuserlist := []repo.User{}

	for _, value := range userlist {
		if !isSameUser(deleteuser, value) {
			newuserlist = append(newuserlist, value)
		}
	}
	userlist = newuserlist
}

func isSameUser(user1, user2 repo.User) bool {
	return user1.Name == user2.Name
}
