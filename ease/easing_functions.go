// Package ease provides default easing functions to be used in a Tween.
package ease

import (
	"math"
)

const backS float32 = 1.70158

var pi = float32(math.Pi)

// TweenFunc provides an interface used for the easing equation. You can use
// one of the provided easing functions or provide your own.
// t = current time, b = begin value, c = change from begin, d = duration
type TweenFunc func(t, b, c, d float32) float32

// Linear is a linear interpolation of some t with respect to a total duration d
// between the values b and b+c
func Linear(t, b, c, d float32) float32 {
	return c*t/d + b
}

// InQuad is a quadratic transition based on the square of t that starts slow
// and speeds up
func InQuad(t, b, c, d float32) float32 {
	return c*pow(t/d, 2) + b
}

// OutQuad is a quadratic transition based on the square of t that starts fast
// and slows down
func OutQuad(t, b, c, d float32) float32 {
	t /= d
	return -c*t*(t-2) + b
}

// InOutQuad is a quadratic transition based on the square of t that starts and
// ends slow, accelerating through the middle
func InOutQuad(t, b, c, d float32) float32 {
	t = t / d * 2
	if t < 1 {
		return c/2*pow(t, 2) + b
	}
	return -c/2*((t-1)*(t-3)-1) + b
}

// OutInQuad is a quadratic transition based on the square of t that starts and
// ends fast, slowing through the middle
func OutInQuad(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutQuad(t*2, b, c/2, d)
	}
	return InQuad((t*2)-d, b+c/2, c/2, d)
}

// InCubic is a cubic transition based on the cube of t that starts slow and
// speeds up
func InCubic(t, b, c, d float32) float32 {
	return c*pow(t/d, 3) + b
}

// OutCubic is a cubic transition based on the cube of t that starts fast and
// slows down
func OutCubic(t, b, c, d float32) float32 {
	return c*(pow(t/d-1, 3)+1) + b
}

// InOutCubic is a cubic transition based on the cube of t that starts and ends
// slow, accelerating through the middle
func InOutCubic(t, b, c, d float32) float32 {
	t = t / d * 2
	if t < 1 {
		return c/2*t*t*t + b
	}
	t -= 2
	return c/2*(t*t*t+2) + b
}

// OutInCubic is a cubic transition based on the cube of t that starts and ends
// fast, slowing through the middle
func OutInCubic(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutCubic(t*2, b, c/2, d)
	}
	return InCubic((t*2)-d, b+c/2, c/2, d)
}

// InQuart is a quartic transition based on the fourth power of t that starts
// slow and speeds up
func InQuart(t, b, c, d float32) float32 {
	return c*pow(t/d, 4) + b
}

// OutQuart is a quartic transition based on the fourth power of t that starts
// fast and slows down
func OutQuart(t, b, c, d float32) float32 {
	return -c*(pow(t/d-1, 4)-1) + b
}

// InOutQuart is a quartic transition based on the fourth power of t that starts
// and ends slow, accelerating through the middle
func InOutQuart(t, b, c, d float32) float32 {
	t = t / d * 2
	if t < 1 {
		return c/2*pow(t, 4) + b
	}
	return -c/2*(pow(t-2, 4)-2) + b
}

// OutInQuart is a quartic transition based on the fourth power of t that starts
// and ends fast, slowing through the middle
func OutInQuart(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutQuart(t*2, b, c/2, d)
	}
	return InQuart((t*2)-d, b+c/2, c/2, d)
}

// InQuint is a quintic transition based on the fifth power of t that starts
// slow and speeds up
func InQuint(t, b, c, d float32) float32 {
	return c*pow(t/d, 5) + b
}

// OutQuint is a quintic transition based on the fifth power of t that starts
// fast and slows down
func OutQuint(t, b, c, d float32) float32 {
	return c*(pow(t/d-1, 5)+1) + b
}

// InOutQuint is a quintic transition based on the fifth power of t that starts
// and ends slow, accelerating through the middle
func InOutQuint(t, b, c, d float32) float32 {
	t = t / d * 2
	if t < 1 {
		return c/2*pow(t, 5) + b
	}
	return c/2*(pow(t-2, 5)+2) + b
}

