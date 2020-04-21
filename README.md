loadcpu
=======

Load a single CPU thread and measure its speed.

#### Why

The `gopsutil` project uses `/proc/cpuinfo` as a fallback for
determining a CPU's capable speed. However, on Amazon EC2
instances only the current speed is made available. There is
really no way to get the real "max" CPU speed (wihout `sudo`) 
other than to load the CPU and take measurements.

This tool is for exploring the idea of what it would take
to get reliable measurements of a loaded CPU.

#### Example

```bash
$ go build
$ for i in {1..10}; do ./loadcpu && sleep 3 && echo ""; done
read current speed: 800.04
loaded max speed:   3900.70

read current speed: 1924.65
loaded max speed:   3901.08

read current speed: 1495.16
loaded max speed:   3900.33

read current speed: 2826.81
loaded max speed:   3900.00

read current speed: 3400.18
loaded max speed:   3902.43

read current speed: 1979.91
loaded max speed:   3900.95

read current speed: 2627.13
loaded max speed:   3900.19

read current speed: 889.96
loaded max speed:   3901.62

read current speed: 3391.65
loaded max speed:   3902.97

read current speed: 906.17
loaded max speed:   3900.63
```
