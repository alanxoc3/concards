<!-- @> Who is the coolest person in the world? | You are :D. | Have a great day! <@ -->
# <img src="logo.svg" />

[![Build Status](https://travis-ci.org/alanxoc3/concards.svg?branch=master)](https://travis-ci.org/alanxoc3/concards)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanxoc3/concards)](https://goreportcard.com/report/github.com/alanxoc3/concards)
[![Coverage Status](https://coveralls.io/repos/github/alanxoc3/concards/badge.svg?branch=master)](https://coveralls.io/github/alanxoc3/concards?branch=master)

Turning notes into flashcards, or should I say concards! Concards is my ongoing
attempt to make flashcards simple and easily embeddable into text document
based notes. Concards is much lighter than other flashcard applications such as
such as [Anki](https://apps.ankiweb.net/) or
[Memrise](https://www.memrise.com/), but is also very powerful by following the
[Unix Philosophy](https://en.wikipedia.org/wiki/Unix_philosophy) of "Do one
thing and do it well".

## Features
- Implements the [SM2](https://www.supermemo.com/english/ol/sm2.htm) Repetition Algorithm.
- Beautiful terminal gui.
- Supports UTF-8!
- Reading in from multiple files & directories.
- Undoing/Redoing.
- Easily editing a card while reviewing your cards.

## Install
Download the latest binary executable from the [release
page](https://github.com/alanxoc3/concards/releases). At the moment, only Linux
and Mac are supported.

### Building From Source
Use the `go install` command, passing in snapshot as the version concards
compiles with.
``` bash
go install -ldflags="-X main.version=snapshot" github.com/alanxoc3/concards
```

## Usage
The complete syntax of embedding your flashcards into text documents consists
of these keywords:
```
'@>' = Starts a concards block and also starts a question.
'|'  = Separates sides.
'<@' = Ends the concards block.
'\'  = Escapes the special tokens above.
```

Here are a few example concards:
```
@> What does this mean 你好世界?
 | Hello World

@> What does "concards" stand for?
 | Console Cards

@> Can a concard have more than 2 sides?
 | Yes. | Yes it can.

@> What does a concard look like?
 | \@> It could look like this \| What does a concard look like? \<@
 | \@> It could also look like this \| with multiple \| answers! \<@

@> How do you escape a concard keyword?
 | Put a backslash before it. Your text file would show "\@>", but the app shows "@>".

@> How do you show a backslash then a keyword in the concards ui?
 | To see \\@> in the ui, your text document must have 2 backslashes: \\\@>

@> The person who created concards.
 | Alan Morgan
<@
```

The easiest way to understand that syntax is by trying it out! Just run
concards on this `README.md` file and see what happens!
``` bash
$ concards README.md
```

## Advanced Usage
### The Meta Data File
Here is an example meta-data file:
```
3dda75cb44ed447186834541475f32e2 2019-01-01T00:00:00Z 0 sm2 2.5
8525b45f883c05eec46b4f7a88e7f7ef 2020-01-01T00:00:00Z 0 sm2 2.5
```

Here is the same file, but annotated:
```
sha256sum cut in half            | review timestamp      | streak | alg | data
---------------------------------+-----------------------+--------+-----+-----
3dda75cb44ed447186834541475f32e2 | 2019-01-01T00:00:00Z  | 0      | sm2 | 2.5
8525b45f883c05eec46b4f7a88e7f7ef | 2020-01-01T00:00:00Z  | 0      | sm2 | 2.5
```

This file is saved to `$CONCARDS_META`. If that environment variable doesn't
exist, then it is saved to `$HOME/.concards-meta`.
