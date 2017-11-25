//
// Define the safe prime groups for different bit-lengths.
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
       0x30, 0x82, 0x01, 0x08, 0x02, 0x82, 0x01, 0x01, 0x00, 0x9E, 0x6B, 0xC0,
       0x62, 0xEB, 0xC2, 0xAD, 0x49, 0x28, 0xFE, 0x3E, 0x3F, 0x6B, 0xE4, 0xCE,
       0x95, 0xF2, 0xF0, 0xCD, 0x29, 0xD5, 0x2E, 0x41, 0x6C, 0xFB, 0x6F, 0x09,
       0x4A, 0x2B, 0x6D, 0xFF, 0x5C, 0x2C, 0xDF, 0xB0, 0xFE, 0x5C, 0x65, 0x1E,
       0xBA, 0xD2, 0x6F, 0xDA, 0x20, 0xEC, 0x6E, 0x51, 0x35, 0x38, 0xF3, 0xFA,
       0xD0, 0x25, 0x0B, 0x7C, 0xA8, 0x53, 0xB8, 0x22, 0xB7, 0x5D, 0xF2, 0xAB,
       0x67, 0xA7, 0x74, 0x1D, 0x7C, 0xFD, 0x88, 0x54, 0xD7, 0x6B, 0xCF, 0x85,
       0x56, 0x07, 0x6F, 0xC5, 0x9A, 0xCF, 0x0F, 0x59, 0x39, 0x45, 0x22, 0xC5,
       0x2E, 0xA6, 0xCC, 0xC1, 0x7D, 0x50, 0x69, 0x6C, 0x94, 0xC3, 0xAC, 0x7C,
       0x9A, 0x0E, 0xFD, 0x5A, 0x75, 0x13, 0x04, 0xCE, 0xF6, 0xD0, 0x16, 0x11,
       0x08, 0x43, 0xFE, 0x72, 0x0C, 0xCA, 0x2A, 0xC7, 0xC1, 0xC5, 0x88, 0x9D,
       0x6E, 0xEC, 0x7F, 0xD7, 0x66, 0x0E, 0xC2, 0x9B, 0xBA, 0xDE, 0x32, 0x74,
       0xB6, 0xCA, 0xED, 0xBD, 0xA7, 0x5E, 0x73, 0x38, 0x11, 0x0F, 0xC3, 0xD7,
       0x38, 0x15, 0x9F, 0x1C, 0x8F, 0x30, 0x10, 0x40, 0xE5, 0xE0, 0xA9, 0xC1,
       0x79, 0xEC, 0x83, 0x63, 0xDE, 0x55, 0x38, 0x74, 0xC8, 0x58, 0x20, 0x92,
       0xEA, 0xF3, 0x30, 0x08, 0x00, 0x00, 0xA6, 0x79, 0xA6, 0x3C, 0x91, 0x41,
       0xA7, 0x2C, 0x0D, 0xAA, 0x1C, 0x4D, 0x9A, 0x35, 0x86, 0x23, 0xA7, 0x99,
       0x78, 0x4D, 0x97, 0xC1, 0xD7, 0x50, 0x2E, 0x3F, 0x49, 0x8E, 0xAB, 0xB0,
       0x46, 0xD9, 0xF2, 0xAB, 0x4C, 0x92, 0xE2, 0x8B, 0xD9, 0x75, 0x6F, 0xDF,
       0x72, 0xB7, 0xFE, 0xB4, 0xD3, 0x94, 0xBE, 0x8C, 0x77, 0x98, 0x48, 0x1B,
       0x1C, 0xB8, 0xDB, 0xD9, 0x1A, 0xAB, 0x53, 0xAF, 0xA8, 0xE6, 0x2E, 0x1E,
       0xDB, 0xA6, 0x34, 0x67, 0xA4, 0xB3, 0x76, 0xCA, 0xA7, 0xB9, 0xF9, 0x80,
       0xAB, 0x02, 0x01, 0x02,
  }}
}

