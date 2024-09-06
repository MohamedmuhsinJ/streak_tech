package main

import "net/url"

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

func (r *Request) validate() url.Values {
	errs := url.Values{}
	//check for non empty numbers
	if len(r.Numbers) == 0 {
		errs.Add("numbers", "numbers should not be empty")
	}

	return errs
}
