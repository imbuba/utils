package alerts

// Alerter is an interface for alerts
type Alerter interface {
	PostMessage(message string, infoLevel string) error
}
