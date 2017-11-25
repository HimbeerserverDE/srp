//
// In SRP, the client is responsible for quite a bit of the cryptographic
// processing. This has advantages and disadvantages, but a noteworthy advantage
// is the inherent trust that comes with generating random values on one's own
// hardware.
//
// At a high level, a client is responsible for the following tasks:
//  1. Generating their salt and verifier
//  2. Initiating an authentication handshake
//  3. Generating a key based on the server's response to (2)
//  4. Providing a proof of K
//  5. Verifying a server's proof of K
//

package srp

import "errors"

// type Client struct {
//   I []byte
//   s []byte
//   v []byte
//   a []byte
//   A []byte
// }

//
// NewClient
//
// Params:
//  p ([]byte) - a client's passphrase
//
// Return:
//  []byte - a client's salt
//  []byte - a client's verifer
//  error
//
func NewClient(p []byte) ([]byte, []byte, error) {
  // Generate a random salt. Default to 16 bytes.
  s, err := randomBytes(16)
  if err != nil {
    return nil, nil, err
  }

  // Compute the secret "x" value
  x := Hash(s, p)

  // Calculate the verifer.
  v := dGrp.exp(dGrp.g, x)

  return s, v, nil
}

//
// InitiateHandshake
//
// Params:
//  None
//
// Return:
//  []byte - the client's public key "A" to be sent to the server
//  []byte - the client's private key "a" to complete the handshake
//  error
//
func InitiateHandshake() ([]byte, []byte, error) {
  // Create a random secret "a" value
  a, err := randomBytes(32)
  if err != nil {
    return nil, nil, err
  }

  // Calculate "A" based on "a"
  A := dGrp.exp(dGrp.g, a)

  return A, a, nil
}

//
// CompleteHandshake
//
// Params:
//  A ([]byte) - the client's session public key
//  a ([]byte) - the client's session private key
//  p ([]byte) - the client's secret passphrase
//  s ([]byte) - the client's salt looked up by the server
//  B ([]byte) - the server's public key for this session
//
// Return:
//  []byte - the client's computed session key
//  error
//
func CompleteHandshake(A, a, p, s, B []byte) ([]byte, error) {
  // "B" cannot be zero
  if isZero(B) {
    return nil, errors.New("\"B\" value is zero. Aborting handshake.")
  }

  // Calculate "u"
  u := Hash(A, B)

  // "u" cannot be zero
  if isZero(u) {
    return nil, errors.New("\"u\" value is zero. Aborting handshake.")
  }

  // Compute the secret "x" value
  x := Hash(s, p)

  // Calculate the SRP-6a version of the multiplier parameter "k"
  k := Hash(dGrp.N, dGrp.g)

  // Compute the pseudo-session key, "S"
  //   S = (B - kg^x) ^ (a + ux)
  //
  //    let l = (B - kg^x)
  //        r = (a + ux)
  //
  //    ... so that S = l ^ r
  l := dGrp.sub(B, dGrp.mul(k, dGrp.exp(dGrp.g, x)))
  r := dGrp.add(a, dGrp.mul(u, x))
  S := dGrp.exp(l, r)

  // The actual session key is the hash of the pseudo-session key "S"
  K := Hash(S)

  // Return K
  return K, nil
}
