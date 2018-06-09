//
// Define safe prime groups along with some primative operations on the groups.
//

package srp

import "math/big"

//
// Define the default group to be 4096 bit. Use of a group should always be done
// via the "dGrp" variable, rather than calling a "group****()" method directly
//
var dGrp = group4096()

type group struct {
	g []byte // Group generator
	N []byte // Group safe prime
}

func group2048() group {
	return group{
		g: []byte{0x02},
		N: []byte{
			0xAC, 0x6B, 0xDB, 0x41, 0x32, 0x4A, 0x9A, 0x9B, 0xF1, 0x66, 0xDE, 0x5E,
			0x13, 0x89, 0x58, 0x2F, 0xAF, 0x72, 0xB6, 0x65, 0x19, 0x87, 0xEE, 0x07,
			0xFC, 0x31, 0x92, 0x94, 0x3D, 0xB5, 0x60, 0x50, 0xA3, 0x73, 0x29, 0xCB,
			0xB4, 0xA0, 0x99, 0xED, 0x81, 0x93, 0xE0, 0x75, 0x77, 0x67, 0xA1, 0x3D,
			0xD5, 0x23, 0x12, 0xAB, 0x4B, 0x03, 0x31, 0x0D, 0xCD, 0x7F, 0x48, 0xA9,
			0xDA, 0x04, 0xFD, 0x50, 0xE8, 0x08, 0x39, 0x69, 0xED, 0xB7, 0x67, 0xB0,
			0xCF, 0x60, 0x95, 0x17, 0x9A, 0x16, 0x3A, 0xB3, 0x66, 0x1A, 0x05, 0xFB,
			0xD5, 0xFA, 0xAA, 0xE8, 0x29, 0x18, 0xA9, 0x96, 0x2F, 0x0B, 0x93, 0xB8,
			0x55, 0xF9, 0x79, 0x93, 0xEC, 0x97, 0x5E, 0xEA, 0xA8, 0x0D, 0x74, 0x0A,
			0xDB, 0xF4, 0xFF, 0x74, 0x73, 0x59, 0xD0, 0x41, 0xD5, 0xC3, 0x3E, 0xA7,
			0x1D, 0x28, 0x1E, 0x44, 0x6B, 0x14, 0x77, 0x3B, 0xCA, 0x97, 0xB4, 0x3A,
			0x23, 0xFB, 0x80, 0x16, 0x76, 0xBD, 0x20, 0x7A, 0x43, 0x6C, 0x64, 0x81,
			0xF1, 0xD2, 0xB9, 0x07, 0x87, 0x17, 0x46, 0x1A, 0x5B, 0x9D, 0x32, 0xE6,
			0x88, 0xF8, 0x77, 0x48, 0x54, 0x45, 0x23, 0xB5, 0x24, 0xB0, 0xD5, 0x7D,
			0x5E, 0xA7, 0x7A, 0x27, 0x75, 0xD2, 0xEC, 0xFA, 0x03, 0x2C, 0xFB, 0xDB,
			0xF5, 0x2F, 0xB3, 0x78, 0x61, 0x60, 0x27, 0x90, 0x04, 0xE5, 0x7A, 0xE6,
			0xAF, 0x87, 0x4E, 0x73, 0x03, 0xCE, 0x53, 0x29, 0x9C, 0xCC, 0x04, 0x1C,
			0x7B, 0xC3, 0x08, 0xD8, 0x2A, 0x56, 0x98, 0xF3, 0xA8, 0xD0, 0xC3, 0x82,
			0x71, 0xAE, 0x35, 0xF8, 0xE9, 0xDB, 0xFB, 0xB6, 0x94, 0xB5, 0xC8, 0x03,
			0xD8, 0x9F, 0x7A, 0xE4, 0x35, 0xDE, 0x23, 0x6D, 0x52, 0x5F, 0x54, 0x75,
			0x9B, 0x65, 0xE3, 0x72, 0xFC, 0xD6, 0x8E, 0xF2, 0x0F, 0xA7, 0x11, 0x1F,
			0x9E, 0x4A, 0xFF, 0x73,
		}}
}

