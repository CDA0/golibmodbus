package golibmodbus

/*
#cgo CFLAGS: -I/usr/local/include/modbus
#cgo LDFLAGS: -L/usr/local/lib/ -lmodbus
#include <errno.h>
#include <stdio.h>
#include <modbus.h>
*/
import "C"

import (
  "unsafe"
  "errors"
  "syscall"
)

type modbus_t C.modbus_t
type modbus_mapping_t C.modbus_mapping_t
type uintt_t C.uint8_t
type uint16_t C.uint16_t
type uint32_t C.uint32_t

type Modbus struct {
  ctx *C.modbus_t
  mb_mapping *C.modbus_mapping_t
}

func (m *Modbus) VersionString() string { return VersionString() }
func VersionString() string {
  return C.LIBMODBUS_VERSION_STRING
}

func (m *Modbus) VersionHex() int { return VersionHex() }
func VersionHex() int {
  return C.LIBMODBUS_VERSION_HEX
}

func (m *Modbus) VersionCheck(major uint, minor uint, micro uint) bool { return VersionCheck(major, minor, micro) }
func VersionCheck(major uint, minor uint, micro uint) bool {
  if major == VersionMajor() && minor == VersionMinor() && micro == VersionMicro() {
    return true
  }
  return false
}

func (m *Modbus) VersionMajor() uint { return VersionMajor() }
func VersionMajor() uint {
  return uint(C.libmodbus_version_major)
}

func (m *Modbus) VersionMinor() uint { return VersionMinor() }
func VersionMinor() uint {
  return uint(C.libmodbus_version_minor)
}

func (m *Modbus) VersionMicro() uint { return VersionMicro() }
func VersionMicro() uint {
  return uint(C.libmodbus_version_micro)
}

func NewRTU(device string, baud int, parity byte, dataBit int, stopBit int) (*Modbus, error) {
  cDevice := C.CString(device)
  cBaud := C.int(baud)
  cParity := C.char(parity)
  cDataBit := C.int(dataBit)
  cStopBit := C.int(stopBit)

  ctx, err := C.modbus_new_rtu(cDevice, cBaud, cParity, cDataBit, cStopBit)

  m := new(Modbus)

  err = CheckError(err)

  if err == nil { m.ctx = ctx }

  return m, err
}

func NewTCP(ip string, port int) (*Modbus, error) {
  cIP := C.CString(ip)
  cPort := C.int(port)

  ctx, err :=  C.modbus_new_tcp(cIP, cPort)

  m := new(Modbus)

  err = CheckError(err)

  if err == nil { m.ctx = ctx }

  return m, err
}

func NewTCPPI(node string, service string) (*Modbus, error) {
  cNode := C.CString(node)
  cService := C.CString(service)

  ctx, err :=  C.modbus_new_tcp_pi(cNode, cService)

  m := new(Modbus)

  err = CheckError(err)

  if err == nil { m.ctx = ctx }

  return m, err
}

func (m *Modbus) RTUGetSerialMode() (int, error) {
  r, err :=  C.modbus_rtu_get_serial_mode(m.ctx)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) RTUSetSerialMode(mode int) (int, error) {
  cMode := C.int(mode)

  r, err :=  C.modbus_rtu_set_serial_mode(m.ctx, cMode)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) RTUGetRTS() (int, error) {
  r, err :=  C.modbus_rtu_get_rts(m.ctx)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) RTUSetRTS(mode int) (int, error) {
  cMode := C.int(mode)
  r, err :=  C.modbus_rtu_set_rts(m.ctx, cMode)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) Free() {
  C.modbus_free(m.ctx)
}

func (m *Modbus) SetSlave(slave int) (int, error) {
  cSlave := C.int(slave)

  r, err :=  C.modbus_set_slave(m.ctx, cSlave)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) SetDebug(flag int) (int, error) {
  cFlag := C.int(flag)

  r, err :=  C.modbus_set_debug(m.ctx, cFlag)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) GetByteTimeout(toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_get_byte_timeout(m.ctx, &cToSec, &cToUsec)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) SetByteTimeout(toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_set_byte_timeout(m.ctx, cToSec, cToUsec)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) GetResponseTimeout(toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_get_response_timeout(m.ctx, &cToSec, &cToUsec)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) SetResponseTimeout(toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_set_response_timeout(m.ctx, cToSec, cToUsec)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) SetErrorRecovery(errorRecovery uint) (int, error) {
  r, err :=  C.modbus_set_error_recovery(m.ctx, C.modbus_error_recovery_mode(errorRecovery))

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) SetSocket(s int) (int, error) {
  cS := C.int(s)

  r, err :=  C.modbus_set_socket(m.ctx, cS)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) GetSocket() (int, error) {
  r, err :=  C.modbus_get_socket(m.ctx)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) GetHeaderLength() (int, error) {
  r, err :=  C.modbus_get_header_length(m.ctx)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) SetBitsFromByte(index int, value uint8) []uint8 { return SetBitsFromByte(index, value) }
