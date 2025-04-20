package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// Simulate a VRF output and proof using ECDSA signature
func main() {
	// Generate key pair (private/public)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// Input message
	input := "Hello, VRF!"
	hash := sha256.Sum256([]byte(input))

	// Sign the hash (simulated proof)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}

	// Simulate VRF Output (can be the hash or something derived from r/s)
	vrfOutput := sha256.Sum256(append(r.Bytes(), s.Bytes()...))

	// Show result
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("VRF Output: %x\n", vrfOutput)
	fmt.Printf("Proof (Signature):\n  r: %x\n  s: %x\n\n", r, s)

	// === Now verify the proof ===
	isValid := ecdsa.Verify(publicKey, hash[:], r, s)
	fmt.Printf("Is the VRF proof valid? %v\n", isValid)
}
