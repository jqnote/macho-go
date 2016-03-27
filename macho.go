package macho

import (
	"debug/macho"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

type MachoUUIDS map[string]string
type MachoFile struct {
	name  string
	uuids MachoUUIDS
}

func NewMacho(name string) (*MachoFile, error) {
	machoFile := &MachoFile{
		name:  name,
		uuids: MachoUUIDS{},
	}
	err := machoFile.load()
	if err != nil {
		return nil, err
	}
	return machoFile, nil
}

func NewMacheFromReader(r io.ReaderAt) (*MachoFile, error) {
	machoFile := &MachoFile{
		name:  "",
		uuids: MachoUUIDS{},
	}
	err := machoFile.loadFromReader(r)
	if err != nil {
		return nil, err
	}
	return machoFile, nil
}

func (m *MachoFile) loadFromReader(r io.ReaderAt) error {
	machoFile, err := macho.NewFile(r)
	if err != nil {
		fatFile, _ := macho.NewFatFile(r)
		defer fatFile.Close()
		for _, arch := range fatFile.Arches {
			archName, _ := m.getArchName(arch.Cpu, arch.SubCpu)
			m.parse_data(arch.Loads, arch.ByteOrder, archName)
		}
	} else {
		defer machoFile.Close()
		archName, _ := m.getArchName(machoFile.Cpu, machoFile.SubCpu)
		m.parse_data(machoFile.Loads, machoFile.ByteOrder, archName)
	}
	return nil
}

//加载符号表文件
func (m *MachoFile) load() error {
	r, err := os.Open(m.name)
	if err != nil {
		return err
	}

	machoFile, err := macho.NewFile(r)
	if err != nil {
		fatFile, _ := macho.NewFatFile(r)
		defer fatFile.Close()
		for _, arch := range fatFile.Arches {
			archName, _ := m.getArchName(arch.Cpu, arch.SubCpu)
			m.parse_data(arch.Loads, arch.ByteOrder, archName)
		}
	} else {
		defer machoFile.Close()
		archName, _ := m.getArchName(machoFile.Cpu, machoFile.SubCpu)
		m.parse_data(machoFile.Loads, machoFile.ByteOrder, archName)
	}
	return nil
}

//补充官方库没解析的 command 信息
//暂时只补充了 LC_UUID command
func (m *MachoFile) parse_data(loads []macho.Load, byteOrder binary.ByteOrder, archName string) {
	for _, l := range loads {
		_, ok := l.(*macho.Segment)
		if !ok { //补充未解的部分
			dat := l.Raw()
			cmd, _ := macho.LoadCmd(byteOrder.Uint32(dat[0:4])), byteOrder.Uint32(dat[4:8])
			if cmd == LC_UUID {
				m.uuids[strings.ToLower(archName)] = m.lc_uuid(dat)
			}
		}
	}
}

//获取符号表的 uuid 字典
func (m *MachoFile) GetUUIDS() MachoUUIDS {
	return m.uuids
}

//根据 cpu 类型，获取符号表的 uuid
func (m *MachoFile) GetUUIDByName(arch string) (string, error) {
	uuid, ok := m.uuids[arch]
	if ok {
		return uuid, nil
	} else {
		return "", nil
	}
}

func (m *MachoFile) getArchName(Cpu macho.Cpu, SubCpu uint32) (string, error) {
	subCpuTypeName := getCpuSubTypeName(Cpu, SubCpu)
	archName := strings.Replace(subCpuTypeName, "CPU_SUBTYPE_", "", 1)
	archName = strings.Replace(archName, "_", "", 1)
	if archName == "ARMALL12" {
		archName = CPU_TYPE_NAMES[Cpu]
	}
	return archName, nil
}

//解析uuid的数据
func (m *MachoFile) lc_uuid(data []byte) string {
	uuidRaw := data[8:]
	uuidHex := make([]byte, hex.EncodedLen(len(uuidRaw)))
	hex.Encode(uuidHex, uuidRaw)
	uuid := fmt.Sprintf("%s-%s-%s-%s-%s\n", uuidHex[0:8], uuidHex[8:12], uuidHex[12:16], uuidHex[16:20], uuidHex[20:])
	return strings.ToUpper(uuid)
}
