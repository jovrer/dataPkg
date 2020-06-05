package core

import "bytes"

type utils struct {
}

var objUtils *utils

func NewUtils() *utils {
	if objUtils == nil {
		objUtils = &utils{}
	}
	return objUtils
}

func (*utils) String2Bytes(data string, bytesLen int) []byte {
	datas := make([]byte, bytesLen)
	datas = []byte(data)

	return datas
}

func (*utils) StringFromBytes(datas []byte) string {
	return string(datas)
}

func (*utils) Int2Bytes(data string, bytesLen int) []byte {
	datas := make([]byte, bytesLen)
	datas = []byte(data)

	return datas
}

func (*utils) IntFromBytes(datas []byte) int {
	return 0
}

func (*utils) BytesAlign(data []byte, bytesLen int) []byte {
	if bytesLen >= len(data) {
		datas := make([]byte, bytesLen)
		orignLen := len(data)
		offPos := bytesLen - orignLen

		for i := offPos; i < bytesLen; i++ {
			datas[i] = data[i-offPos]
		}

		return datas
	}
	return nil
}

func (*utils) BytesCombine(data ...[]byte) []byte {
	return bytes.Join(data, []byte(""))
}
