package geometry

import (
	"errors"
	"math"
)

func CubeVolume(edge float64) (float64, error) {
	if edge > 0 {
		return math.Pow(edge, 3), nil
	}
	return 0, errors.New("edge length must be a positive value")
}
