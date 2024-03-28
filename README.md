[![Go](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml/badge.svg)](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml)

# BoostPi
Raspberry Pi based boost monitor written in Go

![demo image](https://github.com/hamburgertrain/BoostPi/blob/main/images/demo.jpg?raw=true)

## Target System Requirements
- Raspberry Pi (should be compatible with any version as long as it is 64-bit) with i2c enabled running Debian 12 bookworm

- A USB OBD-II ELM327 reader

- An i2c 1602 16x2 Serial LCD

## Build System Requirements
- Linux Host System - even WSL will suffice

- Go 1.21.6

### Optional

- golangci-lint v1.55.2

## Setup

- Enable i2c interface on Raspberry Pi via `sudo raspi-config`

- Connect your display to your RaspberryPi via i2c, install i2cdetect `sudo apt-get install i2c-tools`

- Use `i2cdetect -y 1` to find out the address for the display (your Pi may have a different bus number depending on revision, scanning 0 and 1 covers most cases)

- Once you have your address, replace `i2cBus` and `i2cAddress` in `internal/display/display.go` as appropriate for your configuration

- Connect your USB OBD reader to the Raspberry Pi, find out what the USB address is and replace `elm327DeviceLocation` in `internal/elm327/elm327.go` as appropriate

- Follow Building instructions below

- I recommend setting the BoostPi executable to run on startup of the Raspberry Pi system via systemctl, systemd, etc.

- Connect OBD to your car and if everything is set up correctly, you should be tracking current and max boost pressure in PSI

## Building
Clone Submodule:
`git submodule init`
`git submodule update`

Build:
`env GOOS=linux GOARCH=arm64 go build`

## Credits
- https://wiki.52pi.com/index.php?title=Z-0234#1602_Serial_LCD_Module_Display

- DenisFromHR (Denis Pleic) - https://gist.github.com/DenisFromHR/cc863375a6e19dce359d

- http://www.recantha.co.uk/blog/?p=4849

## License
This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.