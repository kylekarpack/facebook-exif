# fix-fb-meta

This is a small utility for adding back missing metadata from Facebook photo downloads. It works by retrieving metadata from the JSON files Facebook provides and writing it back the photos from which it was stripped.

At the moment, it primarily adds back the "Create Date" for each photo, which is useful for cataloging photos in other applications.

## Usage

```bash
go run *.go
```