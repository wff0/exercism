package highscores

import (
	"slices"
)

type HighScores struct {
	oriScores    []int
	sortedScores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	ori := slices.Clone(scores)
	//copy(ori, scores)
	slices.SortFunc(scores, func(a, b int) int {
		return b - a
	})

	return &HighScores{
		oriScores:    ori,
		sortedScores: scores,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.oriScores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.oriScores[len(s.oriScores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return s.sortedScores[0]
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	maxL := min(3, len(s.sortedScores))
	return s.sortedScores[0:maxL]
}
