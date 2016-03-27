package macho

const (
	MH_OBJECT     = 0x1
	MH_EXECUTE    = 0x2
	MH_FVMLIB     = 0x3
	MH_CORE       = 0x4
	MH_PRELOAD    = 0x5
	MH_DYLIB      = 0x6
	MH_DYLINKER   = 0x7
	MH_BUNDLE     = 0x8
	MH_DYLIB_STUB = 0xa
	MH_DSYM       = 0xb
)
const (
	MAGIC_LEN = 4

	MH_MAGIC    = 0xfeedface
	MH_CIGAM    = 0xcefaedfe
	MH_MAGIC_64 = 0xfeedfacf
	MH_CIGAM_64 = 0xcffaedfe

	LC_REQ_DYLD = 0x80000000

	LC_SEGMENT                  = 0x1
	LC_SYMTAB                   = 0x2
	LC_SYMSEG                   = 0x3
	LC_THREAD                   = 0x4
	LC_UNIXTHREAD               = 0x5
	LC_LOADFVMLIB               = 0x6
	LC_IDFVMLIB                 = 0x7
	LC_IDENT                    = 0x8
	LC_FVMFILE                  = 0x9
	LC_PREPAGE                  = 0xa
	LC_DYSYMTAB                 = 0xb
	LC_LOAD_DYLIB               = 0xc
	LC_ID_DYLIB                 = 0xd
	LC_LOAD_DYLINKER            = 0xe
	LC_ID_DYLINKER              = 0xf
	LC_PREBOUND_DYLIB           = 0x10
	LC_ROUTINES                 = 0x11
	LC_SUB_FRAMEWORK            = 0x12
	LC_SUB_UMBRELLA             = 0x13
	LC_SUB_CLIENT               = 0x14
	LC_SUB_LIBRARY              = 0x15
	LC_TWOLEVEL_HINTS           = 0x16
	LC_PREBIND_CKSUM            = 0x17
	LC_LOAD_WEAK_DYLIB          = 0x80000018
	LC_SEGMENT_64               = 0x19
	LC_ROUTINES_64              = 0x1a
	LC_UUID                     = 0x1b
	LC_RPATH                    = 0x8000001c
	LC_CODE_SIGNATURE           = 0x1d
	LC_CODE_SEGMENT_SPLIT_INFO  = 0x1e
	LC_REEXPORT_DYLIB           = 0x8000001f
	LC_LAZY_LOAD_DYLIB          = 0x20
	LC_ENCRYPTION_INFO          = 0x21
	LC_DYLD_INFO                = 0x22
	LC_DYLD_INFO_ONLY           = 0x80000022
	LC_LOAD_UPWARD_DYLIB        = 0x80000023
	LC_VERSION_MIN_MACOSX       = 0x24
	LC_VERSION_MIN_IPHONEOS     = 0x25
	LC_FUNCTION_STARTS          = 0x26
	LC_DYLD_ENVIRONMENT         = 0x27
	LC_MAIN                     = 0x80000028
	LC_DATA_IN_CODE             = 0x29
	LC_SOURCE_VERSION           = 0x2a
	LC_DYLIB_CODE_SIGN_DRS      = 0x2b
	LC_ENCRYPTION_INFO_64       = 0x2c
	LC_LINKER_OPTION            = 0x2d
	LC_LINKER_OPTIMIZATION_HINT = 0x2e
)

const (
	MH_NOUNDEFS = 1 << iota
	MH_INCRLINK
	MH_DYLDLINK
	MH_BINDATLOAD
	MH_PREBOUND
	MH_SPLIT_SEGS
	MH_LAZY_INIT
	MH_TWOLEVEL
	MH_FORCE_FLAT
	MH_NOMULTIDEFS
	MH_NOFIXPREBINDING
	MH_PREBINDABLE
	MH_ALLMODSBOUND
	MH_SUBSECTIONS_VIA_SYMBOLS
	MH_CANONICAL
	MH_WEAK_DEFINES
	MH_BINDS_TO_WEAK
	MH_ALLOW_STACK_EXECUTION
	MH_ROOT_SAFE
	MH_SETUID_SAFE
	MH_NO_REEXPORTED_DYLIBS
	MH_PIE
	MH_DEAD_STRIPPABLE_DYLIB
	MH_HAS_TLV_DESCRIPTORS
	MH_NO_HEAP_EXECUTION
)

