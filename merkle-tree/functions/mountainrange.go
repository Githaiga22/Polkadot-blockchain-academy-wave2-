package functions

type Mountain struct {
	Leaves []string
	Root   string
}

// Build a complete Merkle Tree from data of size 2^n
func buildMountain(data []string) Mountain {
	var hashed []string
	for _, d := range data {
		hashed = append(hashed, hash(d))
	}
	return Mountain{
		Leaves: data,
		Root:   buildMerkleTree(hashed),
	}
}

// Build MMR from growing data
func BuildMMR(data []string) []Mountain {
	var mmr []Mountain
	n := len(data)
	start := 0

	for n > 0 {
		size := 1
		for size*2 <= n {
			size *= 2
		}
		mountain := buildMountain(data[start : start+size])
		mmr = append(mmr, mountain)
		start += size
		n -= size
	}

	return mmr
}
