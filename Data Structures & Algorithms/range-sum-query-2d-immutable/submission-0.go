type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{prefix: [][]int{{0}}}
	}
	n := len(matrix)
	m := len(matrix[0])
	prefix := make([][]int, n+1)
	for i := range prefix {
		prefix[i] = make([]int, m+1)
	}

	for i := range n {
		for j := range m {
			prefix[i+1][j+1] = matrix[i][j] + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
		}
	}
	return NumMatrix{prefix: prefix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	i, j := row1+1, col1+1
	k, l := row2+1, col2+1
	return this.prefix[k][l] - this.prefix[k][j-1] - this.prefix[i-1][l] + this.prefix[i-1][j-1]
}
