package chuper

type Queue struct {
	*fetchbot.Queue
}

type Enqueuer interface {
	Enqueue(string, string, string) (bool, error)

	EnqueueWithBasicAuth(string, string, string, string) (bool, error)
}

func (q *Queue) Enqueue(method, URL, sourceURL string) (bool, error) {
	return nil
}

func (q *Queue) EnqueueWithBasicAuth(method, URL, sourceURL, user, password string) (bool, error) {
	return nil
}
