package gween

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tanema/gween/ease"
)

func TestNew(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)

	assert.Equal(t, float32(0), tween.d.Begin)
	assert.Equal(t, float32(10), tween.d.End)
	assert.Equal(t, float32(10), tween.d.Change)
	assert.Equal(t, float32(10), tween.d.Duration)
	assert.Equal(t, float32(0), tween.d.Time)
	assert.Equal(t, float32(0), tween.d.Overflow)
	assert.False(t, tween.d.Reverse)
}

func TestTween_Set(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(0), tween.d.Overflow)
	assert.False(t, isFinished)
	current, isFinished = tween.Set(11)
	assert.Equal(t, float32(10), current)
	assert.Equal(t, float32(1), tween.d.Overflow)
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
	tween.d.Reverse = true
	current, isFinished := tween.Set(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(0), tween.d.Overflow)
	assert.False(t, isFinished)
	current, isFinished = tween.Set(11)
	assert.Equal(t, float32(10), current)
	assert.Equal(t, float32(1), tween.d.Overflow)
	assert.False(t, isFinished)
}

func TestTween_SetNegReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.d.Reverse = true
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
	assert.Equal(t, float32(2), tween.d.Time)
	assert.Equal(t, float32(0), tween.d.Overflow)
	assert.False(t, isFinished)
	tween.Reset()
	assert.Equal(t, float32(0), tween.d.Time)
	assert.Equal(t, float32(0), tween.d.Overflow)
}

func TestTween_ResetReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.Set(2)
	tween.d.Reverse = true
	tween.Reset()
	assert.Equal(t, float32(10), tween.d.Time)
	assert.Equal(t, float32(0), tween.d.Overflow)
}

func TestTween_Update(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	current, isFinished := tween.Update(2)
	assert.Equal(t, float32(2), current)
	assert.Equal(t, float32(0), tween.d.Overflow)
	assert.False(t, isFinished)
	current, isFinished = tween.Update(9)
	assert.Equal(t, float32(10), current)
	assert.Equal(t, float32(1), tween.d.Overflow)
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
	tween.d.Reverse = true
	current, isFinished := tween.Update(-1)
	assert.Equal(t, float32(3), current)
	assert.False(t, isFinished)
}

func TestTween_Defaults_Forward(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	assert.False(t, tween.d.Reverse)
}

func TestTween_CanReverse(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.Update(8)
	tween.d.Reverse = true
	current, isFinished := tween.Update(2)
	assert.Equal(t, float32(6), current)
	assert.False(t, isFinished)
}

func TestTween_CanReverseFromFinished(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	_, isFinished := tween.Update(10)
	assert.True(t, isFinished)
	tween.d.Reverse = true
	current, isFinished := tween.Update(2)
	assert.Equal(t, float32(8), current)
	assert.False(t, isFinished)
}

func TestTween_CanReverseFromStart(t *testing.T) {
	tween := New(0, 10, 10, ease.Linear)
	tween.d.Reverse = true
	current, isFinished := tween.Update(0)
	assert.True(t, isFinished)
	assert.Equal(t, float32(0), current)
	assert.Equal(t, float32(0), tween.d.Overflow)
	current, isFinished = tween.Update(1)
	assert.True(t, isFinished)
	assert.Equal(t, float32(0), current)
	assert.Equal(t, float32(-1.0), tween.d.Overflow)
}

func TestTween_Serializes(t *testing.T) {
	tControl := New(0, 10, 10, ease.Linear)
	tControl.Update(1)

	tb, err := json.Marshal(tControl)
	assert.NoError(t, err)

	// Init with something different
	tUnmarshalled := New(-5, 5, 5, ease.InCubic)
	err = json.Unmarshal(tb, tUnmarshalled)
	assert.NoError(t, err)

	assert.True(t, tControl.Equal(tUnmarshalled))
	fp1 := reflect.ValueOf(tControl.d.easing).Pointer()
	fp2 := reflect.ValueOf(tUnmarshalled.d.easing).Pointer()
	assert.Equal(t, fp1, fp2)
	tControl.d.easing = nil
	tUnmarshalled.d.easing = nil
	assert.Equal(t, *tControl, *tUnmarshalled)
}

func MyTestEasingFunc(t, b, c, d float32) float32 {
	return c*t/d + b
}

func TestTween_SerializesCustomEasing(t *testing.T) {
	ease.EasingFunctions["MyTestEasingFunc"] = MyTestEasingFunc
	tControl := New(0, 10, 10, MyTestEasingFunc)
	tControl.Update(1)

	tb, err := json.Marshal(tControl)
	assert.NoError(t, err)

	// Init with something different
	tUnmarshalled := New(-5, 5, 5, ease.InCubic)
	err = json.Unmarshal(tb, tUnmarshalled)
	assert.NoError(t, err)

	assert.True(t, tControl.Equal(tUnmarshalled))
	fp1 := reflect.ValueOf(tControl.d.easing).Pointer()
	fp2 := reflect.ValueOf(tUnmarshalled.d.easing).Pointer()
	assert.Equal(t, fp1, fp2)
	tControl.d.easing = nil
	tUnmarshalled.d.easing = nil
	assert.Equal(t, *tControl, *tUnmarshalled)
}
