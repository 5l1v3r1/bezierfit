package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/unixpickle/bezierfit"
)

func main() {
	fmt.Println("Enter lines of the form 'x y'; EOF to finish:")
	var points []bezierfit.Point
	for {
		x, y, err := readPoint()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read input:", err)
			os.Exit(1)
		}
		points = append(points, bezierfit.Point{x, y})
	}
	fmt.Println("Solving...")
	fit := bezierfit.BestFit(points)
	fmt.Printf("cubic-bezier(%.3f,%.3f,%.3f,%.3f)\n", fit.P1.X, fit.P1.Y, fit.P2.X, fit.P2.Y)
}

func readPoint() (x, y float64, err error) {
	line, err := readLine()
	if err != nil {
		return 0, 0, err
	}
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid input: " + line)
	}
	x, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, err
	}
	y, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, err
	}
	return
}

func readLine() (line string, err error) {
	str := ""
	for {
		buf := make([]byte, 1)
		if _, err := os.Stdin.Read(buf); err != nil {
			return "", err
		}
		if buf[0] == '\n' {
			break
		} else if buf[0] != '\r' {
			str += string(buf[0])
		}
	}
	return str, nil
}
