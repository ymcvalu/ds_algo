package algorithm

import "bytes"

/**
64位整型编解码
压缩的规则：把整数压缩成为不定长的 buffer 。
     [0, 2^7 -1 ] : 1字节 0xxxxxxx
 [2^7, 2^14 - 1 ] : 2字节 1xxxxxxx 0xxxxxxx
[2^14, 2^21 - 1 ] : 3字节 1xxxxxxx 1xxxxxxx 0xxxxxxx
[2^21, 2^28 - 1 ] : 4字节 1xxxxxxx 1xxxxxxx 1xxxxxxx 0xxxxxxx
......
以此类推。

即每个字节只使用低7bit来作为有效负载，最高的第8bit用来标识多字节的情况。
实现Encode和Decode两个函数：
func Encode(value uint64) []byte
说明：假设 buffer 足够大，函数的返回值指明使用了 buffer 多少个字节

func Decode(buf []byte) (value uint64, len int)
说明：假设 buffer 的数据是正常的，解出来的整数值存放在 value 中，函数的返回值指明解压缩使用了 buffer 中多少个字节。

示例：对数字300进行编码，uint64_t的二进制形式为：
00000000 00000000 00000000 00000000 00000000 00000000 00000001 00101100
编码后的为2字节的buffer，buffer[0]和buffer[1]的二进制分别为：10101100 00000010
*/

func Encode(value uint64) []byte {
	buf := bytes.NewBuffer(nil)
	//i := 0
	for value > 0 {
		b := byte(value & 0b01111111)
		value >>= 7
		if value > 0 {
			b |= 0b10000000
		}
		buf.WriteByte(b)
	}
	return buf.Bytes()
}

func Decode(byts []byte) uint64 {
	var n uint64 = 0
	for i, b := range byts {
		n |= (uint64(b) & 0b01111111) << uint(i*7)

		if b&0b10000000 == 0 {
			break
		}
	}
	return n

}
