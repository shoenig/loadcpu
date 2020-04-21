package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// take the current max from /proc/cpuinfo
	fmt.Printf("read current speed: %.2f\n", ReadSpeedCPU())

	// load CPU and take measurements
	fmt.Printf("loaded max speed:   %.2f\n", LoadedSpeedCPU())
}

var (
	cpuInfoRe = regexp.MustCompile(`cpu MHz[\s]+:[\s]+([\.\d]+)`)
)

func ReadSpeedCPU() float64 {
	speed, err := mhz("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}
	return speed
}

func LoadedSpeedCPU() float64 {
	speed, err := measure()
	if err != nil {
		panic(err)
	}
	return speed
}

func measure() (float64, error) {
	timerC := time.After(200 * time.Millisecond)
	stopC := make(chan struct{})

	go load(stopC)
	max := 0.0
	for {
		select {
		case <-timerC:
			stopC <- struct{}{}
			return max, nil
		case <-time.After(20 * time.Millisecond):
			speed, err := mhz("/proc/cpuinfo")
			if err != nil {
				return 0, err
			}
			if speed > max {
				max = speed
			}
		}
	}
}

type measurment struct {
	speed float64
	err   error
}

func load(stopC <-chan struct{}) {
	runtime.LockOSThread()
	for {
		select {
		case <-stopC:
			return
		default:
		}

		if busy() {
			return
		}
	}
}

func busy() bool {
	// trick the compiler into doing some work
	// rand.Int is always positive
	return (0x5353&rand.Int())<<1 < 0
}

func mhz(file string) (float64, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}

	matches := cpuInfoRe.FindStringSubmatch(string(b))

	if len(matches) < 2 {
		return 0, errors.New("cpu mhz not found")
	}

	max := 0.0
	for _, match := range matches[1:] {
		mhz, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return 0, err
		}
		if mhz > max {
			max = mhz
		}
	}
	return max, nil
}
