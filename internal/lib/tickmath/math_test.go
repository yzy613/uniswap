package tickmath

import (
	"runtime"
	"testing"
)

func TestGetRatioAtTickAndInverse(t *testing.T) {
	cpuN := runtime.NumCPU()
	if cpuN&1 == 1 {
		cpuN--
	}

	tasks := make(chan int32, cpuN)
	for i := 0; i < cpuN; i++ {
		go func() {
			for tick := range tasks {
				d, _ := GetRatioAtTick(tick)
				ti, _ := GetTickAtRatio(d)

				if ti != tick {
					t.Errorf("want %d got %d", tick, ti)
				}
			}
		}()
	}

	testLower := MinTick >> 2
	testUpper := MaxTick >> 2

	for tick := int32(testLower); tick < int32(testUpper); tick++ {
		tasks <- tick
	}

	close(tasks)
}
