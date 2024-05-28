package filesystem

type option func(*fileSystem)

func WithPrefix(prefix string) option {
	return func(f *fileSystem) {
		f.prefix = prefix
	}
}

func WithIconURL(iconURL string) option {
	return func(f *fileSystem) {
		f.iconURL = iconURL
	}
}
