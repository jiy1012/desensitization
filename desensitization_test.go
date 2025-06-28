package desensitization

import (
	"fmt"
	"testing"
)

type TestCommonFields struct {
	Phone    string `json:"phone" desensitization:"PHONE"`
	Email    string `json:"email" desensitization:"EMAIL"`
	UserName string `json:"user_name" desensitization:"CHINESE_NAME"`
	IDCard   string `json:"id_card" desensitization:"CHINESE_IDCARD"`
}

type TestArrayFields struct {
	Phones []string `json:"phones" desensitization:"PHONE"`
	Emails []string `json:"emails" desensitization:"EMAIL"`
}

type TestStructFields struct {
	PhoneAndEmail TestCommonFields `json:"phone_and_email"`
}
type TestStructArrayFields struct {
	PhoneAndEmails []TestCommonFields `json:"phone_and_emails"`
}

func TestDesensitizationCommonFields(t *testing.T) {
	p := TestCommonFields{
		Phone:    "11111111111",
		Email:    "example@example.com",
		UserName: "刘亿",
	}
	fmt.Println(p)
	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
	fmt.Println(p)
}

func TestDesensitizationArrayFields(t *testing.T) {
	p := TestArrayFields{
		Phones: []string{"11111111111", "22222222222"},
		Emails: []string{"example@example.com", "example1@example1.com"},
	}
	fmt.Println(p)
	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
	fmt.Println(p)
}
func TestDesensitizationStructFields(t *testing.T) {
	p := TestStructFields{PhoneAndEmail: TestCommonFields{
		Phone: "11111111111",
		Email: "example@example.com",
	}}
	fmt.Println(p)
	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
	fmt.Println(p)
}

func TestDesensitizationStructArrayFields(t *testing.T) {
	p := TestStructArrayFields{
		[]TestCommonFields{
			{
				Phone:    "11111111111",
				Email:    "e@example.com",
				UserName: "刘亿",
				IDCard:   "111111199911091121",
			}, {
				Phone:    "22222222222",
				Email:    "ex@example2.com",
				UserName: "刘千万",
				IDCard:   "222222199911091121",
			}, {
				Phone:    "33333333333",
				Email:    "exa@example3.com",
				UserName: "刘一百块",
				IDCard:   "333333199911091121",
			}, {
				Phone:    "44444444444",
				Email:    "exam@example4.com",
				UserName: "刘不到一百",
				IDCard:   "444444199911091121",
			}, {
				Phone:    "55555555555",
				Email:    "examp@example5.com",
				UserName: "刘不到几十块",
				IDCard:   "555555199911091121",
			}, {
				Phone:    "66666666666",
				Email:    "exampl@example6.com",
				UserName: "刘只有几毛几分",
				IDCard:   "666666199911091121",
			}, {
				Phone:    "77777777777",
				Email:    "example@example7.com",
				UserName: "刘连个屁都没有",
				IDCard:   "777777199911091121",
			}}}
	fmt.Println(p)
	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
	fmt.Println(p)
}
func TestDesensitizationArrayStructArrayFields(t *testing.T) {
	p := []TestStructArrayFields{{
		[]TestCommonFields{
			{
				Phone: "11111111111",
				Email: "example1@example1.com",
			}, {
				Phone: "22222222222",
				Email: "example2@example2.com",
			}}}, {
		[]TestCommonFields{
			{
				Phone: "33333333333",
				Email: "example3@example3.com",
			}, {
				Phone: "44444444444",
				Email: "example4@example4.com",
			}}}, {
		[]TestCommonFields{
			{
				Phone: "55555555555",
				Email: "example5@example5.com",
			}, {
				Phone: "66666666666",
				Email: "example6@example6.com",
			}}}}
	fmt.Println(p)
	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
	fmt.Println(p)
}

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg" desensitization:"MASK"`
	Message string      `json:"message" desensitization:"EMPTY"`
	Data    interface{} `json:"data"`
}

func TestDesensitizationStructInterfaceArrayFields(t *testing.T) {
	p := Response{
		Code:    0,
		Msg:     "msg is ok",
		Message: "this is a message",
		Data:    nil,
	}
	var pData []TestCommonFields
	pData = append(pData, TestCommonFields{
		Phone: "11111111111",
		Email: "example1@example1.com",
	})
	pData = append(pData, TestCommonFields{
		Phone: "22222222222",
		Email: "example2@example2.com",
	})
	p.Data = pData
	fmt.Println(p)
	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
	fmt.Println(p)
}

func TestDesensitizationStructInterfaceStructFields(t *testing.T) {
	p := Response{
		Code:    0,
		Msg:     "msg is ok",
		Message: "this is a message",
		Data:    nil,
	}
	var pData TestCommonFields
	pData = TestCommonFields{
		Phone: "11111111111",
		Email: "example1@example1.com",
	}
	p.Data = pData
	fmt.Println(p)
	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
	fmt.Println(p)
}
