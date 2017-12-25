# waveforms

Simple waveforms generator.

From [Wikipedia](https://en.wikipedia.org/wiki/Waveform): A waveform is the shape and form of a signal such as a 
wave moving in a physical medium or an abstract representation.

In many cases, the medium in which the wave propagates does not permit a direct observation of the true form.
In these cases, the term "waveform" refers to the shape of a graph of the varying quantity against time.
An instrument called an oscilloscope can be used to pictorially represent a wave as a repeating image on a screen.

[This](https://docs.google.com/spreadsheets/d/1yav__Zk_zyuSIr3frNeUTSwELCtd5-JVMRQob9vdnpU/edit?usp=sharing) google spreadsheet will be our oscillator.

## Usage
There are two variants of usage: simple call and flow. Let's see both.

```go
import (
        "fmt"
        "github.com/thecoldwine/waveforms"
)

// simpleExample is a straightforward call to signal generator. Note waveform's parameters cannot be changed
// after construction.
func simpleExample() {
	wv := NewWaveform(50, 1, 0) // wavelength: 50, amplitude: 1, phase: 0

	for i := 1; i < 100; i++ {
		t := float64(i)
		fmt.Printf("Sine: %.5f %.5f\n", t, wv.Sine(t))
	}
}

// flowExample is a channel-based option. Every interval tick the next value is put on out channel.
func flowExample() {
	wv := NewWaveform(50, 1, 0) // wavelength: 50, amplitude: 1, phase: 0
	defer wv.StopFlow()

	ch, err := wv.SineFlow(25)

	if err != nil {
		panic(err)
	}

	go func() {
		i := 0
		for v := range ch {
			i++
			fmt.Printf("Sine: %d %.5f\n", i, v)
		}}()

	time.Sleep(time.Second * 1)
}

```
