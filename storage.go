package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
)

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Storage interface {
	Insert(u *User)
	Get(id uint) (User, error)
	GetAll() ([]User, error)
	Update(id uint, u User) error
	Delete(id uint) error
}

type UserStorage struct {
	Users []User `json:"users"`
	file  *os.File
	sync.Mutex
}

func NewUserStorage(file *os.File) *UserStorage {
	var users []User
	data, _ := ioutil.ReadFile(file.Name())
	if len(data) > 1 {
		_ = json.Unmarshal(data, &users)
	}
	return &UserStorage{
		Users: users,
		file:  file,
	}
}

func (us *UserStorage) Insert(u *User) {
	us.Lock()
	defer us.Unlock()

	us.Users = append(us.Users, *u)
	data, _ := json.MarshalIndent(us.Users, "", " ")
	_ = ioutil.WriteFile(us.file.Name(), data, 0644)
}

func (us UserStorage) Get(id uint) (User, error) {
	us.Lock()
	defer us.Unlock()

	for _, user := range us.Users {
		if id == user.Id {
			return user, nil
		}
	}
	return User{}, errors.New("Can't get user. Wrong id")
}

func (us UserStorage) GetAll() ([]User, error) {
	return us.Users, nil
}

func (us *UserStorage) Update(id uint, u User) error {
	us.Lock()
	defer us.Unlock()

	for i, user := range us.Users {
		if id == user.Id {
			us.Users[i] = u
			data, _ := json.MarshalIndent(us.Users, "", " ")
			_ = ioutil.WriteFile(us.file.Name(), data, 0644)
			return nil
		}
	}
	return errors.New("Can't update user. Wrong id")
}

func (us *UserStorage) Delete(id uint) error {
	us.Lock()
	defer us.Unlock()

	for i, user := range us.Users {
		if id == user.Id {
			us.Users = append(us.Users[:i], us.Users[i+1:]...)
			data, _ := json.MarshalIndent(us.Users, "", " ")
			_ = ioutil.WriteFile(us.file.Name(), data, 0644)
			return nil
		}
	}
	return errors.New("Can't delete user. Wrong id")
}
