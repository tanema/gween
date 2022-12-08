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
	assert.Equal(t, float32(0), tween.time)
	assert.Equal(t, float32(0), tween.Overflow)
	assert.False(t, tween.reverse)
}

func TestTween_Set(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(0), tween.Overflow)
	assert.False(t, isFinished)
	current, isFinished = tween.Set(11)
	assert.Equal(t, float32(10), current)
	assert.Equal(t, float32(1), tween.Overflow)
	assert.True(t, isFinished)
}

func TestTween_SetNeg(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.False(t, isFinished)
	current, isFinished = tween.Set(-1)
	assert.Equal(t, float32(0), current)
	assert.False(t, isFinished)
}

func TestTween_SetReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.reverse = true
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(0), tween.Overflow)
	assert.False(t, isFinished)
	current, isFinished = tween.Set(11)
	assert.Equal(t, float32(10), current)
	assert.Equal(t, float32(1), tween.Overflow)
	assert.False(t, isFinished)
}

func TestTween_SetNegReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.reverse = true
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.False(t, isFinished)
	current, isFinished = tween.Set(-1)
	assert.Equal(t, float32(0), current)
	assert.True(t, isFinished)
}

func TestTween_Reset(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(2), tween.time)
	assert.Equal(t, float32(0), tween.Overflow)
	assert.False(t, isFinished)
	tween.Reset()
	assert.Equal(t, float32(0), tween.time)
	assert.Equal(t, float32(0), tween.Overflow)
}

func TestTween_ResetReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.Set(2)
	tween.reverse = true
	tween.Reset()
	assert.Equal(t, float32(10), tween.time)
	assert.Equal(t, float32(0), tween.Overflow)
}

func TestTween_Update(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Update(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(0), tween.Overflow)
	assert.False(t, isFinished)
	current, isFinished = tween.Update(9)
	assert.Equal(t, float32(10), current)
	assert.Equal(t, float32(1), tween.Overflow)
	assert.True(t, isFinished)
}

func TestTween_UpdateZero(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.Update(2)
	current, isFinished := tween.Update(0)
	assert.Equal(t, float32(2), current)
	assert.False(t, isFinished)
}

func TestTween_UpdateNeg(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.Update(2)
	current, isFinished := tween.Update(-1)
	assert.Equal(t, float32(1), current)
	assert.False(t, isFinished)
}

func TestTween_UpdateNegReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.Update(2)
	tween.reverse = true
	current, isFinished := tween.Update(-1)
	assert.Equal(t, float32(3), current)
	assert.False(t, isFinished)
}

func TestTween_Defaults_Forward(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	assert.False(t, tween.reverse)
}

func TestTween_CanReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.Update(8)
	tween.reverse = true
	current, isFinished := tween.Update(2)
	assert.Equal(t, float32(6), current)
	assert.False(t, isFinished)
}

func TestTween_CanReverseFromFinished(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Update(10)
	assert.True(t, isFinished)
	tween.reverse = true
	current, isFinished = tween.Update(2)
	assert.Equal(t, float32(8), current)
	assert.False(t, isFinished)
}

func TestTween_CanReverseFromStart(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.reverse = true
	current, isFinished := tween.Update(0)
	assert.True(t, isFinished)
	assert.Equal(t, float32(0), current)
	assert.Equal(t, float32(0), tween.Overflow)
	current, isFinished = tween.Update(1)
	assert.True(t, isFinished)
	assert.Equal(t, float32(0), current)
	assert.Equal(t, float32(-1.0), tween.Overflow)
}
