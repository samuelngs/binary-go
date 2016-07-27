package binary

// Tree holds directory information
type Tree struct {
	name   string
	path   string
	parent *Tree
	dirs   []*Tree
	assets []*Asset
}

// NewTree creates directory tree info
func NewTree(name, path string, parent *Tree, dirs []*Tree, assets []*Asset) *Tree {
	return &Tree{name, path, parent, dirs, assets}
}

// Pwd returns the full path of the directory
func (v *Tree) Pwd() string {
	return v.path
}

// Name returns directory name
func (v *Tree) Name() string {
	return v.name
}

// Dirs returns sub-directories
func (v *Tree) Dirs() []string {
	l := make([]string, len(v.assets))
	for i, d := range v.dirs {
		l[i] = d.Pwd()
	}
	return l
}

// Files returns asserts
func (v *Tree) Files() []string {
	l := make([]string, len(v.assets))
	for i, a := range v.assets {
		l[i] = a.Filepath()
	}
	return l
}

// Ls to list files and folders
func (v *Tree) Ls() []string {
	d := v.Dirs()
	f := v.Files()
	l := make([]string, len(d)+len(f))
	offset := len(d)
	for i, p := range d {
		l[i] = p
	}
	for i, p := range f {
		l[offset+i] = p
	}
	return l
}

// Dir to retrieve dir tree
func (v *Tree) Dir(s string) *Tree {
	for _, dir := range v.dirs {
		if dir.Name() == s {
			return dir
		}
	}
	return nil
}

// File to retrieve asset
func (v *Tree) File(s string) *Asset {
	for _, asset := range v.assets {
		if asset.Filename() == s {
			return asset
		}
	}
	return nil
}

// Iter for-loop read assets
func (v *Tree) Iter() <-chan *Asset {
	ch := make(chan *Asset)
	go func() {
		defer close(ch)
		for _, dir := range v.dirs {
			for _, asset := range dir.assets {
				ch <- asset
			}
		}
		for _, asset := range v.assets {
			ch <- asset
		}
	}()
	return ch
}
