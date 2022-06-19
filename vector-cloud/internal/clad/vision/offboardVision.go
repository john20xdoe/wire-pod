// Autogenerated Go message buffer code.
// Source: offboardVision.clad
// Full command line: victor-clad/tools/message-buffers/emitters/Go_emitter.py -C clad/types -I coretech/vision/clad_src coretech/common/clad_src -o generated/cladgo/src/clad/vision offboardVision.clad

package vision

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// ENUM OffboardCommsType
type OffboardCommsType uint8

const (
	OffboardCommsType_FileIO OffboardCommsType = iota
	OffboardCommsType_CLAD
)

// STRUCTURE OffboardImageReady
type OffboardImageReady struct {
	Timestamp    uint32
	NumRows      uint32
	NumCols      uint32
	NumChannels  uint32
	IsCompressed bool
	IsEncrypted  bool
	ProcTypes    []string
	Filename     string
}

func (o *OffboardImageReady) Size() uint32 {
	var result uint32
	result += 4 // Timestamp uint_32
	result += 4 // NumRows uint_32
	result += 4 // NumCols uint_32
	result += 4 // NumChannels uint_32
	result += 1 // IsCompressed bool
	result += 1 // IsEncrypted bool
	result += 1 // ProcTypes length (uint_8)
	for idx := range o.ProcTypes {
		result += 1                             // ProcTypes[idx] length (uint_8)
		result += uint32(len(o.ProcTypes[idx])) // uint_8 array
	}
	result += 1                       // Filename length (uint_8)
	result += uint32(len(o.Filename)) // uint_8 array
	return result
}

func (o *OffboardImageReady) Unpack(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.LittleEndian, &o.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &o.NumRows); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &o.NumCols); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &o.NumChannels); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &o.IsCompressed); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &o.IsEncrypted); err != nil {
		return err
	}
	var ProcTypesLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &ProcTypesLen); err != nil {
		return err
	}
	o.ProcTypes = make([]string, ProcTypesLen)
	for idx := range o.ProcTypes {
		var ProcTypesidxLen uint8
		if err := binary.Read(buf, binary.LittleEndian, &ProcTypesidxLen); err != nil {
			return err
		}
		o.ProcTypes[idx] = string(buf.Next(int(ProcTypesidxLen)))
		if len(o.ProcTypes[idx]) != int(ProcTypesidxLen) {
			return errors.New("string byte mismatch")
		}
	}
	var FilenameLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &FilenameLen); err != nil {
		return err
	}
	o.Filename = string(buf.Next(int(FilenameLen)))
	if len(o.Filename) != int(FilenameLen) {
		return errors.New("string byte mismatch")
	}
	return nil
}

func (o *OffboardImageReady) Pack(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.LittleEndian, o.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, o.NumRows); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, o.NumCols); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, o.NumChannels); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, o.IsCompressed); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, o.IsEncrypted); err != nil {
		return err
	}
	if len(o.ProcTypes) > 255 {
		return errors.New("max_length overflow in field ProcTypes")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(o.ProcTypes))); err != nil {
		return err
	}
	for idx := range o.ProcTypes {
		if len(o.ProcTypes[idx]) > 255 {
			return errors.New("max_length overflow in field ProcTypes[idx]")
		}
		if err := binary.Write(buf, binary.LittleEndian, uint8(len(o.ProcTypes[idx]))); err != nil {
			return err
		}
		if _, err := buf.WriteString(o.ProcTypes[idx]); err != nil {
			return err
		}
	}
	if len(o.Filename) > 255 {
		return errors.New("max_length overflow in field Filename")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(o.Filename))); err != nil {
		return err
	}
	if _, err := buf.WriteString(o.Filename); err != nil {
		return err
	}
	return nil
}

func (o *OffboardImageReady) String() string {
	return fmt.Sprint("Timestamp: {", o.Timestamp, "} ",
		"NumRows: {", o.NumRows, "} ",
		"NumCols: {", o.NumCols, "} ",
		"NumChannels: {", o.NumChannels, "} ",
		"IsCompressed: {", o.IsCompressed, "} ",
		"IsEncrypted: {", o.IsEncrypted, "} ",
		"ProcTypes: {", o.ProcTypes, "} ",
		"Filename: {", o.Filename, "}")
}

// STRUCTURE OffboardResultReady
type OffboardResultReady struct {
	Timestamp  uint32
	JsonResult string
}

func (o *OffboardResultReady) Size() uint32 {
	var result uint32
	result += 4                         // Timestamp uint_32
	result += 2                         // JsonResult length (uint_16)
	result += uint32(len(o.JsonResult)) // uint_8 array
	return result
}

func (o *OffboardResultReady) Unpack(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.LittleEndian, &o.Timestamp); err != nil {
		return err
	}
	var JsonResultLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &JsonResultLen); err != nil {
		return err
	}
	o.JsonResult = string(buf.Next(int(JsonResultLen)))
	if len(o.JsonResult) != int(JsonResultLen) {
		return errors.New("string byte mismatch")
	}
	return nil
}

func (o *OffboardResultReady) Pack(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.LittleEndian, o.Timestamp); err != nil {
		return err
	}
	if len(o.JsonResult) > 65535 {
		return errors.New("max_length overflow in field JsonResult")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(o.JsonResult))); err != nil {
		return err
	}
	if _, err := buf.WriteString(o.JsonResult); err != nil {
		return err
	}
	return nil
}

func (o *OffboardResultReady) String() string {
	return fmt.Sprint("Timestamp: {", o.Timestamp, "} ",
		"JsonResult: {", o.JsonResult, "}")
}