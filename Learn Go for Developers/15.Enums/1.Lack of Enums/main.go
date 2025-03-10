package main

import (
	"fmt"
)

func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return fmt.Errorf("error updating user status: %w", err)
	}
	err1 := a.track(em.status)

	if err1 != nil {
		return fmt.Errorf("error tracking user bounce: %w", err)
	}
	return nil
}
