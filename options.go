package asciimoby

type Options struct {
	Color bool
	Size  int
	Angle int
	Scale float64
}

func WithColor(color bool) func(*Options) {
	return func(o *Options) {
		o.Color = color
	}
}

func WithSize(size int) func(*Options) {
	return func(o *Options) {
		o.Size = size
	}
}

func WithAngle(angle int) func(*Options) {
	return func(o *Options) {
		o.Angle = angle
	}
}

func WithScale(scale float64) func(*Options) {
	return func(o *Options) {
		o.Scale = scale
	}
}
