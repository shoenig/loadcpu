loadcpu
=======

Load a single CPU thread and measure its speed.

#### Why

The `gopsutil` project uses `/proc/cpuinfo` as a fallback for
determining a CPU's capable speed. However, on Amazon EC2
instances only the current speed is made available. There is
really no way to get the real "max" CPU speed other than to
load the CPU and take measurements.

This tool is for exploring the idea of what it would take
to get reliable measurements of a loaded CPU.

#### Example

```bash
$ go build
$ for i in {1..10}; do ./loadcpu && sleep 3 && echo ""; done
read current speed: 796.962000.2
loaded max speed:   3902.771000.2

read current speed: 1193.820000.2
loaded max speed:   3903.541000.2

read current speed: 800.177000.2
loaded max speed:   3916.474000.2

read current speed: 2251.564000.2
loaded max speed:   3900.003000.2

read current speed: 800.826000.2
loaded max speed:   3907.950000.2

read current speed: 709.335000.2
loaded max speed:   3904.464000.2

read current speed: 955.625000.2
loaded max speed:   3903.182000.2

read current speed: 799.782000.2
loaded max speed:   3900.002000.2

read current speed: 916.650000.2
loaded max speed:   3912.479000.2

read current speed: 748.158000.2
loaded max speed:   3902.867000.2
```
