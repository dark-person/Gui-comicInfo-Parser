# Tags Package

This package store source code to manipulate tags.

Stored tags data will be stored by `sqlite`, with filename `storage.db`.

Database connection may not be used frequently, therefore every connection will be closed after used.

There will have serval situation that require connection:

1. When program starts, tags will be loaded to memory
2. When new tag is added as stored tag

## Tag Structure

A sample of tag will be like:

```
{prefix}:{tag name}({suffix})
```

## Basic Schema

The database will contains these table:

| Table       | Purpose                                    |
| ----------- | ------------------------------------------ |
| `tags`      | Master for tags, contains tag relationship |
| `prefix`    | Available prefix for tags                  |
| `suffix`    | Available suffix for tags                  |
| `tag_alias` | Alias for existing tag                     |

All table will rely on `tag_id` as primary key.

In Master table `tags`, `tag_id` should be auto incremented.
