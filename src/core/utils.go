package core

import (
	"bytes"
	"encoding/binary"
)

type utils struct {
}

var objUtils *utils

func NewUtils() *utils {
	if objUtils == nil {
		objUtils = &utils{}
	}
	return objUtils
}

func (*utils) String2Bytes(data string, alignLen int) []byte {
	dataBytes := []byte(data)
	if alignLen >= len(dataBytes) {
		datas := make([]byte, alignLen)
		orignLen := len(dataBytes)
		offPos := alignLen - orignLen

		for i := offPos; i < alignLen; i++ {
			datas[i] = dataBytes[i-offPos]
		}

		return datas
	} else {
		return nil
	}
}

func (*utils) StringFromBytes(datas []byte) string {
	return string(datas)
}

func (*utils) Int2Bytes(data int, alignLen int) []byte {
	x := int32(data)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	dataBytes := bytesBuffer.Bytes()

	if alignLen >= len(dataBytes) {
		datas := make([]byte, alignLen)

		orignLen := len(dataBytes)
		offPos := alignLen - orignLen

		for i := offPos; i < alignLen; i++ {
			datas[i] = dataBytes[i-offPos]
		}

		return datas
	} else {
		return nil
	}
}

func (*utils) IntFromBytes(datas []byte) int {
	bytesBuffer := bytes.NewBuffer(datas)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func (*utils) BytesAlign(data []byte, alignLen int) []byte {
	if alignLen >= len(data) {
		datas := make([]byte, alignLen)
		orignLen := len(data)
		offPos := alignLen - orignLen

		for i := offPos; i < alignLen; i++ {
			datas[i] = data[i-offPos]
		}

		return datas
	}
	return nil
}

func (*utils) BytesCombine(data ...[]byte) []byte {
	return bytes.Join(data, []byte(""))
}

func (*utils) ExistKey(k string, v map[string]VALUE_TYPE) bool {
	_, ok := v[k]
	if ok {
		return true
	}
	return false
}
