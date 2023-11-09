package email

import "strings"

const Type = "EMAIL"

type Operator struct{}

// Desensitization
// 脱敏规则
// 以@符分隔 分为前缀和后缀
// 前缀长度为1，不显示
// 前缀长度为2，显示第一个,
// 前缀长度为3-5，显示第一个和最后一个
// 前缀长度为6以上，显示前两个和后两个
func (Operator) Desensitization(in interface{}) (out interface{}, err error) {
	inStr := in.(string)
	if inStr == "" {
		return "", nil
	}
	pos := strings.Index(inStr, "@")
	//l := len(inStr)
	perfix := inStr[:pos]
	//suffix := inStr[pos+1:]
	if pos < 2 {
		perfix = "*" + inStr[pos:]
	} else if pos < 3 {
		perfix = inStr[:1] + "*" + inStr[pos:]
	} else if pos < 6 {
		perfix = inStr[:1] + strings.Repeat("*", pos-2) + inStr[pos-1:]
	} else {
		perfix = inStr[:2] + strings.Repeat("*", pos-4) + inStr[pos-2:]
	}
	return perfix, nil
}
