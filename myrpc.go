package main

import "errors"

// Keeps a running total
type Adder struct {
	total int
}

// call("Adder.Add", {"Num": 123})
type AddRequest struct {
	Num int
}
type AddRespose struct{}

func (a *Adder) Add(req *AddRequest, resp *AddRespose) error {
	if req.Num <= 0 {
		return errors.New("Negative")
	}

	a.total += req.Num
	return nil
}

// call("Adder.GetTotal", {})
type GetTotalRequest struct{}
type GetTotalResponse struct {
	Total int
}

func (a *Adder) GetTotal(req *GetTotalRequest, resp *GetTotalResponse) error {
	resp.Total = a.total
	return nil
}

// This is public on the server, but not exposed via rpc.
func (a *Adder) Reset() {
	a.total = 0
}
