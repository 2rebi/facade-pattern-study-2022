package facade

import (
	"errors"
	complexsys "facade-pattern-example"
	"fmt"
)

var (
	ErrNeedToUserRegistration = errors.New("유저 등록 필요")
	ErrNeedToGenSecretCode    = errors.New("비밀 코드 생성 필요")
	ErrWrongSecretCode        = errors.New("잘못된 비밀코드 입력")
)

type UserMobileCertificationFacade struct {
	userSystem   *complexsys.UserStore
	smsSystem    *complexsys.SMSSender
	secretSystem *complexsys.SecretCodeStore
}

func NewUserMobileCertificationFacade() *UserMobileCertificationFacade {
	return &UserMobileCertificationFacade{
		userSystem:   complexsys.NewUserStore(),
		smsSystem:    complexsys.NewSMSSender(),
		secretSystem: complexsys.NewSecretCodeStore(),
	}
}

func (f *UserMobileCertificationFacade) StartCertification(mobile string) {
	user, err := f.userSystem.FindByMobile(mobile)
	if err != nil {
		user = f.userSystem.NewUser(mobile)
	}

	code := f.secretSystem.GenCode(user)
	plain, _ := code.PlainCode()
	f.smsSystem.SendSMS(user.Mobile, plain)
	code.Lock()
}

func (f *UserMobileCertificationFacade) ConfirmCertification(mobile, secretCode string) error {
	user, err := f.userSystem.FindByMobile(mobile)
	if err != nil {
		fmt.Println(err)
		return ErrNeedToUserRegistration
	}

	code, err := f.secretSystem.Find(user)
	if err != nil {
		fmt.Println(err)
		return ErrNeedToGenSecretCode
	}

	err = code.Compare([]byte(secretCode))
	if err != nil {
		return ErrWrongSecretCode
	}

	user.Activate()

	return nil
}
