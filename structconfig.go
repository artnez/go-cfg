package structconfig

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/artnez/structconfig/v2/internal"
)

type Field struct {
	Name  string
	Value interface{}
	Tag   string
}

func FromEnviron(config interface{}, environ []string) {
	fields := internal.NewFields(config, true)
	tags := fields.Tags("env")
	values := internal.NewEnviron(environ)
	update := map[string]interface{}{}
	for _, field := range fields {
		if tag, ok := tags[field.Name]; ok {
			if value, ok := values[tag.Name]; ok {
				update[field.Name] = value
			}
		}
	}
	internal.Decode(update, config)
}

func Slice(config interface{}, withSecrets bool) []Field {
	fields := internal.NewFields(config, withSecrets)
	result := []Field{}
	for _, field := range fields {
		result = append(result, Field{field.Name, field.Value.Interface(), field.Tag.Get("env")})
	}
	return result
}

func Map(config interface{}, withSecrets bool) map[string]interface{} {
	fields := internal.NewFields(config, withSecrets)
	result := map[string]interface{}{}
	for _, field := range fields {
		result[field.Name] = field.Value.Interface()
	}
	return result
}

func String(config interface{}, withSecrets bool) string {
	fields := internal.NewFields(config, withSecrets)
	buffer := []string{}
	for _, field := range fields {
		value, _ := json.Marshal(field.Value.Interface())
		buffer = append(buffer, fmt.Sprintf("%s=%s", field.Name, value))
	}
	return strings.Join(buffer, ", ")
}
