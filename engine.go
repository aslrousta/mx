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

import (
	"bufio"
	"errors"
	"io"
	"os"
	"unicode"
)

const (
	// DefaultEscape is the default character which starts an escape sequence.
	DefaultEscape = '\\'
	// DefaultQuote is the default character which quotes a sequence.
	DefaultQuote = '`'
	// DefaultGroupStart is the default character which starts a grouping.
	DefaultGroupStart = '{'
	// DefaultGroupEnd is the default character which ends a grouping.
	DefaultGroupEnd = '}'
	// MinMacroBufSize is the minimum acceptable macro buffer size.
	MinMacroBufSize = 512
	// MinExpBufSize is the minimum acceptable expansion buffer size.
	MinExpBufSize = 512
)

const (
	modeNormal = iota
	modeWhitespace
	modeCommand
)

// Engine is the core of the macro expansion.
type Engine struct {

	// Reader is the input stream from-which characters are read. If Reader is
	// nil, it tries to read from os.Stdin.
	Reader io.Reader

	// Writer is the output stream to-which expanded result is written. If
	// Writer is nil, it tries to write into os.Stdout.
	Writer io.Writer

	// IncludePaths is the list of directory paths which the engine uses to look
	// for included files.
	IncludePaths []string

	// MacroBufSize is the number of unicode characters that the macro buffer
	// can hold. It should be at-least MinMacroBufSize characters long. A zero
	// indicates to use the minimum value.
	MacroBufSize int

	// ExpBufSize is the number of unicode characters that the expansion buffer
	// can hold. It should be at-least MinExpBufSize characters long. A zero
	// indicates to use the minimum value.
	ExpBufSize int

	escape     rune
	quote      rune
	groupStart rune
	groupEnd   rune

	mode int
	ch   rune

	// Macro buffer holds the definition and expantion of the macros. Its size
	// is determined using MacroBufSize field, and mpos keeps the current
	// position in the macro buffer.
	mbuf []rune
	mpos int

	// Expansion buffer hold the intermediate result of the macro expansion
	// before being transmitted to the output stream. Its size is determined
	// using ExpBufSize field, and epos keeps the current position in the
	// expansion buffer.
	ebuf []rune
	epos int
}

type runeReader interface {
	ReadRune() (r rune, size int, err error)
}

type runeWriter interface {
	WriteRune(r rune) (size int, err error)
}

// Execute runs the expansion until there is no character left in the input.
func (e *Engine) Execute() error {
	if e.Reader == nil {
		e.Reader = os.Stdin
	}

	if e.Writer == nil {
		e.Writer = os.Stdout
	}

	if e.MacroBufSize == 0 {
		e.MacroBufSize = MinMacroBufSize
	} else if e.MacroBufSize < MinMacroBufSize {
		return errors.New("too small macro buffer")
	}

	if e.ExpBufSize == 0 {
		e.ExpBufSize = MinExpBufSize
	} else if e.ExpBufSize < MinExpBufSize {
		return errors.New("too small expansion buffer")
	}

	var r runeReader
	if rr, ok := e.Reader.(runeReader); ok {
		r = rr
	} else {
		r = bufio.NewReader(e.Reader)
	}

	var w runeWriter
	if rw, ok := e.Writer.(runeWriter); ok {
		w = rw
	} else {
		w = bufio.NewWriter(e.Writer)
	}

	e.escape = DefaultEscape
	e.quote = DefaultQuote
	e.groupStart = DefaultGroupStart
	e.groupEnd = DefaultGroupEnd

	e.mode = modeNormal

	e.mbuf = make([]rune, e.MacroBufSize)
	e.ebuf = make([]rune, e.ExpBufSize)

	if err := e.expand(w, r); err != nil {
		return err
	}

	if f, ok := w.(interface{ Flush() error }); ok {
		return f.Flush()
	}

	return nil
}

func (e *Engine) expand(w runeWriter, r runeReader) (err error) {
	if err := e.read(r); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	for {
		switch e.mode {
		case modeNormal:
			err = e.doNormal(w, r)
		case modeWhitespace:
			err = e.doWhitespace(w, r)
		case modeCommand:
			err = e.doCommand(w, r)
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (e *Engine) doNormal(w runeWriter, r runeReader) error {
	for {
		if e.ch == e.escape {
			e.mode = modeCommand
			return nil
		} else if unicode.IsSpace(e.ch) {
			e.mode = modeWhitespace
			return nil
		} else if err := e.write(w, e.ch); err != nil {
			return err
		}
		if err := e.read(r); err != nil {
			return err
		}
	}
}

func (e *Engine) doWhitespace(w runeWriter, r runeReader) error {
	newLines := 0
	for {
		if e.ch == '\n' {
			newLines++
		} else if !unicode.IsSpace(e.ch) {
			break
		}
		if err := e.read(r); err != nil {
			return err
		}
	}
	if newLines > 1 {
		if err := e.write(w, '\n'); err != nil {
			return err
		}
		if err := e.write(w, '\n'); err != nil {
			return err
		}
	} else {
		if err := e.write(w, ' '); err != nil {
			return err
		}
	}
	e.mode = modeNormal
	return nil
}

func (e *Engine) doCommand(w runeWriter, r runeReader) error {
	return nil
}

func (e *Engine) read(r runeReader) (err error) {
	e.ch, _, err = r.ReadRune()
	return
}

func (e *Engine) write(w runeWriter, ch rune) (err error) {
	_, err = w.WriteRune(ch)
	return
}