var MH_FILETYPE_SHORTNAMES = map[uint32]string{
	MH_OBJECT:     "object",
	MH_EXECUTE:    "execute",
	MH_FVMLIB:     "fvmlib",
	MH_CORE:       "core",
	MH_PRELOAD:    "preload",
	MH_DYLIB:      "dylib",
	MH_DYLINKER:   "dylinker",
	MH_BUNDLE:     "bundle",
	MH_DYLIB_STUB: "dylib_stub",
	MH_DSYM:       "dsym",
}

var MH_FILETYPE_NAMES = map[uint32]string{
	MH_OBJECT:     "relocatable object",
	MH_EXECUTE:    "demand paged executable",
	MH_FVMLIB:     "fixed vm shared library",
	MH_CORE:       "core",
	MH_PRELOAD:    "preloaded executable",
	MH_DYLIB:      "dynamically bound shared library",
	MH_DYLINKER:   "dynamic link editor",
	MH_BUNDLE:     "dynamically bound bundle",
	MH_DYLIB_STUB: "shared library stub for static linking",
	MH_DSYM:       "symbol information",
}

var MH_FLAGS_NAMES = map[uint32]string{
	MH_NOUNDEFS:                "MH_NOUNDEFS",
	MH_INCRLINK:                "MH_INCRLINK",
	MH_DYLDLINK:                "MH_DYLDLINK",
	MH_BINDATLOAD:              "MH_BINDATLOAD",
	MH_PREBOUND:                "MH_PREBOUND",
	MH_SPLIT_SEGS:              "MH_SPLIT_SEGS",
	MH_LAZY_INIT:               "MH_LAZY_INIT",
	MH_TWOLEVEL:                "MH_TWOLEVEL",
	MH_FORCE_FLAT:              "MH_FORCE_FLAT",
	MH_NOMULTIDEFS:             "MH_NOMULTIDEFS",
	MH_NOFIXPREBINDING:         "MH_NOFIXPREBINDING",
	MH_PREBINDABLE:             "MH_PREBINDABLE",
	MH_ALLMODSBOUND:            "MH_ALLMODSBOUND",
	MH_SUBSECTIONS_VIA_SYMBOLS: "MH_SUBSECTIONS_VIA_SYMBOLS",
	MH_CANONICAL:               "MH_CANONICAL",
	MH_WEAK_DEFINES:            "MH_WEAK_DEFINES",
	MH_BINDS_TO_WEAK:           "MH_BINDS_TO_WEAK",
	MH_ALLOW_STACK_EXECUTION:   "MH_ALLOW_STACK_EXECUTION",
	MH_ROOT_SAFE:               "MH_ROOT_SAFE",
	MH_SETUID_SAFE:             "MH_SETUID_SAFE",
	MH_NO_REEXPORTED_DYLIBS:    "MH_NO_REEXPORTED_DYLIBS",
	MH_PIE:                     "MH_PIE",
	MH_DEAD_STRIPPABLE_DYLIB: "MH_DEAD_STRIPPABLE_DYLIB",
	MH_HAS_TLV_DESCRIPTORS:   "MH_HAS_TLV_DESCRIPTORS",
	MH_NO_HEAP_EXECUTION:     "MH_NO_HEAP_EXECUTION",
}

