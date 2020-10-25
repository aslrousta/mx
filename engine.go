/*
 *  File: engine.go
 *  ---
 *
 *  Macro eXpander
 *  Copyright (C) 2020  Ali AslRousta
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package mx

import "io"

// Engine is the core of the macro expansion.
type Engine struct {

	// Reader is the stream from-which characters are read. If Reader is nil, it
	// tries to read from os.Stdin.
	Reader io.Reader

	// Writer is the stream to-which expanded result is written. If Writer is
	// nil, it tries to write into os.Stdout.
	Writer io.Writer
}

// Execute runs the expansion until there is no character left in the input.
func (e *Engine) Execute() error {
	return nil
}
