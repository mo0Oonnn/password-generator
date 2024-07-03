package main

import (
	"testing"
)

func TestSetupLogger(t *testing.T) {
	logger := setupLogger(envLocal)

	if logger == nil {
		t.Errorf("logger is nil")
	}

	logger = setupLogger(envDev)
	if logger == nil {
		t.Errorf("logger is nil")
	}

	logger = setupLogger(envProd)
	if logger == nil {
		t.Errorf("logger is nil")
	}
}
