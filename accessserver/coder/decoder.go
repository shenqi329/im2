package coder

import (
//"log"
)

type Decoder struct {
	buffer        []byte
	messageType   *int
	messageLength *int
}

type Message struct {
	MessageType int
	MessageBuf  []byte
}

func NEWDecoder() *Decoder {

	decoder := &Decoder{}

	decoder.Reset()

	return decoder
}

func (d *Decoder) Reset() {
	d.buffer = make([]byte, 0)
	d.messageType = nil
	d.messageLength = nil
}

func (d *Decoder) decoder(buf []byte) (messages []*Message, err error) {

	d.buffer = append(d.buffer, buf...)
	//log.Println(d.buffer)

	if d.messageType == nil {
		retType, retCountByte, retNeedMore := decodeByteToType(d.buffer)
		if retNeedMore {
			return
		}
		//log.Println(retType, "+", retCountByte, "+", retNeedMore)
		d.messageType = &retType
		d.buffer = d.buffer[retCountByte:]
	}
	if d.messageLength == nil {
		retLength, retCountByte, retNeedMore := decodeByteToLength(d.buffer)
		if retNeedMore {
			return
		}
		//log.Println(retLength, "+", retCountByte, "+", retNeedMore)
		d.messageLength = &retLength
		d.buffer = d.buffer[retCountByte:]
	}

	if len(d.buffer) < *d.messageLength {
		return
	}

	message := &Message{
		MessageType: *d.messageType,
		MessageBuf:  d.buffer[0:*d.messageLength],
	}
	//log.Println(message.MessageBuf)

	messages = append(messages, message)

	d.buffer = d.buffer[*d.messageLength:]
	d.messageType = nil
	d.messageLength = nil

	temp, err := d.decoder(make([]byte, 0))

	messages = append(messages, temp...)

	return
}

func (d *Decoder) Decode(buf []byte) (messages []*Message, err error) {
	return d.decoder(buf)
}

func decodeByteToType(buffer []byte) (retType int, retCountByte int, retNeedMore bool) {
	return decodeByteToLength(buffer)
}

func decodeByteToLength(buffer []byte) (retLength int, retCountByte int, retNeedMore bool) {

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
