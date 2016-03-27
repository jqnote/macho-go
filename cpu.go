package macho

import (
	"debug/macho"
	"strconv"
)

const _CPU_ARCH_ABI64 = 0x01000000

type cupTypesEnum map[macho.Cpu]string

var CPU_TYPE_NAMES cupTypesEnum = cupTypesEnum{
	1:                    "VAX",
	6:                    "MC680x0",
	7:                    "i386",
	_CPU_ARCH_ABI64 | 7:  "x86_64", //16777223
	8:                    "MIPS",
	10:                   "MC98000",
	11:                   "HPPA",
	12:                   "ARM",
	_CPU_ARCH_ABI64 | 12: "ARM64", //16777228
	13:                   "MC88000",
	14:                   "SPARC",
	15:                   "i860",
	16:                   "Alpha",
	18:                   "PowerPC",
	_CPU_ARCH_ABI64 | 18: "PowerPC64", //16777234
}

type cpuSubTypes map[uint32]string

type cpuSubTypeEnum struct {
	data cpuSubTypes
}

func (c *cpuSubTypeEnum) Get(st uint32, dft uint32) string {
	val, got := c.data[st]
	if got {
		return val
	} else {
		return ""
	}
}

var INTEL64_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		3: "CPU_SUBTYPE_X86_64_ALL",
		4: "CPU_SUBTYPE_X86_ARCH1",
	},
}

var INTEL_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0:   "CPU_SUBTYPE_INTEL_MODEL_ALL",
		1:   "CPU_THREADTYPE_INTEL_HTT",
		3:   "CPU_SUBTYPE_I386_ALL",
		4:   "CPU_SUBTYPE_486",
		5:   "CPU_SUBTYPE_586",
		8:   "CPU_SUBTYPE_PENTIUM_3",
		9:   "CPU_SUBTYPE_PENTIUM_M",
		10:  "CPU_SUBTYPE_PENTIUM_4",
		11:  "CPU_SUBTYPE_ITANIUM",
		12:  "CPU_SUBTYPE_XEON",
		34:  "CPU_SUBTYPE_XEON_MP",
		42:  "CPU_SUBTYPE_PENTIUM_4_M",
		43:  "CPU_SUBTYPE_ITANIUM_2",
		38:  "CPU_SUBTYPE_PENTPRO",
		40:  "CPU_SUBTYPE_PENTIUM_3_M",
		52:  "CPU_SUBTYPE_PENTIUM_3_XEON",
		102: "CPU_SUBTYPE_PENTII_M3",
		132: "CPU_SUBTYPE_486SX",
		166: "CPU_SUBTYPE_PENTII_M5",
		199: "CPU_SUBTYPE_CELERON",
		231: "CPU_SUBTYPE_CELERON_MOBILE",
	},
}

var MC680_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		1: "CPU_SUBTYPE_MC680x0_ALL",
		2: "CPU_SUBTYPE_MC68040",
		3: "CPU_SUBTYPE_MC68030_ONLY",
	},
}

var MIPS_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0: "CPU_SUBTYPE_MIPS_ALL",
		1: "CPU_SUBTYPE_MIPS_R2300",
		2: "CPU_SUBTYPE_MIPS_R2600",
		3: "CPU_SUBTYPE_MIPS_R2800",
		4: "CPU_SUBTYPE_MIPS_R2000a",
		5: "CPU_SUBTYPE_MIPS_R2000",
		6: "CPU_SUBTYPE_MIPS_R3000a",
		7: "CPU_SUBTYPE_MIPS_R3000",
	},
}

var MC98000_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0: "CPU_SUBTYPE_MC98000_ALL",
		1: "CPU_SUBTYPE_MC98601",
	},
}

var HPPA_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0: "CPU_SUBTYPE_HPPA_7100",
		1: "CPU_SUBTYPE_HPPA_7100LC",
	},
}

var MC88_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0: "CPU_SUBTYPE_MC88000_ALL",
		1: "CPU_SUBTYPE_MC88100",
		2: "CPU_SUBTYPE_MC88110",
	},
}

var SPARC_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0: "CPU_SUBTYPE_SPARC_ALL",
	},
}

var I860_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0: "CPU_SUBTYPE_I860_ALL",
		1: "CPU_SUBTYPE_I860_860",
	},
}

