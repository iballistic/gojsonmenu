package gojsonmenu

import (
	"errors"
)

type IConverter interface {
	Convert(value float64, fromUnit string, toUnit string) (float64, error)
}

type UnitMassConverter struct{}

func (umc UnitMassConverter) Convert(value float64, fromUnit string, toUnit string) (float64, error) {
	var coefficient = map[string]float64{

		"kilograms":  1,
		"grams":      0.001,
		"decigrams":  0.0001,
		"centigrams": 0.00001,
		"milligrams": 0.000001,
		"micrograms": 1.00e-09,
		"nanograms":  1.00e-12,
		"picograms":  1.00e-15,
		"ounces":     0.0283495,
		"pounds":     0.453592,
		"stones":     0.157473,
		"metricTons": 1000,
		"shortTons":  907.185,
		"carats":     0.0002,
		"ouncesTroy": 0.03110348,
		"slugs":      14.5939,
		"grains":     0.000064799, // grains to kilograms
		"grain":      0.000064799, // grain to kilograms (same as grains)
	}

	//1kg is 15432.3584 grain
	//6.4799e-5

	baseValue, ok := coefficient[fromUnit]
	if !ok {
		return 0, errors.New("unknown from unit")
	}

	targetValue, ok := coefficient[toUnit]
	if !ok {
		return 0, errors.New("unknown to unit")
	}

	return value * baseValue / targetValue, nil
}

type UnitSpeedConverter struct{}

func (usc UnitSpeedConverter) Convert(value float64, fromUnit string, toUnit string) (float64, error) {
	coefficient := map[string]float64{
		"footPerSecond":     0.3048, // feet per second to meters per second
		"metersPerSecond":   1,      // base unit
		"kilometersPerHour": 0.277778,
		"milesPerHour":      0.44704,
		"knots":             0.514444,
	}

	baseValue, ok := coefficient[fromUnit]
	if !ok {
		return 0, errors.New("unknown from unit")
	}

	targetValue, ok := coefficient[toUnit]
	if !ok {
		return 0, errors.New("unknown to unit")
	}

	return value * baseValue / targetValue, nil
}

type UnitTemperatureConverter struct{}

func (utc UnitTemperatureConverter) Convert(value float64, fromUnit string, toUnit string) (float64, error) {
	switch fromUnit {
	case "celsius":
		switch toUnit {
		case "fahrenheit":
			return value*9/5 + 32, nil
		case "kelvin":
			return value + 273.15, nil
		}
	case "fahrenheit":
		switch toUnit {
		case "celsius":
			return (value - 32) * 5 / 9, nil
		case "kelvin":
			return (value-32)*5/9 + 273.15, nil
		}
	case "kelvin":
		switch toUnit {
		case "celsius":
			return value - 273.15, nil
		case "fahrenheit":
			return (value-273.15)*9/5 + 32, nil
		}
	}
	return 0, errors.New("unknown unit conversion")
}

type UnitLengthConverter struct{}

func (ulc UnitLengthConverter) Convert(value float64, fromUnit string, toUnit string) (float64, error) {
	coefficient := map[string]float64{

		"megameters":        1000000.0,
		"kilometers":        1000.0,
		"hectometers":       100.0,
		"decameters":        10.0,
		"meters":            1.0, // base unit
		"decimeters":        0.1,
		"centimeters":       0.01,
		"millimeters":       0.001,
		"micrometers":       0.000001,
		"nanometers":        1e-9,
		"picometers":        1e-12,
		"inches":            0.0254,
		"feet":              0.3048,
		"yards":             0.9144,
		"miles":             1609.34,
		"scandinavianMiles": 10000,
		"lightyears":        9.461e+15,
		"nauticalMiles":     1852,
		"fathoms":           1.8288,
		"furlongs":          201.168,
		"astronomicalUnits": 1.496e+11,
		"parsecs":           3.086e+16,
	}

	baseValue, ok := coefficient[fromUnit]
	if !ok {
		return 0, errors.New("unknown from unit")
	}

	targetValue, ok := coefficient[toUnit]
	if !ok {
		return 0, errors.New("unknown to unit")
	}

	return value * baseValue / targetValue, nil
}

type UnitPressureConverter struct{}

func (upc UnitPressureConverter) Convert(value float64, fromUnit string, toUnit string) (float64, error) {
	conversionFactors := map[string]float64{
		"newtonsPerMetersSquared":  1, // base unit (Pascal)
		"gigapascals":              1.00e+09,
		"megapascals":              1000000,
		"kilopascals":              1000,
		"hectopascals":             100,
		"inchesOfMercury":          3.39e+03,
		"bars":                     1.00e+05,
		"millibars":                1.00e+02,
		"millimetersOfMercury":     133.322,
		"poundsForcePerSquareInch": 6894.76,
	}

	baseValue, ok := conversionFactors[fromUnit]
	if !ok {
		return 0, errors.New("unknown from unit")
	}

	targetValue, ok := conversionFactors[toUnit]
	if !ok {
		return 0, errors.New("unknown to unit")
	}

	return value * baseValue / targetValue, nil
}
