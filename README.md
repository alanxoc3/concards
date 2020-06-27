<!-- @> concards @ Console Cards @ A cool + simple flashcard app. <@ -->
# <img src="logo.svg" />

[![Build Status](https://travis-ci.org/alanxoc3/concards.svg?branch=master)](https://travis-ci.org/alanxoc3/concards)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanxoc3/concards)](https://goreportcard.com/report/github.com/alanxoc3/concards)
[![Coverage Status](https://coveralls.io/repos/github/alanxoc3/concards/badge.svg?branch=master)](https://coveralls.io/github/alanxoc3/concards?branch=master)

Turning notes into flashcards. Concards is my ongoing attempt to make
flashcards simple and quick to record or embed into a text document. Unlike the
overhead of other flashcard applications, Concards accepts a minimal amount of
rules. This allows tons of freedom in the way you want to create your own
flashcards!

## Install
Install like any other go application.
``` bash
go install github.com/alanxoc3/concards
~/$GOPATH/bin/concards --help
```

## Features
* Implements the [SM2](https://www.supermemo.com/english/ol/sm2.htm) Repetition Algorithm.
* Beautiful terminal gui.
* Supports UTF-8!
* Reading in from multiple files.
* Undoing/Redoing
* Easily editing a card while reviewing your cards.
* And More!!!

## Usage
The file syntax was designed to be very simple and flexible, allowing anyone to
quickly embed flashcards into their text document without extra hassle of other
flashcard apps (like [Anki](https://apps.ankiweb.net/) or
[Memrise](https://www.memrise.com/)).

The syntax to embed your flashcards is like this:
```
@> What is the answer?
 @ Here is the answer! <@
```

Wanna try it out? Run concards on this `README.md` file!
``` bash
concards README.md
```

### A Simple Concard
```
@> This is a question.
 @ Answer #1
 @ Answer #2
 @ Answer #3 <@
```

### 3 Simple Concards
```
@> This is question #1. @ Answer #1
@> This is question #2. @ Answer #1 @ Answer #2
@> This is question #3. @ Answer #1 @ Answer #2 @ Answer #3 <@
```

## All Special Tokens
All the special tokens that are a part of concards syntax are below.
```
@>   = Starts a concards block. Starts a question.
<@   = Ends the concards block.
@    = Separates answers.
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

This file is saved in `$CONCARDS_META`. If that environment variable doesn't
exist, then it is saved in `$HOME/.concards-meta`.

## Dev TODOs
- TODO: Implement the [leitner system](https://en.wikipedia.org/wiki/Leitner_system)
- TODO: Rework the terminal GUI.
- TODO: Add file name to the GUI.
- TODO: Enable CTRL-L (reloads the display).
- TODO: Add ability to change algorithm in GUI.
- TODO: Create a web flashcard front-end too.
- TODO: Create a man page.
- TODO: Create my own version of arg parse.
