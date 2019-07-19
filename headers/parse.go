package headers


import (
     "reflect"
     "strconv"
     "strings"
)

func ParseHeader(headerValue string, H interface{}){
     
     valueMap := make(map[string]string)
     parts := strings.Split(headerValue,",")
     for _, val := range parts {
          val = strings.TrimSpace(val)
          if len(val) > 0 {
               if strings.Contains(val,"=") {
                    part := val[:strings.Index(val,"=")]
                    subval := val[strings.Index(val,"=")+1:]
                    valueMap[part] = subval
               } else {
                    valueMap[val] = "true"
               }
          }
     }
     parseType(valueMap, H)
}

func getBaseType(t reflect.Type) reflect.Type {
     
     switch t.Kind() {
     case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
          return getBaseType(t.Elem())
     
     default: // is struct
          return t
     }
}

func generateFieldMap(T reflect.Type) map[string]string {
     
     out := make(map[string]string)
     T = getBaseType(T)
     if T.Kind() == reflect.Struct {
          for i:= 0; i< T.NumField(); i++ {
               if val, ok := T.Field(i).Tag.Lookup("field"); ok {
                    out[strings.ToUpper(val)] = T.Field(i).Name
               } else {
                    out[strings.ToUpper(T.Field(i).Name)] = T.Field(i).Name
               }
          }
     }
     return out
}

func parseType(valueMap map[string]string, T interface{}) {
     
     vo := reflect.ValueOf(T)
     fieldMap := generateFieldMap(vo.Type())
     
     s := vo.Elem()
     
     if s.Kind() == reflect.Struct {
          
          for key, val := range valueMap {
               fieldName, exists := fieldMap[strings.ToUpper(key)]
               f := s.FieldByName(fieldName)
               
               if exists && f.IsValid() {
                    // A Value can be changed only if it is
                    // addressable and was not obtained by
                    // the use of unexported struct fields.
                    if f.CanSet() {
                         
                         switch f.Kind() {
                         // change value of N
                         case reflect.Int:
                              {
                                   if x, err := strconv.ParseInt(val,10,64); err == nil {
                                        if !f.OverflowInt(x) {
                                             f.SetInt(x)
                                        }
                                   }
                              }
                         case reflect.Bool:
                              {
                                   if x, err := strconv.ParseBool(val); err == nil {
                                        f.SetBool(x)
                                   }
                              }
                         default:
                              {
                                   f.SetString(val)
                              }
                         }
                    }
               }
               
          }
     }
}