package controllers

import (
	"util/errs"
)

///////////////////////////////////////////////////////////////////////////////
//
type httpResult struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

//
type httpError struct {
	Err int `json:"err"`
}

//
func resutls(result interface{}, err *errs.Error) *httpResult {
	ret := &httpResult{Code: 200}
	if nil == err {
		ret.Result = result
	} else {
		ret.Result = &httpError{Err: err.Code}
	}

	return ret
}
