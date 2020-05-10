package enum

const (
	ParamTypeString   = "string"
	ParamTypeInt      = "int"
	ParamTypeFloat    = "float"
	ParamTypeDatetime = "datetime"
	ParamTypeObject   = "object"
	ParamTypeArray    = "array"
)

var paramTypes = map[string]string{
	ParamTypeString:   "string",
	ParamTypeInt:      "int",
	ParamTypeFloat:    "float",
	ParamTypeDatetime: "datetime",
	ParamTypeObject:   "object",
	ParamTypeArray:    "array",
}

func ParamTypes() map[string]string {
	return paramTypes
}
