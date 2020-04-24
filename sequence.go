package gween

// Sequence represents a sequence of Tweens, executed one after the other.
type Sequence struct {
	Tweens []*Tween
	index  int
}

// NewSequence returns a new Sequence object.
func NewSequence(tweens ...*Tween) *Sequence {
	seq := &Sequence{
		Tweens: tweens,
	}
	return seq
}

// Add adds one or more Tweens in order to the Sequence.
func (seq *Sequence) Add(tweens ...*Tween) {
	seq.Tweens = append(seq.Tweens, tweens...)
}

// Remove removes a Tween of the specified index from the Sequence.
func (seq *Sequence) Remove(index int) {
	seq.Tweens = append(seq.Tweens[:index], seq.Tweens[index+1:]...)
}

// Update updates the currently active Tween in the Sequence; once that Tween is done, the Sequence moves onto the next one.
// Update() returns the current Tween's output, whether that Tween is complete, and whether the entire Sequence is complete.
func (seq *Sequence) Update(dt float32) (float32, bool, bool) {

	value := float32(0.0)
	tweenComplete := false
	sequenceComplete := false

	if seq.index < len(seq.Tweens) {

		value, tweenComplete = seq.Tweens[seq.index].Update(dt)

		if tweenComplete {
			seq.Tweens[seq.index].Reset()
			seq.index++
			if seq.index >= len(seq.Tweens) {
				sequenceComplete = true
			}
		}

	} else {
		sequenceComplete = true
	}

	return value, tweenComplete, sequenceComplete

}

// Index returns the current index of the Sequence. Note that this can exceed the number of Tweens in the Sequence.
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
