package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// location with a latitude, longitude in decimal degress.
type location struct {
	lat, long coordinate
}

// String formats a location with latitude, longitude
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type coordinate struct { 
	d, m, s float64
	h rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.2f\" %c", c.d, c.m, c.s, c.h)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	marshalMap := map[string]string {
		"decimal" : fmt.Sprintf("%.1f", c.decimal()),
		"dms": fmt.Sprint(c),
		"degrees": fmt.Sprintf("%.0f", c.d),
		"minutes": fmt.Sprintf("%.0f", c.m),
		"seconds": fmt.Sprintf("%.0f", c.s),
		"hemisphere": fmt.Sprintf("%c", c.h),
	}

	return json.Marshal(marshalMap)
}

func main() {
	// elysium := location{
	// 	coordinate{4,30,0.0,'N'},
	// 	coordinate{135,54,0.0,'E'},
	// }

	marshaled, err := json.MarshalIndent(coordinate{135,54,0.0,'E'}, "", " ")
	if(err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(marshaled))
}
