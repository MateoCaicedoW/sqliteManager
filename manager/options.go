package manager

type option func(*manager)

func WithPrefix(prefix string) option {
	return func(f *manager) {
		f.prefix = prefix
	}
}

func WithIconURL(iconURL string) option {
	return func(f *manager) {
		f.iconURL = iconURL
	}
}
