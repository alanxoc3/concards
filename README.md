# <img src="logo.svg" />

[![Build Status](https://travis-ci.com/alanxoc3/concards.svg?branch=main)](https://travis-ci.com/alanxoc3/concards)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanxoc3/concards)](https://goreportcard.com/report/github.com/alanxoc3/concards)
[![Coverage Status](https://coveralls.io/repos/github/alanxoc3/concards/badge.svg?branch=main)](https://coveralls.io/github/alanxoc3/concards?branch=main)

Turning notes into flashcards, or should I say concards! This is my ongoing attempt to make flashcards more simple and convenient. Concards provides much of same functionality of other mainstream flashcard applications, but with a POSIX inspired twist.

## Features

- Configurable [spaced repetition](https://en.wikipedia.org/wiki/Spaced_repetition) algorithms!
- [Cloze](https://en.wikipedia.org/wiki/Cloze_test) cards, reversible cards, and multi-sided cards all supported!
- Cards embedded in note files, similar to [literate programming](https://en.wikipedia.org/wiki/Literate_programming)!
- [UTF-8](https://en.wikipedia.org/wiki/UTF-8) as a first-class citizen!
- Full-fledged CLI support!

## Install

Download the latest release from the [release page](https://github.com/alanxoc3/concards/releases). At the moment, only Linux and Mac are supported.

You can also build a snapshot from source with the `go` command:

```bash
$ go install github.com/alanxoc3/concards
```

Once installed, you may want to try running concards on this readme:

```bash
$ concards README.md
```

You may also want to review the output of the help command:

```bash
$ concards --help
```

## Basic Syntax

You can learn the full flashcard embedding syntax in just a few minutes! Let's get started.

### Creating a flashcard

To make a flashcard, you must put the flashcard text within a concards block. A concards block looks like this `#: :#`, where text would be placed between the two colons. Ex:

```
#: This is a one sided flashcard. :#
```

The text above will produce a one sided flashcard! But flashcards are normally 2 sided, so let's create a new flashcard with a question and answer:

```
#: What is a great way to decrease the effects of the forgetting curve?
 : Spending time every day to review previously learned information. :#
```

Cards in concards can only have 1 question, but they can have any number of answers:

```
#: What are Newton's 3 laws of motion?
 : 1. an object at rest stays at rest unless acted upon.
 : 2. force is equal to mass times acceleration.
 : 3. for every action, there is an equal and opposite reaction. :#
```

You can use the begin delimiter multiple times before using the end delimiter. This will create multiple cards:

```
#: Who published the first flashcards? : Favell Lee Mortimer
#: When were the first flashcards published? : 1834 :#
```

### Reversible Cards

When learning a language, you might find yourself writing a flashcard that translates a phrase from esperanto to english and rewriting the same flashcard translating the same phrase from english to esperanto. With concards though, you just have to use the `::` operator:

```
#: saluton al la mundo
:: hello world :#

Generates these cards:

#: saluton al la mundo : hello world
#: hello world : saluton al la mundo :#
```

You can use as many `::` operators as you want. It might be useful if you're learning mandarin as well:

```
#: spaghetti
:: spagetoj
:: 意面 :#

Generates these cards:

#: spaghetti : spagetoj : 意面
#: spagetoj : spaghetti : 意面
#: 意面 : spaghetti : spagetoj :#
```

Concards allows you to combine the `::` and `:` operators together. This can be useful if two different concepts in one language translates to the same word in another language:

```
#: apricot
 : almond
:: 杏仁 :#

Generates these cards:

#: apricot : 杏仁
#: almond : 杏仁
#: 杏仁 : apricot : almond :#
```

### Cloze Cards

A lot of cards phrased in question form can be rewritten in cloze form. A cloze in concards is expressed with curly braces `{...}`. Concards will generate cards from the text in the curly braces and replace the text with an empty set of curly braces `{}` which signifies a blank. Here is an example cloze card:

```
#: {Sebastian Leitner} published about the {Leitner System} in {1972}. :#

Generates these card:

#: {} published about the Leitner System in 1972. : Sebastian Leitner
#: Sebastian Leitner published about the {} in 1972. : Leitner System
#: Sebastian Leitner published about the Leitner System in {}. : 1972 :#
```

Cloze nesting is supported:

```
#: {Education is the {kindling of a flame}}, {not the {filling of a vessel}}. :#

Generates this cards:

#: {}, not the filling of a vessel. : Education is the kindling of a flame
#: Education is the {}, not the filling of a vessel. : kindling of a flame
#: Education is the kindling of a flame, {}. : not the filling of a vessel
#: Education is the kindling of a flame, not the {}. : filling of a vessel :#
```

Finally, you can combine the cloze syntax with the `::` and `:` operators:

```
#: {新型}{冠状}{病毒} :: Coronavirus : COVID-19 :#

Generates these cards:

#: {}冠状病毒 : 新型
#: 新型{}病毒 : 冠状
#: 新型冠状{} : 病毒
#: 新型冠状病毒 : Coronavirus : COVID-19
#: Coronavirus : 新型冠状病毒
#: COVID-19 : 新型冠状病毒 :#
```

### Whitespace & Escaping

Concards ignores consecutive whitespace and treats whitespace between delimiters as optional. The following flashcards are equivalent:

```
#: {Piotr A. Woźniak} created the SM-2 spaced repetition algorithm in {1987}.
#: { Piotr A. Woźniak } created the SM-2   spaced   repetition algorithm in { 1987}.
#:{Piotr A. Woźniak }created the
   SM-2 spaced repetition algorithm in{ 1987}.:#

Generates these cards:

#: {} created the SM-2 spaced repetition algorithm in 1987. : Piotr A. Woźniak
#: Piotr A. Woźniak created the SM-2 spaced repetition algorithm in {}. : 1987 :#
```

Backslash any reserved character or whitespace to include it in the card text:

```
#: Which characters are special in concards?
 : \# \: \{ \}

#: Leave my door open just a crack\
Cause I feel like such an insomniac\
Why do I tire of counting sheep?\
When I'm far too tired to fall asleep
 : Fireflies, by Owl City :#
```

## File Structure

Concards has config and data files in 2 directories, following the XDG standard.

The config directory is calculated by following this order of steps until one succeeds:

1. `$CONCARDS_CONFIG_DIR`
2. `$XDG_CONFIG_HOME/concards`
3. `$HOME/.config/concards`
4. `./`

The data directory is calculated by following this order of steps until one succeeds:

1. `$CONCARDS_DATA_DIR`
2. `$XDG_DATA_HOME/concards`
3. `$HOME/.local/share/concards`
4. `./`

### Events & Hooks

Concards events and hooks works similar to git hooks. You must place an executable file with a specific name in the "config directory". Files that begin with `event-` are run in parallel with concards and perform tasks that don't affect concards directly. Files that begin with `hook-` are not run in parallel with concards and can change the behaviour of concards.

Here are all the currently supported events:

- `event-review` - executed right after passing off a card with a pass or fail
- `event-startup` - executed once if/when concards starts the GUI up successfully

Here are all the currently supported hooks:

#### event-review & hook-review

Both the review event and hook are executed right after passing off a card with a pass or fail. Stdout for the event is ignored. Stdout for the hook must be a single date following the date format described in "The Meta File" section below. Here are the environment variables available to the event and hook, as well as example values:

```
CONCARD_HASH = 'YKB4BOBAU5WkkyLdhaah'
```

Program arguments are passed into the event and hook as well. Each argument represents a historical meta file entry in descending order according to the timestamp reviewed. This means a few things:
- there will always be at least 1 argument available
- the first argument is always the metadata of the time the card was reviewed right before executing the event/hook
- the last argument is always the metadata of the earliest time the card was reviewed

Here is an example of what the arguments might look like. For a more detailed explanation, see "The Meta File" section:

```
1 = '2020-12-01T01:47:11Z y AaA6231boaWTNyndaDZZ')
2 = '2020-10-01T01:47:11Z y AaA6231boaWTNyndaDZZ')
3 = '2020-09-01T01:47:11Z y AaA6231boaWTNyndaDZZ')
...
```

You set the new timestamp. Print a date as specified in "The Meta File" section to standard out. All other output will be ignored.

#### The Meta File

The meta file is located at `<data-dir>/meta`. This contains the data concards needs to function.

Here is an example meta file:

```
YKB4BOBAU5WkkyLdhaah 2020-11-21T03:47:12Z
2020-11-21T01:47:11Z n AaA6231boaWTNyndaDZZ
2020-11-21T03:47:12Z y ynda4BOBUa6231boawtn
2020-11-21T03:47:12Z y ynda4BOBUa6231boawtn

nYNDAdzzaAa6231BOAwt 2020-11-21T03:47:12Z
2020-11-21T01:47:11Z n obau5wKKYlDHAAHykb4b
2020-11-21T03:47:12Z y kkyLdhaahYKB4BOBAU5W obau5wKKYlDHAAHykb4b

nYNDAdzzaAa6231BOAwt never
```

There are 3 different data types in this file:

- date: either a date in UTC RFC-3339 format without spaces, or the string "never".
- hash: the sha256 of a string, truncated to the first 120 bits and represented in base64.
- bool: a y for yes or n for no.

And there are a few more definitions:

- block: synonomous with a paragraph.
- entry: synonomous with a line in a block.
- token: synonomous with a word in an entry.

Each block is divided into 2 parts:

- current entry: the first entry in a block.
- historical entries: all entries in a block, excluding the first one.

The current entry of each block contains 2 tokens. The first token is a hash of the card's question. The second token is a date describing when the card should be reviewed next. If the current time is after that date, the card will be marked as reviewable. Otherwise, the card will not show up in a review session.

The historical entries of each block represents a list of historical interactions with the card separated by new lines. Each item in this list has 3 or more tokens. The first token is a date showing the timestamp of when the card was reviewed. The second token is a bool that represents if the user knew the answer or not. And the remaining tokens are the different answers that were available when reviewing the card.

If the date in the first line is "never", then that card is blacklisted and will never show up in your review session. This could be useful to filter out noise, but also to show you know a card good well enough that you never need to review it again.

Invalid blocks, entries, or tokens may get filtered out when concards saves to a file.

- A current entry with an invalid/empty date or hash will filter out the entire block.
- A historical entry that has a "never" date, invalid date, or no hashes will get filtered out individually.
- Invalid hashes within a historical entry will get filtered out without filtering out the historical entry itself.

## Dependencies

Concards currently depends on these libraries:

- [stretchr/testify](https://github.com/stretchr/testify) for unit tests.
- [spf13/cobra](https://github.com/spf13/cobra) for CLI options.
- [nsf/termbox-go](https://github.com/nsf/termbox-go) for the terminal gui.
- [go-cmd/cmd](https://github.com/go-cmd/cmd) for hook support.
- [mattn/go-runewidth](https://github.com/mattn/go-runewidth) to help with Asian characters.

Concards wouldn't be where it is today without those repositories & their contributors, so please check them out too!

## Future Work

The next big things to focus on:

- Improved UIs.
- More hook support (algorithm, filters, etc)
- Server/client architecture.
