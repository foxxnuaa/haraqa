package server

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestWithQueue(t *testing.T) {
	s := &Server{}
	err := WithQueue(nil)(s)
	if err.Error() != "queue cannot be nil" {
		t.Fatal(err)
	}

	q := &MockQueue{}
	err = WithQueue(q)(s)
	if err != nil {
		t.Fatal(err)
	}
	if s.q != q {
		t.Fatal(s.q, q)
	}
}

func TestWithFileQueue(t *testing.T) {
	s := &Server{}
	err := WithFileQueue(nil, true, 5000)(s)
	if err.Error() != "at least one directory must be given" {
		t.Fatal(err)
	}

	err = WithFileQueue([]string{".haraqa_options"}, true, -1)(s)
	if err == nil || err.Error() != "invalid entries, value must not be negative" {
		t.Error(err)
	}
}

func TestWithMetrics(t *testing.T) {
	s := &Server{}
	err := WithMetrics(nil)(s)
	if err.Error() != "metrics cannot be nil" {
		t.Fatal(err)
	}

	m := noOpMetrics{}
	err = WithMetrics(m)(s)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.metrics, m) {
		t.Fatal(s.metrics, m)
	}
}

func TestWithDefaultConsumeLimit(t *testing.T) {
	s := &Server{}
	err := WithDefaultConsumeLimit(0)(s)
	if err != nil {
		t.Fatal(err)
	}
	if s.defaultConsumeLimit != -1 {
		t.Fatal(s.defaultConsumeLimit)
	}

	err = WithDefaultConsumeLimit(-1)(s)
	if err != nil {
		t.Fatal(err)
	}
	if s.defaultConsumeLimit != -1 {
		t.Fatal(s.defaultConsumeLimit)
	}

	err = WithDefaultConsumeLimit(1024)(s)
	if err != nil {
		t.Fatal(err)
	}
	if s.defaultConsumeLimit != 1024 {
		t.Fatal(s.defaultConsumeLimit)
	}
}

func TestWithMiddleware(t *testing.T) {
	s := &Server{}
	mw := func(next http.Handler) http.Handler {
		return next
	}
	err := WithMiddleware(mw)(s)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.middlewares) != 1 && !reflect.DeepEqual(s.middlewares[0], mw) {
		t.Error(s.middlewares)
	}
}

func TestWithLogger(t *testing.T) {
	s := &Server{}
	logger := &noopLogger{}
	err := WithLogger(*logger)(s)
	if err != nil {
		t.Fatal(err)
	}
	if s.logger != *logger {
		t.Error(s.logger, *logger)
	}
}

func TestWithWebsocketInterval(t *testing.T) {
	s := &Server{}
	d := time.Second * 23
	err := WithWebsocketInterval(d)(s)
	if err != nil {
		t.Fatal(err)
	}
	if s.wsPingInterval != d {
		t.Error(s.wsPingInterval)
	}
}
