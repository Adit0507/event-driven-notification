package delivery

import (

	"github.com/Adit0507/event-driven-notification/internal/models"
)

// interface for notification delivery
type Channel interface {
    Send(notification models.Notification) error
}