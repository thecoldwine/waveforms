package waveforms

import (
	"errors"
	"math"
	"time"
)

/**
Waveforms is a simple library to generate wave formed signals with given
amplitude, wavelength and phase for some period of time. This library is
useful for generating of the sample time-series data i.e some metrics or
stub telemetry.
*/

// generator is a periodic waveform generator function
type generator func(wavelength float64, amplitude float64, phase float64, t float64) float64

// sine waveform generator
func sine(wavelength float64, amplitude float64, phase float64, t float64) float64 {
	return amplitude * math.Sin((2*math.Pi*t-phase)/wavelength)
}

// square waveform generator
func square(wavelength float64, amplitude float64, phase float64, t float64) float64 {
	switch {
	case math.Mod(t-phase, wavelength) < (wavelength / 2.0):
		return amplitude
	default:
		return -amplitude
	}
}

// triangle waveform generator
func triangle(wavelength float64, amplitude float64, phase float64, t float64) float64 {
	return 2 * amplitude / math.Pi * math.Asin(math.Sin((2*math.Pi*t-phase)/wavelength))
}

// sawtooth waveform generator
func sawtooth(wavelength float64, amplitude float64, phase float64, t float64) float64 {
	return 2 * amplitude / math.Pi * math.Atan(math.Tan((2*math.Pi*t-phase)/(2*wavelength)))
}

// Waveform is a definition of a waveform.
type Waveform struct {
	wavelength float64
	amplitude  float64
	phase      float64

	ticker      *time.Ticker
	flowRunning bool
}

func NewWaveform(wavelength float64, amplitude float64, phase float64) *Waveform {
	return &Waveform{wavelength: wavelength, amplitude: amplitude, phase: phase}
}

func (wv *Waveform) run(interval int64, g generator) (out chan float64, err error) {
	out = make(chan float64)

	if wv.flowRunning {
		err = errors.New("waveform is running already")
		return out, err
	}

	wv.ticker = time.NewTicker(time.Duration(interval) * time.Millisecond)
	wv.flowRunning = true

	go func() {
		for {
			v, ok := <-wv.ticker.C

			if ok {
				out <- g(wv.wavelength, wv.amplitude, wv.phase, float64(v.UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond))))
			}
		}

		close(out)
	}()

	return out, nil
}

func (wv *Waveform) Sine(t float64) float64 {
	return sine(wv.wavelength, wv.amplitude, wv.phase, t)
}

func (wv *Waveform) Square(t float64) float64 {
	return square(wv.wavelength, wv.amplitude, wv.phase, t)
}

func (wv *Waveform) Triangle(t float64) float64 {
	return triangle(wv.wavelength, wv.amplitude, wv.phase, t)
}

func (wv *Waveform) Sawtooth(t float64) float64 {
	return sawtooth(wv.wavelength, wv.amplitude, wv.phase, t)
}

func (wv *Waveform) SineFlow(interval int64) (out chan float64, err error) {
	return wv.run(interval, sine)
}

func (wv *Waveform) SquareFlow(interval int64) (out chan float64, err error) {
	return wv.run(interval, square)
}

func (wv *Waveform) TriangleFlow(interval int64) (out chan float64, err error) {
	return wv.run(interval, triangle)
}

func (wv *Waveform) SawtoothFlow(interval int64) (out chan float64, err error) {
	return wv.run(interval, sawtooth)
}

// StopFlow ends generation of signal.
func (wv *Waveform) StopFlow() {
	wv.flowRunning = false
	wv.ticker.Stop()
}
