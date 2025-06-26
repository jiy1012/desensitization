package desensitization

import (
	"github.com/jiy1012/desensitization/desensitizer/chinese_idcard"
	"github.com/jiy1012/desensitization/desensitizer/chinese_name"
	"github.com/jiy1012/desensitization/desensitizer/email"
	"github.com/jiy1012/desensitization/desensitizer/empty"
	"github.com/jiy1012/desensitization/desensitizer/mask"
	"github.com/jiy1012/desensitization/desensitizer/phone"
	"sync"
)

func init() {
	_ = RegisterDesensitizer(email.Type, email.Operator{})
	_ = RegisterDesensitizer(phone.Type, phone.Operator{})
	_ = RegisterDesensitizer(chinese_name.Type, chinese_name.Operator{})
	_ = RegisterDesensitizer(chinese_idcard.Type, chinese_idcard.Operator{})
	_ = RegisterDesensitizer(empty.Type, empty.Operator{})
	_ = RegisterDesensitizer(mask.Type, mask.Operator{})
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
