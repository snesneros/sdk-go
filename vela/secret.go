// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// SecretService handles retriving secrets from
// the server methods of the Vela API.
type SecretService service

// Get returns the provided secret.
func (s *SecretService) Get(engine, sType, org, name string, target string) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, target)

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetAll returns a list of all secrets.
func (s *SecretService) GetAll(engine, sType, org, name string) (*[]library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s", engine, sType, org, name)

	// slice library Secret type we want to return
	v := new([]library.Secret)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Add constructs a secret with the provided details.
func (s *SecretService) Add(engine, sType, org, name string, target *library.Secret) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s", engine, sType, org, name)

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := s.client.Call("POST", u, target, v)
	return v, resp, err
}

// Update modifies a secret with the provided details.
func (s *SecretService) Update(engine, sType, org, name string, target *library.Secret) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, *target.Name)

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Remove deletes the provided secret.
func (s *SecretService) Remove(engine, sType, org, name string, target string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, target)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
