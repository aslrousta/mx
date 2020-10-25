/*
 *  File: main.go
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

/*
 *  File: main.go
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

/*
 *  File: main.go
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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aslrousta/mx"
)

func printUsage() {
	fmt.Fprintf(
		flag.CommandLine.Output(),
		"Usage: mx [OPTION]... FILE\n"+
			"Processes and expands macros in FILE and writes the result to stdout.\n"+
			"If no FILE is given, it reads from stdin.\n\n"+
			"Options:\n",
	)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage
	flag.Parse()

	var e mx.Engine

	if len(flag.Args()) > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "! %v", err)
			os.Exit(2)
		}
		defer f.Close()
		e.Reader = f
	}

	if err := e.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "! %v", err)
		os.Exit(2)
	}
}
