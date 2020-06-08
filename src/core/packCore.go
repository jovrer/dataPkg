package core

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type fieldDefine struct {
	Name      string
	Value     interface{}
	Type      string
	ValueType VALUE_TYPE
	Len       int

	Handle interface{}
}

type packCore struct {
	fieldsPreix  []fieldDefine //真实数据包前缀
	dataReal     interface{}
	fieldsSuffix []fieldDefine //真实数据包后缀

	pkgLenAt         int //整个数据包长度字段所在位移(单位byte)
	pkgLen           int //整个数据包长度(单位byte)
	pkgLenLen        int //整个数据包长度字段长度(单位byte)
	pkgExpectdataLen int //整个数据包除数据外长度(单位byte)
	dataLen          int //数据长度(单位byte)

}

var objPackCore *packCore

func NewPackCore(conf string, conf_format CONF_FORMAT) *packCore {
	if objPackCore == nil {
		if conf != "" {
			objPackCore = &packCore{}
			if !objPackCore.parseConf(conf, conf_format) {
				objPackCore = nil
			}
		}

	}
	return objPackCore
}

// type fieldDefine struct {
// 	Name      string
// 	Value     interface{}
// 	Type      string
// 	ValueType VALUE_TYPE
// 	Len       int

// 	Handle interface{}
// }
func (this *packCore) parseConf(conf string, conf_format CONF_FORMAT) bool {
	var err error
	var datas []fieldDefine
	switch conf_format {
	case CONF_FORMAT_JSON:
		err = json.Unmarshal([]byte(conf), &datas)
		if err == nil {
			inPre := true
			indexPkgLen := 0
			for i := 0; i < len(datas); i++ {
				switch datas[i].Name {
				case PKG_DATA_NAME:
					inPre = false
					break
				case PKG_LEN_NAME:
					if datas[i].Len > 0 {
						this.pkgLenLen = datas[i].Len
						this.pkgExpectdataLen += datas[i].Len
					}

					this.pkgLenAt = indexPkgLen

					if !inPre {
						fmt.Println("dataLen需要在data前")
					}
					break
				default:
					if datas[i].Len > 0 {
						this.pkgExpectdataLen += datas[i].Len
					}
					break
				}

				if datas[i].Type != "" {
					if NewUtils().ExistKey(datas[i].Type, VALUE_TYPE_MAP) {
						datas[i].ValueType = VALUE_TYPE_MAP[datas[i].Type]
					} else {
						datas[i].ValueType = VALUE_TYPE_UNKOWN
					}
				} else {
					datas[i].ValueType = VALUE_TYPE_UNKOWN
				}

				if datas[i].Len > 0 {
					if inPre {
						// if datas[i].Len > 0 {

						// }
						this.fieldsPreix = append(this.fieldsPreix, datas[i])
					} else {
						if datas[i].Name != PKG_DATA_NAME {
							this.fieldsSuffix = append(this.fieldsSuffix, datas[i])
						}
					}
				}

				indexPkgLen += datas[i].Len
			}
		}
		break
	default:
		// err = error{}
		fmt.Println("not support")
		break
	}

	if err == nil {
		return true
	}

	return false
}

func (this *packCore) TotalLen() int {
	return this.pkgLen
}

func (this *packCore) DataLen() int {
	return this.dataLen
}

