package bcm2835

/*
#include "bcm2835.h"
*/
import "C"

func Init() int {
	return int(C.bcm2835_init())
}

func Close() int {
	return int(C.bcm2835_close())
}

func SetDebug(debug uint8) {
	C.bcm2835_set_debug(C.uchar(debug))
}

func GpioFsel(pin GpioPin, function FunctionSelect) {
	C.bcm2835_gpio_fsel(C.uchar(pin), C.uchar(function))
}

func GpioLev(pin GpioPin) uint8 {
	return uint8(C.bcm2835_gpio_lev(C.uchar(pin)))
}

func GpioEds(pin GpioPin) uint8 {
	return uint8(C.bcm2835_gpio_eds(C.uchar(pin)))
}

func GpioSetEds(pin GpioPin) {
	C.bcm2835_gpio_eds(C.uchar(pin))
}

func GpioRen(pin GpioPin) {
	C.bcm2835_gpio_ren(C.uchar(pin))
}

func GpioClrRen(pin GpioPin) {
	C.bcm2835_gpio_clr_ren(C.uchar(pin))
}

func GpioFen(pin GpioPin) {
	C.bcm2835_gpio_fen(C.uchar(pin))
}

func GpioClrFen(pin GpioPin) {
	C.bcm2835_gpio_clr_fen(C.uchar(pin))
}

func GpioHen(pin GpioPin) {
	C.bcm2835_gpio_hen(C.uchar(pin))
}

func GpioClrHen(pin GpioPin) {
	C.bcm2835_gpio_clr_hen(C.uchar(pin))
}

func GpioLen(pin GpioPin) {
	C.bcm2835_gpio_len(C.uchar(pin))
}

func GpioClrLen(pin GpioPin) {
	C.bcm2835_gpio_clr_len(C.uchar(pin))
}

func GpioAren(pin GpioPin) {
	C.bcm2835_gpio_aren(C.uchar(pin))
}

func GpioClrAren(pin GpioPin) {
	C.bcm2835_gpio_clr_aren(C.uchar(pin))
}

func GpioAfen(pin GpioPin) {
	C.bcm2835_gpio_afen(C.uchar(pin))
}

func GpioClrAfen(pin GpioPin) {
	C.bcm2835_gpio_clr_afen(C.uchar(pin))
}

func GpioSet(pin GpioPin) {
	C.bcm2835_gpio_set(C.uchar(pin))
}

func GpioClr(pin GpioPin) {
	C.bcm2835_gpio_clr(C.uchar(pin))
}

func GpioPud(pin GpioPin) {
	C.bcm2835_gpio_pud(C.uchar(pin))
}

func GpioPudclk(pin GpioPin, on PudControl) {
	C.bcm2835_gpio_pudclk(C.uchar(pin), C.uchar(on))
}

func Delay(millis uint64) {
	C.bcm2835_delay(C.uint(millis))
}

func GpioWrite(pin GpioPin, on uint8) {
	C.bcm2835_gpio_write(C.uchar(pin), C.uchar(on))
}

func GpioSetPud(pin GpioPin, pud PudControl) {
	C.bcm2835_gpio_set_pud(C.uchar(pin), C.uchar(pud))
}

func SpiBegin() int {
	return int(C.bcm2835_spi_begin())
}

func SpiEnd() {
	C.bcm2835_spi_end()
}

func SpiSetBitOrder(order SpiBitOrder) {
	C.bcm2835_spi_setBitOrder(C.uchar(order))
}

func SpiSetClockDivider(divider SpiClockDivider) {
	C.bcm2835_spi_setClockDivider(C.ushort(divider))
}

func SpiSetDataMode(mode SpiMode) {
	C.bcm2835_spi_setDataMode(C.uchar(mode))
}

func SpiSelectChip(cs SpiChipSelect) {
	C.bcm2835_spi_chipSelect(C.uchar(cs))
}

func SetChipSelectPolaritry(cs SpiChipSelect, active uint8) {
	C.bcm2835_spi_setChipSelectPolarity(C.uchar(cs), C.uchar(active))
}

func SpiTransfer(value uint8) uint8 {
	return uint8(C.bcm2835_spi_transfer(C.uchar(value)))
}

func SpiWrite(data uint16) {
	C.bcm2835_spi_write(C.ushort(data))
}

func AuxSpiBegin() int {
	return int(C.bcm2835_aux_spi_begin())
}

func AuxSpiEnd() {
	C.bcm2835_aux_spi_end()
}

func AuxSpiSetClockDivider(divider SpiClockDivider) {
	C.bcm2835_aux_spi_setClockDivider(C.ushort(divider))
}

func AuxSpiCalcClockDivider(speed uint32) {
	C.bcm2835_aux_spi_CalcClockDivider(C.uint(speed))
}

func AuxSpiWrite(data uint16) {
	C.bcm2835_aux_spi_write(C.ushort(data))
}
