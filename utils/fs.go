package utils

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
)

type Folder struct {
}

func OpenDir(dir string) fs.FS {
	fsys := os.DirFS(dir)
	return fsys 
}

func GetDirContent(fsys fs.FS) ([]fs.DirEntry, error) {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func ListDirContent(dir string) error {
	fsys := OpenDir(dir)

	entries, err := GetDirContent(fsys)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("[DIR] ", entry.Name())
		} else {
			fmt.Println("[FILE] ", entry.Name())
		}
	}

	return nil
}

func GetDirSize(fsys fs.FS) (Int, error) {
	var totalSize Int

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		//
		if err != nil {
			return err 
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		totalSize += Int(info.Size())
		return nil
	})
	
	return totalSize, err
}

type Entry struct {
	Path string 
	Size float64
}

func GetBiggestFilesSorted(fsys fs.FS) ([]Entry, error) {
	var entries []Entry 

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		//
		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("Access denied:", path)
				return nil 
			}
			return err 
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		entry := Entry{}
		entry.Size = Int(info.Size()).ToMegabytes()
		entry.Path = path 
		entries = append(entries, entry)

		return nil
	})
	
	sort.Sort(Entries(entries))
	return entries, err 
}

func GetSubDirSize(fsys fs.FS, name string) (Int, error) {
	var totalSize Int

	err := fs.WalkDir(fsys, name, func(path string, d fs.DirEntry, err error) error {
		//
		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("Access denied:", path)
				return nil 
			}
			return err 
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		totalSize += Int(info.Size())
		return nil
	})
	
	return totalSize, err
}

// Implementing sort interface. Sorts in descending order
type Entries []Entry 
func (entries Entries) Len() int {
	return len(entries)
}
func (entries Entries) Less(a, b int) bool {
	return entries[a].Size > entries[b].Size
}
func (entries Entries) Swap(a, b int) {
	entries[a], entries[b] = entries[b], entries[a] 
}

type Int int64
func (i Int) ToMegabytes() float64 {
	return float64(i) / (1024 * 1024)
}


func GetBiggestDirSorted(fsys fs.FS) ([]Entry, error) {
	entries, err := GetDirContent(fsys)
	if err != nil {
		return nil, err 
	}

	var dirEntries Entries
	for _, entry := range entries {
		if entry.IsDir() {
			info, _ := entry.Info()
			
			totalSize, err := GetSubDirSize(fsys, info.Name())
			if err != nil {
				return nil, err 
			}
			dirEntries = append(dirEntries, Entry{Size: totalSize.ToMegabytes(), Path: info.Name()})

		}
	}

	sort.Sort(dirEntries)
	return dirEntries, nil 
}