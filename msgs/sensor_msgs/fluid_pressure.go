// Autogenerated with msg-import, do not edit.
package sensor_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type FluidPressure struct {
	Header        std_msgs.Header
	FluidPressure msgs.Float64
	Variance      msgs.Float64
}