package client

import "testing"

func TestNew(t *testing.T) {

	c := New("x", "y")

	c.GetAccessToken()
}
