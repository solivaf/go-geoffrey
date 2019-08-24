package git

type Fetcher interface {
	Fetch() error
}
