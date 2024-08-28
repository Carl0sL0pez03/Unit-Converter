package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var lengthConversion = map[string]float64{
	"millimeter": 1,
	"centimeter": 10,
	"meter":      1000,
	"kilometer":  1e6,
	"inch":       25.4,
	"foot":       304.8,
	"yard":       914.4,
	"mile":       1.609e6,
}

var weightConversion = map[string]float64{
	"milligram": 1,
	"gram":      1000,
	"kilogram":  1e6,
	"ounce":     28349.5,
	"pound":     453592,
}

func convertUnits(value float64, fromUnit string, toUnit string, conversionMap map[string]float64) float64 {
	return value * conversionMap[fromUnit] / conversionMap[toUnit]
}

func convertLength(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		valueStr := r.FormValue("value")
		fromUnit := r.FormValue("from_unit")
		toUnit := r.FormValue("to_unit")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err == nil {
			result := convertUnits(value, fromUnit, toUnit, lengthConversion)
			tmpl := template.Must(template.ParseFiles("templates/length.html"))
			tmpl.Execute(w, result)
			return
		}
	}
	tmpl := template.Must(template.ParseFiles("templates/length.html"))
	tmpl.Execute(w, nil)
}

func converWeight(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		valueStr := r.FormValue("value")
		fromUnit := r.FormValue("from_unit")
		toUnit := r.FormValue("to_unit")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err == nil {
			result := convertUnits(value, fromUnit, toUnit, weightConversion)
			tmpl := template.Must(template.ParseFiles("templates/weight.html"))
			tmpl.Execute(w, result)
			return
		}
	}
	tmpl := template.Must(template.ParseFiles("templates/weight.html"))
	tmpl.Execute(w, nil)
}

func converTemperature(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		valueStr := r.FormValue("value")
		fromUnit := r.FormValue("from_unit")
		toUnit := r.FormValue("to_unit")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err == nil {
			var result float64
			if fromUnit == "Celsius" && toUnit == "Fahrenheit" {
				result = value*9/5 + 32
			} else if fromUnit == "Celsius" && toUnit == "Kelvin" {
				result = value + 273.15
			} else if fromUnit == "Fahrenheit" && toUnit == "Celsius" {
				result = (value - 32) * 5 / 9
			} else if fromUnit == "Fahrenheit" && toUnit == "Kelvin" {
				result = (value-32)*5/9 + 273.15
			} else if fromUnit == "Kelvin" && toUnit == "Celsius" {
				result = value - 273.15
			} else if fromUnit == "Kelvin" && toUnit == "Fahrenheit" {
				result = (value-273.15)*9/5 + 32
			} else {
				result = value
			}
			tmpl := template.Must(template.ParseFiles("templates/temperature.html"))
			tmpl.Execute(w, result)
			return
		}
	}
	tmpl := template.Must(template.ParseFiles("templates/temperature.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/convert-length", convertLength)
	http.HandleFunc("/convert-weight", converWeight)
	http.HandleFunc("/convert-temperature", converTemperature)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
