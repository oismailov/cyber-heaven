package rabbitmq

type Message struct {
	Exchange    string
	RoutingKey  string
	ContentType string
	Body        []byte
}
