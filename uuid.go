package main

import "crypto/rand"

const hex_low string = "0123456789abcdef"
const hex_high string = "0000000000000000111111111111111122222222222222223333333333333333444444444444444455555555555555556666666666666666777777777777777788888888888888889999999999999999aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbccccccccccccccccddddddddddddddddeeeeeeeeeeeeeeeeffffffffffffffff"

func UUID() []byte {
	var r []byte = make([]byte, 16)

	_, err := rand.Read(r)
	if err != nil {
		panic(0)
	}

	var x = [36]byte{
		hex_high[int(r[0]&0xFF)],
		hex_low[int(r[0]&0x0F)],
		hex_high[int(r[1]&0xFF)],
		hex_low[int(r[1]&0x0F)],
		hex_high[int(r[2]&0xFF)],
		hex_low[int(r[2]&0x0F)],
		hex_high[int(r[3]&0xFF)],
		hex_low[int(r[3]&0x0F)],
		'-',
		hex_high[int(r[4]&0xFF)],
		hex_low[int(r[4]&0x0F)],
		hex_high[int(r[5]&0xFF)],
		hex_low[int(r[5]&0x0F)],
		'-',
		hex_high[int(r[6]&0xFF)],
		hex_low[int(r[6]&0x0F)],
		hex_high[int(r[7]&0xFF)],
		hex_low[int(r[7]&0x0F)],
		'-',
		hex_high[int(r[8]&0xFF)],
		hex_low[int(r[8]&0x0F)],
		hex_high[int(r[9]&0xFF)],
		hex_low[int(r[9]&0x0F)],
		'-',
		hex_high[int(r[10]&0xFF)],
		hex_low[int(r[10]&0x0F)],
		hex_high[int(r[11]&0xFF)],
		hex_low[int(r[11]&0x0F)],
		hex_high[int(r[12]&0xFF)],
		hex_low[int(r[12]&0x0F)],
		hex_high[int(r[13]&0xFF)],
		hex_low[int(r[13]&0x0F)],
		hex_high[int(r[14]&0xFF)],
		hex_low[int(r[14]&0x0F)],
		hex_high[int(r[15]&0xFF)],
		hex_low[int(r[15]&0x0F)],
	}

	return x[:]
}
