<!-- Ignore comments for concards... !> --> <!-- * ## <@ -->
# <img src="logo.svg" />

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

<!-- @> Secret Card!!! Name all the Concards -->
## Features <!-- @ -->
* Implements the [SM2](https://www.supermemo.com/english/ol/sm2.htm) Repetition Algorithm. <!-- @ -->
* Beautiful terminal gui. <!-- @ -->
* Supports UTF-8! <!-- @ -->
* Reading in from multiple files. <!-- @ -->
* Undoing/Redoing <!-- @ -->
* Easily editing a card while reviewing your cards. <!-- @ -->
* And More!!! <!-- <@ -->

## Usage
The file syntax was designed to be very simple and flexible, allowing anyone to
quickly embed flashcards into their text document without extra hassle of other
flashcard apps (like [Anki](https://apps.ankiweb.net/) or
[Memrise](https://www.memrise.com/)).

The syntax to embed your flashcards is like this:
```
@> What is the answer?
 @ This is the answer! <@
```

Wanna try it out? Run concards on this README.md file!
``` bash
concards README.md
```
