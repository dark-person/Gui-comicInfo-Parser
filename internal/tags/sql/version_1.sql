-- List for available prefix
CREATE TABLE "list_prefix" (
	"prefix_name" TEXT NOT NULL,
	PRIMARY KEY("prefix_name")
);

-- List for available suffix
CREATE TABLE "list_suffix" (
	"suffix_name" TEXT NOT NULL,
	PRIMARY KEY("suffix_name")
);

-- Mapping for alias, allow one tags_id has multiple alias.
CREATE TABLE "tag_alias" (
	"alias_id" INTEGER,
	"tag_id" INTEGER,
	"alias_name" TEXT NOT NULL,
	PRIMARY KEY("alias_id" AUTOINCREMENT)
);

-- Master table for tags.
CREATE TABLE "tags" (
	"tag_id" INTEGER,
	"prefix" TEXT,
	"name" TEXT,
	"suffix" TEXT,
	PRIMARY KEY("tag_id" AUTOINCREMENT)
);

-- View for alias with original tag
CREATE VIEW view_alias AS
SELECT
	alias_id,
	alias_name,
	tags.prefix,
	tags.name,
	tags.suffix
FROM
	tag_alias
	LEFT JOIN tags ON tags.tag_id = tag_alias.tag_id;

-- SQLite version
PRAGMA user_version = 1