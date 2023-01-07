# Titles command line tool
### Allows to easily organize a bunch of files into their specific folders/seasons.
---
### Moved files are organized accordingly:
```
OUTPUT_DIR/SHOW_NAME/SEASON/EPISODE_FILE
```
### If the tool doesn't detect what `season` episode belongs to it will default to `Season 1`.  
### If the tool can't match Regex for `title`, error will be displayed and file ignored.
---
## Format tool searches for examples
Tool works best with common `eztv` or `nyaa` site torrent file formats.
```
Doom.Patrol.S02E05.1080p.HEVC.x265-MeGusta[eztv.re].mkv -> OUTPUT_DIR/Doom Patrol/Season 2/Doom.Patrol.S02E05.1080p.HEVC.x265-MeGusta[eztv.re].mkv
[ASW] Arknights - Reimei Zensou - 08 [1080p HEVC][EA45AE4C].mkv -> OUTPUT_DIR/Arknights  Reimei Zensou/Season 1/[ASW] Arknights - Reimei Zensou - 08 [1080p HEVC][EA45AE4C].mkv
```
---
## Usage help
```
Usage of titles:
  -b string
        blacklisted directories separated by ','. Example: './dir1,./dir2'
  -d    do a dry run without affecting files
  -e string
        file extension to look for separated by ',' (default ".mkv,.mp4") 
  -i string
        input directory (default ".")
  -o string
        output directory (default "./output")
  -r    recursively search for all files
```
---
## Running with [Task](https://taskfile.dev/)
Just build
```sh
task build
```
Build and run
```sh
task run
```