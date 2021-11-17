package message_queue

import (
	"encoding/json"

	"github.com/huypher/kit/rabbitmq"
)

type Message = rabbitmq.Message
type Header = rabbitmq.Header
type Body = rabbitmq.Body

type Msg struct {
	Header     Header
	Body       Body
	RoutingKey string
	Priority   int
}

func (m *Msg) MessageHeaderInit()        { m.Header = make(Header) }
func (m *Msg) MessageHeader() Header     { return m.Header }
func (m *Msg) MessageBody() Body         { return m.Body }
func (m *Msg) MessageRoutingKey() string { return m.RoutingKey }
func (m *Msg) MessagePriority() int      { return m.Priority }

var _ Message = (*Msg)(nil)

type PriorityProducer interface {
	rabbitmq.Producer
}

type priorityProducer struct {
	producer PriorityProducer
}

func NewPriorityProducer(cfg *Config, rabbitmqProducer *RabbitmqProducer) *priorityProducer {
	p := new(priorityProducer)

	p.producer = rabbitmqProducer.RabbitmqConn.CreateProducer(
		rabbitmq.ExchangeName(cfg.DelayExchangeName),
		rabbitmq.ExchangeKind(rabbitmq.ExchangeType(cfg.DelayExchangeType)),
		rabbitmq.RegisterMarshalFunc(json.Marshal),
	)

	return p
}

func (p *priorityProducer) Publish(message Message) error {
	return p.producer.Publish(message)
}

func (p *priorityProducer) PublishWithRetry(message Message, numOfRetries int64) error {
	return p.producer.PublishWithRetry(message, numOfRetries)
}