func SetBitsFromByte(index int, value uint8) []uint8 {
  cIndex := C.int(index)
  cValue := C.uint8_t(value)

  output := make([]uint8, 8)
  cOutput := (*C.uint8_t)(&output[0])

  C.modbus_set_bits_from_byte(cOutput,cIndex, cValue)

  return output
}

func (m *Modbus) SetBitsFromBytes(index int, nbBits uint, tabByte *uint8) []uint8 { return SetBitsFromBytes(index, nbBits, tabByte) }
func SetBitsFromBytes(index int, nbBits uint, tabByte *uint8) []uint8 {
  cIndex := C.int(index)
  cNbBits := C.uint(nbBits)
  cTabByte := (*C.uint8_t)(unsafe.Pointer(tabByte))

  output := make([]uint8, nbBits)
  cOutput := (*C.uint8_t)(&output[0])

  C.modbus_set_bits_from_bytes(cOutput, cIndex, cNbBits, cTabByte)

  return output
}

func (m *Modbus) GetByteFromBits(source *uint8, index int, nbBits uint) byte { return GetByteFromBits(source, index, nbBits) }
func GetByteFromBits(source *uint8, index int, nbBits uint) byte {
  cSource := (*C.uint8_t)(unsafe.Pointer(source))
  cIndex := C.int(index)
  cNbBits := C.uint(nbBits)

  return byte(C.modbus_get_byte_from_bits(cSource, cIndex,cNbBits))
}

func (m *Modbus) GetFloat(src *uint16) float32 { return GetFloat(src) }
func GetFloat(src *uint16) float32 {
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))

  return float32(C.modbus_get_float(cSrc))
}

func (m *Modbus) SetFloat(f float32) []uint16 { return SetFloat(f) }
func SetFloat(f float32) []uint16 {
  cF := C.float(f)

  output := make([]uint16, 2)
  cOutput := (*C.uint16_t)(&output[0])

  C.modbus_set_float(cF, cOutput)

  return output
}

func (m *Modbus) GetFloatDcba(src *uint16) float32 { return GetFloatDcba(src) }
func GetFloatDcba(src *uint16) float32 {
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))

  return float32(C.modbus_get_float_dcba(cSrc))
}

func (m *Modbus) SetFloatDcba(f float32) []uint16 { return SetFloatDcba(f) }
func SetFloatDcba(f float32) []uint16 {
  cF := C.float(f)

  output := make([]uint16, 2)
  cOutput := (*C.uint16_t)(&output[0])

  C.modbus_set_float(cF, cOutput)

  return output
}

func (m *Modbus) Connect() (int, error) {
  r, err :=  C.modbus_connect(m.ctx)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) Close() {
  C.modbus_close(m.ctx)
}

