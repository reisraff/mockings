package assert

import "testing"
import "reflect"

var calls map[string]interface{}

func AddCall(_struct interface{}, method string, with []interface{}){
    if v, ok := calls[reflect.TypeOf(_struct).String()]; ok {
        v.(map[string]interface{})[method] = append(v.(map[string]interface{})[method].([]interface{}), with)
    } else {
        calls[reflect.TypeOf(_struct).String()] = make(map[string]interface{}, 0)
        calls[reflect.TypeOf(_struct).String()].(map[string]interface{})[method] = []interface{}{with}
    }
}

func ResetAsserts() {
    calls = make(map[string]interface{}, 0)
}

func AssertCalledWith(t *testing.T, _struct interface{}, method string, with []interface{}) {
    if s, ok := calls[reflect.TypeOf(_struct).String()]; ok {
        if m, ok := s.(map[string]interface{})[method]; ok {
            for _, callargs := range m.([]interface{}) {
                if len(callargs.([]interface{})) != len(with) {
                    continue
                }

                match := true
                for i, expectedarg := range with {
                    if callargs.([]interface{})[i] != expectedarg {
                        match = false
                    }
                }

                if match {
                    return
                }
            }
        }
    }

    t.Errorf("%s::%s was not called with expected args '%v'", reflect.TypeOf(_struct).String(), method, with)
}

