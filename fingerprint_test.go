package adyen

import (
    "testing"
    "zhaojunlike/logger"
)

func TestFingerprint_String(t *testing.T) {
    fp := NewDF()
    logger.Info("fp: %v", fp.String())
}
