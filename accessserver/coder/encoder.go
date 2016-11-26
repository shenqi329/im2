package coder

import (
//proto "github.com/golang/protobuf/proto"
)

type Encoder struct {
	message []byte
}

// func (e *Encoder) encoder(message *proto.Message) error {

// 	b, err := proto.Marshal(message)
// 	if err != nil {
// 		return err
// 	}

// 	e.message = b

// 	len := len(e.message)

// 	return nil
// }

func getByteLen(buffer []byte) (retLength int, retCountByte int, retNeedMore bool) {

	retLength = 0
	retNeedMore = true
	retCountByte = 0

	for retCountByte = 0; retCountByte < len(buffer); retCountByte++ {
		byt := buffer[retCountByte]
		retLength |= int(byt & 0x7F)
		byt = buffer[retCountByte]
		if (byt & 0x80) == 0x80 {
			retLength = retLength << 7
		} else {
			retNeedMore = false
			retCountByte++
			break
		}
	}
	if retNeedMore {
		retLength = 0
	}

	return retLength, retCountByte, retNeedMore
}

func encodeLengthToByte(length int) []byte {

	needByte := 1
	tempLen := length

	for tempLen >= 128 {
		tempLen = tempLen >> 7
		needByte++
	}

	tempLen = length
	buffer := make([]byte, needByte)

	for index := 0; index < needByte; index++ {
		buffer[needByte-index-1] = (byte)(tempLen & 0x7F)
		if index != 0 {
			buffer[needByte-index-1] |= 0x80
		}
		tempLen = tempLen >> 7
	}
	return buffer
}

func encodeTypeToByte(typ int) []byte {

	return encodeLengthToByte(typ)

}
