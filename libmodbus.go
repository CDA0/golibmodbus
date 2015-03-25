package libmodbus

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

func VersionString() string {
  return C.LIBMODBUS_VERSION_STRING
}

func VersionHex() int {
  return C.LIBMODBUS_VERSION_HEX
}

func VersionCheck(major uint, minor uint, micro uint) bool {
  if major == VersionMajor() && minor == VersionMinor() && micro == VersionMicro() {
    return true
  }
  return false
}

func VersionMajor() uint {
  return uint(C.libmodbus_version_major)
}

func VersionMinor() uint {
  return uint(C.libmodbus_version_minor)
}

func VersionMicro() uint {
  return uint(C.libmodbus_version_micro)
}

func NewRTU(device string, baud int, parity byte, dataBit int, stopBit int) (*C.modbus_t, error) {
  cDevice := C.CString(device)
  cBaud := C.int(baud)
  cParity := C.char(parity)
  cDataBit := C.int(dataBit)
  cStopBit := C.int(stopBit)

  ctx, err := C.modbus_new_rtu(cDevice, cBaud, cParity, cDataBit, cStopBit)

  err = CheckError(err)

  return ctx, err
}

func NewTCP(ip string, port int) (*C.modbus_t, error) {
  cIP := C.CString(ip)
  cPort := C.int(port)

  ctx, err :=  C.modbus_new_tcp(cIP, cPort)

  err = CheckError(err)

  return ctx, err
}

func NewTCPPI(node string, service string) (*C.modbus_t, error) {
  cNode := C.CString(node)
  cService := C.CString(service)

  ctx, err :=  C.modbus_new_tcp_pi(cNode, cService)

  err = CheckError(err)

  return ctx, err
}

func RTUGetSerialMode(ctx *C.modbus_t) (int, error) {
  r, err :=  C.modbus_rtu_get_serial_mode(ctx)

  err = CheckError(err)

  return int(r), err
}

func RTUSetSerialMode(ctx *C.modbus_t, mode int) (int, error) {
  cMode := C.int(mode)

  r, err :=  C.modbus_rtu_set_serial_mode(ctx, cMode)

  err = CheckError(err)

  return int(r), err
}

func RTUGetRTS(ctx *C.modbus_t) (int, error) {
  r, err :=  C.modbus_rtu_get_rts(ctx)

  err = CheckError(err)

  return int(r), err
}

func RTUSetRTS(ctx *C.modbus_t, mode int) (int, error) {
  cMode := C.int(mode)
  r, err :=  C.modbus_rtu_set_rts(ctx, cMode)

  err = CheckError(err)

  return int(r), err
}

func Free(ctx *C.modbus_t) {
  C.modbus_free(ctx)
}

func SetSlave(ctx *C.modbus_t, slave int) (int, error) {
  cSlave := C.int(slave)

  r, err :=  C.modbus_set_slave(ctx, cSlave)

  err = CheckError(err)

  return int(r), err
}

func SetDebug(ctx *C.modbus_t, flag int) (int, error) {
  cFlag := C.int(flag)

  r, err :=  C.modbus_set_debug(ctx, cFlag)

  err = CheckError(err)

  return int(r), err
}

func GetByteTimeout(ctx *C.modbus_t, toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_get_byte_timeout(ctx, &cToSec, &cToUsec)

  err = CheckError(err)

  return int(r), err
}

func SetByteTimeout(ctx *C.modbus_t, toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_set_byte_timeout(ctx, cToSec, cToUsec)

  err = CheckError(err)

  return int(r), err
}

func GetResponseTimeout(ctx *C.modbus_t, toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_get_response_timeout(ctx, &cToSec, &cToUsec)

  err = CheckError(err)

  return int(r), err
}

func SetResponseTimeout(ctx *C.modbus_t, toSec uint32, toUsec uint32) (int, error) {
  cToSec := C.uint32_t(toSec)
  cToUsec := C.uint32_t(toUsec)

  r, err :=  C.modbus_set_response_timeout(ctx, cToSec, cToUsec)

  err = CheckError(err)

  return int(r), err
}

func SetErrorRecovery(ctx *C.modbus_t, errorRecovery uint) (int, error) {
  r, err :=  C.modbus_set_error_recovery(ctx, C.modbus_error_recovery_mode(errorRecovery))

  err = CheckError(err)

  return int(r), err
}

