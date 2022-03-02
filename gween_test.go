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
