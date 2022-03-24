package complexsys

import (
	"errors"
)

type User struct {
	Id     int64
	active bool
	Mobile string
}

type UserStore struct {
	autoIncrementId int64
	store           map[int64]*User
	mobileIdx       map[string]int64
}

func NewUserStore() *UserStore {
	return &UserStore{
		autoIncrementId: 0,
		store:           make(map[int64]*User),
		mobileIdx:       make(map[string]int64),
	}
}

func (s *UserStore) Find(id int64) (*User, error) {
	u, ok := s.store[id]
	if !ok {
		return nil, errors.New("찾을 수 없는 유저")
	}

	return u, nil
}

func (s *UserStore) FindByMobile(mobile string) (*User, error) {
	id, ok := s.mobileIdx[mobile]
	if !ok {
		return nil, errors.New("휴대폰 번호로 찾을 수 없는 유저")
	}

	return s.Find(id)
}

func (s *UserStore) NewUser(mobile string) (u *User) {
	s.autoIncrementId++
	u = &User{
		Id:     s.autoIncrementId,
		active: false,
		Mobile: mobile,
	}
	s.store[u.Id] = u
	s.mobileIdx[u.Mobile] = u.Id
	return u
}

func (s *User) Activate() {
	s.active = true
}
