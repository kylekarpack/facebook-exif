[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=kylekarpack_fix-fb-meta&metric=bugs)](https://sonarcloud.io/dashboard?id=kylekarpack_fix-fb-meta) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=kylekarpack_fix-fb-meta&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=kylekarpack_fix-fb-meta)

# fix-fb-meta

This is utility for adding back missing metadata from Facebook photo downloads. It works by retrieving metadata from the JSON files Facebook provides and writing it back to the photos from which it was stripped.

At the moment, it primarily adds back the created date for each photo, which is useful for cataloging photos in other applications.

## Usage

1. Download the utility. From the `bin/` folder, download the executable matching your operating system and architecture:
   - [Windows x86](bin/win-x86/fix-fb-meta.exe)
   - [Windows x64](bin/win-x64/fix-fb-meta.exe)
   - [MacOS AMD x64](bin/darwin-amd64/fix-fb-meta)
   - [MacOS ARM x64](bin/darwin-arm64/fix-fb-meta)
   - [Linux AMD x64](bin/linux-amd64/fix-fb-meta)
   - If you need another version, please open an issue or feel free to compile your own!
2. To start, [download your photos](https://www.facebook.com/dyi/) from Facebook.
3. Extract the downloaded archive. Note the `photos_and_videos` directory - you'll need to target this later
4. Run the CLI per instructions below

Example usage:
```
fix-fb-meta path/to/download/photos_and_videos
```

```
Usage:
   fix-fb-meta [dir] {flags}
   fix-fb-meta <command> {flags}

Commands: 
   help                          displays usage informationn
   info                          perform a 'dry run' and return info without modifying files
   version                       displays version number

Arguments: 
   dir                           directory to photos (default: ./photos_and_videos)

Flags: 
   -h, --help                    displays usage information of the application or a command (default: false)
   -v, --version                 displays version number (default: false)
```

## Help

Please feel free to open an issue and I will attempt to provide basic support.

## Development

`fix-fb-meta` is written in Go. To develop, you'll need Go version 1.16 or higher. First, check out this repository:
```
git clone https://github.com/kylekarpack/fix-fb-meta.git
```

You can run with:
``` 
make run
```
or to pass arguments:
```
go run *.go <args>
```

Or build for all platforms with
```
make compile
```
