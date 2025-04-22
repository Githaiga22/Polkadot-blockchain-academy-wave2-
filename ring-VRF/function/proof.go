package function

import (
	"crypto/sha256"
	"encoding/hex"
)

type PrivateKey struct {
	ID string // simulate private key with a string
}

type PublicKey struct {
	ID string // simulate public key with a string
}

// Simulate VRF generation: hash(message + privateKey.ID)
func GenerateVRF(privateKey PrivateKey, message []byte) (output []byte, proof []byte) {
	data := append(message, []byte(privateKey.ID)...)
	hash := sha256.Sum256(data)

	// Simulate "proof" as SHA256 of hash + privateKey.ID
	proofData := append(hash[:], []byte(privateKey.ID)...)
	proofHash := sha256.Sum256(proofData)

	return hash[:], proofHash[:]
}

// Simulate a ring signature: hash(proof + all public keys)
func CreateRingSignature(proof []byte, ring []PublicKey, privateKey PrivateKey) []byte {
	data := append([]byte{}, proof...)
	for _, pub := range ring {
		data = append(data, []byte(pub.ID)...)
	}
	// Add privateKey.ID to simulate the signer being part of the ring
	data = append(data, []byte(privateKey.ID)...)

	sig := sha256.Sum256(data)
	return sig[:]
}

// Simulate verification by recomputing what the ring signature should be
func VerifyRingVRF(message []byte, output []byte, ring []PublicKey, ringSignature []byte) bool {
	// We don’t know which private key was used, so we brute-force all possibilities
	for _, pub := range ring {
		// Simulate trying each pub.ID as if it was the private key's ID
		fakePrivate := PrivateKey{ID: pub.ID}
		_, fakeProof := GenerateVRF(fakePrivate, message)
		expectedSig := CreateRingSignature(fakeProof, ring, fakePrivate)

		// Compare
		if hex.EncodeToString(expectedSig) == hex.EncodeToString(ringSignature) {
			return true
		}
	}
	return false
}
