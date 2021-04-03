package dsim

type (
	// Blockchain is an interface for blockchains
	Blockchain interface {
		Height() uint64
		AddBlock(blockTime uint64)
		Name() string
		AlgorithmName() string
		Statistics() (sd, mean float64)
		Difficulty() uint64
	}
	// Algorithm is an interface for difficulty algorithms
	Algorithm interface {
		Name() string
		NextDifficulty(chain []*Block) uint64
	}
)
