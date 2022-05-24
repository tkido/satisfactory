package main

import "fmt"

func main() {
	fmt.Println("Hello Satisfactory Calculator")
	out := Items{
		NuclearPasta:            2,
		ThermalPropulsionRocket: 2,
		AssemblyDirectorSystem:  8,
		MagneticFieldGenerator:  8,
	}
	fmt.Println(out)
	in := parse(out)
	fmt.Println(in)
}

func parse(out Items) (in Items) {
	in = Items{}
	flag := false
	for k, v := range out {
		if recipe, ok := recipes[k]; ok {
			flag = true
			for k2, v2 := range recipe {
				if inv, ok := in[k2]; ok {
					in[k2] = inv + v*v2
				} else {
					in[k2] = v * v2
				}
			}
		} else {
			if inv, ok := in[k]; ok {
				in[k] = inv + v
			} else {
				in[k] = v
			}
		}
	}
	if flag {
		in = parse(in)
	}
	return in
}
