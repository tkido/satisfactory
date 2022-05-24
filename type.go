package main

type ItemType int

//go:generate stringer -type=ItemType
const (
	Null ItemType = iota
	IronOre
	IronIngot
	IronPlate
	CopperOre
	CopperIngot
	HeavyModularFrame
	AluminumCasing
	NitrogenGas
	NuclearPasta
	PressureConversionCube
	FusedModularFrame
	RadioControlUnit
	CopperPowder
	ThermalPropulsionRocket
	AssemblyDirectorSystem
	MagneticFieldGenerator
	AluminumIngot
	ModularFrame
	EncasedIndustrialBeam
	Concrete
	SteelPipe
	SteelIngot
	ReinforcedIronPlate
	Rubber
	Water
	Limestone
	CrystalOscillator
	Computer
	QuartzCrystal
	Cable
	RawQuartz
	Wire
	CircuitBoard
	Quickwire
	Plastic
	CateriumIngot
	CateriumOre
	ModularEngine
	TurboMotor
	CoolingSystem
	Motor
	HeatSink
	SmartPlating
	Rotor
	Stator
	Supercomputer
	AdaptiveControlUnit
	ElectromagneticControlRod
	Battery
	AiLimiter
	CopperSheet
	AutomatedWiring
	Sulfur
	AlcladAluminumSheet
	VersatileFramework
	SteelBeam
	Coal
	Fuel
	HeavyOilResidue
	CrudeOil
	PetroleumCoke
	Screw
)

type Items map[ItemType]float64
