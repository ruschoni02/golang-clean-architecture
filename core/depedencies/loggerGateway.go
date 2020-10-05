package depedencies

type Event map[string]interface{}

type LoggerGateway interface {
	Debug(msg string, event ...Event)
	Info(msg string, event ...Event)
	Warn(msg string, event ...Event)
	Error(msg string, err error, event ...Event)
	Fatal(msg string, err error, event ...Event)
}
