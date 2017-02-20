package gateway

// FetchRequest is a request to fetch a branch.
type FetchRequest struct {
	Remote    string // name of the remote
	RemoteRef string // ref to fetch
	LocalRef  string // name of the ref locally
}

// PushRequest is a request to push a branch.
type PushRequest struct {
	Remote    string
	LocalRef  string
	RemoteRef string
	Force     bool
}

// TODO: All operations can automatically be scoped to a single remote.

// Git is a gateway to access git locally.
type Git interface {
	// Determines the name of the current branch.
	CurrentBranch() (string, error)

	// Determines if a local branch with the given name exists.
	DoesBranchExist(name string) bool

	// Deletes the given branch.
	DeleteBranch(name string) error

	// Deletes the remote tracking branch with the given name.
	DeleteRemoteTrackingBranch(remote, name string) error

	// Create a branch with the given name and head but don't switch to it.
	CreateBranch(name, head string) error

	// Creates and switches to a local branch with the given name at the given
	// ref.
	CreateBranchAndSwitch(name, head string) error

	// Switches branches.
	Checkout(name string) error

	// Fetch a ref
	Fetch(*FetchRequest) error

	// Push a branch
	Push(*PushRequest) error

	// Get the SHA1 hash for the given ref.
	SHA1(ref string) (string, error)

	// Pulls a branch from a specific remote.
	Pull(remote, name string) error

	// Applies the given patch using git-am.
	ApplyPatches(patches string) error

	// RemoteURL gets the URL for the given remote.
	RemoteURL(name string) (string, error)
}