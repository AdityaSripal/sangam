package types

// Content
type Content struct {
	Name          string
	Parent        Domain
	ContentHashes [][]byte
}

// GetName returns the name of the content.
func (c Content) GetName() string {
	return c.Name
}

// Path returns the full path to this content entry.
func (c Content) Path() string {
	return c.Path() + c.GetName()
}

// GetParent returns the parent domain of this content.
func (c Content) GetParent() Domain {
	return c.Parent
}

// GetContentHashes returns all the content hashes that have been created for this content.
//
// NOTE: Older hash versions may have been deleted and thus the resulting entry will be
// returned as nil.
func (c Content) GetContentHashes() [][]byte {
	return c.ContentHashes
}

// GetContent returns content hash corresponding to the lastest version of this content.
func (c Content) GetContent() []byte {
	return c.ContentHashes[len(c.ContentHashes)-1]
}

// GetContentAtVersion returns the content hash corresponding to the specified version.
func (c Content) GetContentAtVersion(version uint64) []byte {
	return c.ContentHashes[version]
}
