package validation

import (
	"net"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func ValidateCIDR(field validator.FieldLevel) bool {
	isValid := false

	if e, ok := field.Field().Interface().(string); ok {
		_, _, err := net.ParseCIDR(e)
		isValid = err == nil
	}

	return isValid
}

func ValidateIP(field validator.FieldLevel) bool {
	isValid := false

	if e, ok := field.Field().Interface().(string); ok {
		ip := net.ParseIP(e)
		isValid = len(ip) != 0
	}

	return isValid
}

func RegisterCustomValidation(v *validator.Validate) error {
	err := v.RegisterValidation("is-cidr", ValidateCIDR)
	if err != nil {
		log.WithError(err).Fatal("set is-ip-prefic error")
	}
	err = v.RegisterValidation("is-ip", ValidateIP)
	if err != nil {
		log.WithError(err).Fatal("set is-ip error")
	}
	return err
}
