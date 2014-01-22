package storeadapter

type StoreAdapter interface {
	// Intiailize connection to server. For a store with no
	// persistent connection, this effectively just tests connectivity.
	Connect() error

	// Create a node and fail if it already exists.
	Create(StoreNode) error

	// Set multiple nodes at once. If any of them fail,
	// it will return the first error.
	SetMulti(nodes []StoreNode) error

	// Retrieve a node from the store at the given key.
	// Returns an error if it does not exist.
	Get(key string) (StoreNode, error)

	// Recursively get the contents of a key.
	ListRecursively(key string) (StoreNode, error)

	// Delete a set of keys from the store. If any fail to be
	// deleted or don't actually exist, an error is returned.
	Delete(keys ...string) error

	// Watch a given key recursively for changes. Events will come in on one channel, and watching will stop when a value is sent over the stop channel.
	//
	// Events may be missed, but the watcher will do its best to continue.
	//
	// Returns an error if the watcher cannot initially "attach" to the stream.
	//
	// Otherwise, the caller can assume that the watcher will continue attempting to stream events.
	Watch(key string) (events <-chan WatchEvent, stop chan<- bool, errors <-chan error)

	// Close any live persistent connection, and cleans up any running state.
	Disconnect() error

	// Grab a lock and send a notification when it is lost. Blocks until the lock can be acquired.
	//
	// To release the lock, send any value to the releaseLock channel.
	//
	// If the store times out, returns an error.
	GetAndMaintainLock(lockName string, lockTTL uint64) (lostLock <-chan bool, releaseLock chan<- bool, err error)
}
