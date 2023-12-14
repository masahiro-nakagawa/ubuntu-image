package imagedefinition

import "errors"

var (
	ErrKeepEnabledNil = errors.New("KeepEnabled is nil. Thi value cannot be properly used.")
	ErrExpireNil      = errors.New("Expire is nil. This value cannot be properly used.")
)
