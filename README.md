[![Go](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml/badge.svg)](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml)

# BoostPi
Raspberry pi based boost monitor written in go

## Build Requirements
Linux Host System - even WSL will suffice

Go 1.22.6

## Target System Requirements
RaspberryPi (should be compatible with any version) running Debian 12 bookworm

i2c-tools/stable,now 4.3-2+b3 arm64

An OBD-II ELM327 reader

An i2c Serial LCD

## Building
Build using linux, even WSL will work.
Recommended build command is: `env GOOS=linux GOARCH=arm64 go build`

## Credits
Based off of Python code examples in https://wiki.52pi.com/index.php?title=Z-0234#1602_Serial_LCD_Module_Display

https://gist.github.com/DenisFromHR/cc863375a6e19dce359d

http://www.recantha.co.uk/blog/?p=4849

DenisFromHR (Denis Pleic)
