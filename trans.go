package levenshtein

// Operation represents the type of edit operation
type Operation int

const (
	OpMatch   Operation = iota // 0: Characters match
	OpReplace                  // 1: Character substitution
	OpDelete                   // 2: Character deletion
	OpInsert                   // 3: Character insertion
)

func (o Operation) String() string {
	switch o {
	case OpMatch:
		return "match"
	case OpReplace:
		return "replace"
	case OpDelete:
		return "delete"
	case OpInsert:
		return "insert"
	default:
		return "unknown"
	}
}

// Step
type EditOp struct {
	Op Operation // Operation type (integer based)
	//Char string    // Character from the target string (or the deleted one)
	//Pos  int       // Position in the source string (0-indexed)

	SourceIdx int
	TargetIdx int
	Cost      int
}

// GetTransformations reconstructs the transformation path using a completed Levenshtein matrix.
// todo
func getTransformations(v Interface, cs Costs) []EditOp {

	dp := MakeMatrix(v, cs)

	// Обратный ход (Backtracking) в едином стиле
	i, j := v.Lens()

	var eos []EditOp

	for i > 0 || j > 0 {
		diagonalCost := cs.ReplaceCost
		currentOp := OpReplace
		if i > 0 && j > 0 && v.Match(i-1, j-1) {
			diagonalCost = cs.MatchCost
			currentOp = OpMatch
		}

		if i > 0 && j > 0 && (dp[i][j] == (dp[i-1][j-1] + diagonalCost)) {

			eo := EditOp{
				Op:        currentOp,
				SourceIdx: i - 1,
				TargetIdx: j - 1,
				Cost:      diagonalCost,
			}

			eos = append(eos, eo)
			i--
			j--
		} else if j > 0 && (dp[i][j] == dp[i][j-1]+cs.InsertCost) {

			eo := EditOp{
				Op:        OpInsert,
				SourceIdx: -1,
				TargetIdx: j - 1,
				Cost:      cs.InsertCost,
			}

			eos = append(eos, eo)
			j--
		} else if i > 0 && (dp[i][j] == (dp[i-1][j] + cs.DeleteCost)) {

			eo := EditOp{
				Op:        OpDelete,
				SourceIdx: i - 1,
				TargetIdx: -1,
				Cost:      cs.DeleteCost,
			}

			eos = append(eos, eo)
			i--
		}
	}

	flipSlice(eos)

	return eos
}

func flipSlice[T any](as []T) {
	i, j := 0, (len(as) - 1)
	for i < j {
		as[i], as[j] = as[j], as[i]
		i, j = i+1, j-1
	}
}