var POWERPC_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0:   "CPU_SUBTYPE_POWERPC_ALL",
		1:   "CPU_SUBTYPE_POWERPC_601",
		2:   "CPU_SUBTYPE_POWERPC_602",
		3:   "CPU_SUBTYPE_POWERPC_603",
		4:   "CPU_SUBTYPE_POWERPC_603e",
		5:   "CPU_SUBTYPE_POWERPC_603ev",
		6:   "CPU_SUBTYPE_POWERPC_604",
		7:   "CPU_SUBTYPE_POWERPC_604e",
		8:   "CPU_SUBTYPE_POWERPC_620",
		9:   "CPU_SUBTYPE_POWERPC_750",
		10:  "CPU_SUBTYPE_POWERPC_7400",
		11:  "CPU_SUBTYPE_POWERPC_7450",
		100: "CPU_SUBTYPE_POWERPC_970",
	},
}

var ARM_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0:  "CPU_SUBTYPE_ARM_ALL12",
		5:  "CPU_SUBTYPE_ARM_V4T",
		6:  "CPU_SUBTYPE_ARM_V6",
		7:  "CPU_SUBTYPE_ARM_V5TEJ",
		8:  "CPU_SUBTYPE_ARM_XSCALE",
		9:  "CPU_SUBTYPE_ARM_V7",
		10: "CPU_SUBTYPE_ARM_V7F",
		11: "CPU_SUBTYPE_ARM_V7S",
		12: "CPU_SUBTYPE_ARM_V7K",
		14: "CPU_SUBTYPE_ARM_V6M",
		15: "CPU_SUBTYPE_ARM_V7M",
		16: "CPU_SUBTYPE_ARM_V7EM",
	},
}

var VAX_SUBTYPE cpuSubTypeEnum = cpuSubTypeEnum{
	data: cpuSubTypes{
		0:  "CPU_SUBTYPE_VAX_ALL",
		1:  "CPU_SUBTYPE_VAX780",
		2:  "CPU_SUBTYPE_VAX785",
		3:  "CPU_SUBTYPE_VAX750",
		4:  "CPU_SUBTYPE_VAX730",
		5:  "CPU_SUBTYPE_UVAXI",
		6:  "CPU_SUBTYPE_UVAXII",
		7:  "CPU_SUBTYPE_VAX8200",
		8:  "CPU_SUBTYPE_VAX8500",
		9:  "CPU_SUBTYPE_VAX8600",
		10: "CPU_SUBTYPE_VAX8650",
		11: "CPU_SUBTYPE_VAX8800",
		12: "CPU_SUBTYPE_UVAXIII",
	},
}

func getCpuSubTypeName(cpuType macho.Cpu, cpuSubType uint32) string {
	st := cpuSubType & 0x0fffffff
	var subTypeName string

	if cpuType == 1 {
		subTypeName = VAX_SUBTYPE.Get(st, st)
	} else if cpuType == 6 {
		subTypeName = MC680_SUBTYPE.Get(st, st)
	} else if cpuType == 7 {
		subTypeName = INTEL_SUBTYPE.Get(st, st)
	} else if cpuType == 7|_CPU_ARCH_ABI64 {
		subTypeName = INTEL64_SUBTYPE.Get(st, st)
	} else if cpuType == 8 {
		subTypeName = MIPS_SUBTYPE.Get(st, st)
	} else if cpuType == 10 {
		subTypeName = MC98000_SUBTYPE.Get(st, st)
	} else if cpuType == 11 {
		subTypeName = HPPA_SUBTYPE.Get(st, st)
	} else if cpuType == 12 {
		subTypeName = ARM_SUBTYPE.Get(st, st)
	} else if cpuType == 12|_CPU_ARCH_ABI64 {
		subTypeName = ARM_SUBTYPE.Get(st, st)
	} else if cpuType == 13 {
		subTypeName = MC88_SUBTYPE.Get(st, st)
	} else if cpuType == 14 {
		subTypeName = SPARC_SUBTYPE.Get(st, st)
	} else if cpuType == 15 {
		subTypeName = I860_SUBTYPE.Get(st, st)
	} else if cpuType == 16 {
		subTypeName = MIPS_SUBTYPE.Get(st, st)
	} else if cpuType == 18 {
		subTypeName = POWERPC_SUBTYPE.Get(st, st)
	} else if cpuType == 18|_CPU_ARCH_ABI64 {
		subTypeName = POWERPC_SUBTYPE.Get(st, st)
	} else {
		subTypeName = strconv.FormatUint(uint64(st), 10)
	}
	return subTypeName
}
