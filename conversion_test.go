package gojsonmenu

import (
	"testing"
)

func TestUnitMassConverter(t *testing.T) {
	value := 50.00
	fromUnit := "grain"
	toUnit := "grams"
	if conv, ok := provider[fromUnit]; ok {
		massResult, err := conv.Convert(value, fromUnit, toUnit)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		expectedResult := 3.23995 // Adjust the expected result based on your conversion logic
		if Round(massResult, 4) != Round(expectedResult, 4) {
			t.Errorf("Mass Conversion: expected %.4f grams, got %.6f grams", expectedResult, massResult)
		} else {
			t.Logf("Mass Conversion: from %.4f %s to %s = %.6f %s", value, fromUnit, toUnit, massResult, toUnit)
		}
	} else {
		t.Fatalf("No converter found for unit: %s", fromUnit)
	}
}

func TestUnitSpeedConverter(t *testing.T) {
	value := 100.0
	fromUnit := "footPerSecond"
	toUnit := "metersPerSecond"
	if conv, ok := provider[fromUnit]; ok {
		speedResult, err := conv.Convert(value, fromUnit, toUnit)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		expectedResult := 30.48 // Adjust the expected result based on your conversion logic
		if speedResult != expectedResult {
			t.Errorf("Speed Conversion: expected %.6f m/s, got %.6f m/s", expectedResult, speedResult)
		} else {
			t.Logf("Speed Conversion: %.4f %s to %s = %.6f %s", value, fromUnit, toUnit, speedResult, toUnit)
		}
	} else {
		t.Fatalf("No converter found for unit: %s", fromUnit)
	}
}

func TestUnitTemperatureConverter(t *testing.T) {
	value := 100.0
	fromUnit := "celsius"
	toUnit := "fahrenheit"
	if conv, ok := provider[fromUnit]; ok {
		tempResult, err := conv.Convert(value, fromUnit, toUnit)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		expectedResult := 212.0 // Adjust the expected result based on your conversion logic
		if tempResult != expectedResult {
			t.Errorf("Temperature Conversion: expected %.2f F, got %.2f F", expectedResult, tempResult)
		} else {
			t.Logf("Temperature Conversion: 100 %s to %s = %.2f %s", fromUnit, toUnit, tempResult, toUnit)
		}
	} else {
		t.Fatalf("No converter found for unit: %s", fromUnit)
	}
}

func TestUnitLengthConverter(t *testing.T) {
	value := 100.0
	fromUnit := "kilometers"
	toUnit := "miles"
	if conv, ok := provider[fromUnit]; ok {
		lengthResult, err := conv.Convert(value, fromUnit, toUnit)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		expectedResult := 62.1371 // Adjust the expected result based on your conversion logic
		if Round(lengthResult, 3) != Round(expectedResult, 3) {
			t.Errorf("Length Conversion: expected %.3f miles, got %.3f miles", expectedResult, lengthResult)
		} else {
			t.Logf("Length Conversion: %f %s to %s = %.3f %s", value, fromUnit, toUnit, lengthResult, toUnit)
		}
	} else {
		t.Fatalf("No converter found for unit: %s", fromUnit)
	}
}

func TestUnitPressureConverter(t *testing.T) {
	value := 100.0
	fromUnit := "bars"
	toUnit := "poundsForcePerSquareInch"
	if conv, ok := provider[fromUnit]; ok {
		pressureResult, err := conv.Convert(value, fromUnit, toUnit)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		expectedResult := 1450.38 // Adjust the expected result based on your conversion logic
		if Round(pressureResult, 2) != Round(expectedResult, 2) {
			t.Errorf("Pressure Conversion: expected %.2f psi, got %.2f psi", expectedResult, pressureResult)
		} else {
			t.Logf("Pressure Conversion: 100 %s to %s = %.2f %s", fromUnit, toUnit, pressureResult, toUnit)
		}
	} else {
		t.Fatalf("No converter found for unit: %s", fromUnit)
	}
}
