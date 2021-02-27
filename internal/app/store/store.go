package store

type Storer interface {
	User() UserRepoHandler
}
