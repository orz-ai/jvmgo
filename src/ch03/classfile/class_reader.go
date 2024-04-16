package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

/**
 * @Description: 读取 u1 数据类型
 * @receiver self
 * @return uint8
 */
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]

	// 跳过已经读取的数据
	// 将剩下的重新赋值
	self.data = self.data[1:]
	return val
}

/**
 * @Description: 读取u2类型
 * @receiver self
 * @return uint8
 */
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
 * @Description: 读取uint16表
 * @receiver self
 * @return []uint16
 */
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

/**
 * @Description: 读取指定数量的字节
 * @receiver self
 * @param length
 * @return []byte
 */
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