func group4096() group {
	return group{
		g: []byte{0x05},
		N: []byte{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC9, 0x0F, 0xDA, 0xA2,
			0x21, 0x68, 0xC2, 0x34, 0xC4, 0xC6, 0x62, 0x8B, 0x80, 0xDC, 0x1C, 0xD1,
			0x29, 0x02, 0x4E, 0x08, 0x8A, 0x67, 0xCC, 0x74, 0x02, 0x0B, 0xBE, 0xA6,
			0x3B, 0x13, 0x9B, 0x22, 0x51, 0x4A, 0x08, 0x79, 0x8E, 0x34, 0x04, 0xDD,
			0xEF, 0x95, 0x19, 0xB3, 0xCD, 0x3A, 0x43, 0x1B, 0x30, 0x2B, 0x0A, 0x6D,
			0xF2, 0x5F, 0x14, 0x37, 0x4F, 0xE1, 0x35, 0x6D, 0x6D, 0x51, 0xC2, 0x45,
			0xE4, 0x85, 0xB5, 0x76, 0x62, 0x5E, 0x7E, 0xC6, 0xF4, 0x4C, 0x42, 0xE9,
			0xA6, 0x37, 0xED, 0x6B, 0x0B, 0xFF, 0x5C, 0xB6, 0xF4, 0x06, 0xB7, 0xED,
			0xEE, 0x38, 0x6B, 0xFB, 0x5A, 0x89, 0x9F, 0xA5, 0xAE, 0x9F, 0x24, 0x11,
			0x7C, 0x4B, 0x1F, 0xE6, 0x49, 0x28, 0x66, 0x51, 0xEC, 0xE4, 0x5B, 0x3D,
			0xC2, 0x00, 0x7C, 0xB8, 0xA1, 0x63, 0xBF, 0x05, 0x98, 0xDA, 0x48, 0x36,
			0x1C, 0x55, 0xD3, 0x9A, 0x69, 0x16, 0x3F, 0xA8, 0xFD, 0x24, 0xCF, 0x5F,
			0x83, 0x65, 0x5D, 0x23, 0xDC, 0xA3, 0xAD, 0x96, 0x1C, 0x62, 0xF3, 0x56,
			0x20, 0x85, 0x52, 0xBB, 0x9E, 0xD5, 0x29, 0x07, 0x70, 0x96, 0x96, 0x6D,
			0x67, 0x0C, 0x35, 0x4E, 0x4A, 0xBC, 0x98, 0x04, 0xF1, 0x74, 0x6C, 0x08,
			0xCA, 0x18, 0x21, 0x7C, 0x32, 0x90, 0x5E, 0x46, 0x2E, 0x36, 0xCE, 0x3B,
			0xE3, 0x9E, 0x77, 0x2C, 0x18, 0x0E, 0x86, 0x03, 0x9B, 0x27, 0x83, 0xA2,
			0xEC, 0x07, 0xA2, 0x8F, 0xB5, 0xC5, 0x5D, 0xF0, 0x6F, 0x4C, 0x52, 0xC9,
			0xDE, 0x2B, 0xCB, 0xF6, 0x95, 0x58, 0x17, 0x18, 0x39, 0x95, 0x49, 0x7C,
			0xEA, 0x95, 0x6A, 0xE5, 0x15, 0xD2, 0x26, 0x18, 0x98, 0xFA, 0x05, 0x10,
			0x15, 0x72, 0x8E, 0x5A, 0x8A, 0xAA, 0xC4, 0x2D, 0xAD, 0x33, 0x17, 0x0D,
			0x04, 0x50, 0x7A, 0x33, 0xA8, 0x55, 0x21, 0xAB, 0xDF, 0x1C, 0xBA, 0x64,
			0xEC, 0xFB, 0x85, 0x04, 0x58, 0xDB, 0xEF, 0x0A, 0x8A, 0xEA, 0x71, 0x57,
			0x5D, 0x06, 0x0C, 0x7D, 0xB3, 0x97, 0x0F, 0x85, 0xA6, 0xE1, 0xE4, 0xC7,
			0xAB, 0xF5, 0xAE, 0x8C, 0xDB, 0x09, 0x33, 0xD7, 0x1E, 0x8C, 0x94, 0xE0,
			0x4A, 0x25, 0x61, 0x9D, 0xCE, 0xE3, 0xD2, 0x26, 0x1A, 0xD2, 0xEE, 0x6B,
			0xF1, 0x2F, 0xFA, 0x06, 0xD9, 0x8A, 0x08, 0x64, 0xD8, 0x76, 0x02, 0x73,
			0x3E, 0xC8, 0x6A, 0x64, 0x52, 0x1F, 0x2B, 0x18, 0x17, 0x7B, 0x20, 0x0C,
			0xBB, 0xE1, 0x17, 0x57, 0x7A, 0x61, 0x5D, 0x6C, 0x77, 0x09, 0x88, 0xC0,
			0xBA, 0xD9, 0x46, 0xE2, 0x08, 0xE2, 0x4F, 0xA0, 0x74, 0xE5, 0xAB, 0x31,
			0x43, 0xDB, 0x5B, 0xFC, 0xE0, 0xFD, 0x10, 0x8E, 0x4B, 0x82, 0xD1, 0x20,
			0xA9, 0x21, 0x08, 0x01, 0x1A, 0x72, 0x3C, 0x12, 0xA7, 0x87, 0xE6, 0xD7,
			0x88, 0x71, 0x9A, 0x10, 0xBD, 0xBA, 0x5B, 0x26, 0x99, 0xC3, 0x27, 0x18,
			0x6A, 0xF4, 0xE2, 0x3C, 0x1A, 0x94, 0x68, 0x34, 0xB6, 0x15, 0x0B, 0xDA,
			0x25, 0x83, 0xE9, 0xCA, 0x2A, 0xD4, 0x4C, 0xE8, 0xDB, 0xBB, 0xC2, 0xDB,
			0x04, 0xDE, 0x8E, 0xF9, 0x2E, 0x8E, 0xFC, 0x14, 0x1F, 0xBE, 0xCA, 0xA6,
			0x28, 0x7C, 0x59, 0x47, 0x4E, 0x6B, 0xC0, 0x5D, 0x99, 0xB2, 0x96, 0x4F,
			0xA0, 0x90, 0xC3, 0xA2, 0x23, 0x3B, 0xA1, 0x86, 0x51, 0x5B, 0xE7, 0xED,
			0x1F, 0x61, 0x29, 0x70, 0xCE, 0xE2, 0xD7, 0xAF, 0xB8, 0x1B, 0xDD, 0x76,
			0x21, 0x70, 0x48, 0x1C, 0xD0, 0x06, 0x91, 0x27, 0xD5, 0xB0, 0x5A, 0xA9,
			0x93, 0xB4, 0xEA, 0x98, 0x8D, 0x8F, 0xDD, 0xC1, 0x86, 0xFF, 0xB7, 0xDC,
			0x90, 0xA6, 0xC0, 0x8F, 0x4D, 0xF4, 0x35, 0xC9, 0x34, 0x06, 0x31, 0x99,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		}}
}

