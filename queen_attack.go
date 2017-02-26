package queenattack

import (
	"errors"
	"math"
	"strconv"
)

const testVersion = 2

var maxX = int64(17)
var minX = int64(10)
var maxY = int64(8)
var minY = int64(1)

func CanQueenAttack(w, b string) (attrack bool, err error) {
	wX, wY, errW := GetPosition(w)
	bX, bY, errB := GetPosition(b)
	if err = errW; errB != nil {
		err = errB
	}
	if err != nil {
		return false, err
	}
	if w == b {
		return false, errors.New("same square")
	}
	if wX == bX || wY == bY || math.Abs(wX-bX) == math.Abs(wY-bY) {
		return true, nil
	}
	return false, nil
}

func GetPosition(w string) (float64, float64, error) {
	if len(w) != 2 {
		return 0, 0, errors.New(w + " invalid")
	}
	x, errWX := strconv.ParseInt(w[0:1], 32, 32)
	y, errWY := strconv.ParseInt(w[1:], 32, 32)
	if errWX != nil || errWY != nil {
		return 0, 0, errors.New(w + " invalid")
	}
	if x > maxX || y > maxY || x < minX || y < minY {
		return 0, 0, errors.New(w + " off board")
	}
	return float64(x), float64(y), nil
}
