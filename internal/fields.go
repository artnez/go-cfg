package internal

import (
	"reflect"
	"strings"
)

type Tag struct {
	Name   string
	Secret bool
}

type Tags map[string]Tag

type Field struct {
	Name  string
	Tag   reflect.StructTag
	Value reflect.Value
}

type Fields []Field

func NewFields(structPtr interface{}, withSecrets bool) Fields {
	fields := []Field{}
	elem := reflect.TypeOf(structPtr).Elem()
	val := reflect.ValueOf(structPtr).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		tag := parseTag(field.Tag.Get("env"))
		if !tag.Secret || withSecrets {
			value := val.Field(i)
			fields = append(fields, Field{field.Name, field.Tag, value})
		}
	}
	return fields
}

func (f Fields) Tags(tag string) Tags {
	tags := map[string]Tag{}
	for _, field := range f {
		tags[field.Name] = parseTag(field.Tag.Get(tag))
	}
	return tags
}

func parseTag(value string) Tag {
	parts := strings.Split(value, ",")
	secret := len(parts) > 1 && parts[1] == "secret"
	return Tag{parts[0], secret}
}
