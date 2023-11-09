package desensitization

import (
	"desensitization/desensitizer/chinese_idcard"
	"desensitization/desensitizer/chinese_name"
	"desensitization/desensitizer/email"
	"desensitization/desensitizer/phone"
	"sync"
)

func init() {
	RegisterDesensitizer(email.Type, email.Operator{})
	RegisterDesensitizer(phone.Type, phone.Operator{})
	RegisterDesensitizer(chinese_name.Type, chinese_name.Operator{})
	RegisterDesensitizer(chinese_idcard.Type, chinese_idcard.Operator{})
}

var desensitizers map[string]Desensitizer
var mu sync.RWMutex

type Desensitizer interface {
	Desensitization(in interface{}) (out interface{}, err error)
}

func RegisterDesensitizer(desensitizerType string, desensitizer Desensitizer) error {
	mu.Lock()
	defer mu.Unlock()
	if desensitizers == nil {
		desensitizers = make(map[string]Desensitizer)
	}
	desensitizers[desensitizerType] = desensitizer
	return nil
}

func OperateByRule(desensitizerType string, in interface{}) (interface{}, error) {
	mu.RLock()
	operator, ok := desensitizers[desensitizerType]
	mu.RUnlock()
	if !ok {
		return nil, ErrTypeNotFound
	}
	return operator.Desensitization(in)
}
