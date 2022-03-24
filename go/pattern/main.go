package main

import (
	"facade-pattern-example/pattern/facade"
	"fmt"
)

func main() {
	f := facade.NewUserMobileCertificationFacade()
	f.StartCertification("01012345678")
	err := f.ConfirmCertification("01012345678", "1234")
	if err == nil {
		fmt.Println("비밀 코드 입력 성공 및 유저 활성화 성공")
	}
}
