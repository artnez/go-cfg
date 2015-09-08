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
	Int64Value           int64   `env:"INT64_VALUE"`
	Float32Value         float32 `env:"FLOAT32_VALUE"`
	Float64Value         float64 `env:"FLOAT64_VALUE"`
	BoolValue            bool    `env:"BOOL_VALUE"`
	BoolValueWithDefault bool    `env:"BOOL_VALUE_WITH_DEFAULT"`
}

func NewConfig() *Config {
	return &Config{
		StringValue:          "default string value",
		IntValue:             1000,
		Float64Value:         5.0,
		BoolValueWithDefault: true,
	}
}

func TestString(t *testing.T) {
	c := NewConfig()
	s := String(c)
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
}

func TestMap(t *testing.T) {
	c := NewConfig()
	m := Map(c)

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
}

func TestUpdateStringFromEnviron(t *testing.T) {
	os.Setenv("STRING_VALUE", "new string value")
	c := NewConfig()
	assert.Equal(t, "default string value", c.StringValue)
	FromEnviron(c, os.Environ())
	assert.Equal(t, "new string value", c.StringValue)
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
