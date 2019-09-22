package structconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	StringValue          string `env:"STRING_VALUE"`
	IntValue             int    `env:"INT_VALUE"`
	Int8Value            int8   `env:"INT8_VALUE"`
	Int16Value           int16  `env:"INT16_VALUE"`
	Int32Value           int32
	Int64Value           int64   `env:"INT64_VALUE",json:"IGNORE_THIS_FIELD"`
	Float32Value         float32 `env:"FLOAT32_VALUE"`
	Float64Value         float64 `env:"FLOAT64_VALUE"`
	BoolValue            bool    `env:"BOOL_VALUE"`
	BoolValueWithDefault bool    `env:"BOOL_VALUE_WITH_DEFAULT"`
	SecretStringValue    string  `env:"SECRET_STRING_VALUE,secret"`
}

func NewConfig() *Config {
	return &Config{
		StringValue:          "default string value",
		IntValue:             1000,
		Float64Value:         5.0,
		BoolValueWithDefault: true,
		SecretStringValue:    "SECRET!",
	}
}

func TestString(t *testing.T) {
	c := NewConfig()
	s := String(c, true)
	assert.Contains(t, s, "BoolValue=false")
	assert.Contains(t, s, "BoolValueWithDefault=true")
	assert.Contains(t, s, "IntValue=1000")
	assert.Contains(t, s, "Int8Value=0")
	assert.Contains(t, s, "Int16Value=0")
	assert.Contains(t, s, "Int32Value=0")
	assert.Contains(t, s, "Int64Value=0")
	assert.Contains(t, s, "Float32Value=0")
	assert.Contains(t, s, "Float64Value=5")
	assert.Contains(t, s, `StringValue="default string value"`)
	assert.Contains(t, s, `SecretStringValue="SECRET!"`)
	assert.NotContains(t, s, "IGNORE_THIS_FIELD")

	s = String(c, false)
	assert.NotContains(t, s, `SecretStringValue="SECRET!"`)
}

func TestMap(t *testing.T) {
	c := NewConfig()
	m := Map(c, true)

	var ok bool
	_, ok = m["BoolValue"].(bool)
	assert.True(t, ok)
	_, ok = m["IntValue"].(int)
	assert.True(t, ok)
	_, ok = m["Int8Value"].(int8)
	assert.True(t, ok)
	_, ok = m["Int16Value"].(int16)
	assert.True(t, ok)
	_, ok = m["Int32Value"].(int32)
	assert.True(t, ok)
	_, ok = m["Int64Value"].(int64)
	assert.True(t, ok)
	_, ok = m["Float32Value"].(float32)
	assert.True(t, ok)
	_, ok = m["Float64Value"].(float64)
	assert.True(t, ok)
	_, ok = m["StringValue"].(string)
	assert.True(t, ok)

	m = Map(c, false)
	_, ok = m["SecretStringValue"].(string)
	assert.True(t, !ok)
}

func TestSlice(t *testing.T) {
	c := NewConfig()
	m := Slice(c, true)

	var ok bool
	var found bool

	assert.Equal(t, "StringValue", m[0].Name)
	assert.Equal(t, "STRING_VALUE", m[0].Tag)
	_, ok = m[0].Value.(string)
	assert.True(t, ok)

	assert.Equal(t, "IntValue", m[1].Name)
	assert.Equal(t, "INT_VALUE", m[1].Tag)
	_, ok = m[1].Value.(int)
	assert.True(t, ok)

	assert.Equal(t, "Int8Value", m[2].Name)
	assert.Equal(t, "INT8_VALUE", m[2].Tag)
	_, ok = m[2].Value.(int8)
	assert.True(t, ok)

	assert.Equal(t, "Int16Value", m[3].Name)
	assert.Equal(t, "INT16_VALUE", m[3].Tag)
	_, ok = m[3].Value.(int16)
	assert.True(t, ok)

	found = false
	for _, field := range m {
		if field.Name == "SecretStringValue" {
			found = true
			break
		}
	}
	assert.True(t, found)

	m = Slice(c, false)
	found = false
	for _, field := range m {
		if field.Name == "SecretStringValue" {
			found = true
			break
		}
	}
	assert.True(t, !found)
}

func TestUpdateStringFromEnviron(t *testing.T) {
	os.Setenv("STRING_VALUE", "new string value")
	os.Setenv("SECRET_STRING_VALUE", "ssshhh")
	c := NewConfig()
	assert.Equal(t, "default string value", c.StringValue)
	assert.Equal(t, "SECRET!", c.SecretStringValue)
	FromEnviron(c, os.Environ())
	assert.Equal(t, "new string value", c.StringValue)
	assert.Equal(t, "ssshhh", c.SecretStringValue)
}

func TestUpdateIntFromEnviron(t *testing.T) {
	os.Setenv("INT_VALUE", "2000")
	os.Setenv("INT8_VALUE", "40")
	os.Setenv("INT16_VALUE", "1000")
	os.Setenv("INT32_VALUE", "100000")
	os.Setenv("INT64_VALUE", "10000000")
	c := NewConfig()
	assert.Equal(t, 1000, c.IntValue)
	assert.Equal(t, int8(0), c.Int8Value)
	assert.Equal(t, int16(0), c.Int16Value)
	assert.Equal(t, int32(0), c.Int32Value)
	assert.Equal(t, int64(0), c.Int64Value)
	FromEnviron(c, os.Environ())
	assert.Equal(t, 2000, c.IntValue)
	assert.Equal(t, int8(40), c.Int8Value)
	assert.Equal(t, int16(1000), c.Int16Value)
	assert.Equal(t, int32(0), c.Int32Value)
	assert.Equal(t, int64(10000000), c.Int64Value)
}

func TestUpdateFloatFromEnviron(t *testing.T) {
	os.Setenv("FLOAT32_VALUE", "1000.5")
	os.Setenv("FLOAT64_VALUE", "10000.555")
	c := NewConfig()
	assert.Equal(t, float32(0), c.Float32Value)
	assert.Equal(t, float64(5.0), c.Float64Value)
	FromEnviron(c, os.Environ())
	assert.Equal(t, float32(1000.5), c.Float32Value)
	assert.Equal(t, float64(10000.555), c.Float64Value)
}

func TestUpdateBoolNumberFromEnviron(t *testing.T) {
	os.Setenv("BOOL_VALUE", "1")
	os.Setenv("BOOL_VALUE_WITH_DEFAULT", "0")
	c := NewConfig()
	assert.Equal(t, false, c.BoolValue)
	assert.Equal(t, true, c.BoolValueWithDefault)
	FromEnviron(c, os.Environ())
	assert.Equal(t, true, c.BoolValue)
	assert.Equal(t, false, c.BoolValueWithDefault)
}

func TestUpdateBoolStringFromEnviron(t *testing.T) {
	os.Setenv("BOOL_VALUE", "true")
	os.Setenv("BOOL_VALUE_WITH_DEFAULT", "false")
	c := NewConfig()
	assert.Equal(t, false, c.BoolValue)
	assert.Equal(t, true, c.BoolValueWithDefault)
	FromEnviron(c, os.Environ())
	assert.Equal(t, true, c.BoolValue)
	assert.Equal(t, false, c.BoolValueWithDefault)
}

func TestUpdateBoolEmptyStringFromEnviron(t *testing.T) {
	os.Setenv("BOOL_VALUE", "")
	c := NewConfig()
	assert.Equal(t, false, c.BoolValue)
	FromEnviron(c, os.Environ())
	assert.Equal(t, true, c.BoolValue)
}
