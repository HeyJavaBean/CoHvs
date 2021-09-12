package client

import "testing"

func TestSniff(t *testing.T) {


	newsniff()

}

func TestCall(t *testing.T) {


	call([]byte("Let's Play CoH!"))

}

