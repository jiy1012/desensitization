# desensitization
golang版脱敏，通过结构体tag配置的方式，执行脱敏。可在api返回json时统一调用处理，减少业务逻辑代码侵入。

使用方法：
```
type TestCommonFields struct {
	Phone    string `json:"phone" desensitization:"PHONE"`
	Email    string `json:"email" desensitization:"EMAIL"`
	UserName string `json:"user_name" desensitization:"CHINESE_NAME"`
	IDCard   string `json:"id_card" desensitization:"CHINESE_IDCARD"`
}

	if err := Desensitization(&p); err != nil {
		t.Errorf("Desensitization() error = %v", err)
	}
```

自定义脱敏规则:
参考desensitizer文件夹下的规则，实现接口
```
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

```
然后再调用之前注册即可
```
RegisterDesensitizer(email.Type, email.Operator{})
```