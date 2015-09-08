package internal

import "reflect"

type Tags map[string]string

type Field struct {
	Type  reflect.StructField
	Value reflect.Value
}

type Fields map[string]Field

func NewFields(structPtr interface{}) Fields {
	fields := Fields{}
	elem := reflect.TypeOf(structPtr).Elem()
	val := reflect.ValueOf(structPtr).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		value := val.Field(i)
		fields[field.Name] = Field{field, value}
	}
	return fields
}

func (f Fields) Tags(tag string) Tags {
	tags := map[string]string{}
	for name, field := range f {
		tags[name] = field.Type.Tag.Get(tag)
	}
	return tags
}
