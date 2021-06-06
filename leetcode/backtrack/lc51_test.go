package backtrack

import "testing"

func TestSolveNQueens(t *testing.T) {
	data := map[int][][]string{
		4: {{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
	}

	for key, value := range data {
		actual := SolveNQueens(4)
		if len(actual) != len(value) {
			t.Errorf("NQueens(%d) = %v; expected %v", key, actual, value)
		}
		for i := 0; i < len(actual); i++ {
			for j := 0; j < key; j++ {
				if actual[i][j] != value[i][j] {
					t.Errorf("NQueens(%d) = %v; expected %v", key, actual, value)
				}
			}
		}
	}
}
