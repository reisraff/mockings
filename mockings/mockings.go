package mockings

import "testing"
import "reflect"
import "fmt"

var calls map[string]interface{}
var returns map[string]interface{}

func AddCall(_struct interface{}, method string, with []interface{}) []interface{} {
    if v, ok := calls[reflect.TypeOf(_struct).String()]; ok {
        if method_calls, ok := v.(map[string]interface{})[method]; ok {
            method_calls_val := method_calls.([]interface{})
            calls[reflect.TypeOf(_struct).String()].(map[string]interface{})[method] = append(method_calls_val, with)
        } else {
            v.(map[string]interface{})[method] = []interface{}{with}
        }
    } else {
        calls[reflect.TypeOf(_struct).String()] = make(map[string]interface{}, 0)
        calls[reflect.TypeOf(_struct).String()].(map[string]interface{})[method] = []interface{}{with}
    }

    return getReturn(_struct, method, with)
}

func Reset() {
    calls = make(map[string]interface{}, 0)
    returns = make(map[string]interface{}, 0)
}

func Print() {
    fmt.Printf("%v\n", calls)
    fmt.Printf("%v\n", returns)
}

func AssertCalledWith(t *testing.T, _struct interface{}, method string, with []interface{}) {
    if s, ok := calls[reflect.TypeOf(_struct).String()]; ok {
        if m, ok := s.(map[string]interface{})[method]; ok {
            for _, callargs := range m.([]interface{}) {
                if len(callargs.([]interface{})) != len(with) {
                    continue
                }

                if reflect.DeepEqual(with, callargs.([]interface{})) {
                    return
                }
            }
        }
    }

    t.Errorf("%s::%s was not called with expected args '%v'", reflect.TypeOf(_struct).String(), method, with)
}

func GetErrorOrNil(_error interface{}) error {
    var result error
    if _error != nil {
        result = _error.(error)
    } else {
        result = nil
    }

    return result
}

func getReturn(_struct interface{}, method string, with []interface{}) []interface{} {
    if s, ok := returns[reflect.TypeOf(_struct).String()]; ok {
        if m, ok := s.(map[string]interface{})[method]; ok {
            for _, call := range m.([]interface{}) {
                if reflect.DeepEqual(with, call.(map[string]interface{})["with"].([]interface{})) {
                    return call.(map[string]interface{})["return"].([]interface{})
                }
            }
        }
    }

    return []interface{}{}
}

func Mock(_struct interface{}, method string, with []interface{}, _return []interface{}) {
    if v, ok := returns[reflect.TypeOf(_struct).String()]; ok {
        if method_returns, ok := v.(map[string]interface{})[method]; ok {
            method_returns_val := method_returns.([]interface{})
            returns[reflect.TypeOf(_struct).String()].(map[string]interface{})[method] = append(method_returns_val, map[string]interface{}{"with": with, "return": _return})
        } else {
            v.(map[string]interface{})[method] = []interface{}{map[string]interface{}{"with": with, "return": _return}}
        }
    } else {
        returns[reflect.TypeOf(_struct).String()] = make(map[string]interface{}, 0)
        returns[reflect.TypeOf(_struct).String()].(map[string]interface{})[method] = []interface{}{map[string]interface{}{"with": with, "return": _return}}
    }
}
