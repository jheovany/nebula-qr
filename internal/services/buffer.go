package services

import "bytes"

type writeCloserBuffer struct {
	*bytes.Buffer
}

func (wcb *writeCloserBuffer) Close() error {
	return nil
}