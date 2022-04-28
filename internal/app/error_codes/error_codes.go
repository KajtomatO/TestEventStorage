package error_codes

import "errors"

//not working ;-/ not seen from server_lib_test
var ErrRecordNotFound = errors.New("cannot find data")
