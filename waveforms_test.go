package waveforms

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"time"
)

func TestWaveform_Sine(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(50, amp, 0.0)

	assert.True(t, math.Abs(wv.Sine(1)-0.12533) < 1.0e-5)
	assert.True(t, math.Abs(wv.Sine(13)-0.99802) < 1.0e-5)
	assert.True(t, math.Abs(wv.Sine(24)-0.12533) < 1.0e-5)
	assert.True(t, math.Abs(wv.Sine(40) - -0.95105) < 1.0e-5)
}

func TestWaveform_Square(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(50, amp, 0.0)

	assert.True(t, math.Abs(wv.Square(1)-1) < 1.0e-5)
	assert.True(t, math.Abs(wv.Square(13)-1) < 1.0e-5)
	assert.True(t, math.Abs(wv.Square(24)-1) < 1.0e-5)
	assert.True(t, math.Abs(wv.Square(40) - -1) < 1.0e-5)
	assert.True(t, math.Abs(wv.Square(50)-1) < 1.0e-5)
}

func TestWaveform_Triangle(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(50, amp, 0.0)

	assert.True(t, math.Abs(wv.Triangle(1)-0.08) < 1.0e-5)
	assert.True(t, math.Abs(wv.Triangle(13)-0.96) < 1.0e-5)
	assert.True(t, math.Abs(wv.Triangle(24)-0.08) < 1.0e-5)
	assert.True(t, math.Abs(wv.Triangle(40) - -0.8) < 1.0e-5)
}

func TestWaveform_Sawtooth(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(50, amp, 0.0)

	assert.True(t, math.Abs(wv.Sawtooth(1)-0.04) < 1.0e-5)
	assert.True(t, math.Abs(wv.Sawtooth(13)-0.52) < 1.0e-5)
	assert.True(t, math.Abs(wv.Sawtooth(24)-0.96) < 1.0e-5)
	assert.True(t, math.Abs(wv.Sawtooth(40) - -0.4) < 1.0e-5)
}

func TestWaveform_Stopped(t *testing.T) {
	wv := Waveform{}

	assert.False(t, wv.flowRunning)
}

func TestWaveform_SineFlow(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(1000, amp, math.Pi)

	ch, err := wv.SineFlow(10)
	assert.Nil(t, err)
	min := math.MaxFloat32
	max := -math.MaxFloat32

	go func() {
		for v := range ch {
			min = math.Min(min, v)
			max = math.Max(max, v)
		}
	}()

	time.Sleep(1 * time.Second)
	wv.StopFlow()

	assert.True(t, min > -amp)
	assert.True(t, max < amp)
}

func TestWaveform_SquareFlow(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(1000, amp, math.Pi)

	ch, err := wv.SquareFlow(10)
	assert.Nil(t, err)
	min := math.MaxFloat32
	max := -math.MaxFloat32

	go func() {
		for v := range ch {
			min = math.Min(min, v)
			max = math.Max(max, v)
		}
	}()

	time.Sleep(1 * time.Second)
	wv.StopFlow()

	assert.True(t, min >= -amp)
	assert.True(t, max <= amp)
}

func TestWaveform_SawtoothFlow(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(1000, amp, math.Pi)

	ch, err := wv.SawtoothFlow(10)
	assert.Nil(t, err)
	min := math.MaxFloat32
	max := -math.MaxFloat32

	go func() {
		for v := range ch {
			min = math.Min(min, v)
			max = math.Max(max, v)
		}
	}()

	time.Sleep(1 * time.Second)
	wv.StopFlow()

	assert.True(t, min > -amp)
	assert.True(t, max < amp)
}

func TestWaveform_TriangleFlow(t *testing.T) {
	amp := 1.0
	wv := NewWaveform(1000, amp, math.Pi)

	ch, err := wv.TriangleFlow(10)
	assert.Nil(t, err)
	min := math.MaxFloat32
	max := -math.MaxFloat32

	go func() {
		for v := range ch {
			min = math.Min(min, v)
			max = math.Max(max, v)
		}
	}()

	time.Sleep(1 * time.Second)
	wv.StopFlow()

	assert.True(t, min > -amp)
	assert.True(t, max < amp)
}

// after construction.
func Test_simpleExample(t *testing.T) {
	wv := NewWaveform(50, 1, 0) // wavelength: 50, amplitude: 1, phase: 0

	for i := 1; i < 100; i++ {
		t := float64(i)
		fmt.Printf("Sine: %.5f %.5f\n", t, wv.Sine(t))
	}
}

func Test_flowExample(t *testing.T) {
	wv := NewWaveform(100, 1, 0) // wavelength: 50, amplitude: 1, phase: 0
	defer wv.StopFlow()

	ch, err := wv.SineFlow(1)

	if err != nil {
		panic(err)
	}

	go func() {
		i := 0
		for v := range ch {
			i++
			fmt.Printf("Sine: %d %.5f\n", i, v)
		}
	}()

	time.Sleep(time.Second * 1)
}
