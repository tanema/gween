package gween

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tanema/gween/ease"
)

func TestSequenceNew(t *testing.T) {
	seq := NewSequence(New(0, 1, 1, ease.Linear))

	current, finishedTween, seqFinished := seq.Update(0.0)
	assert.Equal(t, float32(0), current)
	assert.False(t, finishedTween)
	assert.False(t, seqFinished)
	assert.Equal(t, 0, seq.index)
}

func TestSequence_Update(t *testing.T) {
	seq := NewSequence(
		New(0, 1, 1, ease.Linear),
		New(1, 2, 1, ease.Linear),
	)

	current, finishedTween, seqFinished := seq.Update(0.5)
	assert.Equal(t, float32(0.5), current)
	assert.False(t, finishedTween)
	assert.False(t, seqFinished)
	assert.Equal(t, 0, seq.index)
}

func TestSequence_Reset(t *testing.T) {
	seq := NewSequence(
		New(0, 1, 1, ease.Linear),
		New(1, 2, 1, ease.Linear),
	)

	seq.Update(1.5)
	seq.Reset()
	assert.Equal(t, 0, seq.index)
	assert.Equal(t, float32(0.0), seq.Tweens[0].time)
	assert.Equal(t, float32(0.0), seq.Tweens[0].Overflow)
	assert.Equal(t, float32(0.0), seq.Tweens[1].time)
	assert.Equal(t, float32(0.0), seq.Tweens[1].Overflow)
}

func TestSequence_CompleteFirst(t *testing.T) {
	seq := NewSequence(
		New(0, 1, 1, ease.Linear),
		New(1, 2, 1, ease.Linear),
	)

	current, finishedTween, seqFinished := seq.Update(1.0)
	assert.Equal(t, float32(1.0), current)
	assert.True(t, finishedTween)
	assert.False(t, seqFinished)
	assert.Equal(t, 1, seq.index)
}

func TestSequence_OverflowSecond(t *testing.T) {
	seq := NewSequence(
		New(0, 1, 1, ease.Linear),
		New(1, 2, 1, ease.Linear),
	)

	current, finishedTween, seqFinished := seq.Update(1.5)
	assert.Equal(t, float32(1.5), current)
	assert.True(t, finishedTween)
	assert.False(t, seqFinished)
	assert.Equal(t, 1, seq.index)
}

func TestSequence_OverflowAndComplete(t *testing.T) {
	seq := NewSequence(
		New(0, 1, 1, ease.Linear),
		New(1, 2, 1, ease.Linear),
		New(2, 3, 1, ease.Linear),
	)

	current, finishedTween, seqFinished := seq.Update(3.5)
	assert.Equal(t, float32(3.0), current)
	assert.True(t, finishedTween)
	assert.True(t, seqFinished)
	assert.Equal(t, 3, seq.index)
}

func TestSequence_Remove(t *testing.T) {
	seq := NewSequence(
		New(0, 1, 1, ease.Linear),
		New(1, 2, 1, ease.Linear),
		New(2, 3, 1, ease.Linear),
		New(3, 4, 1, ease.Linear),
		New(4, 5, 1, ease.Linear),
	)
	assert.Equal(t, 5, len(seq.Tweens))
	seq.Remove(2)
	assert.Equal(t, 4, len(seq.Tweens))
	current, finishedTween, seqFinished := seq.Update(2.5)
	assert.Equal(t, float32(3.5), current)
	assert.True(t, finishedTween)
	assert.False(t, seqFinished)
	assert.Equal(t, 2, seq.index)
	seq.Remove(0)
	assert.Equal(t, 3, len(seq.Tweens))
	seq.Remove(0)
	assert.Equal(t, 2, len(seq.Tweens))
	seq.Remove(0)
	assert.Equal(t, 1, len(seq.Tweens))
	// Out of bound checking
	seq.Remove(0)
	assert.Equal(t, 0, len(seq.Tweens))
	seq.Remove(2)
	assert.Equal(t, 0, len(seq.Tweens))
}

func TestSequence_Has(t *testing.T) {
	seq := NewSequence()
	assert.False(t, seq.HasTweens())
	seq.Add(New(0, 5, 1, ease.Linear))
	assert.True(t, seq.HasTweens())
	seq.Remove(0)
	assert.False(t, seq.HasTweens())
}

func TestSequence_SetIndex(t *testing.T) {
	seq := NewSequence(
		New(0, 1, 1, ease.Linear),
		New(1, 2, 1, ease.Linear),
	)
	seq.SetIndex(1)
	current, finishedTween, seqFinished := seq.Update(1.5)
	assert.Equal(t, float32(2), current)
	assert.True(t, finishedTween)
	assert.True(t, seqFinished)
	assert.Equal(t, 2, seq.index)
}

func TestSequence_RealWorld(t *testing.T) {

	seq := NewSequence(
		New(0, 5, 1, ease.Linear),
		New(5, 0, 1, ease.Linear),
		New(0, 2, 2, ease.Linear),
		New(2, 0, 2, ease.Linear),
		New(0, 1, 100, ease.Linear),
	)

	assert.True(t, len(seq.Tweens) == 5)
	seq.Remove(0)
	seq.Remove(0)
	assert.True(t, len(seq.Tweens) == 3)

	current, finishedTween, sequenceFinished := seq.Update(1)
	// Half-way through first tween
	assert.Equal(t, float32(1), current)
	assert.False(t, finishedTween)
	assert.False(t, sequenceFinished)

	current, finishedTween, sequenceFinished = seq.Update(1)
	// Now at the start of the second tween
	assert.Equal(t, float32(2), current)
	assert.Equal(t, seq.Index(), 1)
	assert.True(t, finishedTween)
	assert.False(t, sequenceFinished)

	current, finishedTween, sequenceFinished = seq.Update(2)
	// Now at the start of the third Tween
	assert.Equal(t, seq.Index(), 2)
	assert.False(t, sequenceFinished)

	seq.Remove(2)
	current, finishedTween, sequenceFinished = seq.Update(1)
	// Now finished because we removed the third tween and then called Sequence.Update()
	assert.False(t, finishedTween)
	assert.True(t, sequenceFinished)
}
