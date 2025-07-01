package mask

const Type = "MASK"

type Operator struct{}

// Desensitization
// 脱敏规则
// <5位 显示全部隐藏
// 否则 显示前三后四
func (Operator) Desensitization(in interface{}) (out interface{}, err error) {
	inStr := in.(string)
	if inStr == "" {
		return "", nil
	}
	l := len(inStr)
	if l < 5 {
		return "****", nil
	}
	if l < 12 {
		return inStr[:2] + "****" + inStr[l-2:], nil
	}
	return inStr[:4] + "****" + inStr[l-4:], nil
}
