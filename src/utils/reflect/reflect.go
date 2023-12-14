package reflect

import (
  "reflect"
)

// ================================================
func IsComparable(rawValue interface{}) bool {
  v := reflect.ValueOf(rawValue)

  switch v.Kind() {
  case reflect.Array, reflect.Slice, reflect.Map:
    return false
  default:
    return v.IsValid() && v.Type().Comparable()
  }
}