func group4096() group {
  return group{
    g: []byte{0x05},
    N: []byte{
       0x30, 0x82, 0x02, 0x08, 0x02, 0x82, 0x02, 0x01, 0x00, 0x90, 0x99, 0xB7,
       0x60, 0x7F, 0x46, 0xA9, 0x53, 0xEA, 0x4E, 0xC7, 0x99, 0x86, 0x47, 0x1B,
       0x00, 0x8F, 0x66, 0xD3, 0xFF, 0x94, 0xE4, 0x0C, 0x57, 0x3D, 0x9F, 0xCB,
       0xE5, 0x1B, 0xDD, 0x36, 0xC0, 0xE1, 0x50, 0x4D, 0x37, 0xE1, 0x24, 0xE9,
       0xAA, 0x3E, 0x33, 0xD8, 0x84, 0x0C, 0x0A, 0x58, 0xF6, 0x1A, 0x1A, 0x38,
       0x7B, 0x4E, 0x4E, 0x6C, 0x16, 0x5F, 0x9B, 0x44, 0x5B, 0xDE, 0xEB, 0xCD,
       0xC8, 0x36, 0x56, 0x9D, 0xBF, 0x78, 0xFB, 0xBF, 0x41, 0xE1, 0x3F, 0xA0,
       0x8F, 0xB4, 0xA6, 0x0A, 0x44, 0x48, 0xAF, 0x5E, 0xE4, 0xA0, 0xEF, 0xA3,
       0x1A, 0x52, 0x6B, 0x35, 0xDF, 0x98, 0x60, 0x4F, 0xC0, 0x27, 0x41, 0x73,
       0x3A, 0x29, 0xAE, 0xD7, 0x59, 0x33, 0x3D, 0x70, 0x1D, 0xD8, 0xCC, 0xDE,
       0xF5, 0xD2, 0x60, 0xB2, 0x2C, 0x47, 0x08, 0x1B, 0x75, 0x34, 0x26, 0xE7,
       0xB8, 0xE2, 0x18, 0x78, 0xAF, 0x9D, 0x30, 0x4E, 0x4D, 0x17, 0x3C, 0xC9,
       0x24, 0x1E, 0x7B, 0x16, 0x98, 0x71, 0xE8, 0x3F, 0xCC, 0xE3, 0x77, 0x4D,
       0xA0, 0x10, 0x97, 0xEC, 0xE3, 0xA7, 0x10, 0x75, 0x76, 0xF6, 0xF6, 0x84,
       0xCF, 0x33, 0x35, 0x70, 0x3F, 0x97, 0xC9, 0xA6, 0x45, 0x8F, 0x65, 0xC8,
       0x90, 0x31, 0xBE, 0xE9, 0xC6, 0xC8, 0x69, 0x3A, 0x92, 0xD1, 0x64, 0xC8,
       0x96, 0x00, 0x5F, 0x93, 0x5D, 0x29, 0x24, 0xA4, 0x91, 0x34, 0xFC, 0x14,
       0x0D, 0x91, 0x2F, 0xA0, 0x37, 0x0A, 0x3B, 0xD6, 0xB2, 0x2B, 0xF6, 0x17,
       0xF3, 0x8F, 0xC1, 0x32, 0xB3, 0x20, 0x86, 0x4B, 0xBC, 0x22, 0x75, 0x06,
       0x29, 0x3F, 0x1B, 0xF2, 0xD1, 0xA7, 0xA7, 0x8F, 0x43, 0xDC, 0x6B, 0x10,
       0x4D, 0xAB, 0xF1, 0x0F, 0x1F, 0xEF, 0x76, 0x81, 0x8E, 0x0E, 0xDA, 0xC1,
       0xCB, 0x6D, 0x5F, 0x3A, 0x4B, 0xB0, 0x58, 0xE6, 0x82, 0x22, 0xAA, 0x59,
       0x8C, 0x65, 0x97, 0xF2, 0xAD, 0xA0, 0xDE, 0x96, 0xA8, 0x2F, 0xDE, 0x1D,
       0x5D, 0x89, 0x92, 0x6A, 0x0C, 0xEA, 0x58, 0x6F, 0xCE, 0x97, 0x40, 0x3A,
       0xED, 0x44, 0xA9, 0x2F, 0xE3, 0x64, 0xEE, 0x48, 0xED, 0x79, 0x10, 0xF6,
       0xA6, 0x80, 0x7D, 0x16, 0x16, 0xC1, 0x9D, 0xE5, 0x0A, 0xEB, 0xAB, 0xF4,
       0x9B, 0xB2, 0x9C, 0x2B, 0x98, 0x06, 0xEA, 0xCE, 0xD2, 0xFC, 0x8A, 0xE1,
       0x4E, 0x84, 0x44, 0x10, 0x46, 0x59, 0x33, 0x46, 0x86, 0x64, 0xEA, 0x72,
       0xEA, 0x01, 0x36, 0x95, 0x2E, 0xE4, 0x5B, 0xB2, 0x74, 0x16, 0x82, 0xE2,
       0xB0, 0x30, 0x21, 0xE7, 0x63, 0x73, 0xD4, 0xFD, 0xDB, 0xF0, 0xB5, 0xA7,
       0x17, 0x0A, 0x3B, 0x4B, 0x73, 0x51, 0x05, 0x85, 0xF3, 0x4A, 0xFF, 0xCE,
       0xDA, 0xD0, 0x3C, 0x4C, 0x16, 0x1F, 0x21, 0xF8, 0x2F, 0x47, 0x4D, 0x93,
       0x77, 0x79, 0xCE, 0xD8, 0x45, 0xFC, 0x66, 0x13, 0x7E, 0xB5, 0x9A, 0xF6,
       0xCD, 0xB3, 0x60, 0x4E, 0xE1, 0x7D, 0x06, 0xAF, 0x7A, 0x89, 0xC1, 0x1B,
       0xCA, 0xF5, 0x12, 0x4D, 0x9B, 0xB0, 0x1C, 0x72, 0x3F, 0x05, 0xD4, 0x2C,
       0xE8, 0x49, 0x43, 0x70, 0x38, 0xBD, 0xD1, 0x06, 0x68, 0xC3, 0xD0, 0x26,
       0x2A, 0xA2, 0x41, 0xA1, 0xB2, 0xDC, 0x85, 0x25, 0x4F, 0x2E, 0x45, 0xAC,
       0xEB, 0x8E, 0x3E, 0x64, 0x92, 0x20, 0xEA, 0x1E, 0x3C, 0xD2, 0x70, 0xE0,
       0x11, 0x38, 0x1F, 0x4D, 0x3F, 0x48, 0xE0, 0x15, 0x36, 0xA6, 0xA1, 0x88,
       0x06, 0xA3, 0x6E, 0xA0, 0xA1, 0x12, 0x68, 0xEA, 0x41, 0xC2, 0x11, 0x85,
       0xFF, 0x8D, 0x6D, 0xB9, 0xB2, 0x22, 0xCF, 0xF2, 0x1C, 0x65, 0xEB, 0xDD,
       0x2E, 0xB6, 0x50, 0x42, 0xE2, 0x26, 0x51, 0x39, 0x61, 0xB2, 0xC5, 0x18,
       0x2C, 0xF1, 0x8D, 0x76, 0xDB, 0x9E, 0xAA, 0x71, 0xC1, 0x1A, 0x41, 0xB6,
       0x20, 0x2D, 0xFE, 0x8A, 0x6F, 0x02, 0x01, 0x05,
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
