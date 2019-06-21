package yamltmpl

import (
	"testing"
	"reflect"
)

func TestApplyOverrides(t *testing.T) {
	ref := map[string]interface{}{
		"life": 42,
		"leet": 1337,
		"apple": "banana",
		"true": false,
	}
	arg := map[string]interface{}{
		"life": 42,
		"apple": "orange",
		"true": false,
	}
	override1 := map[string]interface{} {
		"leet": 1337,
		"true": true,
	}
	override2 := map[string]interface{}{
		"apple": "banana",
		"true": false,
	}
	applyOverrides(arg, override1, override2)
	if !reflect.DeepEqual(arg, ref) {
		t.Errorf("expected: %v, got: %v", ref, arg)
	}
}