func (g *group) exp(x1, x2 []byte) []byte {
	// Convert x1 from []byte -> *Int
	x1BigInt := big.NewInt(0).SetBytes(x1)

	// Convert x2 from []byte -> *Int
	x2BigInt := big.NewInt(0).SetBytes(x2)

	// Convert mod from []byte -> *Int
	modBigInt := big.NewInt(0).SetBytes(g.N)

	// Calculate (x1^x2) modulo N
	resultBigInt := big.NewInt(0).Exp(x1BigInt, x2BigInt, modBigInt)

	// Convert the resulting big int to []byte
	resultBytes := resultBigInt.Bytes()

	// Finally, return the resulting bytes.
	return resultBytes
}

func (g *group) mul(x1, x2 []byte) []byte {
	// Convert x1 from []byte -> *Int
	x1BigInt := big.NewInt(0).SetBytes(x1)

	// Convert x2 from []byte -> *Int
	x2BigInt := big.NewInt(0).SetBytes(x2)

	// Convert mod from []byte -> *Int
	modBigInt := big.NewInt(0).SetBytes(g.N)

	// Calculate (x1 * x2) modulo N
	resultBigInt := big.NewInt(0).Mul(x1BigInt, x2BigInt)
	resultBigInt.Mod(resultBigInt, modBigInt)

	// Convert the resulting big int to []byte
	resultBytes := resultBigInt.Bytes()

	// Finally, return the resulting bytes.
	return resultBytes
}

func (g *group) add(x1, x2 []byte) []byte {
	// Convert x1 from []byte -> *Int
	x1BigInt := big.NewInt(0).SetBytes(x1)

	// Convert x2 from []byte -> *Int
	x2BigInt := big.NewInt(0).SetBytes(x2)

	// Convert mod from []byte -> *Int
	modBigInt := big.NewInt(0).SetBytes(g.N)

	// Calculate (x1 + x2) modulo mod
	resultBigInt := big.NewInt(0).Add(x1BigInt, x2BigInt)
	resultBigInt.Mod(resultBigInt, modBigInt)

	// Convert the resulting big int to []byte
	resultBytes := resultBigInt.Bytes()

	// Finally, return the resulting bytes.
	return resultBytes
}

func (g *group) sub(x1, x2 []byte) []byte {
	// Convert x1 from []byte -> *Int
	x1BigInt := big.NewInt(0).SetBytes(x1)

	// Convert x2 from []byte -> *Int
	x2BigInt := big.NewInt(0).SetBytes(x2)

	// Convert mod from []byte -> *Int
	modBigInt := big.NewInt(0).SetBytes(g.N)

	// Calculate (x1 - x2) modulo mod
	resultBigInt := big.NewInt(0).Sub(x1BigInt, x2BigInt)
	resultBigInt.Mod(resultBigInt, modBigInt)

	// Convert the resulting big int to []byte
	resultBytes := resultBigInt.Bytes()

	// Finally, return the resulting bytes.
	return resultBytes
}
