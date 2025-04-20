# Verifiable Random Functions (VRF) in Go

This project demonstrates a simple simulation of **Verifiable Random Functions (VRF)** using Go's `crypto/ecdsa` and `crypto/sha256` packages. VRFs are essential in blockchain and cryptographic systems where randomness needs to be **deterministic, verifiable, and tamper-proof**.

---

## 🔍 What is a VRF?

A **Verifiable Random Function** is a cryptographic function that:
- Takes an input (e.g., a message or seed)
- Produces a **pseudorandom output**
- Generates a **proof** (a digital signature) that can be verified by others using a public key

VRFs ensure that the randomness is **generated fairly and cannot be tampered with**, which is essential for:
- Blockchain validator selection
- Unbiased random events in games
- Secure lotteries or rare NFT generation

---

###  ⚙️ How it Works
### 1. 🔐 Key Generation

    An ECDSA private-public key pair is generated on the fly using the P256 curve.

### 2. 📥 Input
```go
input := "Hello, VRF!"
```
The input string is hashed using SHA-256:
```go
hash := sha256.Sum256([]byte(input))
```
### 3. ✍️ VRF Signature (Proof Generation)

The hash is signed using the private key to produce a digital signature composed of two values:
```go
r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
```
 -   r and s are the core of the VRF proof.

### 4. 📤 VRF Output

The VRF output is computed by hashing the r and s values:
```go
output := sha256.Sum256(append(r.Bytes(), s.Bytes()...))
```
This gives a pseudorandom output that is deterministic and verifiable.
### 5. ✅ VRF Verification

The proof is verified using the public key:
```go
isValid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
```
If isValid == true, then the signature is confirmed to be valid, and the output is trustworthy.
#### 📦 Example Output
```go
Input: Hello, VRF!
VRF Output: 3068bed222b08d6a23e759f1afbfaea7cf18ce211ad06b2934a1b2a1c7b3bb22
Proof (Signature):
  r: abb86abe6f6c2c77cc7f259905f5f532943170646bac2af5db0748a0ce44ee9c
  s: 5d32e4f964a46404a3938e950c8636cef22dad496362ee1c36983f9a1cf5bb8

Is the VRF proof valid? true
```
### 🧠 Concepts Learned

- 📘 ECDSA: Elliptic Curve Digital Signature Algorithm used for signing messages

- 🔐 SHA-256: Cryptographic hash function for creating fixed-length hashes

- 🧾 r and s: Components of an ECDSA digital signature (used as the VRF proof)

- 🧪 Verification: Publicly checking that the signature is valid and was generated with the correct private key

- ⚙️ Verifiability: Anyone can verify the randomness without needing access to the private key

### 🚀 Running the Project

Make sure you have Go installed. Then run:
```go
go run main.go
```
### 📌 Dependencies

    Standard Go crypto libraries (crypto/ecdsa, crypto/elliptic, crypto/sha256, crypto/rand)
```go
    Go 1.18+
```
### 🛠️ Future Improvements

- Store and reuse persistent key pairs

-   Support custom input strings

-   Implement multiple VRF rounds

-   Use curves like Ed25519 for more efficient signatures

## 👤 Author

Developed by Allan Robinson as part of the Polkadot Blockchain Academy (Wave 2) learning journey.