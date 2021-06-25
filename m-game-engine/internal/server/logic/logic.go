package logic

// Inital scores to play the game
var pastFourScores = []float64{5.0, 4.0, 2.0, 1.0}

const MIN_SIZE = 10.0
const MAX_SIZE = 2000.0

// Get size of next figure
func GetSize() float64 {
	oldScores := pastFourScores[0] + pastFourScores[1]
	newScores := pastFourScores[2] + pastFourScores[3]

	diff := newScores - oldScores

	// check if the game is too hard for current user; then make it easier by upsizing figures
	if diff > 0.0 {
		size := 600.0 + diff*60.0
		if size < MAX_SIZE {
			return size
		}
		return MAX_SIZE
	}

	// check if the figures is too small and it's about to get smaller; then make it be in a reasonable size
	if diff > -5.0 && diff <= 0.0 {
		return 100.0 + 18.0*diff
	}
	return MIN_SIZE
}

// Update "pastFourScores" when new score comes in
func SetScore(newScore float64) bool {
	pastFourScores = append(pastFourScores, newScore)
	pastFourScores = pastFourScores[1:]
	return true
}
