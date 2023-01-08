package utils

import (
	"reflect"
	"strings"
	"sync"

	"gorm.io/gorm/schema"
)

func GetColumnName(e interface{}, col string) string {
	val := reflect.ValueOf(e)

	// var columnName string
OuterLoop:
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name

		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
			// fmt.Println("kosong ", fieldName, " ", t.Name)
		default:
			parts := strings.Split(jsonTag, ",")
			name := parts[0]
			if name == "" {
				name = fieldName
			}
			if col == name {
				col = t.Name
				break OuterLoop
			}
		}

	}

	s, err := schema.Parse(e, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		panic("failed to parse schema")
	}

	m := make(map[string]string)
	for _, field := range s.Fields {
		dbName := field.DBName
		modelName := field.Name
		m[modelName] = dbName
	}
	return m[col]
}
