package gween

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tanema/gween/ease"
)

func TestNew(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)

	assert.Equal(t, float32(0), tween.begin)
	assert.Equal(t, float32(10), tween.end)
	assert.Equal(t, float32(10), tween.change)
	assert.Equal(t, float32(10), tween.duration)
}

func TestTween_Set(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.False(t, isFinished)
	current, isFinished = tween.Set(11)
	assert.Equal(t, float32(10), current)
	assert.True(t, isFinished)
}

func TestTween_Reset(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(2), tween.time)
	assert.False(t, isFinished)
	tween.Reset()
	assert.Equal(t, float32(0), tween.time)
}

func TestTween_Update(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Update(2)
	assert.Equal(t, float32(2), current)
	assert.False(t, isFinished)
	current, isFinished = tween.Update(9)
	assert.Equal(t, float32(10), current)
	assert.True(t, isFinished)

}

func TestSequence(t *testing.T) {

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
