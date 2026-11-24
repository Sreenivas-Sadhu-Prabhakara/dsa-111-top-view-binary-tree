package main

import (
	"reflect"
	"testing"
)

func TestTopView(t *testing.T) {
	// Columns: 1:0, 2:-1, 3:1, 4:-2, 5:0(hidden under 1), 6:0(hidden), 7:2
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, tn(6, nil, nil), tn(7, nil, nil)))
	if got := TopView(root); !reflect.DeepEqual(got, []int{4, 2, 1, 3, 7}) {
		t.Errorf("TopView = %v, want [4 2 1 3 7]", got)
	}

	if got := TopView(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("TopView(nil) = %v", got)
	}
	// Deep nodes never shadow shallow ones in the same column.
	r2 := tn(1, tn(2, nil, tn(3, nil, nil)), nil) // 3 is at column 0, below 1
	if got := TopView(r2); !reflect.DeepEqual(got, []int{2, 1}) {
		t.Errorf("TopView = %v, want [2 1]", got)
	}
}
