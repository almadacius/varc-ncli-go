package obj

import (

)

type JsonData map[string] interface{}

// ================================================
func EmptyMap() JsonData {
  out := make(JsonData)
  return out
}
