package stringid

import (
	"math/rand"
	"sync"
	"time"
)

const (
	// pushChars are the lexiographically correct base 64 characters used for
	// push-style IDs.
	pushChars = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
)

// PushGenerator is a push-style ID generator that satisifies the Generator
// interface.
type PushGenerator struct {
	// mu is the mutex lock.
	mu sync.Mutex

	// r is the random source.
	r *rand.Rand

	// stamp is the timestamp of the last ID creation, used to prevent
	// collisions if called multiple times during the same millisecond.
	stamp int64

	// stamp is comprised of 72 bits of entropy converted to 12 characters.
	// this is appended to the generated id to prevent collisions.
	// the numeric value is incremented in the event of a collision.
	last [12]int
}

// NewPushGenerator creates a new push ID generator.
func NewPushGenerator(r *rand.Rand) *PushGenerator {
	// ensure rand source
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	// create generator and random entropy
	pg := &PushGenerator{r: r}
	for i := 0; i < 12; i++ {
		pg.last[i] = r.Intn(64)
	}

	return pg
}

// Generate generates a unique, 20-character push-style ID.
func (pg *PushGenerator) Generate() string {
	var i int

	id := make([]byte, 20)

	// grab last characters
	pg.mu.Lock()
	now := time.Now().UTC().UnixNano() / 1e6
	if pg.stamp == now {
		for i = 0; i < 12; i++ {
			pg.last[i]++
			if pg.last[i] < 64 {
				break
			}
			pg.last[i] = 0
		}
	}
	pg.stamp = now

	// set last 12 characters
	for i = 0; i < 12; i++ {
		id[19-i] = pushChars[pg.last[i]]
	}
	pg.mu.Unlock()

	// set id to first 8 characters
	for i = 7; i >= 0; i-- {
		id[i] = pushChars[int(now%64)]
		now /= 64
	}

	return string(id)
}
