package task_errors

import "errors"

var FoundNothingErr = errors.New("found nothing")
var EpmtyStringErr = errors.New("empty inserted string")
var WrongStatusErr = errors.New("wrong status")