const (
	_MH_EXECUTE_SYM  = "__mh_execute_header"
	MH_EXECUTE_SYM   = "_mh_execute_header"
	_MH_BUNDLE_SYM   = "__mh_bundle_header"
	MH_BUNDLE_SYM    = "_mh_bundle_header"
	_MH_DYLIB_SYM    = "__mh_dylib_header"
	MH_DYLIB_SYM     = "_mh_dylib_header"
	_MH_DYLINKER_SYM = "__mh_dylinker_header"
	MH_DYLINKER_SYM  = "_mh_dylinker_header"
)

const (
	SEG_PAGEZERO      = "__PAGEZERO"
	SEG_TEXT          = "__TEXT"
	SECT_TEXT         = "__text"
	SECT_FVMLIB_INIT0 = "__fvmlib_init0"
	SECT_FVMLIB_INIT1 = "__fvmlib_init1"
	SEG_DATA          = "__DATA"
	SECT_DATA         = "__data"
	SECT_BSS          = "__bss"
	SECT_COMMON       = "__common"
	SEG_OBJC          = "__OBJC"
	SECT_OBJC_SYMBOLS = "__symbol_table"
	SECT_OBJC_MODULES = "__module_info"
	SECT_OBJC_STRINGS = "__selector_strs"
	SECT_OBJC_REFS    = "__selector_refs"
	SEG_ICON          = "__ICON"
	SECT_ICON_HEADER  = "__header"
	SECT_ICON_TIFF    = "__tiff"
	SEG_LINKEDIT      = "__LINKEDIT"
	SEG_UNIXSTACK     = "__UNIXSTACK"
)

const (
	SECTION_ATTRIBUTES_USR   = 0xff000000
	S_ATTR_PURE_INSTRUCTIONS = 0x80000000
	S_ATTR_NO_TOC            = 0x40000000
	S_ATTR_STRIP_STATIC_SYMS = 0x20000000
	SECTION_ATTRIBUTES_SYS   = 0x00ffff00
	S_ATTR_SOME_INSTRUCTIONS = 0x00000400
	S_ATTR_EXT_RELOC         = 0x00000200
	S_ATTR_LOC_RELOC         = 0x00000100
)

var FLAG_SECTION_ATTRIBUTES = map[uint32]string{
	0x80000000: "S_ATTR_PURE_INSTRUCTIONS",
	0x40000000: "S_ATTR_NO_TOC",
	0x20000000: "S_ATTR_STRIP_STATIC_SYMS",
	0x10000000: "S_ATTR_NO_DEAD_STRIP",
	0x08000000: "S_ATTR_LIVE_SUPPORT",
	0x04000000: "S_ATTR_SELF_MODIFYING_CODE",
	0x02000000: "S_ATTR_DEBUG",
	0x00000400: "S_ATTR_SOME_INSTRUCTIONS",
	0x00000200: "S_ATTR_EXT_RELOC",
	0x00000100: "S_ATTR_LOC_RELOC",
}

var FLAG_SECTION_TYPES = map[uint32]string{
	0x0:  "S_REGULAR",
	0x1:  "S_ZEROFILL",
	0x2:  "S_CSTRING_LITERALS",
	0x3:  "S_4BYTE_LITERALS",
	0x4:  "S_8BYTE_LITERALS",
	0x5:  "S_LITERAL_POINTERS",
	0x6:  "S_NON_LAZY_SYMBOL_POINTERS",
	0x7:  "S_LAZY_SYMBOL_POINTERS",
	0x8:  "S_SYMBOL_STUBS",
	0x9:  "S_MOD_INIT_FUNC_POINTERS",
	0xa:  "S_MOD_TERM_FUNC_POINTERS",
	0xb:  "S_COALESCED",
	0xc:  "S_GB_ZEROFILL",
	0xd:  "S_INTERPOSING",
	0xe:  "S_16BYTE_LITERALS",
	0xf:  "S_DTRACE_DOF",
	0x10: "S_LAZY_DYLIB_SYMBOL_POINTERS",
	0x11: "S_THREAD_LOCAL_REGULAR",
	0x12: "S_THREAD_LOCAL_ZEROFILL",
	0x13: "S_THREAD_LOCAL_VARIABLES",
	0x14: "S_THREAD_LOCAL_VARIABLE_POINTERS",
	0x15: "S_THREAD_LOCAL_INIT_FUNCTION_POINTERS",
}