func SetSocket(ctx *C.modbus_t, s int) (int, error) {
  cS := C.int(s)

  r, err :=  C.modbus_set_socket(ctx, cS)

  err = CheckError(err)

  return int(r), err
}

func GetSocket(ctx *C.modbus_t) (int, error) {
  r, err :=  C.modbus_get_socket(ctx)

  err = CheckError(err)

  return int(r), err
}

func GetHeaderLength(ctx *C.modbus_t) (int, error) {
  r, err :=  C.modbus_get_header_length(ctx)

  err = CheckError(err)

  return int(r), err
}

func SetBitsFromByte(index int, value uint8) []uint8 {
  cIndex := C.int(index)
  cValue := C.uint8_t(value)

  output := make([]uint8, 8)
  cOutput := (*C.uint8_t)(&output[0])

  C.modbus_set_bits_from_byte(cOutput,cIndex, cValue)

  return output
}

func SetBitsFromBytes(index int, nbBits uint, tabByte *uint8) []uint8 {
  cIndex := C.int(index)
  cNbBits := C.uint(nbBits)
  cTabByte := (*C.uint8_t)(unsafe.Pointer(tabByte))

  output := make([]uint8, nbBits)
  cOutput := (*C.uint8_t)(&output[0])

  C.modbus_set_bits_from_bytes(cOutput, cIndex, cNbBits, cTabByte)

  return output
}

func GetByteFromBits(source *uint8, index int, nbBits uint) byte {
  cSource := (*C.uint8_t)(unsafe.Pointer(source))
  cIndex := C.int(index)
  cNbBits := C.uint(nbBits)

  return byte(C.modbus_get_byte_from_bits(cSource, cIndex,cNbBits))
}

func GetFloat(src *uint16) float32 {
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))

  return float32(C.modbus_get_float(cSrc))
}

func SetFloat(f float32) []uint16 {
  cF := C.float(f)

  output := make([]uint16, 2)
  cOutput := (*C.uint16_t)(&output[0])

  C.modbus_set_float(cF, cOutput)

  return output
}

func GetFloatDcba(src *uint16) float32 {
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))

  return float32(C.modbus_get_float_dcba(cSrc))
}

func SetFloatDcba(f float32) []uint16 {
  cF := C.float(f)

  output := make([]uint16, 2)
  cOutput := (*C.uint16_t)(&output[0])

  C.modbus_set_float(cF, cOutput)

  return output
}

func Connect(ctx *C.modbus_t) (int, error) {
  r, err :=  C.modbus_connect(ctx)

  err = CheckError(err)

  return int(r), err
}

func Close(ctx *C.modbus_t) {
  C.modbus_close(ctx)
}

func Flush(ctx *C.modbus_t) (int, error) {
  r, err := C.modbus_flush(ctx)

  err = CheckError(err)

  return int(r), err
}

func ReadBits(ctx *C.modbus_t, addr int, nb int) ([]uint8, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint8, nb)
  cOutput := (*C.uint8_t)(&output[0])

  _, err :=  C.modbus_read_bits(ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func ReadInputBits(ctx *C.modbus_t, addr int, nb int) ([]uint8, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint8, nb)
  cOutput := (*C.uint8_t)(&output[0])

  _, err := C.modbus_read_input_bits(ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func ReadRegisters(ctx *C.modbus_t, addr int, nb int) ([]uint16, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint16, nb)
  cOutput := (*C.uint16_t)(&output[0])

  _, err := C.modbus_read_registers(ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func ReadInputRegisters(ctx *C.modbus_t, addr int, nb int, dest *uint16) ([]uint16, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)

  output := make([]uint16, nb)
  cOutput := (*C.uint16_t)(&output[0])

  _, err := C.modbus_read_input_registers(ctx, cAddr, cNb, cOutput)

  err = CheckError(err)

  return output, err
}

func ReportSlaveId(ctx *C.modbus_t, maxDest int, dest *uint8) ([]uint8, error) {
  cMaxDest := C.int(maxDest)

  output := make([]uint8, maxDest)
  cOutput := (*C.uint8_t)(&output[0])

  _, err := C.modbus_report_slave_id(ctx, cMaxDest, cOutput)

  err = CheckError(err)

  return output, err
}

func WriteBit(ctx *C.modbus_t, addr int, status int) (int, error) {
  cAddr := C.int(addr)
  cStatus := C.int(status)

  r, err :=  C.modbus_write_bit(ctx, cAddr, cStatus)

  err = CheckError(err)

  return int(r), err
}

func WriteRegister(ctx *C.modbus_t, addr int, value int) (int, error) {
  cAddr := C.int(addr)
  cValue := C.int(value)

  r, err :=  C.modbus_write_register(ctx, cAddr, cValue)

  err = CheckError(err)

  return int(r), err
}

func WriteBits(ctx *C.modbus_t, addr int, nb int, src *uint8) (int, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)
  cSrc := (*C.uint8_t)(unsafe.Pointer(src))

  r, err :=  C.modbus_write_bits(ctx, cAddr, cNb, cSrc)

  return int(r), err
}

func WriteRegisters(ctx *C.modbus_t, addr int, nb int, src *uint16) (int, error) {
  cAddr := C.int(addr)
  cNb := C.int(nb)
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))

  r, err :=  C.modbus_write_registers(ctx, cAddr, cNb, cSrc)

  return int(r), err
}

