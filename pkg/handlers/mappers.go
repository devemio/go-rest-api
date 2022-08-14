package handlers

import "errors"

func MapToErr(err, from, to error) error {
	if errors.Is(err, from) {
		return to
	}

	return err
}
