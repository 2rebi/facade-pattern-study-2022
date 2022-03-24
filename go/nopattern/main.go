package main

import (
	complexsys "facade-pattern-example"
	"fmt"
)

func main() {
	inputMobile := "01012345678"
	secretCode := "1234"

	userSystem := complexsys.NewUserStore()
	secretSystem := complexsys.NewSecretCodeStore()
	smsSystem := complexsys.NewSMSSender()

	// create routine
	{
		user, err := userSystem.FindByMobile(inputMobile)
		if err != nil {
			user = userSystem.NewUser(inputMobile)
		}

		code := secretSystem.GenCode(user)
		plain, _ := code.PlainCode()
		smsSystem.SendSMS(user.Mobile, plain)
		code.Lock()
	}

	// confirm routine
	{
		user, err := userSystem.FindByMobile(inputMobile)
		if err != nil {
			fmt.Println(err)
			fmt.Println("유저 등록 필요")
			return
		}

		code, err := secretSystem.Find(user)
		if err != nil {
			fmt.Println(err)
			fmt.Println("비밀 코드 생성 필요")
			return
		}

		err = code.Compare([]byte(secretCode))
		if err != nil {
			fmt.Println("잘못된 비밀코드 입력")
			return
		}

		user.Activate()
		fmt.Println("비밀 코드 입력 성공 및 유저 활성화 성공")
	}
}
