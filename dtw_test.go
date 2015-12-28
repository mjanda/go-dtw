package dtw

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func init() {
	log.Println("DTW test init")
}

func TestBasicPath(t *testing.T) {
	a := []float64{1, 1, 1, 2, 2, 2, 3, 3, 3, 2, 2, 4, 4, 4, 4}
	b := []float64{1, 1, 2, 2, 3, 3, 2, 4, 4, 4}
	expectedResult := [][]int{
		[]int{0, 0},
		[]int{1, 0},
		[]int{2, 1},
		[]int{3, 2},
		[]int{4, 2},
		[]int{5, 3},
		[]int{6, 4},
		[]int{7, 4},
		[]int{8, 5},
		[]int{9, 6},
		[]int{10, 6},
		[]int{11, 7},
		[]int{12, 7},
		[]int{13, 8},
		[]int{14, 9},
	}

	dtw := Dtw{}
	dtw.ComputeOptimalPathWithWindow(a, b, len(a)+len(b))

	if dtw.similarity != 0 {
		t.Error("Similarity should be 0")
	}
	path := dtw.RetrieveOptimalPath()

	if !reflect.DeepEqual(path, expectedResult) {
		for i := 0; i < len(dtw.distanceCostMatrix); i++ {
			fmt.Println(dtw.distanceCostMatrix[i])
		}
		fmt.Println(path)
		fmt.Println(expectedResult)

		t.Error("Computed path != expected path")
	}
}

func TestFloatPath(t *testing.T) {
	a := []float64{0.4125044, 0.1827033, 0.7174426, 0.5938232, 0.4614635,
		0.5900535, 0.8329995, 0.2489138, 0.0204920, 0.9778591,
		0.5764358, 0.4740868, 0.4325138, 0.3667226, 0.9619953,
		0.4576290, 0.2961572, 0.1273494, 0.1837332, 0.6646660,
		0.0533731, 0.7448532, 0.2209947, 0.1150104, 0.0697953,
		0.2262660, 0.6347957, 0.6412367, 0.0032259, 0.2817307,
		0.9574800, 0.5865984, 0.4622795, 0.7135204, 0.6929830,
		0.7052597, 0.9643922, 0.1590985, 0.8196736, 0.0813144,
		0.5112076, 0.1490992, 0.6219234, 0.4254558, 0.1539139,
		0.6556169, 0.5459852, 0.3675036, 0.6331521, 0.8443600}
	b := []float64{0.954327, 0.371023, 0.305392, 0.917947, 0.100184,
		0.636795, 0.301041, 0.726715, 0.850064, 0.362574,
		0.634449, 0.241995, 0.470016, 0.187247, 0.080302,
		0.164183, 0.337284, 0.721616, 0.228075, 0.049611,
		0.401937, 0.599079, 0.365990, 0.883565, 0.444008,
		0.879180, 0.165539, 0.220239, 0.318087, 0.356081,
		0.769599, 0.301509, 0.247175, 0.201820, 0.243712,
		0.531967, 0.682490, 0.028431, 0.627859, 0.350267,
		0.751287, 0.658828, 0.115952, 0.449262, 0.697056,
		0.479946, 0.017637, 0.727200, 0.153417, 0.467764,
		0.315294, 0.165611}
	expectedResult := [][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{0, 4},
		[]int{0, 5},
		[]int{0, 6},
		[]int{1, 7},
		[]int{2, 8},
		[]int{3, 9},
		[]int{4, 10},
		[]int{5, 11},
		[]int{6, 12},
		[]int{7, 13},
		[]int{8, 14},
		[]int{9, 15},
		[]int{10, 16},
		[]int{11, 17},
		[]int{12, 18},
		[]int{13, 19},
		[]int{14, 20},
		[]int{15, 21},
		[]int{16, 22},
		[]int{17, 23},
		[]int{18, 24},
		[]int{19, 25},
		[]int{20, 26},
		[]int{21, 27},
		[]int{22, 28},
		[]int{23, 29},
		[]int{24, 30},
		[]int{25, 31},
		[]int{26, 32},
		[]int{27, 33},
		[]int{28, 34},
		[]int{29, 34},
		[]int{30, 35},
		[]int{31, 35},
		[]int{32, 35},
		[]int{33, 36},
		[]int{34, 36},
		[]int{35, 36},
		[]int{36, 36},
		[]int{37, 37},
		[]int{38, 38},
		[]int{39, 39},
		[]int{40, 40},
		[]int{40, 41},
		[]int{41, 42},
		[]int{42, 43},
		[]int{42, 44},
		[]int{43, 45},
		[]int{44, 46},
		[]int{45, 47},
		[]int{46, 48},
		[]int{47, 49},
		[]int{48, 50},
		[]int{49, 51},
	}

	dtw := Dtw{}
	dtw.ComputeOptimalPathWithWindow(a, b, 5)

	if !approximatelyEquals(dtw.similarity, 10.38, 0.01) {
		t.Error(fmt.Sprintf("Similarity should be roughly 10.38 [is %f]", dtw.similarity))
	}
	path := dtw.RetrieveOptimalPath()

	if !reflect.DeepEqual(path, expectedResult) {
		t.Error("Computed path != expected path")
	}
}

func approximatelyEquals(a, b, epsilon float64) bool {
	if (a-b) < epsilon || (b-a) < epsilon {
		return true
	}
	return false
}
