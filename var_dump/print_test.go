package var_dump

import (
  "testing"
  "unsafe"
)

func TestExportBool(t *testing.T) {
  value := true
  actual := Export(value)
  expected := "bool(true)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportInt(t *testing.T) {
  value := int(123)
  actual := Export(value)
  expected := "int(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportInt8(t *testing.T) {
  value := int8(123)
  actual := Export(value)
  expected := "int8(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportInt16(t *testing.T) {
  value := int16(123)
  actual := Export(value)
  expected := "int16(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportInt32(t *testing.T) {
  value := int32(123)
  actual := Export(value)
  expected := "int32(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportInt64(t *testing.T) {
  value := int64(123)
  actual := Export(value)
  expected := "int64(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportUint(t *testing.T) {
  value := uint(123)
  actual := Export(value)
  expected := "uint(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportUint8(t *testing.T) {
  value := uint8(123)
  actual := Export(value)
  expected := "uint8(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportUint16(t *testing.T) {
  value := uint16(123)
  actual := Export(value)
  expected := "uint16(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportUint32(t *testing.T) {
  value := uint32(123)
  actual := Export(value)
  expected := "uint32(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportUint64(t *testing.T) {
  value := uint64(123)
  actual := Export(value)
  expected := "uint64(123)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportFloat32(t *testing.T) {
  value := float32(1.23)
  actual := Export(value)
  expected := "float32(1.23)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportFloat64(t *testing.T) {
  value := float64(1.23)
  actual := Export(value)
  expected := "float64(1.23)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportComplex64(t *testing.T) {
  value := complex64(1.23 + 4.56i)
  actual := Export(value)
  expected := "complex64(1.23+4.56i)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportComplex128(t *testing.T) {
  value := complex128(1.23 + 4.56i)
  actual := Export(value)
  expected := "complex128(1.23+4.56i)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportString(t *testing.T) {
  value := "foo"
  actual := Export(value)
  expected := "string(\"foo\")"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportArray(t *testing.T) {
  value := [...]int{1, 2, 3}
  actual := Export(value)
  expected := "" + // for gofmt
    "[3]int{\n" +
    "  int(1),\n" +
    "  int(2),\n" +
    "  int(3),\n" +
    "}"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportSlice(t *testing.T) {
  value := []int{1, 2, 3}
  actual := Export(value)
  expected := "" + // for gofmt
    "[]int{\n" +
    "  int(1),\n" +
    "  int(2),\n" +
    "  int(3),\n" +
    "}"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportMap(t *testing.T) {
  value := map[string]int{"C": 1, "B": 2, "A": 3}
  actual := Export(value)
  expected := "" + // for gofmt
    "map[string]int{\n" +
    "  string(\"C\"): int(1),\n" +
    "  string(\"B\"): int(2),\n" +
    "  string(\"A\"): int(3),\n" +
    "}"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportPtr(t *testing.T) {
  base_value := "foo"
  value := &base_value
  actual := Export(value)
  expected := "&string(\"foo\")"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportNilPtr(t *testing.T) {
  var value *string
  actual := Export(value)
  expected := "(*string)nil"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func FunctionExample() {}

func TestExportFunction(t *testing.T) {
  value := FunctionExample
  actual := Export(value)
  expected := "<func is not supported>"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportChan(t *testing.T) {
  var value chan int
  actual := Export(value)
  expected := "(chan int)0x0"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

func TestExportUnsafePointer(t *testing.T) {
  value := unsafe.Pointer(uintptr(0x1234))
  actual := Export(value)
  expected := "unsafe.Pointer(0x1234)"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}

type InterfaceExample interface {
  PublicMethod()
}

type StructExampleChild struct {
  ChildPublicValue    string
  child_private_value int
}

type StructExample struct {
  PublicValue   string
  Self          *StructExample
  private_value int
  child         StructExampleChild
}

func (s StructExample) PublicMethod()   {}
func (s StructExample) private_method() {}

func TestExportInterfaceAndStruct(t *testing.T) {
  value := []InterfaceExample{
    StructExample{
      PublicValue:   "foo",
      private_value: 12345,
      child: StructExampleChild{
        ChildPublicValue:    "bar",
        child_private_value: 67890,
      },
      Self: &StructExample{
        PublicValue:   "hoge",
        private_value: 11111,
        child: StructExampleChild{
          ChildPublicValue:    "piyo",
          child_private_value: 22222,
        },
      },
    },
  }
  value[0].(StructExample).Self.Self = value[0].(StructExample).Self
  actual := Export(value)
  expected := "" + // for gofmt
    "[]var_dump.InterfaceExample{\n" +
    "  var_dump.StructExample{\n" +
    "    PublicValue: string(\"foo\"),\n" +
    "    Self: &var_dump.StructExample{\n" +
    "      PublicValue: string(\"hoge\"),\n" +
    "      Self: <infinite loop is detected>,\n" +
    "      private_value: int(11111),\n" +
    "      child: var_dump.StructExampleChild{\n" +
    "        ChildPublicValue: string(\"piyo\"),\n" +
    "        child_private_value: int(22222),\n" +
    "      },\n" +
    "    },\n" +
    "    private_value: int(12345),\n" +
    "    child: var_dump.StructExampleChild{\n" +
    "      ChildPublicValue: string(\"bar\"),\n" +
    "      child_private_value: int(67890),\n" +
    "    },\n" +
    "  },\n" +
    "}"
  if actual != expected {
    t.Errorf("Wrong format: %#v vs %#v", expected, actual)
  }
}
