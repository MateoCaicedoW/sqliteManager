package manager

import "github.com/jmoiron/sqlx"

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

func WithConnection(connection *sqlx.DB) option {
	return func(f *manager) {
		f.connection = connection
	}
}
