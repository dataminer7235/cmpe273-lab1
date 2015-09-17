package main
import "testing"
func Testtime(t *testing.T) {
  var v int = Time_delay()
  if v != 5 {t.Error("Expected 5 for sleep time , got ", v)}
}
