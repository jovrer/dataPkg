package core

const (
	PackStructName string = "PackCore"

	PKG_DATA_NAME = "data"    //真实数据字段名
	PKG_LEN_NAME  = "dataLen" //整个数据包长度字段名
)

type CONF_FORMAT int

const (
	CONF_FORMAT_JSON CONF_FORMAT = 1
	CONF_FORMAT_XML  CONF_FORMAT = 2
	CONF_FORMAT_YAML CONF_FORMAT = 3
)

type VALUE_TYPE int

// var VALUE_TYPE_MAP map[VALUE_TYPE]string
var VALUE_TYPE_MAP map[string]VALUE_TYPE

const (
	VALUE_TYPE_UNKOWN VALUE_TYPE = 0
	VALUE_TYPE_STRING VALUE_TYPE = 1
	VALUE_TYPE_INT    VALUE_TYPE = 2
	VALUE_TYPE_BYTE   VALUE_TYPE = 3
)

func init() {

	VALUE_TYPE_MAP = make(map[string]VALUE_TYPE)
	// VALUE_TYPE_MAP = map[string]VALUE_TYPE{}
	// VALUE_TYPE_MAP[VALUE_TYPE_STRING] = "string"
	// VALUE_TYPE_MAP[VALUE_TYPE_INT] = "int"
	// VALUE_TYPE_MAP[VALUE_TYPE_BYTE] = "byte"

	VALUE_TYPE_MAP["string"] = VALUE_TYPE_STRING
	VALUE_TYPE_MAP["int"] = VALUE_TYPE_INT
	VALUE_TYPE_MAP["byte"] = VALUE_TYPE_BYTE

}
