// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Universally_unique_identifier
//
package uuid

import (
	"crypto/rand"
	"fmt"
)

const hex_low string = "0123456789abcdef"
const hex_high string = "0000000000000000111111111111111122222222222222223333333333333333444444444444444455555555555555556666666666666666777777777777777788888888888888889999999999999999aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbccccccccccccccccddddddddddddddddeeeeeeeeeeeeeeeeffffffffffffffff"

// Generate uuid version 4 variant DCE 1.1, ISO/IEC 11578:1996
func Generate() ([]byte, error) {
	r := make([]byte, 36)

	_, err := rand.Read(r)
	if err != nil {
		return nil, fmt.Errorf("unable to read from /dev/urandom")
	}

	r[35] = hex_low[int(r[15]&0x0F)]
	r[34] = hex_high[int(r[15]&0xFF)]
	r[33] = hex_low[int(r[14]&0x0F)]
	r[32] = hex_high[int(r[14]&0xFF)]
	r[31] = hex_low[int(r[13]&0x0F)]
	r[30] = hex_high[int(r[13]&0xFF)]
	r[29] = hex_low[int(r[12]&0x0F)]
	r[28] = hex_high[int(r[12]&0xFF)]
	r[27] = hex_low[int(r[11]&0x0F)]
	r[26] = hex_high[int(r[11]&0xFF)]
	r[25] = hex_low[int(r[10]&0x0F)]
	r[24] = hex_high[int(r[10]&0xFF)]
	r[23] = '-'
	r[22] = hex_low[int(r[9]&0x0F)]
	r[21] = hex_high[int(r[9]&0xFF)]
	r[20] = hex_low[int(r[8]&0x0F)]
	r[19] = '8'
	r[18] = '-'
	r[17] = hex_low[int(r[7]&0x0F)]
	r[16] = hex_high[int(r[7]&0xFF)]
	r[15] = hex_low[int(r[6]&0x0F)]
	r[14] = '4'
	r[13] = '-'
	r[12] = hex_low[int(r[5]&0x0F)]
	r[11] = hex_high[int(r[5]&0xFF)]
	r[10] = hex_low[int(r[4]&0x0F)]
	r[9] = hex_high[int(r[4]&0xFF)]
	r[8] = '-'
	r[7] = hex_low[int(r[3]&0x0F)]
	r[6] = hex_high[int(r[3]&0xFF)]
	r[5] = hex_low[int(r[2]&0x0F)]
	r[4] = hex_high[int(r[2]&0xFF)]
	r[3] = hex_low[int(r[1]&0x0F)]
	r[2] = hex_high[int(r[1]&0xFF)]
	r[1] = hex_low[int(r[0]&0x0F)]
	r[0] = hex_high[int(r[0]&0xFF)]
	
	return r, nil
}
