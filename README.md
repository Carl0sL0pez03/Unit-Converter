# Unit-Converter

This is a simple web application built with Go that allows users to convert between different units of measurement, including length, weight, and temperature. The application provides an easy-to-use interface where users can input a value, select the units to convert from and to, and view the converted value.

## Features

- **Length Conversion**: Convert between millimeters, centimeters, meters, kilometers, inches, feet, yards, and miles.
- **Weight Conversion**: Convert between milligrams, grams, kilograms, ounces, and pounds.
- **Temperature Conversion**: Convert between Celsius, Fahrenheit, and Kelvin.
- Simple and responsive web interface.

## Getting Started

### Prerequisites

- Go 1.16+ installed on your system. You can download it from the [official website](https://golang.org/dl/).

### Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/Carl0sL0pez03/Unit-Converter
   cd unit-converter
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

3. Open your web browser and navigate to `http://localhost:8080`.

### Project Structure

- `main.go`: The main Go file that contains the server logic and conversion functions.
- `templates/`: Directory containing HTML templates for the different pages (index, length, weight, temperature).
- `static/`: Directory containing static files like CSS for styling.

### Usage

1. Navigate to the homepage (`http://localhost:8080`).
2. Select the type of conversion you want to perform (Length, Weight, Temperature).
3. Input the value you want to convert, select the units to convert from and to, and click "Convert".
4. The converted value will be displayed on the same page.

### Example Conversions

- **Length**: Convert 10 meters to feet.
- **Weight**: Convert 500 grams to pounds.
- **Temperature**: Convert 32 Fahrenheit to Celsius.

# Challange roadmap

- https://roadmap.sh/projects/unit-converter
