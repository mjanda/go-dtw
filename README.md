# DTW

Dynamic time warping in golang

ported from [https://github.com/langholz/dtw](https://github.com/langholz/dtw)

# Usage

```go
// prepare arrays
a := []float64{1, 1, 1, 2, 2, 2, 3, 3, 3, 2, 2, 4, 4, 4, 4}
b := []float64{1, 1, 2, 2, 3, 3, 2, 4, 4, 4}

dtw := dtw.Dtw{}

// optionally set your own distance function
dtw.DistanceFunction = func(x float64, y float64) float64 {
    difference := x - y
    return math.Sqrt(difference * difference)
}
dtw.ComputeOptimalPathWithWindow(a, b, 5) // 5 = window size
path := dtw.RetrieveOptimalPath()
```