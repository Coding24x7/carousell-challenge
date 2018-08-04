package lib

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Topic represents topic
type Topic struct {
	ID            string    `json:"id"`            // id of the topic
	Content       string    `json:"content"`       // content of the topic
	Author        string    `json:"author"`        // author
	Upvotes       int       `json:"upvotes"`       // upvote count
	UpvoteUsers   []string  `json:"upvoteUsers"`   // upvote users list
	Downvotes     int       `json:"downvotes"`     // downvote count
	DownvoteUsers []string  `json:"downvoteUsers"` // downvote users list
	CreatedAt     time.Time `json:"createdAt"`     // create time
}

func (t *Topic) comapre(t2 *Topic) int {
	if t.Upvotes > t2.Upvotes {
		return 1
	}
	if t.Upvotes < t2.Upvotes {
		return -1
	}
	if t.CreatedAt.Before(t2.CreatedAt) {
		return -1
	}
	if t.CreatedAt.Equal(t2.CreatedAt) {
		return 0
	}
	return 1
}

func (t *Topic) equals(t2 *Topic) bool {
	if t.ID != t2.ID || t.Content != t2.Content || t.Author != t2.Author || t.Upvotes != t2.Upvotes || t.Downvotes != t2.Downvotes || !t.CreatedAt.Equal(t2.CreatedAt) {
		return false
	}
	return true
}

var mutex sync.Mutex
var id int

func nextID() string {
	mutex.Lock()
	defer mutex.Unlock()

	id++
	return strconv.Itoa(id)
}

// ShowTopics shows topics
func ShowTopics() []*Topic {
	return TopicRegistry.ShowTop()
}

// PostTopic posts new topic
func PostTopic(user, content string) (*Topic, error) {
	u := UserRegistry.get(user)
	if u == nil {
		return nil, NewBadReqErr(ErrTopicWrongKey, fmt.Sprintf("user not available with name %s", user))
	}

	t := &Topic{
		ID:            nextID(),
		Author:        user,
		Content:       content,
		Upvotes:       0,
		UpvoteUsers:   make([]string, 0),
		Downvotes:     0,
		DownvoteUsers: make([]string, 0),
		CreatedAt:     time.Now(),
	}
	err := TopicRegistry.Set(t)
	return t, err
}

// UpvoteTopic upvotes topic
func UpvoteTopic(id string, user string) error {
	u := UserRegistry.get(user)
	if u == nil {
		return NewBadReqErr(ErrTopicWrongKey, fmt.Sprintf("user not available with name %s", user))
	}
	return TopicRegistry.Upvote(id, user)
}

// DownvoteTopic downvotes topic
func DownvoteTopic(id string, user string) error {
	u := UserRegistry.get(user)
	if u == nil {
		return NewBadReqErr(ErrTopicWrongKey, fmt.Sprintf("user not available with name %s", user))
	}
	return TopicRegistry.Downvote(id, user)
}
