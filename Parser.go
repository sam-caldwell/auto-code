package arguments

import (
	"fmt"
	"regexp"
)

// Parser - A configurable multi-source configuration parser solution.
//
// The Parser object can be configured to parse various data sources and generate a Configuration object.
// The sources may include files, environment variables, command-line arguments or even a secrets manager.
//
// Parser must define the parameters of a program which will be compiled into the final configuration state.
// This information is stored in a map.
type Parser struct {
	programName        string
	programDescription string
	programCopyright   string
	data               map[ParameterName]ParameterList
	err                error
}

// Parameter - Create a new Parameter definition and return a Parser object (by reference).
func (parser *Parser) Parameter(name, help string, pDef ...*ParameterDefinition) *Parser {
	//ToDo: define the parameter
	return parser
}

// Value - This struct holds the default (and later final) value and validator object for a given object.
//
//	The validator is used against all data sources.
type Value struct {
	/*
	 * An Argument is a value and its validator
	 */
	value     any // default value at first, overwritten with final value.
	validator *Validator
}

// String - Creates a string value.
func String(defaultValue string, validator *Validator) *Value {
	return &Value{
		value:     defaultValue,
		validator: validator,
	}
}

// Int - Creates a signed integer value
func Int(defaultValue int, validator *Validator) *Value {
	return &Value{
		value:     defaultValue,
		validator: validator,
	}
}

// Int64 - Creates a signed 64-bit integer value
func Int64(defaultValue int64, validator *Validator) *Value {
	return &Value{
		value:     defaultValue,
		validator: validator,
	}
}

// Float64 - Creates a floating-point value
func Float64(defaultValue float64, validator *Validator) *Value {
	return &Value{
		value:     defaultValue,
		validator: validator,
	}
}

// Uint - Creates an unsigned integer value
func Uint(defaultValue uint, validator *Validator) *Value {
	return &Value{
		value:     defaultValue,
		validator: validator,
	}
}

// Uint64 - Creates an unsigned 64-bit integer value
func Uint64(defaultValue uint64, validator *Validator) *Value {
	return &Value{
		value:     defaultValue,
		validator: validator,
	}
}

func (arg *Argument) Environment(name string, eName string, validator *Validator) *Argument {
	return arg
}

func (arg *Argument) Positional(name, help string, validator *Validator) *Argument {
	return arg
}

func (arg *Argument) Optional(name, help string, validator *Validator) *Argument {
	return arg
}

type FileSource struct {
}

func YamlFile(fileName, recordName string) *FileSource {
	return &FileSource{}
}

func JsonFile(fileName, recordName string) *FileSource {
	return &FileSource{}
}

func KeyValueFile(fileName, recordName string) *FileSource {
	return &FileSource{}
}

type ArgumentValue struct{}

func (v *ArgumentValue) String(defaultValue, pattern string) *ArgumentValue {
	return v
}

func (v *ArgumentValue) Int(defaultValue, min, max int64) *ArgumentValue {
	return v
}

func (v *ArgumentValue) Uint(defaultValue, min, max int64) *ArgumentValue {
	return v
}

func (v *ArgumentValue) Float(defaultValue, min, max int64) *ArgumentValue {
	return v
}

type Validator struct {
	vFunc func(v any) bool
}

func Pattern(pattern string) *Validator {
	return &Validator{
		vFunc: func(v any) bool {
			re := regexp.MustCompile(pattern)
			return re.Match([]byte(fmt.Sprintf("%v", v)))
		},
	}
}

func (v *Validator) MinMaxInt(min, max int64) *Validator {
	return v
}
func (v *Validator) MinMaxFloat(min, max float64) *Validator {
	return v
}