var LC_NAMES = map[uint32]string{
	LC_SEGMENT:                 "LC_SEGMENT",
	LC_IDFVMLIB:                "LC_IDFVMLIB",
	LC_LOADFVMLIB:              "LC_LOADFVMLIB",
	LC_ID_DYLIB:                "LC_ID_DYLIB",
	LC_LOAD_DYLIB:              "LC_LOAD_DYLIB",
	LC_LOAD_WEAK_DYLIB:         "LC_LOAD_WEAK_DYLIB",
	LC_SUB_FRAMEWORK:           "LC_SUB_FRAMEWORK",
	LC_SUB_CLIENT:              "LC_SUB_CLIENT",
	LC_SUB_UMBRELLA:            "LC_SUB_UMBRELLA",
	LC_SUB_LIBRARY:             "LC_SUB_LIBRARY",
	LC_PREBOUND_DYLIB:          "LC_PREBOUND_DYLIB",
	LC_ID_DYLINKER:             "LC_ID_DYLINKER",
	LC_LOAD_DYLINKER:           "LC_LOAD_DYLINKER",
	LC_THREAD:                  "LC_THREAD",
	LC_UNIXTHREAD:              "LC_UNIXTHREAD",
	LC_ROUTINES:                "LC_ROUTINES",
	LC_SYMTAB:                  "LC_SYMTAB",
	LC_DYSYMTAB:                "LC_DYSYMTAB",
	LC_TWOLEVEL_HINTS:          "LC_TWOLEVEL_HINTS",
	LC_PREBIND_CKSUM:           "LC_PREBIND_CKSUM",
	LC_SYMSEG:                  "LC_SYMSEG",
	LC_IDENT:                   "LC_IDENT",
	LC_FVMFILE:                 "LC_FVMFILE",
	LC_SEGMENT_64:              "LC_SEGMENT_64",
	LC_ROUTINES_64:             "LC_ROUTINES_64",
	LC_UUID:                    "LC_UUID",
	LC_RPATH:                   "LC_RPATH",
	LC_CODE_SIGNATURE:          "LC_CODE_SIGNATURE",
	LC_CODE_SEGMENT_SPLIT_INFO: "LC_CODE_SEGMENT_SPLIT_INFO",
	LC_REEXPORT_DYLIB:          "LC_REEXPORT_DYLIB",
	LC_LAZY_LOAD_DYLIB:         "LC_LAZY_LOAD_DYLIB",
	LC_ENCRYPTION_INFO:         "LC_ENCRYPTION_INFO",
	LC_DYLD_INFO:               "LC_DYLD_INFO",
	LC_DYLD_INFO_ONLY:          "LC_DYLD_INFO_ONLY",
	LC_LOAD_UPWARD_DYLIB:       "LC_LOAD_UPWARD_DYLIB",
	LC_VERSION_MIN_MACOSX:      "LC_VERSION_MIN_MACOSX",
	LC_VERSION_MIN_IPHONEOS:    "LC_VERSION_MIN_IPHONEOS",
	LC_FUNCTION_STARTS:         "LC_FUNCTION_STARTS",
	LC_DYLD_ENVIRONMENT:        "LC_DYLD_ENVIRONMENT",
	LC_MAIN:                    "LC_MAIN",
	LC_DATA_IN_CODE:            "LC_DATA_IN_CODE",
	LC_SOURCE_VERSION:          "LC_SOURCE_VERSION",
	LC_DYLIB_CODE_SIGN_DRS:     "LC_DYLIB_CODE_SIGN_DRS",
}
