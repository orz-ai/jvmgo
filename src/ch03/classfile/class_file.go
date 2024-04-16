package classfile

import "fmt"

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

/**
 * @Description: 把[]byte 解析成ClassFile结构体
 * @param classData
 * @return cf
 * @return err
 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readMembers(reader, self.constantPool)
}

/**
 * @Description: 读取魔数
 * @receiver self
 * @param reader
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic")
	}
}

/**
 * @Description: 读取版本号
 * @receiver self
 * @param reader
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.majorVersion == 0 {
			return
		}
	}

	panic("java.long.UnsupportedClassVersionError!")
}

/**
 * @Description: 相当于get方法
 * @receiver self
 * @param reader
 */
func (self *ClassFile) MinorVersion(reader *ClassReader) uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion(reader *ClassReader) uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool(reader *ClassReader) {

}

func (self *ClassFile) AccessFlag(reader *ClassReader) {

}

func (self *ClassFile) Fields(reader *ClassReader) {

}

func (self *ClassFile) Methods(reader *ClassReader) {

}

func (self *ClassFile) ClassName(reader *ClassReader) string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName(reader *ClassReader) string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceName(reader *ClassReader) []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