func WriteAndReadRegisters(ctx *C.modbus_t, writeAddr int, writeNb int, src *uint16, readAddr int, readNb int) ([]uint16, error) {
  cWriteAddr := C.int(writeAddr)
  cWriteNb := C.int(writeNb)
  cSrc := (*C.uint16_t)(unsafe.Pointer(src))
  cReadAddr := C.int(readAddr)
  cReadNb := C.int(readNb)

  output := make([]uint16, readNb)
  cOutput := (*C.uint16_t)(&output[0])

  _, err := C.modbus_write_and_read_registers(ctx, cWriteAddr, cWriteNb, cSrc, cReadAddr, cReadNb, cOutput)

  err = CheckError(err)

  return output, err
}

func SendRawRequest(ctx *C.modbus_t, rawReq *uint8, rawReqLen int) (int, error) {
  cRawReq := (*C.uint8_t)(unsafe.Pointer(rawReq))
  cRawReqLen := C.int(rawReqLen)

  r, err :=  C.modbus_send_raw_request(ctx, cRawReq, cRawReqLen)

  err = CheckError(err)

  return int(r), err
}

func ReceiveConfirmation(ctx *C.modbus_t, rsp *uint8) (int, error) {
  cRsp := (*C.uint8_t)(unsafe.Pointer(rsp))

  r, err :=  C.modbus_receive_confirmation(ctx, cRsp)

  err = CheckError(err)

  return int(r), err
}

func ReplyException(ctx *C.modbus_t, req *uint8, exceptionCode int) (int, error) {
  cReq := (*C.uint8_t)(unsafe.Pointer(req))
  cExceptionCode := C.uint(exceptionCode)

  r, err :=  C.modbus_reply_exception(ctx, cReq, cExceptionCode)

  err = CheckError(err)

  return int(r), err
}

func MappingNew(nbBits int, nbInputBits int, nbRegisters int, nbInputRegisters int) (*C.modbus_mapping_t, error) {
  cNbBits := C.int(nbBits)
  cNbInputBits := C.int(nbInputBits)
  cNbRegisters := C.int(nbRegisters)
  cNbInputRegisters := C.int(nbInputRegisters)

  mb_mapping, err := C.modbus_mapping_new(cNbBits, cNbInputBits, cNbRegisters, cNbInputRegisters)

  err = CheckError(err)

  return mb_mapping, err
}

func MappingFree(mbMapping *C.modbus_mapping_t) {
  C.modbus_mapping_free(mbMapping)
}

func Receive(ctx *C.modbus_t, req *uint8) (int, error) {
  r, err :=  C.modbus_receive(ctx, (*C.uint8_t)(unsafe.Pointer(req)))

  err = CheckError(err)

  return int(r), err
}

func Reply(ctx *C.modbus_t, req *uint8, reqLength int, mb_mapping *C.modbus_mapping_t) (int, error) {
  cReq := (*C.uint8_t)(unsafe.Pointer(req))
  cReqLength := C.int(reqLength)

  r, err :=  C.modbus_reply(ctx,  cReq, cReqLength, mb_mapping)

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