// OutInQuint is a quintic transition based on the fifth power of t that starts
// and ends fast, slowing through the middle
func OutInQuint(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutQuint(t*2, b, c/2, d)
	}
	return InQuint((t*2)-d, b+c/2, c/2, d)
}

// InSine is a sinusoidal transition based on the cosine of t that starts slow
// and speeds up
func InSine(t, b, c, d float32) float32 {
	return -c*cos(t/d*(pi/2)) + c + b
}

// OutSine is a sinusoidal transition based on the sine or cosine of t that
// starts fast and slows down
func OutSine(t, b, c, d float32) float32 {
	return c*sin(t/d*(pi/2)) + b
}

// InOutSine is a sinusoidal transition based on the cosine of t that starts and
// ends slow, accelerating through the middle
func InOutSine(t, b, c, d float32) float32 {
	return -c/2*(cos(pi*t/d)-1) + b
}

// OutInSine is a sinusoidal transition based on the sine or cosine of t that
// starts and ends fast, slowing through the middle
func OutInSine(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutSine(t*2, b, c/2, d)
	}
	return InSine((t*2)-d, b+c/2, c/2, d)
}

// InExpo is a exponential transition based on the 2 to power 10*t that starts
// slow and speeds up
func InExpo(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}
	return c*pow(2, 10*(t/d-1)) + b - c*0.001
}

// OutExpo is a exponential transition based on the 2 to power 10*t that starts
// fast and slows down
func OutExpo(t, b, c, d float32) float32 {
	if t == d {
		return b + c
	}
	return c*1.001*(-pow(2, -10*t/d)+1) + b
}

// InOutExpo is a exponential transition based on the 2 to power 10*t that
// starts and ends slow, accelerating through the middle
func InOutExpo(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}
	if t == d {
		return b + c
	}
	t = t / d * 2
	if t < 1 {
		return c/2*pow(2, 10*(t-1)) + b - c*0.0005
	}
	return c/2*1.0005*(-pow(2, -10*(t-1))+2) + b
}

// OutInExpo is a exponential transition based on the 2 to power 10*t that
// starts and ends fast, slowing through the middle
func OutInExpo(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutExpo(t*2, b, c/2, d)
	}
	return InExpo((t*2)-d, b+c/2, c/2, d)
}

// InCirc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts slow and speeds up
func InCirc(t, b, c, d float32) float32 {
	return -c*(sqrt(1-pow(t/d, 2))-1) + b
}

// OutCirc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts fast and slows down
func OutCirc(t, b, c, d float32) float32 {
	return c*sqrt(1-pow(t/d-1, 2)) + b
}

// InOutCirc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts and ends slow, accelerating through
// the middle
func InOutCirc(t, b, c, d float32) float32 {
	t = t / d * 2
	if t < 1 {
		return -c/2*(sqrt(1-t*t)-1) + b
	}
	t -= 2
	return c/2*(sqrt(1-t*t)+1) + b
}

// OutInCirc is a circular transition based on the equation for half of a circle,
// taking the square root of t, that starts and ends fast, slowing through the
// middle
func OutInCirc(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutCirc(t*2, b, c/2, d)
	}
	return InCirc((t*2)-d, b+c/2, c/2, d)
}

// InElastic is an elastic transition that wobbles around from the start value,
// extending past start and away from end, and then accelerates towards the end
// value at the end of the transition.
func InElastic(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}
	t /= d
	if t == 1 {
		return b + c
	}
	p, a, s := calculatePAS(c, d)
	t--
	return -(a * pow(2, 10*t) * sin((t*d-s)*(2*pi)/p)) + b
}

// OutElastic is an elastic transition that accelerates quickly away from the
// start and beyond the end value and then wobbles towards the end value at the
// end of the transition.
func OutElastic(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}
	t /= d
	if t == 1 {
		return b + c
	}
	p, a, s := calculatePAS(c, d)
	return a*pow(2, -10*t)*sin((t*d-s)*(2*pi)/p) + c + b
}

