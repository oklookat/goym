package vantuz

type Logger interface {
	Debugf(msg string, args ...any)
	Err(msg string, err error)
}

type dummyLogger struct {
}

func (d *dummyLogger) Debugf(msg string, args ...any) {

}
func (d *dummyLogger) Err(msg string, err error) {

}
