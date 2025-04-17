package model

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type Status string

const (
	CREATED   Status = "CREATED"
	SHIPPED   Status = "SHIPPED"
	DELIVERED Status = "DELIVERED"
)

var STATUSES = []Status{CREATED, SHIPPED, DELIVERED}

type UpdateStatusRequest struct {
	Status Status `json:"status" valid:"required"`
}

func (sr *UpdateStatusRequest) Validate() error {
	valid, err := govalidator.ValidateStruct(sr)
	if !valid {
		return err
	}

	if !sr.Status.isAllowedStatus() {
		return fmt.Errorf("status %s is not allowed", sr.Status)
	}

	return nil
}

func (s Status) isAllowedStatus() bool {
	for _, status := range STATUSES {
		if strings.EqualFold(string(s), string(status)) {
			return true
		}
	}
	return false
}
