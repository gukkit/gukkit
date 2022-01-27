package nbt

func (tag *CompoundTag) FindTags() (tags []Tag) {
	tags = make([]Tag, len(tag.Payload))
	copy(tags, tag.Payload)
	return tags
}

func (tag *CompoundTag) TagLength() int {
	return len(tag.Payload)
}

func (tag *CompoundTag) FindTag(idx int) Tag {
	if idx >= tag.TagLength() {
		panic("idx out of bounds in tag array")
	}

	return tag.Payload[idx]
}

func (tag *CompoundTag) AddTag(t Tag) bool {
	tag.Payload = append(tag.Payload, t)
	return true
}

func (tag *CompoundTag) DeleteTag(idx int) bool {
	if idx >= tag.TagLength() {
		panic("idx out of bounds in tag array")
	}

	tag.Payload = append(tag.Payload[:idx], tag.Payload[idx:]...)
	return true
}

func (tag *CompoundTag) UpdateTag(idx int, t Tag) bool {
	if idx >= tag.TagLength() {
		panic("idx out of bounds in tag array")
	}

	tag.Payload[idx] = t
	return true
}
