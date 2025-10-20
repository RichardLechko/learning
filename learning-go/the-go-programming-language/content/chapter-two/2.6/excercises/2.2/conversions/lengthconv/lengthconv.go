package lengthconversions

// Millimeters to Centimeters
func MMToCM(mm Millimeter) Centimeter {
	return Centimeter(mm / 10)
}

// Millimeters to Meters
func MMToM(mm Millimeter) Meter {
	return Meter( (MMToCM(mm)) / 100 )
}

// Millimeters to Kilometers
func MMToKM(mm Millimeter) Kilometer {
	return Kilometer( (MMToM(mm)) / 1000 )
}

// Centimeters to Millimeters
func CMToMM(cm Centimeter) Millimeter {
	return Millimeter(cm * 10)
}

// Centimeters to Meters
func CMToM(cm Centimeter) Meter {
	return Meter(cm / 100)
}

// Centimeters to Kilometers
func CMToKM(cm Centimeter) Kilometer {
	return Kilometer( (CMToM(cm)) / 1000 )
}

// Meters to Millimeters
func MToMM(m Meter) Millimeter {
	return Millimeter( MToCM(m) * 10)
}

// Meters to Centimeters
func MToCM(m Meter) Centimeter {
	return Centimeter(m * 100)
}

// Meters to Kilometers
func MToKM(m Meter) Kilometer {
	return Kilometer(m / 1000)
}

// Kilometers to Millimeters
func KMToMM(km Kilometer) Millimeter {
	return Millimeter( (KMToCM(km)) * 10 )
}

// Kilometers to Centimeters
func KMToCM(km Kilometer) Centimeter {
	return Centimeter( (KMToM(km) * 100) )
}

// Kilometers to Meters
func KMToM(km Kilometer) Meter {
	return Meter(km * 1000)
}