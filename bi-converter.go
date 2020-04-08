package main

import (
	"fmt"
	"math"
)

//Converter Struct
type Converter struct{}

//Centimeter Feet type
type (
	Minutes    float64
	Seconds    float64
	Centimeter float64
	Feet       float64
	Celsius    float64
	Fahrenheit float64
	Radian     float64
	Degree     float64
	Kilogram   float64
	Pounds     float64
)

func (convMinToSec Converter) MinutestoSecond(min Minutes) Seconds {
	result := Seconds(min * 60)
	return result
}
func (convSecToMin Converter) SecondtoMinutes(sec Seconds) Minutes {
	result := Minutes(sec / 60)
	return result
}

func (convCmToFt Converter) CentimetertoFeet(cm Centimeter) Feet {
	result := Feet(cm / 30.48)
	return result
}
func (convFtToCm Converter) FeettoCentimeter(ft Feet) Centimeter {
	result := Centimeter(ft * 30.48)
	return result
}

func (convCelToFaht Converter) CelsiustoFahrenheit(cel Celsius) Fahrenheit {
	result := Fahrenheit((cel * 9 / 5) + 32)
	return result
}
func (convFahtToCel Converter) FahrenheittoCelsius(faht Fahrenheit) Celsius {
	result := Celsius((faht - 32) * 5 / 9)
	return result
}

func (convRadToDeg Converter) RadiantoDegree(rad Radian) Degree {
	result := Degree(rad * 180 / math.Pi)
	return result
}
func (convDegToRad Converter) DegreetoRadian(deg Degree) Radian {
	result := Radian(deg * math.Pi / 180)
	return result
}

func (convKgToPd Converter) KilogramtoPounds(kg Kilogram) Pounds {
	result := Pounds(kg * 2.2046226218)
	return result
}
func (convPdToKg Converter) PoundstoKilogram(pd Pounds) Kilogram {
	result := Kilogram(pd / 2.2046226218)
	return result
}

func main() {
	//Convert Minutes to Seconds
	var minSec Minutes
	fmt.Println("Please enter minutes to convert to seconds: ")
	fmt.Scan(&minSec)
	convMinToSec := Converter{}
	fmt.Println(convMinToSec.MinutestoSecond(minSec))
	fmt.Println()

	//Convert Seconds to Minutes
	var secMin Seconds
	fmt.Println("Please enter seconds to convert to minutes: ")
	fmt.Scan(&secMin)
	convSecToMin := Converter{}
	fmt.Println(convSecToMin.SecondtoMinutes(secMin))
	fmt.Println()

	//Convert Centimeter to Feet
	var cmFt Centimeter
	fmt.Println("Please enter centimeter to convert to feet: ")
	fmt.Scan(&cmFt)
	convCmToFt := Converter{}
	fmt.Println(convCmToFt.CentimetertoFeet(cmFt))
	fmt.Println()

	//Convert Feet to Centimeter
	var ftCm Feet
	fmt.Println("Please enter feet to convert to centimeter: ")
	fmt.Scan(&ftCm)
	convFtToCm := Converter{}
	fmt.Println(convFtToCm.FeettoCentimeter(ftCm))
	fmt.Println()

	//Convert Celsius to Fahrenheit
	var celFaht Celsius
	fmt.Println("Please enter celsius to convert to fahrenheit: ")
	fmt.Scan(&celFaht)
	convCelToFaht := Converter{}
	fmt.Println(convCelToFaht.CelsiustoFahrenheit(celFaht))
	fmt.Println()

	//Convert Fahrenheit to Celsius
	var fahtCel Fahrenheit
	fmt.Println("Please enter fahrenheit to convert to celsius: ")
	fmt.Scan(&fahtCel)
	convFahtToCel := Converter{}
	fmt.Println(convFahtToCel.FahrenheittoCelsius(fahtCel))
	fmt.Println()

	//Convert Radian to Degree
	var radDeg Radian
	fmt.Println("Please enter radian to convert to degree: ")
	fmt.Scan(&radDeg)
	convRadToDeg := Converter{}
	fmt.Println(convRadToDeg.RadiantoDegree(radDeg))
	fmt.Println()

	//Convert Degree to Radian
	var degRad Degree
	fmt.Println("Please enter degree to convert to radian: ")
	fmt.Scan(&degRad)
	convDegToRad := Converter{}
	fmt.Println(convDegToRad.DegreetoRadian(degRad))
	fmt.Println()

	//Convert Kilogram to Pounds
	var kgPd Kilogram
	fmt.Println("Please enter kilogram to convert to pounds: ")
	fmt.Scan(&kgPd)
	convKgToPd := Converter{}
	fmt.Println(convKgToPd.KilogramtoPounds(kgPd))
	fmt.Println()

	//Convert Pounds to Kilogram
	var pdKg Pounds
	fmt.Println("Please enter pounds to convert to kilogram: ")
	fmt.Scan(&pdKg)
	convPdToKg := Converter{}
	fmt.Println(convPdToKg.PoundstoKilogram(pdKg))
	fmt.Println()
}
