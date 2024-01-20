[![Go](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml/badge.svg)](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml)

# BoostPi
Raspberry pi based boost monitor written in go

## Target System Requirements
RaspberryPi (should be compatible with any version) with i2c enabled running Debian 12 bookworm

i2c-tools/stable,now 4.3-2+b3 arm64

A USB OBD-II ELM327 reader

An i2c 16x2 Serial LCD

## Build System Requirements
Linux Host System - even WSL will suffice

Go 1.22.6

## Building
`env GOOS=linux GOARCH=arm64 go build`

## Credits
https://wiki.52pi.com/index.php?title=Z-0234#1602_Serial_LCD_Module_Display

DenisFromHR (Denis Pleic) - https://gist.github.com/DenisFromHR/cc863375a6e19dce359d

http://www.recantha.co.uk/blog/?p=4849
