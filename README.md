# golibmodbus
This is a binding to [libmodbus](https://github.com/stephane/libmodbus) (http://libmodbus.org)

cgo flags are set for OSX with libmodbus installed from homebrew.

#Examples
##Create TCP client
```go
mb, err := golibmodbus.NewTCP("127.0.0.1", 1502)

mb.SetDebug(1)
mb.Connect()
mb.SetSlave(1)
mb.SetErrorRecovery(2)

reg, err := mb.ReadRegisters(1, 26)
fmt.Println(reg)

mb.Close()
mb.Free()
```
