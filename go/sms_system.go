package complexsys

import "fmt"

type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) SendSMS(mobile, message string) {
	fmt.Println("메시지 전송", mobile)
	fmt.Println("메세지 전송 내용", message)
	fmt.Println("메세지 전송 완료", mobile)
}
