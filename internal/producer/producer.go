package producer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Adit0507/event-driven-notification/internal/models"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	writer *kafka.Writer
	logger *logrus.Logger
}

func NewProducer(brokers []string, topic string, logger *logrus.Logger) *Producer{
	return &Producer{
		writer: &kafka.Writer{
			Addr: kafka.TCP(brokers...),
			Topic: topic,
			Balancer: &kafka.LeastBytes{}, //so cool
		},

		logger: logger,
	}
}

func (p *Producer) SendNotification(ctx context.Context, notification models.Notification) error {
	msgBytes, err := json.Marshal(notification)
	if err != nil {
		p.logger.Errorf("Failed to marshal notification: %v", err)
		return err
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Key: []byte(notification.UserID),
		Value: msgBytes,
		Time: time.Now(),
	})
	if err != nil {
		p.logger.Errorf("Failed to send notification to Kafka: %v", err)
		return  err
	}

	p.logger.Infof("Notification sent: %s", notification.ID)

	return nil
}

func (p *Producer)Close()error{
	return p.writer.Close()
}