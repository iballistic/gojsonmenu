package gojsonmenu

var provider = map[string]IConverter{
	// Mass
	"kilograms":  UnitMassConverter{},
	"grams":      UnitMassConverter{},
	"decigrams":  UnitMassConverter{},
	"milligrams": UnitMassConverter{},
	"micrograms": UnitMassConverter{},
	"nanograms":  UnitMassConverter{},
	"picograms":  UnitMassConverter{},
	"ounces":     UnitMassConverter{},
	"pounds":     UnitMassConverter{},
	"stones":     UnitMassConverter{},
	"metricTons": UnitMassConverter{},
	"shortTons":  UnitMassConverter{},
	"carats":     UnitMassConverter{},
	"ouncesTroy": UnitMassConverter{},
	"slugs":      UnitMassConverter{},
	"grain":      UnitMassConverter{},

	// Speed
	"metersPerSecond":   UnitSpeedConverter{},
	"footPerSecond":     UnitSpeedConverter{},
	"kilometersPerHour": UnitSpeedConverter{},
	"milesPerHour":      UnitSpeedConverter{},
	"knots":             UnitSpeedConverter{},

	// Temperature
	"kelvin":     UnitTemperatureConverter{},
	"celsius":    UnitTemperatureConverter{},
	"fahrenheit": UnitTemperatureConverter{},

	// Length
	"megameters":  UnitLengthConverter{},
	"kilometers":  UnitLengthConverter{},
	"meters":      UnitLengthConverter{},
	"centimeters": UnitLengthConverter{},
	"millimeters": UnitLengthConverter{},
	"micrometers": UnitLengthConverter{},
	"nanometers":  UnitLengthConverter{},
	"picometers":  UnitLengthConverter{},
	"inches":      UnitLengthConverter{},
	"feet":        UnitLengthConverter{},
	"yards":       UnitLengthConverter{},
	"miles":       UnitLengthConverter{},

	// Pressure
	"inchesOfMercury":          UnitPressureConverter{},
	"bars":                     UnitPressureConverter{},
	"millibars":                UnitPressureConverter{},
	"millimetersOfMercury":     UnitPressureConverter{},
	"poundsForcePerSquareInch": UnitPressureConverter{},
	"kilopascals":              UnitPressureConverter{},
	"hectopascals":             UnitPressureConverter{},
}
