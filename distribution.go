package stat

import (
	"math"

	"code.google.com/p/go-fn/fn"
)

func Freedom(data Interface, twoSided bool) float64 {
	if twoSided {
		return float64(data.Len()) - 2
	}
	return float64(data.Len()) - 1
}

func PDF(df, t float64) float64 {
	pdfconst := (math.Gamma((df + 1) / 2)) / (math.Sqrt(df*math.Pi) * math.Gamma(df/2))
	return pdfconst * math.Pow(1+((t*t)/df), -(df+1)/2)
}

//Largely based on http://www.math.ucla.edu/~tom/distributions/tDist.html
func CDF(df, t float64) float64 {

	z := df / (t*t + df)
	a := df / 2

	num := fn.IB(a, .5, z)

	if t < 0 {
		return 1 - num/2
	}
	return num / 2
}
