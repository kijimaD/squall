// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/4xUXYscRRT9K83Vx3a6Vwkb+skFFx1IZNiJedF5KHtqe2rTXdWpqhl2CA12F0iiUYJf",
	"hKgsG5SJWTIKQYga8MfcTMyjP0Gqej4yHxn3Zbem+ta559x77r0BschywSnXCqIbIKnKBVfU/bgQhk2u",
	"qeQkbVM5oHJfSiHtl1hwTbm2R5LnKYuJZoIHR0pwe0ePSZantI7sUoguhKEPGVWKJBQimKF6NaxX4xY+",
	"qLhHM2LfvS7pIUTwWrDgF9RfVXC5BjqYcoWiKHzoUhVLllseEFnq3sYs3rMnXzy/+yWWD7F6gOYXNE+h",
	"mGd2svf5wP5bBpyMT16c3kZzH81TNLfAh4wcX6I80T2Idqw8PcytNqUl44kV8y7VB0JoS9Pi5VLkVGpW",
	"15bWSbbJtDwsM010X/1fcLuOKnwYUKlY3YdtD65Ow6x4Sa/3maRdiD6cpfMdwwVaZy5QfHxEY20zrfZh",
	"TWTd/NVSvnflSsur03guYg7NuKYJdU6Ym2Xb61nQWvFXNE2zzMI3aWnPq7ycDqvfXcfvYPUHmk/tufrb",
	"nj+pJt//ORnfO48Tri5asoy+12paaAv6BM0IzeMNWgofGD8Us7kjsZs7mhGWQgRcSJaRa4wnu7u7byf2",
	"thGLDHzgJLM419gRy0gX1makfb1P0hTLsWNxC6tHaH5Gc4rmMZqbH3FLhWk7xtNQb6/V9N4RsXrJFxEM",
	"wkbY2LHwIqec5AwieKux0wjBh5zonitqYP8kVG8osLmL5sxVt8RqhMZg9Ss4MOm2SrML0WyWwF/eUG+G",
	"4TnW0fm2ysvjumGjPL95Z/LZSV3FQ9JP9avw5gSDV+xPt276WUbkcIt+TRJlvXuZqRg6jpByIPZ2vYiP",
	"nInO0Jz989Vvk/tmr9UEH/rSWqSndR4FQSpikvaE0tHF8GIIRWeeZBVuB8ufsBx/cHAJy9GL0wdYfovV",
	"baw+R3OvNuvCXvtcyyH4cPxGl6k8JcP36/t/T77+wa1YayY0D6Hw11iXP2I5fvbXdws0p3YT2DejRXTR",
	"Kf4LAAD//4NCfGC4BgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}