# Merkle Mountain Range (MMR) in Go

This project is an educational implementation of a **Merkle Mountain Range (MMR)** in Go. It demonstrates how to construct multiple Merkle Trees from sequential data and group them into a "mountain range" structure — a concept useful in blockchain and cryptographic proofs.


## 🚀 How it Works

1. **Input**: A list of string data (e.g. "A", "B", "C", ...).
2. **Splitting**: The data is split into chunks of powers of two.
3. **Merkle Trees**: Each chunk builds an individual Merkle Tree.
4. **Mountain Range**: Each tree becomes a "mountain" in the range.
5. **Output**: The roots and leaves of each mountain are printed.

## 🧠 Key Concepts

- **Merkle Trees**: Binary trees where each node is the hash of its children.
- **MMRs**: A collection of several Merkle Trees constructed from chunks of data, allowing for efficient appends and verification.
- **Hashing**: SHA256 is used for consistent and secure hashing.

## 📦 Run the Program

```bash
go run main.go
```

You should see an output like:

```bash
Merkle Mountain Range:
Mountain 1 → Root: abc123de | Leaves: [A B]
Mountain 2 → Root: def456gh | Leaves: [C D E F]
...

