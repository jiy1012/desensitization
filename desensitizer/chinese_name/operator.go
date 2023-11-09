package chinese_name

import "strings"

const Type = "CHINESE_NAME"

type Operator struct{}

// Desensitization
// 脱敏规则
// 2个字 显示最后一个字
// 3-5个字 显示首尾各一个字
// 6个字及以上显示首尾各两个字
func (Operator) Desensitization(in interface{}) (out interface{}, err error) {
	inStr := in.(string)
	if inStr == "" {
		return "", nil
	}
	l := len(inStr)
	if l < 3*3 {
		return "*" + inStr[l-3:], nil
	} else if l < 3*6 {
		return inStr[:3] + strings.Repeat("*", l/3-2) + inStr[l-3:], nil
	}
	return inStr[:6] + strings.Repeat("*", l/3-4) + inStr[l-6:], nil
}
