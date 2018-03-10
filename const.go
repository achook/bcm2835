package bcm2835

/*
#include "bcm2835.h"
*/
import "C"

type GpioPin uint8

// FunctionSelect as declared in bcm2835/bcm2835.h:731
type FunctionSelect uint8

// FunctionSelect enumeration from bcm2835/bcm2835.h:731
const (
	GpioFselInpt FunctionSelect = iota
	GpioFselOutp FunctionSelect = 1
	GpioFselAlt0 FunctionSelect = 4
	GpioFselAlt1 FunctionSelect = 5
	GpioFselAlt2 FunctionSelect = 6
	GpioFselAlt3 FunctionSelect = 7
	GpioFselAlt4 FunctionSelect = 3
	GpioFselAlt5 FunctionSelect = 2
	GpioFselMask FunctionSelect = 7
)

// PudControl as declared in bcm2835/bcm2835.h:741
type PudControl uint8

// PudControl enumeration from bcm2835/bcm2835.h:741
const (
	GpioPudOff  PudControl = iota
	GpioPudDown PudControl = 1
	GpioPudUp   PudControl = 2
)

// PadGroup as declared in bcm2835/bcm2835.h:769
type PadGroup uint8

// PadGroup enumeration from bcm2835/bcm2835.h:769
const (
	PadGroupGpio027  PadGroup = iota
	PadGroupGpio2845 PadGroup = 1
	PadGroupGpio4653 PadGroup = 2
)

// SpiBitOrder as declared in bcm2835/bcm2835.h:970
type SpiBitOrder uint8

// SpiBitOrder enumeration from bcm2835/bcm2835.h:970
const (
	SpiBitOrderLsbFirst SpiBitOrder = iota
	SpiBitOrderMsbFirst SpiBitOrder = 1
)

// SpiMode as declared in bcm2835/bcm2835.h:981
type SpiMode uint8

// SpiMode enumeration from bcm2835/bcm2835.h:981
const (
	SpiMode0 SpiMode = iota
	SpiMode1 SpiMode = 1
	SpiMode2 SpiMode = 2
	SpiMode3 SpiMode = 3
)

// SpiChipSelect as declared in bcm2835/bcm2835.h:992
type SpiChipSelect uint8

// SpiChipSelect enumeration from bcm2835/bcm2835.h:992
const (
	SpiCs0    SpiChipSelect = iota
	SpiCs1    SpiChipSelect = 1
	SpiCs2    SpiChipSelect = 2
	SpiCsNone SpiChipSelect = 3
)

// SpiClockDivider as declared in bcm2835/bcm2835.h:1027
type SpiClockDivider uint16

// SpiClockDivider enumeration from bcm2835/bcm2835.h:1027
const (
	SpiClockDivider65536 SpiClockDivider = iota
	SpiClockDivider32768 SpiClockDivider = 32768
	SpiClockDivider16384 SpiClockDivider = 16384
	SpiClockDivider8192  SpiClockDivider = 8192
	SpiClockDivider4096  SpiClockDivider = 4096
	SpiClockDivider2048  SpiClockDivider = 2048
	SpiClockDivider1024  SpiClockDivider = 1024
	SpiClockDivider512   SpiClockDivider = 512
	SpiClockDivider256   SpiClockDivider = 256
	SpiClockDivider128   SpiClockDivider = 128
	SpiClockDivider64    SpiClockDivider = 64
	SpiClockDivider32    SpiClockDivider = 32
	SpiClockDivider16    SpiClockDivider = 16
	SpiClockDivider8     SpiClockDivider = 8
	SpiClockDivider4     SpiClockDivider = 4
	SpiClockDivider2     SpiClockDivider = 2
	SpiClockDivider1     SpiClockDivider = 1
)

// I2cClockDivider as declared in bcm2835/bcm2835.h:1076
type I2cClockDivider uint16

// I2cClockDivider enumeration from bcm2835/bcm2835.h:1076
const (
	I2cClockDivider2500 I2cClockDivider = 2500
	I2cClockDivider626  I2cClockDivider = 626
	I2cClockDivider150  I2cClockDivider = 150
	I2cClockDivider148  I2cClockDivider = 148
)

// I2cReasonCodes as declared in bcm2835/bcm2835.h:1087
type I2cReasonCodes uint8

// I2cReasonCodes enumeration from bcm2835/bcm2835.h:1087
const (
	I2cReasonOk        I2cReasonCodes = iota
	I2cReasonErrorNack I2cReasonCodes = 1
	I2cReasonErrorClkt I2cReasonCodes = 2
	I2cReasonErrorData I2cReasonCodes = 4
)

// PwmClockDivider as declared in bcm2835/bcm2835.h:1159
type PwmClockDivider uint16

// PwmClockDivider enumeration from bcm2835/bcm2835.h:1159
const (
	PwmClockDivider2048 PwmClockDivider = 2048
	PwmClockDivider1024 PwmClockDivider = 1024
	PwmClockDivider512  PwmClockDivider = 512
	PwmClockDivider256  PwmClockDivider = 256
	PwmClockDivider128  PwmClockDivider = 128
	PwmClockDivider64   PwmClockDivider = 64
	PwmClockDivider32   PwmClockDivider = 32
	PwmClockDivider16   PwmClockDivider = 16
	PwmClockDivider8    PwmClockDivider = 8
	PwmClockDivider4    PwmClockDivider = 4
	PwmClockDivider2    PwmClockDivider = 2
	PwmClockDivider1    PwmClockDivider = 1
)
