package dtw

import (
	"math"
)

type distanceFunction func(float64, float64) float64

type Dtw struct {
	m                  int
	n                  int
	distanceCostMatrix [][]float64
	similarity         float64
	DistanceFunction   distanceFunction
}

func distanceEuclidean(x float64, y float64) float64 {
	difference := x - y
	return math.Sqrt(difference * difference)
}

func (dtw *Dtw) ComputeOptimalPath(s []float64, t []float64) {
	dtw.ComputeOptimalPathWithWindow(s, t, len(s)+len(t))
}

func (dtw *Dtw) ComputeOptimalPathWithWindow(s []float64, t []float64, w int) {
	dtw.m = len(s)
	dtw.n = len(t)

	if dtw.DistanceFunction == nil {
		dtw.DistanceFunction = distanceEuclidean
	}

	var window int = dtw.m - dtw.n
	if w > window {
		window = w
	}

	distanceCostMatrix := createFloatMatrix(dtw.m+1, dtw.n+1, math.Inf(1))
	distanceCostMatrix[0][0] = 0

	for rowIndex := 1; rowIndex <= dtw.m; rowIndex++ {
		columnIndexStart := 1
		if (rowIndex - window) > 1 {
			columnIndexStart = (rowIndex - window)
		}
		columnIndexEnd := dtw.n
		if (rowIndex + window) < columnIndexEnd {
			columnIndexEnd = (rowIndex + window)
		}
		for columnIndex := columnIndexStart; columnIndex <= columnIndexEnd; columnIndex++ {
			cost := dtw.DistanceFunction(s[rowIndex-1], t[columnIndex-1])
			distanceCostMatrix[rowIndex][columnIndex] = cost + findSmallestNumber([]float64{
				distanceCostMatrix[rowIndex-1][columnIndex],
				distanceCostMatrix[rowIndex][columnIndex-1],
				distanceCostMatrix[rowIndex-1][columnIndex-1],
			})
		}

	}

	// copy into new matrix
	returnMatrix := createFloatMatrix(dtw.m, dtw.n, math.Inf(1))
	for i := 1; i <= dtw.m; i++ {
		returnMatrix[i-1] = distanceCostMatrix[i][1:(dtw.n + 1)]
	}

	dtw.distanceCostMatrix = returnMatrix
	dtw.similarity = returnMatrix[dtw.m-1][dtw.n-1]
}

func (dtw *Dtw) RetrieveOptimalPath() [][]int {
	rowIndex := dtw.m - 1
	columnIndex := dtw.n - 1
	distanceCostMatrix := dtw.distanceCostMatrix
	epsilon := 1e-14

	var path [][]int
	path = createIntMatrix(dtw.m+dtw.n, 2, 0)
	arrayIndex := len(path) - 1

	path[arrayIndex][0] = rowIndex
	path[arrayIndex][1] = columnIndex
	arrayIndex--

	for rowIndex > 0 || columnIndex > 0 {
		if rowIndex > 0 && columnIndex > 0 {
			min := findSmallestNumber([]float64{
				distanceCostMatrix[rowIndex-1][columnIndex],
				distanceCostMatrix[rowIndex][columnIndex-1],
				distanceCostMatrix[rowIndex-1][columnIndex-1]})

			if nearlyEqual(min, distanceCostMatrix[rowIndex-1][columnIndex-1], epsilon) {
				rowIndex--
				columnIndex--
			} else if nearlyEqual(min, distanceCostMatrix[rowIndex-1][columnIndex], epsilon) {
				rowIndex--
			} else if nearlyEqual(min, distanceCostMatrix[rowIndex][columnIndex-1], epsilon) {
				columnIndex--
			}
		} else if rowIndex > 0 {
			rowIndex--
		} else if columnIndex > 0 {
			columnIndex--
		}

		path[arrayIndex][0] = rowIndex
		path[arrayIndex][1] = columnIndex
		arrayIndex--
	}

	return path[arrayIndex+1 : cap(path)]
}

func createFloatMatrix(n int, m int, value float64) [][]float64 {
	matrix := make([][]float64, n)
	points := make([]float64, m*n)
	for i := range points {
		points[i] = value
	}

	for i := range matrix {
		matrix[i], points = points[:m], points[m:]
	}

	return matrix
}

func createIntMatrix(n int, m int, value int) [][]int {
	matrix := make([][]int, n)
	points := make([]int, m*n)
	for i := range points {
		points[i] = value
	}

	for i := range matrix {
		matrix[i], points = points[:m], points[m:]
	}

	return matrix
}

func findSmallestNumber(x []float64) float64 {
	var number float64 = x[0]

	for i := 0; i < len(x); i++ {
		if x[i] < number {
			number = x[i]
		}
	}
	return number
}

func nearlyEqual(i float64, j float64, epsilon float64) bool {
	iAbsolute := i
	if iAbsolute < 0 {
		iAbsolute = -iAbsolute
	}
	jAbsolute := j
	if jAbsolute < 0 {
		jAbsolute = -jAbsolute
	}
	difference := jAbsolute - iAbsolute
	if difference < 0 {
		difference = -difference
	}

	var equal bool = (i == j)
	if !equal {
		equal = difference < 2.2204460492503130808472633361816E-16
		if !equal {
			equal = difference <= math.Max(iAbsolute, jAbsolute)*epsilon
		}
	}
	return equal
}
