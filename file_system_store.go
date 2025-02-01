package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// 0 (no movements) SeekStart (Start from the beginning of a file)
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}
