package coder

import (
	"bytes"
	"testing"
)

type encoderArray struct {
	length int
	result []byte
}

func TestEncoder(t *testing.T) {

	var test_data = []encoderArray{
		{0, []byte{0}},
		{1, []byte{1}},

		{128 - 1, []byte{127}},
		{128 + 1, []byte{129, 1}},
		{128 + 10, []byte{129, 10}},

		{128*128 - 1, []byte{255, 127}},
		{128 * 128, []byte{129, 128, 0}},
		{128*128 + 1, []byte{129, 128, 1}},
		{128*128 + 100, []byte{129, 128, 100}},

		{128*128*128 - 1, []byte{255, 255, 127}},
		{128 * 128 * 128, []byte{129, 128, 128, 0}},
		{128*128*128 + 1, []byte{129, 128, 128, 1}},
		{128*128*128 + 100, []byte{129, 128, 128, 100}},
	}

	for _, v := range test_data {
		ret := encodeLengthToByte(v.length)
		if !bytes.Equal(v.result, ret) {
			t.Error("encodeLengthToByte(%d)", v.length, "falil", "ret=", ret)
		}
	}

}

type decoderArray struct {
	retLength    int
	retCountByte int
	retNeedMore  bool
	byt          []byte
}

func TestDecoder(t *testing.T) {
	var test_data = []decoderArray{

		//刚刚好
		{0, 1, false, []byte{0}},
		{1, 1, false, []byte{1}},

		{128 - 1, 1, false, []byte{127}},
		{128 + 1, 2, false, []byte{129, 1}},
		{128 + 10, 2, false, []byte{129, 10}},

		{128*128 - 1, 2, false, []byte{255, 127}},
		{128 * 128, 3, false, []byte{129, 128, 0}},
		{128*128 + 1, 3, false, []byte{129, 128, 1}},
		{128*128 + 100, 3, false, []byte{129, 128, 100}},

		{128*128*128 - 1, 3, false, []byte{255, 255, 127}},
		{128 * 128 * 128, 4, false, []byte{129, 128, 128, 0}},
		{128*128*128 + 1, 4, false, []byte{129, 128, 128, 1}},
		{128*128*128 + 100, 4, false, []byte{129, 128, 128, 100}},

		//多了
		{0, 1, false, []byte{0, 1}},
		{1, 1, false, []byte{1, 1}},

		{128 - 1, 1, false, []byte{127, 1}},
		{128 + 1, 2, false, []byte{129, 1, 1}},
		{128 + 10, 2, false, []byte{129, 10, 1}},

		{128*128 - 1, 2, false, []byte{255, 127, 1}},
		{128 * 128, 3, false, []byte{129, 128, 0, 1}},
		{128*128 + 1, 3, false, []byte{129, 128, 1, 1}},
		{128*128 + 100, 3, false, []byte{129, 128, 100, 1}},

		{128*128*128 - 1, 3, false, []byte{255, 255, 127, 1}},
		{128 * 128 * 128, 4, false, []byte{129, 128, 128, 0, 1}},
		{128*128*128 + 1, 4, false, []byte{129, 128, 128, 1, 1}},
		{128*128*128 + 100, 4, false, []byte{129, 128, 128, 100, 1}},

		//少了
		{0, 0, true, []byte{}},

		{0, 1, true, []byte{129}},
		{0, 1, true, []byte{129}},

		{0, 1, true, []byte{255}},
		{0, 2, true, []byte{129, 128}},
		{0, 2, true, []byte{129, 128}},
		{0, 2, true, []byte{129, 128}},

		{0, 2, true, []byte{255, 255}},
		{0, 3, true, []byte{129, 128, 128}},
		{0, 3, true, []byte{129, 128, 128}},
		{0, 3, true, []byte{129, 128, 128}},
	}

	for _, v := range test_data {
		retLength, retCountByte, retNeedMore := decodeByteToLength(v.byt)
		if !(v.retLength == retLength &&
			v.retCountByte == retCountByte &&
			v.retNeedMore == retNeedMore) {
			t.Error(v.byt, "retLength=", retLength, "retCountByte=", retCountByte, "retCountByte=", retNeedMore)
		}
	}
}
