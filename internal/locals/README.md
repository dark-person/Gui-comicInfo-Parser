# Locals Package

Package for loading local storage, includes configuration, data for stored tags.

All related files are stored in `{directory}/comicinfo-parser`. Value of `{directory}` will be `$HOME` directory of current user.

## Configuration

Config filename is `config.yaml`. The config name is STRICT and not allow rename.

When no config is found, a default config will be created at
`{directory}/comicinfo-parser/config.yaml`.

## Stored Tags

Stored tags data will be stored by `sqlite`, with filename `storage.db`.

Database connection may not be used frequently, therefore every connection will be closed after used.

There will have serval situation that require connection:

1. When program starts, tags will be loaded to memory
2. When new tag is added as stored tag
