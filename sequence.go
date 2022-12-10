package gween

// Sequence represents a sequence of Tweens, executed one after the other.
type Sequence struct {
	Tweens []*Tween
	index  int
	// yoyo makes the sequence "yoyo" back to the beginning after it reaches the end
	yoyo bool
	// reverse runs the sequence backwards when true
	reverse bool
	// loop is the initial number of loops for this sequence to make
	loop int
	// loopRemaining is the remaining number of times to loop through the sequence
	loopRemaining int
}

// NewSequence returns a new Sequence object.
func NewSequence(tweens ...*Tween) *Sequence {
	seq := &Sequence{
		Tweens:        tweens,
		yoyo:          false,
		reverse:       false,
		loopRemaining: 1,
		loop:          1,
	}
	return seq
}

// Add adds one or more Tweens in order to the Sequence.
func (seq *Sequence) Add(tweens ...*Tween) {
	seq.Tweens = append(seq.Tweens, tweens...)
}

// Remove removes a Tween of the specified index from the Sequence.
func (seq *Sequence) Remove(index int) {
	if index >= 0 && index < len(seq.Tweens) {
		seq.Tweens = append(seq.Tweens[:index], seq.Tweens[index+1:]...)
	}
}

// Update updates the currently active Tween in the Sequence; once that Tween is done, the Sequence moves onto the next one.
// Update() returns the current Tween's output, whether that Tween is complete, and whether the entire Sequence was completed
// during this Update.
func (seq *Sequence) Update(dt float32) (value float32, tweenComplete, sequenceComplete bool) {
	if !seq.HasTweens() {
		return 0, false, true
	}
	var completed []int
	remaining := dt

	for {
		if seq.yoyo {
			if seq.index < 0 {
				// Out of bounds at beginnning, loop
				seq.reverse = false
				seq.index = seq.clampIndex(seq.index)
				if seq.loopRemaining >= 1 {
					seq.loopRemaining--
				}
				if seq.loopRemaining == 0 || remaining == 0 {
					return seq.Tweens[seq.index].begin, len(completed) > 0, true
				}
				seq.Tweens[seq.index].reverse = seq.Reverse()
				seq.Tweens[seq.index].Reset()
			}
			if seq.index >= len(seq.Tweens) {
				// Out of bounds at end, yoyo
				seq.reverse = true
				seq.index = seq.clampIndex(seq.index)

				seq.Tweens[seq.index].reverse = seq.Reverse()
				seq.Tweens[seq.index].Reset()
			}
		} else if seq.index >= len(seq.Tweens) || seq.index <= -1 {
			// out of bounds at either end, loop
			if seq.loopRemaining >= 1 {
				seq.loopRemaining--
			}
			if seq.loopRemaining == 0 || remaining == 0 {
				if seq.reverse {
					return seq.Tweens[seq.clampIndex(seq.index)].begin, len(completed) > 0, true
				}
				return seq.Tweens[seq.clampIndex(seq.index)].end, len(completed) > 0, true
			}
			seq.index = seq.wrapIndex(seq.index)
			seq.Tweens[seq.index].reverse = seq.Reverse()
			seq.Tweens[seq.index].Reset()
		}
		v, tc := seq.Tweens[seq.index].Update(remaining)
		if !tc {
			return v, len(completed) > 0, false
		}
		remaining = seq.Tweens[seq.index].Overflow
		completed = append(completed, seq.index)
		if remaining < 0 {
			remaining *= -1
		}
		if seq.reverse {
			seq.index--
		} else {
			seq.index++
		}
		// On the way back, tweens need to be configured to not go forward
		if seq.index < len(seq.Tweens) && seq.index >= 0 {
			seq.Tweens[seq.index].reverse = seq.Reverse()
			seq.Tweens[seq.index].Reset()
		}
	}
}

// Index returns the current index of the Sequence. Note that this can exceed the number of Tweens in the Sequence.
func (seq *Sequence) Index() int {
	return seq.index
}

// SetIndex sets the current index of the Sequence, influencing which Tween is active at any given time.
func (seq *Sequence) SetIndex(index int) {
	seq.Tweens[seq.index].reverse = seq.Reverse()
	seq.Tweens[seq.index].Reset()
	seq.index = index
}

// SetLoop sets the default loop and the current remaining loops
func (seq *Sequence) SetLoop(amount int) {
	seq.loop = amount
	seq.loopRemaining = seq.loop
}

// SetYoyo sets whether the Sequence should yoyo off of the end of the last Tween and complete at the beginning of the first Tween
func (seq *Sequence) SetYoyo(willYoyo bool) {
	seq.yoyo = willYoyo
}

// Reset resets the Sequence, resetting all Tweens and setting the Sequence's index back to 0.
func (seq *Sequence) Reset() {
	seq.loopRemaining = seq.loop
	for _, tween := range seq.Tweens {
		tween.Reset()
	}
	seq.index = 0
}

// HasTweens returns whether the Sequence is populated with Tweens or not.
func (seq *Sequence) HasTweens() bool {
	return len(seq.Tweens) > 0
}

// Reverse returns whether the Sequence currently running in reverse.
func (seq *Sequence) Reverse() bool {
	return seq.reverse
}

// SetReverse sets whether the Sequence will start running in reverse.
func (seq *Sequence) SetReverse(r bool) {
	if seq.index >= len(seq.Tweens) || seq.index < 0 {
		seq.index = seq.clampIndex(seq.index)
	}
	seq.Tweens[seq.index].reverse = r
	seq.reverse = r
}

// clampIndex clamps the provided index to the bounds of the Tweens slice
func (seq *Sequence) clampIndex(index int) int {
	if index >= len(seq.Tweens) {
		index = len(seq.Tweens) - 1
	}
	if index < 0 {
		index = 0
	}
	return index
}

// wrapIndex wraps the provided index when it is out of bounds, otherwise returns index
func (seq *Sequence) wrapIndex(index int) int {
	if index >= len(seq.Tweens) {
		index = 0
	}
	if index < 0 {
		index = len(seq.Tweens) - 1
	}
	return index
}
