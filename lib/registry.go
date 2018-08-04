package lib

import (
	"fmt"
	"sync"
)

/********** Users **********/

// synchronized access to a map of username to User
type userRegistry struct {
	mutex sync.RWMutex
	m     map[string]*User
}

// UserRegistry is storage for user accounts
var UserRegistry = userRegistry{m: make(map[string]*User)}

// Get user account info
func (ur *userRegistry) get(name string) *User {
	ur.mutex.RLock()
	defer ur.mutex.RUnlock()

	return ur.m[name]
}

// Login
func (ur *userRegistry) Login(name, password string) (*User, error) {
	u := ur.get(name)
	if u == nil {
		return nil, NewBadReqErr(ErrUserNotFound, fmt.Sprintf("no user registerd with name %s", name))
	}
	if u.password != password {
		return nil, NewAuthErr(ErrUserWrongCredential, "wrong password")
	}
	return u, nil
}

// Set the user
func (ur *userRegistry) Set(u *User) error {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()
	_, ok := ur.m[u.Name]
	if ok {
		return NewBadReqErr(ErrUserWrongKey, fmt.Sprintf("User already registerd with name %s", u.Name))
	}
	ur.m[u.Name] = u
	return nil
}

// Delete the user
func (ur *userRegistry) Delete(name string) (*User, error) {
	ur.mutex.Lock()
	defer ur.mutex.Unlock()
	u, ok := ur.m[name]
	if !ok {
		return nil, NewBadReqErr(ErrUserNotFound, fmt.Sprintf("no user registerd with name %s", name))
	}
	delete(ur.m, name)
	return u, nil
}

/********** Topics **********/

var maxTopTopics = 20

// synchronized access to a map of topicid to Topic
type topicRegistry struct {
	mutex        sync.RWMutex
	m            map[string]*Topic
	top          []string
	topSize      int
	maxTopTopics int
}

// TopicRegistry is storage for topics
var TopicRegistry = topicRegistry{m: make(map[string]*Topic), top: make([]string, maxTopTopics+1), topSize: 0, maxTopTopics: maxTopTopics}

// Get topic
func (r *topicRegistry) Get(id string) *Topic {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.m[id]
}

// Get topic
func (r *topicRegistry) ShowTop() []*Topic {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	topics := make([]*Topic, r.topSize)

	for i := 0; i < r.topSize; i++ {
		topics[i] = r.m[r.top[i]]
	}
	return topics
}

// Set the topic
func (r *topicRegistry) Set(t *Topic) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	_, ok := r.m[t.ID]
	if ok {
		return NewBadReqErr(ErrTopicWrongKey, fmt.Sprintf("topic already posted with ID %s", t.ID))
	}
	r.m[t.ID] = t
	if r.topSize < r.maxTopTopics {
		r.top[r.topSize] = t.ID
		r.topSize++
	}
	return nil
}

func (r *topicRegistry) findTopPosition(id string) int {
	for i := r.topSize - 1; i >= 0; i-- {
		if (r.top[i]) == id {
			return i
		}
	}
	return -1
}

func (r *topicRegistry) updateTop(id string) {
	lastTopic := r.m[r.top[r.topSize-1]]
	t := r.m[id]
	if t.comapre(lastTopic) < 0 {
		// no need to change top list
		return
	}

	// need to update top list
	pos := r.findTopPosition(id)
	if pos < 0 || pos > r.maxTopTopics {
		pos = r.maxTopTopics
	}

	for i := pos - 1; i >= 0; i-- {
		currTopic := r.m[r.top[i]]
		if t.comapre(currTopic) <= 0 {
			r.top[i+1] = id
			break
		}
		r.top[i+1] = r.top[i]
		if i == 0 {
			r.top[i] = id
		}
	}
}

// Upvote the topic
func (r *topicRegistry) Upvote(id string, username string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	t, ok := r.m[id]
	if !ok {
		return NewBadReqErr(ErrTopicNotFound, fmt.Sprintf("no topic posted with ID %s", id))
	}
	t.Upvotes = t.Upvotes + 1
	t.UpvoteUsers = append(t.UpvoteUsers, username)
	r.m[id] = t
	r.updateTop(id)
	return nil
}

// Downvote the topic
func (r *topicRegistry) Downvote(id string, username string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	t, ok := r.m[id]
	if !ok {
		return NewBadReqErr(ErrTopicNotFound, fmt.Sprintf("no topic posted with ID %s", id))
	}
	t.Downvotes = t.Downvotes + 1
	t.DownvoteUsers = append(t.DownvoteUsers, username)
	r.m[id] = t
	return nil
}