func (m *Modbus) Flush() (int, error) {
  r, err := C.modbus_flush(m.ctx)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) ReadBits(addr int, nb int) ([]uint8, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint8, nb)
  cOutput := (*C.uint8_t)(&output[0])

  _, err :=  C.modbus_read_bits(m.ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func (m *Modbus) ReadInputBits(addr int, nb int) ([]uint8, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint8, nb)
  cOutput := (*C.uint8_t)(&output[0])

  _, err := C.modbus_read_input_bits(m.ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func (m *Modbus) ReadRegisters(addr int, nb int) ([]uint16, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint16, nb)
  cOutput := (*C.uint16_t)(&output[0])

  _, err := C.modbus_read_registers(m.ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func (m *Modbus) ReadInputRegisters(addr int, nb int, dest *uint16) ([]uint16, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint16, nb)
  cOutput := (*C.uint16_t)(&output[0])

  _, err := C.modbus_read_input_registers(m.ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func (m *Modbus) ReportSlaveId(maxDest int, dest *uint8) ([]uint8, error) {
  cMaxDest := C.int(maxDest)

  output := make([]uint8, maxDest)
  cOutput := (*C.uint8_t)(&output[0])

  _, err := C.modbus_report_slave_id(m.ctx, cMaxDest, cOutput)

  err = CheckError(err)

  return output, err
}

func (m *Modbus) WriteBit(addr int, status int) (int, error) {
  cAddr := C.int(addr)
  cStatus := C.int(status)

  r, err :=  C.modbus_write_bit(m.ctx, cAddr, cStatus)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) WriteRegister(addr int, value int) (int, error) {
  cAddr := C.int(addr)
  cValue := C.int(value)

  r, err :=  C.modbus_write_register(m.ctx, cAddr, cValue)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) WriteBits(addr int, nb int, src *uint8) (int, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)
  cSrc := (*C.uint8_t)(unsafe.Pointer(src))

  r, err :=  C.modbus_write_bits(m.ctx, cAddr, cNb, cSrc)

  return int(r), err
}

func (m *Modbus) WriteRegisters(addr int, nb int, src *uint16) (int, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))

  r, err :=  C.modbus_write_registers(m.ctx, cAddr, cNb, cSrc)

  return int(r), err
}

func (m *Modbus) WriteAndReadRegisters(writeAddr int, writeNb int, src *uint16, readAddr int, readNb int) ([]uint16, error) {
  cWriteAddr := C.int(writeAddr)
  cWriteNb := C.int(writeNb)
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))
  cReadAddr := C.int(readAddr)
  cReadNb := C.int(readNb)

  output := make([]uint16, readNb)
  cOutput := (*C.uint16_t)(&output[0])

  _, err := C.modbus_write_and_read_registers(m.ctx, cWriteAddr, cWriteNb, cSrc, cReadAddr, cReadNb, cOutput)

  err = CheckError(err)

  return output, err
}

func (m *Modbus) SendRawRequest(rawReq *uint8, rawReqLen int) (int, error) {
  cRawReq := (*C.uint8_t)(unsafe.Pointer(rawReq))
  cRawReqLen := C.int(rawReqLen)

  r, err :=  C.modbus_send_raw_request(m.ctx, cRawReq, cRawReqLen)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) ReceiveConfirmation(rsp *uint8) (int, error) {
  cRsp := (*C.uint8_t)(unsafe.Pointer(rsp))

  r, err :=  C.modbus_receive_confirmation(m.ctx, cRsp)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) ReplyException(req *uint8, exceptionCode int) (int, error) {
  cReq := (*C.uint8_t)(unsafe.Pointer(req))
  cExceptionCode := C.uint(exceptionCode)

  r, err :=  C.modbus_reply_exception(m.ctx, cReq, cExceptionCode)

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) MappingNew(nbBits int, nbInputBits int, nbRegisters int, nbInputRegisters int) error {
  cNbBits := C.int(nbBits)
  cNbInputBits := C.int(nbInputBits)
  cNbRegisters := C.int(nbRegisters)
  cNbInputRegisters := C.int(nbInputRegisters)

  mb_mapping, err := C.modbus_mapping_new(cNbBits, cNbInputBits, cNbRegisters, cNbInputRegisters)

  err = CheckError(err)

  if err == nil { m.mb_mapping = mb_mapping }

  return err
}

func (m *Modbus) MappingFree() {
  C.modbus_mapping_free(m.mb_mapping)
}

func (m *Modbus) Receive(req *uint8) (int, error) {
  r, err :=  C.modbus_receive(m.ctx, (*C.uint8_t)(unsafe.Pointer(req)))

  err = CheckError(err)

  return int(r), err
}

func (m *Modbus) Reply(req *uint8, reqLength int) (int, error) {
  cReq := (*C.uint8_t)(unsafe.Pointer(req))
  cReqLength := C.int(reqLength)

  r, err :=  C.modbus_reply(m.ctx,  cReq, cReqLength, m.mb_mapping)

  err = CheckError(err)

  return int(r), err
}

func CheckError(err error) error {
  var errno syscall.Errno

  if err != nil {
    errno = err.(syscall.Errno)
    err = errors.New(C.GoString(C.modbus_strerror(C.int(errno))))
  }

  return err
}
