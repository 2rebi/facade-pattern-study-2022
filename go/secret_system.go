package complexsys

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type SecretCodeStore struct {
	store map[int64]*SecretCode
}

func NewSecretCodeStore() *SecretCodeStore {
	return &SecretCodeStore{store: make(map[int64]*SecretCode)}
}

func (s *SecretCodeStore) GenCode(u *User) *SecretCode {
	code := genCode()
	s.store[u.Id] = code
	return code
}

func (s *SecretCodeStore) Find(u *User) (*SecretCode, error) {
	code, ok := s.store[u.Id]
	if !ok {
		return nil, errors.New("시크릿 코드 없음")
	}

	return code, nil
}

type SecretCode struct {
	plainCode *string
	cryptCode []byte
}

func genCode() *SecretCode {
	plainCode := fmt.Sprintf("%04d", 1234) //rand.Int()%10000)
	fmt.Println("[유사로그] 코드 생성 :", plainCode)

	cryptCode, _ := bcrypt.GenerateFromPassword([]byte(plainCode), bcrypt.DefaultCost)

	return &SecretCode{
		plainCode: &plainCode,
		cryptCode: cryptCode,
	}
}

func (s *SecretCode) Lock() {
	s.plainCode = nil
}

func (s *SecretCode) PlainCode() (string, error) {
	if s.plainCode == nil {
		return "", errors.New("잠겨진 코드임 접근 불가함")
	}

	return *s.plainCode, nil
}

func (s *SecretCode) Compare(plain []byte) error {
	return bcrypt.CompareHashAndPassword(s.cryptCode, plain)
}
