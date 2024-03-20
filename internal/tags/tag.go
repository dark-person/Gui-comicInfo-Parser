package tags

// File to control tags in database.

// Data Type for storing Tag data.
type TagData struct {
	Name   string
	Prefix string
}

func (l *LocalTags) SaveTag(tag TagData) {
	panic("SaveTag not implemented")
}

func (l *LocalTags) GetTagNames() []string {
	panic("GetTagNames not implemented")
}

func (l *LocalTags) GetPrefixes() []string {
	panic("GetCategory not implemented")
}

func (l *LocalTags) SavePrefix(prefix, meaning string) {
	panic("SavePrefix not implemented")
}
