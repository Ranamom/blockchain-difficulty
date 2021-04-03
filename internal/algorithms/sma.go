package algorithms

import (
	"fmt"

	"github.com/seanvaleo/dsim/internal/config"
	"github.com/seanvaleo/dsim/pkg/dsim"
)

// SMA implements a Simple Moving Average equation, using the average
// block time of the most recent X blocks to estimate a more suitable
// difficulty
type SMA struct {
	name   string
	window uint64
}

// NewSMA instantiates and returns a new SMA
func NewSMA(window uint64) *SMA {
	return &SMA{
		name:   "SMA-" + fmt.Sprint(window),
		window: window,
	}
}

// Name returns the algorithm name
func (s *SMA) Name() string {
	return s.name
}

// NextDifficulty calculates the next difficulty
func (s *SMA) NextDifficulty(chain []*dsim.Block) uint64 {

	var sumBT, meanBT, sumD, meanD uint64

	i := uint64(len(chain))
	if i < s.window {
		return chain[i-1].Difficulty
	}

	j := i - s.window

	for i > j {
		i--
		sumBT += chain[i].BlockTime
		sumD += chain[i].Difficulty
	}
	meanBT = sumBT / s.window
	meanD = sumD / s.window

	return (meanD * config.Cfg.TargetBlockTime) / meanBT
}
