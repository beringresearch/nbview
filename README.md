[![Go Report Card](https://goreportcard.com/badge/github.com/beringresearch/nbview)](https://goreportcard.com/report/github.com/beringresearch/nbview)

## NBVIEW

`nbview` is a lightweight utility for viewing Jupyter Notebook files in your terminal.

![nbview](screenshot.gif)

## Features

- [x] Cross-platform. Compiles to a multi-platform executable.
- [x] No dependencies. `nbview` binary is self contained.
- [x] No need for a Notebook server. `*.ipynb` files are rendered as text in a terminal buffer.
- [x] Fully-fledged Python syntax highlighting based on Pygments.

## Installation

The simplest way to install `nbview` is to [download the binary for your platform](https://github.com/beringresearch/nbview/releases). Rename the binary to `nbview` and add it to your `$PATH`, for example by copying it to `usr/local/bin`.

### Build from source

`make PLATFORM`

where `PLATFORM` is either `ubuntu`, `linux`, `darwin`, or `windows`. The executable will appear in `bin/PLATFORM/nbview` and from there it can be added to your `$PATH` variable.

## Usage

`nbview FILENAME`

`nbview` loads a Jupyter Notebook into a scrollable buffer, which can be navigated using arrow keys. To exit, press `q` or `esc`.