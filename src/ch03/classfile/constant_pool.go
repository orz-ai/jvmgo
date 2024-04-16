package classfile

/*
1.常量池也是一个表
2.表头给出的常量池的大小比实际大1，有效的常量池索引是1~n-1，0是无效索引
3.CONSTANT_Long_info和CONSTANT_Double_info各占两个位置

*/
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantPool, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

/**
 * @Description: 按索引查找常量
 * @receiver self ConstantPool
 * @param index
 * @return ConstantPool
 */
func (self ConstantPool) getConstantInfo(index uint16) ConstantPool {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}

	panic("invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