// InOutElastic is an elastic transition that wobbles around from the start
// value, towards the middle of the transition extending beyond start away from
// end, then rapidly toward, and beyond end value, then wobbling toward end
func InOutElastic(t, b, c, d float32) float32 {
	if t == 0 {
		return b
	}
	t = t / d * 2
	if t == 2 {
		return b + c
	}
	p, a, s := calculatePAS(c, d)
	t--
	if t < 0 {
		return -0.5*(a*pow(2, 10*t)*sin((t*d-s)*(2*pi)/p)) + b
	}
	return a*pow(2, -10*t)*sin((t*d-s)*(2*pi)/p)*0.5 + c + b
}

// OutInElastic is an elastic transition that accelerates towards and beyond the
// average of the start and end values, wobbles toward the average, wobbles out
// and slight away from end before accelerating toward the end value
func OutInElastic(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutElastic(t*2, b, c/2, d)
	}
	return InElastic((t*2)-d, b+c/2, c/2, d)
}

// InBack is a much like InQuint, but extends beyond the start away from end
// before snapping quickly to the end
func InBack(t, b, c, d float32) float32 {
	t /= d
	return c*t*t*((backS+1)*t-backS) + b
}

// OutBack is a much like OutQuint, but extends beyond the end away from start
// before easing toward end
func OutBack(t, b, c, d float32) float32 {
	t = t/d - 1
	return c*(t*t*((backS+1)*t+backS)+1) + b
}

// InOutBack is a much like InOutQuint, but extends beyond both start and end
// values on both sides of the transition
func InOutBack(t, b, c, d float32) float32 {
	s := backS * 1.525
	t = t / d * 2
	if t < 1 {
		return c/2*(t*t*((s+1)*t-s)) + b
	}
	t -= 2
	return c/2*(t*t*((s+1)*t+s)+2) + b
}

// OutInBack is a much like OutInQuint, but extends beyond the average of start
// and end during the middle of the transition
func OutInBack(t, b, c, d float32) float32 {
	if t < (d / 2) {
		return OutBack(t*2, b, c/2, d)
	}
	return InBack((t*2)-d, b+c/2, c/2, d)
}

// OutBounce is a bouncing transition that accelerates toward the end value and
// then bounces back slightly in decreasing amounts until coming to reset at end
func OutBounce(t, b, c, d float32) float32 {
	t /= d
	if t < 1/2.75 {
		return c*(7.5625*t*t) + b
	}
	if t < 2/2.75 {
		t -= 1.5 / 2.75
		return c*(7.5625*t*t+0.75) + b
	} else if t < 2.5/2.75 {
		t -= 2.25 / 2.75
		return c*(7.5625*t*t+0.9375) + b
	}
	t -= 2.625 / 2.75
	return c*(7.5625*t*t+0.984375) + b
}

// InBounce is a bouncing transition that slowly bounces away from start at
// increasing amounts before finally accelerating toward end
func InBounce(t, b, c, d float32) float32 {
	return c - OutBounce(d-t, 0, c, d) + b
}

// InOutBounce is a bouncing transition that bounces off of the start value,
// then accelerates toward the average of start and end, then does the opposite
// toward the end value
func InOutBounce(t, b, c, d float32) float32 {
	if t < d/2 {
		return InBounce(t*2, 0, c, d)*0.5 + b
	}
	return OutBounce(t*2-d, 0, c, d)*0.5 + c*.5 + b
}

// OutInBounce is a bouncing transition that accelerates toward the average of
// start and end, bouncing off of the average toward start, then flips and
// bounces off of average toward end in increasing amounts before accelerating
// toward end
func OutInBounce(t, b, c, d float32) float32 {
	if t < d/2 {
		return OutBounce(t*2, b, c/2, d)
	}
	return InBounce((t*2)-d, b+c/2, c/2, d)
}

func calculatePAS(c, d float32) (p, a, s float32) {
	p = d * 0.3
	return p, c, p / 4
}

func pow(x, y float32) float32 {
	return float32(math.Pow(float64(x), float64(y)))
}

func cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}
