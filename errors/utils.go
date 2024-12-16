package errors

import nerrors "errors"

func Is(err, target error) bool {
	return nerrors.Is(err, target)
}
