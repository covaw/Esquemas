// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Phase.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

// import "C"

type UnionNullPhaseDataTypeEnum int

const (
	UnionNullPhaseDataTypeEnumPhaseData UnionNullPhaseDataTypeEnum = 1
)

type UnionNullPhaseData struct {
	Null      *types.NullVal
	PhaseData PhaseData
	UnionType UnionNullPhaseDataTypeEnum
}

func writeUnionNullPhaseData(r *UnionNullPhaseData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullPhaseDataTypeEnumPhaseData:
		return writePhaseData(r.PhaseData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullPhaseData")
}

func NewUnionNullPhaseData() *UnionNullPhaseData {
	return &UnionNullPhaseData{}
}

func (r *UnionNullPhaseData) Serialize(w io.Writer) error {
	return writeUnionNullPhaseData(r, w)
}

func DeserializeUnionNullPhaseData(r io.Reader) (*UnionNullPhaseData, error) {
	t := NewUnionNullPhaseData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullPhaseDataFromSchema(r io.Reader, schema string) (*UnionNullPhaseData, error) {
	t := NewUnionNullPhaseData()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullPhaseData) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EndDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"StartDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"AdvanceNotice\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DecouplingCause\",\"type\":[\"null\",\"string\"]},{\"name\":\"Compensation\",\"type\":\"boolean\"},{\"name\":\"Holidays\",\"type\":\"boolean\"},{\"name\":\"Real\",\"type\":\"boolean\"},{\"name\":\"RecognizedStartDate\",\"type\":\"boolean\"},{\"name\":\"Salary\",\"type\":\"boolean\"},{\"name\":\"Status\",\"type\":\"boolean\"},{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"EmployeeId\",\"type\":\"int\"}],\"name\":\"PhaseData\",\"namespace\":\"Andreani.RHpro.Events.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullPhaseData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullPhaseData) SetLong(v int64) {

	r.UnionType = (UnionNullPhaseDataTypeEnum)(v)
}

func (r *UnionNullPhaseData) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.PhaseData = NewPhaseData()
		return &types.Record{Target: (&r.PhaseData)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullPhaseData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullPhaseData) Finalize()                        {}

func (r *UnionNullPhaseData) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullPhaseDataTypeEnumPhaseData:
		return json.Marshal(map[string]interface{}{"Andreani.RHpro.Events.Common.PhaseData": r.PhaseData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullPhaseData")
}

func (r *UnionNullPhaseData) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.RHpro.Events.Common.PhaseData"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.PhaseData)
	}
	return fmt.Errorf("invalid value for *UnionNullPhaseData")
}
