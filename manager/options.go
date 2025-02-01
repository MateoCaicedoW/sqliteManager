package manager

import "database/sql"

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

func WithConnection(connection *sql.DB) option {
	return func(f *manager) {
		f.db = connection
	}
}
