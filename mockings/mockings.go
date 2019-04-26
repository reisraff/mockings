package mockings

import "testing"
import "reflect"
import "fmt"

var calls map[string]interface{}
var returns map[string]interface{}

func AddCall(_struct interface{}, method string, with []interface{}) []interface{} {
    ptraddr := fmt.Sprintf("%p", _struct)

    if v, ok := calls[ptraddr]; ok {
        if method_calls, ok := v.(map[string]interface{})[method]; ok {
            method_calls_val := method_calls.([]interface{})
            calls[ptraddr].(map[string]interface{})[method] = append(method_calls_val, with)
        } else {
            v.(map[string]interface{})[method] = []interface{}{with}
        }
    } else {
        calls[ptraddr] = make(map[string]interface{}, 0)
        calls[ptraddr].(map[string]interface{})[method] = []interface{}{with}
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
    ptraddr := fmt.Sprintf("%p", _struct)

    if s, ok := calls[ptraddr]; ok {
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

    t.Errorf("%s(%s)::%s was not called with expected args '%v'", reflect.TypeOf(_struct).String(), ptraddr, method, with)
}

func ErrorOrNil(_error interface{}) error {
    var result error
    if _error != nil {
        result = _error.(error)
    } else {
        result = nil
    }

    return result
}

func getReturn(_struct interface{}, method string, with []interface{}) []interface{} {
    ptraddr := fmt.Sprintf("%p", _struct)

    if s, ok := returns[ptraddr]; ok {
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
    ptraddr := fmt.Sprintf("%p", _struct)

    if v, ok := returns[ptraddr]; ok {
        if method_returns, ok := v.(map[string]interface{})[method]; ok {
            method_returns_val := method_returns.([]interface{})
            returns[ptraddr].(map[string]interface{})[method] = append(method_returns_val, map[string]interface{}{"with": with, "return": _return})
        } else {
            v.(map[string]interface{})[method] = []interface{}{map[string]interface{}{"with": with, "return": _return}}
        }
    } else {
        returns[ptraddr] = make(map[string]interface{}, 0)
        returns[ptraddr].(map[string]interface{})[method] = []interface{}{map[string]interface{}{"with": with, "return": _return}}
    }
}
