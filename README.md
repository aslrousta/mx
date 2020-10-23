# Macro eXpander

MX is a macro processor which copies the input to the output and expands the
macros as it scans through. Despite the great M4 macro processor, MX is designed
to be much easier to use and bring less ambiguities.

## Getting Started

MX at its basics, copies the input to the output character by character, e.g:

```tex
Hello, World
```

will produce `Hello, World`, as expected. It supports uncode texts as well,
e.g.:

```tex
こんにちは世界
```

produces the same result. However, it differs from M4 as it skips multiple
whitespaces by collapsing a sequence of whitespace characters (spaces, tabs and
new-lines) into one, e.g:

```tex
Hello,
    World
```

will produce `Hello, World`, as described. The only exception is for two (or
more) consequent new-lines that mark the start of a new paragraph, so:

```tex
This is a
    paragrah.

And, this is another one.
```

is turned by MX into:

```
This is a paragraph.

And, this is another one.
```

MX macro definition is mostly inspired by TeX, as it uses `\` to start an escape
sequence, `` ` `` for quotation, and `{}` for grouping. Similarly, macros are
defined by `def` built-in command, e.g.:

```tex
\def\mymacro{#1}{Hello, #1}
\mymacro World
```

that also expands to `Hello, World`.

 Macros can accept upto 9 arguments (namely, `#1` to `#9`), and are expanded
using pattern matching, e.g:

```tex
\def\reorder{#2, #1}{#1 #2}
\reorder Smith, John
```

produces `John Smith`, as arguments `#1` and `#2` match to `John` and `Smith`.
Note that how MX recognizes `,` in the pattern correctly and eliminates it from
the output. MX performs the pattern matching over _tokens_ (or _words_). A token
is a sequence of visible characters delimited by whitespaces, or grouped by
`{}`.

## License

MX is published under GNU GPLv3 license.
