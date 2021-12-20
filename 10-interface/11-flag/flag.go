package myflag

import (
	"flag"
	"fmt"
)

type Value interface {
	String() string
	Set(string) error
}

type Sheshi float64
type Huashi float64

func SoH(c Sheshi) Huashi { return Huashi(c*9.0/5.0 + 32.0) }
func HoS(f Huashi) Sheshi { return Sheshi((f - 32.0) * 5.0 / 9.0) }

// 摄氏温度
type sheshiFlag struct {
	Sheshi
}

func (f *sheshiFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Sheshi = Sheshi(value)
		return nil
	case "F", "°F":
		f.Sheshi = HoS(Huashi(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %v", s)
}

func (c Sheshi) String() string {
	return fmt.Sprintf("%gC", c)
}

func SheshiFlag(name string, value Sheshi, usage string) *Sheshi {
	f := sheshiFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Sheshi
}
