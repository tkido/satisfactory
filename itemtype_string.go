// Code generated by "stringer -type=ItemType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Null-0]
	_ = x[IronOre-1]
	_ = x[IronIngot-2]
	_ = x[IronPlate-3]
	_ = x[CopperOre-4]
	_ = x[CopperIngot-5]
	_ = x[HeavyModularFrame-6]
	_ = x[AluminumCasing-7]
	_ = x[NitrogenGas-8]
	_ = x[NuclearPasta-9]
	_ = x[PressureConversionCube-10]
	_ = x[FusedModularFrame-11]
	_ = x[RadioControlUnit-12]
	_ = x[CopperPowder-13]
	_ = x[ThermalPropulsionRocket-14]
	_ = x[AssemblyDirectorSystem-15]
	_ = x[MagneticFieldGenerator-16]
	_ = x[AluminumIngot-17]
	_ = x[ModularFrame-18]
	_ = x[EncasedIndustrialBeam-19]
	_ = x[Concrete-20]
	_ = x[SteelPipe-21]
	_ = x[SteelIngot-22]
	_ = x[ReinforcedIronPlate-23]
	_ = x[Rubber-24]
	_ = x[Water-25]
	_ = x[Limestone-26]
	_ = x[CrystalOscillator-27]
	_ = x[Computer-28]
	_ = x[QuartzCrystal-29]
	_ = x[Cable-30]
	_ = x[RawQuartz-31]
	_ = x[Wire-32]
	_ = x[CircuitBoard-33]
	_ = x[Quickwire-34]
	_ = x[Plastic-35]
	_ = x[CateriumIngot-36]
	_ = x[CateriumOre-37]
	_ = x[ModularEngine-38]
	_ = x[TurboMotor-39]
	_ = x[CoolingSystem-40]
	_ = x[Motor-41]
	_ = x[HeatSink-42]
	_ = x[SmartPlating-43]
	_ = x[Rotor-44]
	_ = x[Stator-45]
	_ = x[Supercomputer-46]
	_ = x[AdaptiveControlUnit-47]
	_ = x[ElectromagneticControlRod-48]
	_ = x[Battery-49]
	_ = x[AiLimiter-50]
	_ = x[CopperSheet-51]
	_ = x[AutomatedWiring-52]
	_ = x[Sulfur-53]
	_ = x[AlcladAluminumSheet-54]
	_ = x[VersatileFramework-55]
	_ = x[SteelBeam-56]
	_ = x[Coal-57]
	_ = x[Fuel-58]
	_ = x[HeavyOilResidue-59]
	_ = x[CrudeOil-60]
	_ = x[PetroleumCoke-61]
	_ = x[Screw-62]
}

const _ItemType_name = "NullIronOreIronIngotIronPlateCopperOreCopperIngotHeavyModularFrameAluminumCasingNitrogenGasNuclearPastaPressureConversionCubeFusedModularFrameRadioControlUnitCopperPowderThermalPropulsionRocketAssemblyDirectorSystemMagneticFieldGeneratorAluminumIngotModularFrameEncasedIndustrialBeamConcreteSteelPipeSteelIngotReinforcedIronPlateRubberWaterLimestoneCrystalOscillatorComputerQuartzCrystalCableRawQuartzWireCircuitBoardQuickwirePlasticCateriumIngotCateriumOreModularEngineTurboMotorCoolingSystemMotorHeatSinkSmartPlatingRotorStatorSupercomputerAdaptiveControlUnitElectromagneticControlRodBatteryAiLimiterCopperSheetAutomatedWiringSulfurAlcladAluminumSheetVersatileFrameworkSteelBeamCoalFuelHeavyOilResidueCrudeOilPetroleumCokeScrew"

var _ItemType_index = [...]uint16{0, 4, 11, 20, 29, 38, 49, 66, 80, 91, 103, 125, 142, 158, 170, 193, 215, 237, 250, 262, 283, 291, 300, 310, 329, 335, 340, 349, 366, 374, 387, 392, 401, 405, 417, 426, 433, 446, 457, 470, 480, 493, 498, 506, 518, 523, 529, 542, 561, 586, 593, 602, 613, 628, 634, 653, 671, 680, 684, 688, 703, 711, 724, 729}

func (i ItemType) String() string {
	if i < 0 || i >= ItemType(len(_ItemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
