package models

const (
	defaultSampleTime = 100
)

type PID struct {
	// The value to aim for.
	Setpoint float64

	// PID algorithm tunings (post-mangling).
	Kp float64
	Ki float64
	Kd float64

	InputChanel  chan float64
	OutputChanel chan float64
	lastInput    float64

	// PID tunings (pre-mangling)
	dispKp float64
	dispKi float64
	dispKd float64

	// Millis between samples.
	sampleTime int64

	// Minimum and maximum output values.
	outMin float64
	outMax float64

	// Whether a positive output moves the input higher or lower.
	direction int16
}

func NewPID(direction int16) *PID {
	p := new(PID)
	p.SetSampleTime(defaultSampleTime)

	return p
}

func (p *PID) SetSampleTime(sampleTime int64) {
	if sampleTime > 0 {
		ratio := float64(sampleTime) / float64(p.sampleTime)
		p.Ki *= ratio
		p.Kd /= ratio
		p.sampleTime = int64(sampleTime)
	}
}
