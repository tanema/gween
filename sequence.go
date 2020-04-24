package gween

// Sequence represents a sequence of Tweens, executed one after the other.
type Sequence struct {
    Tweens []*Tween
    index  int
}

// NewSequence returns a new Sequence object.
func NewSequence() *Sequence {
	seq := &Sequence{
		Tweens: []*Tween{},
	}
	return seq
}

// Add adds one or more tweens in order to the Sequence.
func (seq *Sequence) Add(tweens ...*Tween) {
	for _, tween := range tweens {
		seq.Tweens = append(seq.Tweens, tween)
	}
}

// Remove removes one or more specified tweens from the Sequence.
func (seq *Sequence) Remove(tweens ...*Tween) {
	for _, tween := range tweens {
		for i, t := range seq.Tweens {
			if t == tween {
				seq.Tweens = append(seq.Tweens[:i], seq.Tweens[i+1:]...)
				break
			}
		}
	}
}

// Update updates the currently active Tween in the Sequence; once it's done, the Sequence moves onto the next one.
func (seq *Sequence) Update(dt float32) (float32, bool) {

	value, done := float32(0.0), false
	allComplete := false

	if seq.index < len(seq.Tweens) {

		value, done = seq.Tweens[seq.index].Update(dt)

		if done {
			seq.Tweens[seq.index].Reset()
			seq.index++
			if seq.index >= len(seq.Tweens) {
				allComplete = true
			}
		}

	}

	return value, allComplete

}

// Index returns the current index of the Sequence.
func (seq *Sequence) Index() int {
	return seq.index
}

// SetIndex sets the current index of the Sequence, influencing which Tween is active at any given time.
func (seq *Sequence) SetIndex(index int) {
	seq.Tweens[seq.index].Reset()
	seq.index = index
}

// Reset resets the Sequence, resetting all Tweens and setting the Sequence's index back to 0.
func (seq *Sequence) Reset() {
	for _, tween := range seq.Tweens {
		tween.Reset()
	}
	seq.index = 0
}

// HasTweens returns whether the Sequence is populated with Tweens or not.
func (seq *Sequence) HasTweens() bool {
	return len(seq.Tweens) > 0
}
