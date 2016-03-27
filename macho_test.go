package macho

import (
	"log"
	"testing"
)

func testMacho(name string) error {
	log.Printf("File name : %s\n", name)
	macho, err := NewMacho(name)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	uuids := macho.GetUUIDS()
	log.Printf("UUIDS\n%v\n", uuids)
	var uuid string
	uuid, _ = macho.GetUUIDByName("arm64")
	log.Printf("uuid of arm64\n%v\n", uuid)

	uuid, _ = macho.GetUUIDByName("armv7")
	log.Printf("uuid of armv7\n%v\n", uuid)
	return nil
}
func TestMachoDecode(t *testing.T) {

	testMacho("test/001")

	testMacho("test/003") //003 not exists
	testMacho("test/002")
}
