package main

import "testing"

func TestInputs(t *testing.T) {
	values := map[string]string{
		"{":                  "{}",
		"}":                  "{}",
		"}}":                 "{{}}",
		"}}{{}":              "{{}}{{}}",
		"}}{}":               "{{}}{}",
		"}}{{}}}}":           "{{{{}}{{}}}}",
		"}}}}{{}{{{{}{}}}":   "{{{{}}}}{{}{{{{}{}}}}}",
		"}}}}{{}{{{{}{}}}}}": "{{{{}}}}{{}{{{{}{}}}}}",
		"}}}}{}}}}}":         "{{{{{{{{}}}}{}}}}}",
		"}}}}{{}}}}}":        "{{{{{{{}}}}{{}}}}}",
	}

	for k, v := range values {
		if balance(k) != v {
			t.Errorf("Expected %v to balance to %v, but got: %v", k, v, balance(k))
		}
	}
}
