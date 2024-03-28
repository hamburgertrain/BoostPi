// Houses utility functions
// Copyright (C) 2024 Tyler Bialoblocki
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package utilities

import (
	"strconv"
)

// Converts a string to uint8
func ConvertToUint8(input string) (uint8, error) {
	parsedValue, err := strconv.ParseUint(input, 0, 8)
	if err != nil {
		return 0, err
	}

	return uint8(parsedValue), nil
}
