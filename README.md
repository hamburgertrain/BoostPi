[![Go](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml/badge.svg)](https://github.com/hamburgertrain/BoostPi/actions/workflows/go.yml)

# BoostPi
RaspberryPi based boost monitor written in Go

## Target System Requirements
- RaspberryPi (should be compatible with any version as long as it is 64-bit) with i2c enabled running Debian 12 bookworm

- A USB OBD-II ELM327 reader

- An i2c 1602 16x2 Serial LCD

## Build System Requirements
- Linux Host System - even WSL will suffice

- Go 1.21.6

### Optional

- golangci-lint v1.55.2

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