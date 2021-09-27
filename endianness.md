# Endian Ness

In computing, endianness is the order or sequence of bytes of a word of digital data in computer memory. 

Endianness is primarily expressed as big-endian (BE) or little-endian (LE). 

A big-endian system stores the most significant byte of a word at the smallest memory address and the least significant byte at the largest. 

A little-endian system, in contrast, stores the least-significant byte at the smallest address.

least-significant byte  = LSB

most significant byte = MSB

https://www.iar.com/contentassets/e85ab3798e204378b0effc7275b2dea3/endianness_figure1.png

http://www.mathcs.emory.edu/~cheung/Courses/255/Syl-ARM/7-ARM/FIGS/endian01b.gif

lscpu | grep "Byte Order"

## How to check whether a system is big endian or little endian?

```golang

var nativeEndian binary.ByteOrder

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		panic("Could not determine native endianness.")
	}
}
  
//true = big endian, false = little endian
func getEndian() (ret bool) {
	var i int = 0x1
	bs := (*[2]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		return true
	} else {
		return false
	}
}
```


## Reference
* https://stackoverflow.com/questions/51332658/any-better-way-to-check-endianness-in-go/51332762
