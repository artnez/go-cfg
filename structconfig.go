package structconfig

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/artnez/structconfig/internal"
)

func FromEnviron(config interface{}, environ []string) {
	fields := internal.NewFields(config)
	tags := fields.Tags("env")
	values := internal.NewEnviron(environ)
	update := map[string]interface{}{}
	for name := range fields {
		if tag, ok := tags[name]; ok {
			if value, ok := values[tag]; ok {
				update[name] = value
			}
		}
	}
	internal.Decode(update, config)
}

func Map(config interface{}) map[string]interface{} {
	fields := internal.NewFields(config)
	result := map[string]interface{}{}
	for name, field := range fields {
		result[name] = field.Value.Interface()
	}
	return result
}

func String(config interface{}) string {
	fields := internal.NewFields(config)
	buffer := []string{}
	for name, field := range fields {
		value, _ := json.Marshal(field.Value.Interface())
		buffer = append(buffer, fmt.Sprintf("%s=%s", name, value))
	}
	return strings.Join(buffer, ", ")
}
