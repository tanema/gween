// Package gween provides the Tween struct that allows an easing function to be
// animated over time. This can be used in tandem with the ease package to provide
// the easing functions.
package gween

import (
	"encoding/json"
	"reflect"
	"runtime"
	"strings"

	"github.com/tanema/gween/ease"
)

type (
	// tweenData holds and hides the core information for tween, while still
	// allowing it to be serialized
	tweenData struct {
		Duration   float32
		Time       float32
		Begin      float32
		End        float32
		Change     float32
		Overflow   float32
		easing     ease.TweenFunc
		EasingName string
		Reverse    bool
	}
	// Tween encapsulates the easing function along with timing data. This allows
	// a ease.TweenFunc to be used to be easily animated.
	Tween struct {
		d tweenData
	}
)

// New will return a new Tween when passed a beginning and end value, the duration
// of the tween and the easing function to animate between the two values. The
// easing function can be one of the provided easing functions from the ease package
// or you can provide one of your own.
func New(begin, end, duration float32, easing ease.TweenFunc) *Tween {
	// TODO: This doesn't allow for anonymous / curried easing funcs
	f := strings.Split(runtime.FuncForPC(reflect.ValueOf(easing).Pointer()).Name(), ".")
	return &Tween{
		d: tweenData{
			Begin:      begin,
			End:        end,
			Change:     end - begin,
			Duration:   duration,
			easing:     easing,
			EasingName: f[len(f)-1],
			Overflow:   0,
			Reverse:    false,
		},
	}
}

// Set will set the current time along the duration of the tween. It will then return
// the current value as well as a boolean to determine if the tween is finished.
func (tween *Tween) Set(time float32) (current float32, isFinished bool) {
	switch {
	case time <= 0:
		tween.d.Overflow = time
		tween.d.Time = 0
		current = tween.d.Begin
	case time >= tween.d.Duration:
		tween.d.Overflow = time - tween.d.Duration
		tween.d.Time = tween.d.Duration
		current = tween.d.End
	default:
		tween.d.Overflow = 0
		tween.d.Time = time
		current = tween.d.easing(tween.d.Time, tween.d.Begin, tween.d.Change, tween.d.Duration)
	}

	if tween.d.Reverse {
		return current, tween.d.Time <= 0
	}
	return current, tween.d.Time >= tween.d.Duration
}

// Reset will set the Tween to the beginning of the two values.
func (tween *Tween) Reset() {
	if tween.d.Reverse {
		tween.Set(tween.d.Duration)
	} else {
		tween.Set(0)
	}
}

// Overflow return the current overflow value of the tween
func (tween *Tween) Overflow() float32 {
	return tween.d.Overflow
}

// Update will increment the timer of the Tween and ease the value. It will then
// return the current value as well as a bool to mark if the tween is finished or not.
func (tween *Tween) Update(dt float32) (current float32, isFinished bool) {
	if tween.d.Reverse {
		return tween.Set(tween.d.Time - dt)
	}
	return tween.Set(tween.d.Time + dt)
}

func (tween *Tween) MarshalJSON() ([]byte, error) {
	return json.Marshal(tween.d)
}

func (tween *Tween) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &tween.d); err != nil {
		return err
	}
	tween.d.easing = ease.EasingFunctions[tween.d.EasingName]
	return nil
}

func (tween *Tween) Equal(other *Tween) bool {
	fp1 := reflect.ValueOf(tween.d.easing).Pointer()
	fp2 := reflect.ValueOf(other.d.easing).Pointer()
	if fp1 != fp2 {
		return false
	}
	if tween.d.Duration != other.d.Duration {
		return false
	}
	if tween.d.Time != other.d.Time {
		return false
	}
	if tween.d.Begin != other.d.Begin {
		return false
	}
	if tween.d.End != other.d.End {
		return false
	}
	if tween.d.Change != other.d.Change {
		return false
	}
	if tween.d.Overflow != other.d.Overflow {
		return false
	}
	if tween.d.Reverse != other.d.Reverse {
		return false
	}
	return true
}
