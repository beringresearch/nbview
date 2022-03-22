## NBVIEW

`nbview` is a lightweight utility for viewing Jupyter Notebook files on the command line.

![nbview](screenshot.gif)

## Installation

To build from source:

`make PLATFORM`

where `PLATFORM` is either `ubuntu`, `linux`, `darwin`, or `windows`. The executable will appear in `bin/PLATFORM/nbview` and from there it can be added to you`$PATH` variable.

## Usage

`nbview FILENAME`

`nbview` loads a Jupyter Notebook into a scrollable buffer, which can be navigated using arrow keys. To exit, press `q` or `esc`.