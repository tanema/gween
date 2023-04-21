package gween

import (
	"encoding/json"
)

type (
	seqData struct {
		Tweens []Tween
		// Index of the current tween
		Index int
		// Yoyo makes the sequence "yoyo" back to the beginning after it reaches the end
		Yoyo bool
		// Reverse runs the sequence backwards when true
		Reverse bool
		// Loop is the initial number of loops for this sequence to make
		Loop int
		// LoopRemaining is the remaining number of times to loop through the sequence
		LoopRemaining int
	}
	// Sequence represents a sequence of Tweens, executed one after the other.
	Sequence struct {
		d seqData
	}
)

// NewSequence returns a new Sequence object.
func NewSequence(tweens ...*Tween) *Sequence {
	sd := seqData{
		Tweens:        []Tween{},
		Yoyo:          false,
		Reverse:       false,
		LoopRemaining: 1,
		Loop:          1,
	}
	for _, tween := range tweens {
		sd.Tweens = append(sd.Tweens, *tween)
	}
	seq := &Sequence{
		d: sd,
	}
	return seq
}

// Add adds one or more Tweens in order to the Sequence.
func (seq *Sequence) Add(tweens ...*Tween) {
	for _, tween := range tweens {
		seq.d.Tweens = append(seq.d.Tweens, *tween)
	}
}

// Remove removes a Tween of the specified index from the Sequence.
func (seq *Sequence) Remove(index int) {
	if index >= 0 && index < len(seq.d.Tweens) {
		seq.d.Tweens = append(seq.d.Tweens[:index], seq.d.Tweens[index+1:]...)
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
		if seq.d.Yoyo {
			if seq.d.Index < 0 {
				// Out of bounds at beginnning, loop
				seq.d.Reverse = false
				seq.d.Index = seq.clampIndex(seq.d.Index)
				if seq.d.LoopRemaining >= 1 {
					seq.d.LoopRemaining--
				}
				if seq.d.LoopRemaining == 0 || remaining == 0 {
					return seq.d.Tweens[seq.d.Index].d.Begin, len(completed) > 0, true
				}
				seq.d.Tweens[seq.d.Index].d.Reverse = seq.Reverse()
				seq.d.Tweens[seq.d.Index].Reset()
			}
			if seq.d.Index >= len(seq.d.Tweens) {
				// Out of bounds at end, yoyo
				seq.d.Reverse = true
				seq.d.Index = seq.clampIndex(seq.d.Index)

				seq.d.Tweens[seq.d.Index].d.Reverse = seq.Reverse()
				seq.d.Tweens[seq.d.Index].Reset()
			}
		} else if seq.d.Index >= len(seq.d.Tweens) || seq.d.Index <= -1 {
			// out of bounds at either end, loop
			if seq.d.LoopRemaining >= 1 {
				seq.d.LoopRemaining--
			}
			if seq.d.LoopRemaining == 0 || remaining == 0 {
				if seq.d.Reverse {
					return seq.d.Tweens[seq.clampIndex(seq.d.Index)].d.Begin, len(completed) > 0, true
				}
				return seq.d.Tweens[seq.clampIndex(seq.d.Index)].d.End, len(completed) > 0, true
			}
			seq.d.Index = seq.wrapIndex(seq.d.Index)
			seq.d.Tweens[seq.d.Index].d.Reverse = seq.Reverse()
			seq.d.Tweens[seq.d.Index].Reset()
		}
		v, tc := seq.d.Tweens[seq.d.Index].Update(remaining)
		if !tc {
			return v, len(completed) > 0, false
		}
		remaining = seq.d.Tweens[seq.d.Index].Overflow()
		completed = append(completed, seq.d.Index)
		if remaining < 0 {
			remaining *= -1
		}
		if seq.d.Reverse {
			seq.d.Index--
		} else {
			seq.d.Index++
		}
		// On the way back, tweens need to be configured to not go forward
		if seq.d.Index < len(seq.d.Tweens) && seq.d.Index >= 0 {
			seq.d.Tweens[seq.d.Index].d.Reverse = seq.Reverse()
			seq.d.Tweens[seq.d.Index].Reset()
		}
	}
}

// Index returns the current index of the Sequence. Note that this can exceed the number of Tweens in the Sequence.
func (seq *Sequence) Index() int {
	return seq.d.Index
}

// SetIndex sets the current index of the Sequence, influencing which Tween is active at any given time.
func (seq *Sequence) SetIndex(index int) {
	seq.d.Tweens[seq.d.Index].d.Reverse = seq.Reverse()
	seq.d.Tweens[seq.d.Index].Reset()
	seq.d.Index = index
}

// SetLoop sets the default loop and the current remaining loops
func (seq *Sequence) SetLoop(amount int) {
	seq.d.Loop = amount
	seq.d.LoopRemaining = seq.d.Loop
}

// SetYoyo sets whether the Sequence should yoyo off of the end of the last Tween and complete at the beginning of the first Tween
func (seq *Sequence) SetYoyo(willYoyo bool) {
	seq.d.Yoyo = willYoyo
}

// Reset resets the Sequence, resetting all Tweens and setting the Sequence's index back to 0.
func (seq *Sequence) Reset() {
	seq.d.LoopRemaining = seq.d.Loop
	for i := range seq.d.Tweens {
		seq.d.Tweens[i].Reset()
	}
	seq.d.Index = 0
}

// HasTweens returns whether the Sequence is populated with Tweens or not.
func (seq *Sequence) HasTweens() bool {
	return len(seq.d.Tweens) > 0
}

// Reverse returns whether the Sequence currently running in reverse.
func (seq *Sequence) Reverse() bool {
	return seq.d.Reverse
}

// SetReverse sets whether the Sequence will start running in reverse.
func (seq *Sequence) SetReverse(r bool) {
	if seq.d.Index >= len(seq.d.Tweens) || seq.d.Index < 0 {
		seq.d.Index = seq.clampIndex(seq.d.Index)
	}
	seq.d.Tweens[seq.d.Index].d.Reverse = r
	seq.d.Reverse = r
}

// Equal returns true when sequences match by value, otherwise false
func (seq *Sequence) Equal(other *Sequence) bool {
	if seq.d.Reverse != other.d.Reverse {
		return false
	}
	if seq.d.Yoyo != other.d.Yoyo {
		return false
	}
	if seq.d.Index != other.d.Index {
		return false
	}
	if seq.d.Loop != other.d.Loop {
		return false
	}
	if seq.d.LoopRemaining != other.d.LoopRemaining {
		return false
	}
	for i, t1 := range seq.d.Tweens {
		t2 := other.d.Tweens[i]
		if !t1.Equal(&t2) {
			return false
		}
	}
	return true
}

// clampIndex clamps the provided index to the bounds of the Tweens slice
func (seq *Sequence) clampIndex(index int) int {
	if index >= len(seq.d.Tweens) {
		index = len(seq.d.Tweens) - 1
	}
	if index < 0 {
		index = 0
	}
	return index
}

// wrapIndex wraps the provided index when it is out of bounds, otherwise returns index
func (seq *Sequence) wrapIndex(index int) int {
	if index >= len(seq.d.Tweens) {
		index = 0
	}
	if index < 0 {
		index = len(seq.d.Tweens) - 1
	}
	return index
}

func (seq *Sequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(&seq.d)
}

func (seq *Sequence) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &seq.d); err != nil {
		return err
	}
	return nil
}
