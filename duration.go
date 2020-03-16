package goconfig

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

type dr string

func (d dr) String() string {
	return string(d)
}

func (d dr) Duration() (time.Duration, error) {
	reg := regexp.MustCompile(`(\d+)(\w+)`)
	if reg == nil {
		return time.Second, errors.New("not a valid duration")
	}
	spe := reg.FindAllStringSubmatch(d.String(), -1)
	// fmt.Println(len(spe[0]))
	if len(spe) == 0 {
		return time.Second, errors.New("not a valid duration")
	}
	if len(spe[0]) < 3 {
		return time.Second, nil
	}
	t, err := strconv.ParseInt(spe[0][1], 10, 64)
	if err != nil {
		return time.Second, err
	}
	dt := time.Duration(t)
	switch spe[0][2] {
	case "ms", "MS":
		return dt * time.Millisecond, nil
	case "s", "S":
		return dt * time.Second, nil
	case "m", "M":
		return dt * time.Minute, nil
	case "h", "H":
		return dt * time.Hour, nil
	case "w", "W":
		return dt * time.Hour * 7, nil
	case "d", "D":
		return dt * time.Hour * 24, nil
	default:
		return time.Second, errors.New("invalid duration")
	}
}
