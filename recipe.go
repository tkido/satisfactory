package main

var recipes = map[ItemType]Items{
	// IronIngot: {
	// 	IronOre: 1,
	// },
	IronPlate: {
		IronIngot: 3.0 / 2.0,
	},
	// CopperIngot: {
	// 	CopperOre: 1,
	// },
	NuclearPasta: {
		PressureConversionCube: 1,
		CopperPowder:           200,
	},
	CopperPowder: {
		CopperIngot: 6,
	},
	PressureConversionCube: {
		FusedModularFrame: 1,
		RadioControlUnit:  2,
	},
	FusedModularFrame: {
		HeavyModularFrame: 1,
		NitrogenGas:       25,
		AluminumCasing:    50,
	},
	AluminumCasing: {
		AluminumIngot: 3.0 / 2.0,
	},
	HeavyModularFrame: {
		ModularFrame:          8.0 / 3.0,
		EncasedIndustrialBeam: 10.0 / 3.0,
		Concrete:              22.0 / 3.0,
		SteelPipe:             12.0,
	},
	EncasedIndustrialBeam: {
		Concrete:  5,
		SteelPipe: 7,
	},
	SteelPipe: {
		SteelIngot: 3.0 / 2.0,
	},
	ModularFrame: {
		ReinforcedIronPlate: 2.0 / 3.0,
		SteelPipe:           10.0 / 3.0,
	},
	ReinforcedIronPlate: {
		Screw:     50.0 / 3.0,
		IronPlate: 6,
	},
	// ReinforcedIronPlate: {
	// 	Rubber:    1,
	// 	IronPlate: 3,
	// },
	Screw: {
		SteelBeam: 1.0 / 52.0,
	},
	Concrete: {
		Water:     5.0 / 4.0,
		Limestone: 6.0 / 4.0,
	},
	RadioControlUnit: {
		Computer:          1.0 / 2.0,
		CrystalOscillator: 1.0 / 2.0,
		AluminumCasing:    16,
	},
	CrystalOscillator: {
		QuartzCrystal:       18,
		Cable:               14,
		ReinforcedIronPlate: 5.0 / 2.0,
	},
	QuartzCrystal: {
		RawQuartz: 5.0 / 3.0,
	},
	Cable: {
		Wire: 2,
	},
	Wire: {
		IronIngot: 5.0 / 9.0,
	},
	Computer: {
		CircuitBoard: 7,
		Quickwire:    28,
		Rubber:       12,
	},
	CircuitBoard: {
		Plastic:   10.0 / 7.0,
		Quickwire: 30.0 / 7.0,
	},
	// CircuitBoard: {
	// 	Rubber:        6,
	// 	PetroleumCoke: 9,
	// },
	PetroleumCoke: {
		HeavyOilResidue: 4.0 / 12.0,
	},
	Quickwire: {
		CateriumIngot: 1.0 / 5.0,
	},
	CateriumIngot: {
		CateriumOre: 3,
	},
	ThermalPropulsionRocket: {
		TurboMotor:        1,
		FusedModularFrame: 1,
		ModularEngine:     5.0 / 2.0,
		CoolingSystem:     3,
	},
	TurboMotor: {
		RadioControlUnit: 2,
		CoolingSystem:    4,
		Motor:            4,
		Rubber:           24,
	},
	CoolingSystem: {
		HeatSink:    2,
		Rubber:      2,
		Water:       5,
		NitrogenGas: 25,
	},
	HeatSink: {
		AluminumCasing: 3,
		Rubber:         3,
	},
	ModularEngine: {
		Motor:        2,
		SmartPlating: 2,
		Rubber:       15,
	},
	SmartPlating: {
		ReinforcedIronPlate: 1,
		Rotor:               1,
	},
	Motor: {
		Rotor:  2,
		Stator: 2,
	},
	Rotor: {
		SteelPipe: 2,
		Wire:      6,
	},
	Stator: {
		SteelPipe: 3,
		Wire:      8,
	},
	AssemblyDirectorSystem: {
		AdaptiveControlUnit: 2,
		Supercomputer:       1,
	},
	Supercomputer: {
		ElectromagneticControlRod: 1,
		Computer:                  3.0 / 2.0,
		Battery:                   10,
		Wire:                      45.0 / 2.0,
	},
	ElectromagneticControlRod: {
		AiLimiter: 1,
		Stator:    3.0 / 2.0,
	},
	AiLimiter: {
		CopperSheet: 5,
		Quickwire:   20,
	},
	CopperSheet: {
		CopperIngot: 2,
	},
	AdaptiveControlUnit: {
		HeavyModularFrame: 1,
		Computer:          1,
		CircuitBoard:      2,
		AutomatedWiring:   15.0 / 2.0,
	},
	AutomatedWiring: {
		Stator: 1,
		Cable:  20,
	},
	Battery: {
		Sulfur:              6.0 / 4.0,
		AlcladAluminumSheet: 7.0 / 4.0,
		Plastic:             2,
		Wire:                3,
	},
	AlcladAluminumSheet: {
		AluminumIngot: 1,
		CopperIngot:   1.0 / 3.0,
	},
	MagneticFieldGenerator: {
		VersatileFramework:        5.0 / 2.0,
		ElectromagneticControlRod: 1,
		Battery:                   5,
	},
	VersatileFramework: {
		ModularFrame: 1.0 / 2.0,
		SteelBeam:    3,
		Rubber:       4,
	},
	SteelBeam: {
		SteelIngot: 4,
	},
	SteelIngot: {
		IronIngot: 2.0 / 3.0,
		Coal:      2.0 / 3.0,
	},
	Rubber: {
		Fuel: 1,
	},
	Plastic: {
		Fuel: 1,
	},
	Fuel: {
		HeavyOilResidue: 5.0 / 10.0,
		Water:           1,
	},
	HeavyOilResidue: {
		CrudeOil: 3.0 / 4.0,
	},
}
