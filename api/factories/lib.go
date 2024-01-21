package factories

import (
	"math/rand"
	"time"
)

type Fields map[string]interface{}

type Dependency interface{}

var (
	seed   = time.Now().UnixNano()
	random = rand.New(rand.NewSource(seed))
)

func random1000() int {
	const i = 1000

	return random.Intn(i)
}

const (
	maxID = 1000000000000
	minID = 10
)

func randomID() int {
	return rand.Intn(maxID-minID) + minID
}
