package misc

import (
	"os"
	"testing"
)

type GameConfig struct {
	ReadTimeout int
	LPort       string
	LGMPort     string

	LSSDBHost  string
	LSSDBPass  string
	LDBConnNum int

	LBasicSSDBHost  string
	LBasicSSDBPass  string
	LBasicDBConnNum int

	LDumpDBHost    string
	LDumpDBPass    string
	LDumpDBConnNum int

	LCsvPath        string
	LSeqStart       string
	LLogServer      string
	LSessionTimeout int
	LSessionPath    string
	LFightServer    string
	LConNum         int

	MMPidFilePath      string
	ServerId           int
	ChecksumPrivateKey string
	CheckKeysum        bool
	ChecksumTimeout    int64
	IsDebugCommand     bool
	InitDiamond        int32
	PrivateKey         string
	EnhancedKey        string
	PayServerHost      string
}

func TestReadFromJsonFile(t *testing.T) {
	ds := &GameConfig{}
	err := ReadFromJsonFile("../realm/game.cfg", ds)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteToJsonFile(t *testing.T) {
	ds := &GameConfig{}
	err := ReadFromJsonFile("../realm/game.cfg", ds)
	if err != nil {
		t.Fatal(err)
	}
	err = WriteToJsonFile("./game.cfg", ds)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove("./game.cfg")
}
