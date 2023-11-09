package chinese_idcard

import "strings"

const Type = "CHINESE_IDCARD"

type Operator struct{}

// Desensitization
// 脱敏规则
// 15位 显示前三后四
// 18位 显示前三后四
func (Operator) Desensitization(in interface{}) (out interface{}, err error) {
	inStr := in.(string)
	if inStr == "" {
		return "", nil
	}
	l := len(inStr)
	if l == 15 {
		return inStr[:3] + strings.Repeat("*", l-3-4) + inStr[l-4:], nil
	} else if l == 18 {
		return inStr[:3] + strings.Repeat("*", l-3-4) + inStr[l-4:], nil
	}
	return inStr, nil
}
