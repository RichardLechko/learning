// Package length conversion performs length conversions using the metric system
package lengthconversions

import "fmt"

type Millimeter float64
type Centimeter float64
type Meter      float64
type Kilometer  float64

func (mm Millimeter) String() string {
	return fmt.Sprintf("%gmm", mm)
}

func (cm Centimeter) String() string {
	return fmt.Sprintf("%gcm", cm)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (km Kilometer) String() string {
	return fmt.Sprintf("%gkm", km)
}