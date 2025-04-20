package functions

// Build a basic Merkle Tree and return root
func buildMerkleTree(leaves []string) string {
	if len(leaves) == 1 {
		return leaves[0]
	}

	var parents []string
	for i := 0; i < len(leaves); i += 2 {
		if i+1 < len(leaves) {
			parents = append(parents, hash(leaves[i]+leaves[i+1]))
		} else {
			// odd number of nodes, duplicate last
			parents = append(parents, hash(leaves[i]+leaves[i]))
		}
	}
	return buildMerkleTree(parents)
}
