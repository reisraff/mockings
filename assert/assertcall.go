package assert

import "testing"
import "reflect"


var asserts map[string]interface{}

func AddAssert(_struct interface{}, method string, with []interface{}){
    if v, ok := asserts[reflect.TypeOf(_struct).String()]; ok {
        v.(map[string]interface{})[method] = append(v.(map[string]interface{})[method].([]interface{}), with)
    } else {
        asserts[reflect.TypeOf(_struct).String()] = make(map[string]interface{}, 0)
        asserts[reflect.TypeOf(_struct).String()].(map[string]interface{})[method] = []interface{}{with}
    }
}

func ResetAsserts() {
    asserts = make(map[string]interface{}, 0)
}

func AssertCalledWith(t *testing.T, _struct interface{}, method string, with []interface{}) {
    if s, ok := asserts[reflect.TypeOf(_struct).String()]; ok {
        if m, ok := s.(map[string]interface{})[method]; ok {
            for _, callargs := range m.([]interface{}) {
                if len(callargs.([]interface{})) == len(with) {
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
        return
    }

    t.Errorf("%s::%s was not called with expected args '%v'", reflect.TypeOf(_struct).String(), method, with)
}