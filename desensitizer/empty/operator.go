package empty

const Type = "EMPTY"

type Operator struct{}

// Desensitization
// 脱敏规则
// 全部清空
func (Operator) Desensitization(in interface{}) (out interface{}, err error) {
	return "", nil
}
