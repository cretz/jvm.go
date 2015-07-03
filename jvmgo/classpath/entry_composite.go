package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry struct {
	Entries []Entry
}

func newCompositeEntry(pathList string) *CompositeEntry {
	compoundEntry := &CompositeEntry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compoundEntry.addEntry(entry)
	}

	return compoundEntry
}

func (self *CompositeEntry) addEntry(entry Entry) {
	self.Entries = append(self.Entries, entry)
}

func (self *CompositeEntry) readClass(className string) (Entry, []byte, error) {
	for _, entry := range self.Entries {
		entry, data, err := entry.readClass(className)
		if err == nil {
			return entry, data, nil
		}
	}

	return self, nil, errors.New("class not found: " + className)
}

func (self *CompositeEntry) String() string {
	strs := make([]string, len(self.Entries))

	for i, entry := range self.Entries {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
