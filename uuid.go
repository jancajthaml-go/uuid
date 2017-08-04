package main

import "crypto/rand"

const hex string = "0123456789abcdef"

func UUID() []byte {
	var r []byte = make([]byte, 16)

	_, err := rand.Read(r)
	if err != nil {
		panic(0)
	}

	var x = [...]byte{
		hex[int(r[0]&0xff)>>4],
		hex[int(r[0]&0x0f)],
		hex[int(r[1]&0xff)>>4],
		hex[int(r[1]&0x0f)],
		hex[int(r[2]&0xff)>>4],
		hex[int(r[2]&0x0f)],
		hex[int(r[3]&0xff)>>4],
		hex[int(r[3]&0x0f)],
		'-',
		hex[int(r[4]&0xff)>>4],
		hex[int(r[4]&0x0f)],
		hex[int(r[5]&0xff)>>4],
		hex[int(r[5]&0x0f)],
		'-',
		hex[int(r[6]&0xff)>>4],
		hex[int(r[6]&0x0f)],
		hex[int(r[7]&0xff)>>4],
		hex[int(r[7]&0x0f)],
		'-',
		hex[int(r[8]&0xff)>>4],
		hex[int(r[8]&0x0f)],
		hex[int(r[9]&0xff)>>4],
		hex[int(r[9]&0x0f)],
		'-',
		hex[int(r[10]&0xff)>>4],
		hex[int(r[10]&0x0f)],
		hex[int(r[11]&0xff)>>4],
		hex[int(r[11]&0x0f)],
		hex[int(r[12]&0xff)>>4],
		hex[int(r[12]&0x0f)],
		hex[int(r[13]&0xff)>>4],
		hex[int(r[13]&0x0f)],
		hex[int(r[14]&0xff)>>4],
		hex[int(r[14]&0x0f)],
		hex[int(r[15]&0xff)>>4],
		hex[int(r[15]&0x0f)],
	}

	return x[:]
}