func (this *packCore) Pack(v []byte) []byte {
	var datas []byte
	// if this.checkPackStruct() {
	if true {
		this.dataLen = len(v)
		this.pkgLen = this.dataLen + this.pkgExpectdataLen

		i := 0

		for i = 0; i < len(this.fieldsPreix); i++ {
			switch this.fieldsPreix[i].Name {
			case PKG_LEN_NAME:
				this.fieldsPreix[i].Value = this.pkgLen
				break
			default:
				break
			}
		}

		for i = 0; i < len(this.fieldsPreix); i++ {
			var datasTemp []byte
			switch this.fieldsPreix[i].ValueType {
			case VALUE_TYPE_STRING:
				datasTemp = NewUtils().String2Bytes(this.fieldsPreix[i].Value.(string), this.fieldsPreix[i].Len)
				break
			case VALUE_TYPE_INT:
				datasTemp = NewUtils().Int2Bytes(this.fieldsPreix[i].Value.(int), this.fieldsPreix[i].Len)
				break
			case VALUE_TYPE_BYTE:
				datasTemp = NewUtils().BytesAlign(this.fieldsPreix[i].Value.([]byte), this.fieldsPreix[i].Len)
				break
			default:
				break
			}

			datas = NewUtils().BytesCombine(datas, datasTemp)
		}

		datas = NewUtils().BytesCombine(datas, v)
		// realDatas, err := m_proto_handle.Marshal(this.dataReal)

		for i = 0; i < len(this.fieldsSuffix); i++ {
			var datasTemp []byte
			switch this.fieldsSuffix[i].ValueType {
			case VALUE_TYPE_STRING:
				datasTemp = NewUtils().String2Bytes(this.fieldsSuffix[i].Value.(string), this.fieldsSuffix[i].Len)
				break
			case VALUE_TYPE_INT:
				datasTemp = NewUtils().Int2Bytes(this.fieldsSuffix[i].Value.(int), this.fieldsSuffix[i].Len)
				break
			case VALUE_TYPE_BYTE:
				datasTemp = NewUtils().BytesAlign(this.fieldsSuffix[i].Value.([]byte), this.fieldsSuffix[i].Len)
				// break
			default:
				break
			}

			datas = NewUtils().BytesCombine(datas, datasTemp)
		}
	}
	return datas
}

func (this *packCore) Unpack(datas []byte) (interface{}, error) {
	// if this.checkPackStruct() {
	if true {
		if len(datas) >= (this.pkgLenAt + this.pkgLenLen) {
			this.pkgLen = NewUtils().IntFromBytes(datas[this.pkgLenAt:(this.pkgLenAt + this.pkgLenLen)])
		}

		if this.pkgLen == len(datas) {
			i := 0
			bPos := 0
			for i = 0; i < len(this.fieldsPreix); i++ {
				// var datasTemp []byte
				switch this.fieldsPreix[i].ValueType {
				case VALUE_TYPE_STRING:
					this.fieldsPreix[i].Value = NewUtils().StringFromBytes(datas[bPos:(bPos + this.fieldsPreix[i].Len)])
					bPos += this.fieldsPreix[i].Len
					break
				case VALUE_TYPE_INT:
					this.fieldsPreix[i].Value = NewUtils().IntFromBytes(datas[bPos:(bPos + this.fieldsPreix[i].Len)])
					bPos += this.fieldsPreix[i].Len
					break
				case VALUE_TYPE_BYTE:
					// datasTemp = NewUtils().BytesAlign(this.fieldsPreix[i].Value.([]byte), (bPos + this.fieldsPreix[i].Len))
					break
				default:
					break
				}

				// if this.fieldsPreix[i].Name == PKG_LEN_NAME {
				// 	this.pkgLen = (this.fieldsPreix[i].Value).(int)
				// }
			}

			this.dataLen = this.pkgLen - this.pkgExpectdataLen
			this.dataReal = datas[bPos:(bPos + this.dataLen)]

			for i = 0; i < len(this.fieldsSuffix); i++ {
				// var datasTemp []byte
				switch this.fieldsSuffix[i].ValueType {
				case VALUE_TYPE_STRING:
					this.fieldsPreix[i].Value = NewUtils().StringFromBytes(datas[bPos:(bPos + this.fieldsSuffix[i].Len)])
					bPos += this.fieldsSuffix[i].Len
					break
				case VALUE_TYPE_INT:
					this.fieldsSuffix[i].Value = NewUtils().IntFromBytes(datas[bPos:(bPos + this.fieldsSuffix[i].Len)])
					bPos += this.fieldsSuffix[i].Len
					break
				case VALUE_TYPE_BYTE:
					// datasTemp = NewUtils().BytesAlign(this.fieldsSuffix[i].Value.([]byte), (bPos + this.fieldsSuffix[i].Len))

					break
				default:
					break
				}

			}

			return this.dataReal, nil
		}
	}

	return nil, nil
}

func (this *packCore) checkPackStruct() bool {
	pName := reflect.TypeOf(PackStructName).Name()

	dd := reflect.TypeOf(PackStructName)

	if dd != nil {

	}

	dd1 := dd.Kind()

	dd2 := reflect.New(dd)

	fmt.Println(dd1)

	fmt.Println(dd2.IsValid())

	dd4 := reflect.New(reflect.TypeOf("ww"))

	fmt.Println(dd4.IsValid())

	// if dd2 != nil {

	// }

	if pName != PackStructName {
		return false
		// return true
	}
	return true
}
