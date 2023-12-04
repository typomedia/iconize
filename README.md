# `iconize` - Convert images to Windows icon files

`iconize` is a command line tool to convert an images to a multisize Windows icon file.

[![Go Report Card](https://goreportcard.com/badge/github.com/typomedia/iconize)](https://goreportcard.com/report/github.com/typomedia/iconize)
[![Go Reference](https://pkg.go.dev/badge/github.com/typomedia/iconize.svg)](https://pkg.go.dev/github.com/typomedia/iconize)

## Install

    go install github.com/typomedia/iconize@latest

## Usage

    iconize [options] [file]

## Options

    -o, --out string   output ico file
    -h, --help         display help
    -V, --version      display version

## Example

    iconize icon.png -o icon.ico

    iconize < icon.png > test.ico

## Test

    icotool -l icon.ico 
    --icon --index=1 --width=256 --height=256 --bit-depth=32 --palette-size=0
    --icon --index=2 --width=128 --height=128 --bit-depth=32 --palette-size=0
    --icon --index=3 --width=64 --height=64 --bit-depth=32 --palette-size=0
    --icon --index=4 --width=48 --height=48 --bit-depth=32 --palette-size=0
    --icon --index=5 --width=32 --height=32 --bit-depth=32 --palette-size=0
    --icon --index=6 --width=24 --height=24 --bit-depth=32 --palette-size=0
    --icon --index=7 --width=16 --height=16 --bit-depth=32 --palette-size=0

---
Copyright © 2023 Typomedia Foundation. All rights reserved.