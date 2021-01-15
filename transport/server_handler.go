package transport

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/webdevliteral/FractalGeneration/fractal"
)

var numBands = runtime.NumCPU()

func FractalImageHandler(w http.ResponseWriter, req *http.Request) {
	timeStart := time.Now()
	defer func() {
		timeSince := time.Since(timeStart)
		fmt.Printf("took %d ms\n", timeSince.Milliseconds())
	}()

	err := req.ParseForm()
	if err != nil {
		fmt.Printf("could not parse form: %v\n", err)
		return
	}

	width, err := getParamInt(req, "width")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	height, err := getParamInt(req, "height")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fractalType, err := getParam(req, "type")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var f fractal.Fractal

	switch fractalType {
	case "mandelbrot":
		f = &fractal.Mandelbrot{MaxIterations: 100}

	case "julia":
		offsetX, err := getParamFloat(req, "offsetX")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		offsetY, err := getParamFloat(req, "offsetY")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		f = &fractal.Julia{MaxIterations: 100, Z: complex(offsetX, offsetY)}

	default:
		fmt.Printf("unknown fractal type %q\n", fractalType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	img := fractal.GenFractalImage(f, int(width), int(height), numBands)

	buffer := &bytes.Buffer{}
	if err = jpeg.Encode(buffer, img, nil); err != nil {
		fmt.Println(err)
		return
	}

	responseBytes := buffer.Bytes()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(responseBytes)))
	if _, err = w.Write(responseBytes); err != nil {
		fmt.Println(err)
		return
	}
}

func getParam(req *http.Request, param string) (string, error) {
	params, ok := req.Form[param]
	if !ok {
		return "", fmt.Errorf("did not receive parameters %q in query", param)
	}

	if len(params) != 1 {
		return "", fmt.Errorf("received %d %q parameters in query, should receive exactly one", len(params), param)
	}

	return params[0], nil
}

func getParamInt(req *http.Request, param string) (int, error) {
	params, ok := req.Form[param]
	if !ok {
		return 0, fmt.Errorf("did not receive parameters %q in query", param)
	}

	if len(params) != 1 {
		return 0, fmt.Errorf("received %d %q parameters in query, should receive exactly one", len(params), param)
	}

	paramInt, err := strconv.ParseInt(params[0], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("could not parse parameter %q: %v", params[0], err)
	}

	return int(paramInt), nil
}

func getParamFloat(req *http.Request, param string) (float64, error) {
	params, ok := req.Form[param]
	if !ok {
		return 0, fmt.Errorf("did not receive parameters %q in query", param)
	}

	if len(params) != 1 {
		return 0, fmt.Errorf("received %d %q parameters in query, should receive exactly one", len(params), param)
	}

	paramFloat, err := strconv.ParseFloat(params[0], 32)
	if err != nil {
		return 0, fmt.Errorf("could not parse parameter %q: %v", params[0], err)
	}

	return paramFloat, nil
}
