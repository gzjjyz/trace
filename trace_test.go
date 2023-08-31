package trace

import (
	"log"
	"testing"
)

func TestGenTraceId(t *testing.T) {
	for i := 0; i < 10; i++ {
		log.Println(GenTraceId())
	}
}
