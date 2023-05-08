package web

import (
	"errors"
	"net/http"
	"regexp"

	util "github.com/mrpsousa/api/pkg"
)

func DecodeSinceDateFromURL(r *http.Request) (*string, error) {
	if r == nil {
		err := errors.New("decoding_error : since_date")
		return nil, err
	}
	since := r.URL.Query().Get("since")
	if !util.IsNumeric(since) {
		err := errors.New("type_error : must_be_numeric")
		return nil, err
	}
	return &since, nil
}

func DecodeUserNameToReposFromURL(r *http.Request) (*string, error) {
	if r == nil {
		err := errors.New("decoding_error : user_name")
		return nil, err
	}
	regx := regexp.MustCompile(`users\/(.*)\/repos`)
	url := r.URL.Path
	userName := regx.FindStringSubmatch(url)[1]

	if userName == "" {
		err := errors.New("request_error : request_can't_be_empty")
		return nil, err
	}

	return &userName, nil
}

func DecodeUserNameToDetailFromURL(r *http.Request) (*string, error) {
	if r == nil {
		err := errors.New("decoding_error : user_name")
		return nil, err
	}
	regx := regexp.MustCompile(`users\/(.*)\/details`)
	url := r.URL.Path
	userName := regx.FindStringSubmatch(url)[1]

	if userName == "" {
		err := errors.New("request_error : request_can't_be_empty")
		return nil, err
	}

	return &userName, nil
}
