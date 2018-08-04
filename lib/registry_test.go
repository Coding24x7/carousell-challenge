package lib

import (
	"testing"
	"time"
)

func newTopic(id string) *Topic {
	return &Topic{
		ID:            id,
		Author:        "user" + id,
		Content:       "this is test content" + id,
		Upvotes:       0,
		UpvoteUsers:   make([]string, 0),
		Downvotes:     0,
		DownvoteUsers: make([]string, 0),
		CreatedAt:     time.Now(),
	}
}

func newUser(username, password string) *User {
	return &User{
		Name:     username,
		password: password,
		Country:  "Country" + username,
	}
}

func isEqual(arr1, arr2 []*Topic) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if !arr1[i].equals(arr2[i]) {
			return false
		}
	}
	return true
}

func TestTopicSetGet(t *testing.T) {
	testMaxTopTopics := 5
	testTopicRegistry := &topicRegistry{m: make(map[string]*Topic), top: make([]string, testMaxTopTopics+1), topSize: 0, maxTopTopics: testMaxTopTopics}
	t1 := newTopic("1")
	t2 := newTopic("2")
	t3 := newTopic("3")
	testTopicRegistry.Set(t1)
	testTopicRegistry.Set(t2)
	testTopicRegistry.Set(t3)

	result := testTopicRegistry.Get("1")
	if result == nil || testTopicRegistry.topSize != 3 || result.ID != t1.ID || result.Author != t1.Author || result.Content != t1.Content {
		t.Fatalf("topic set-get failed... Expected %v but got %v", t1, result)
	}
}

func TestTopicUpvote(t *testing.T) {
	testMaxTopTopics := 5
	testTopicRegistry := &topicRegistry{m: make(map[string]*Topic), top: make([]string, testMaxTopTopics+1), topSize: 0, maxTopTopics: testMaxTopTopics}
	t1 := newTopic("1")
	testTopicRegistry.Set(t1)

	testTopicRegistry.Upvote("1", "user2")
	testTopicRegistry.Upvote("1", "user2")
	testTopicRegistry.Upvote("1", "user3")
	result := testTopicRegistry.Get("1")
	if result == nil || result.Upvotes != 3 {
		t.Fatalf("topic upvote failed... got %v", result)
	}
}

func TestTopicDownvote(t *testing.T) {
	testMaxTopTopics := 5
	testTopicRegistry := &topicRegistry{m: make(map[string]*Topic), top: make([]string, testMaxTopTopics+1), topSize: 0, maxTopTopics: testMaxTopTopics}
	t1 := newTopic("1")
	testTopicRegistry.Set(t1)

	testTopicRegistry.Downvote("1", "user2")
	testTopicRegistry.Downvote("1", "user2")
	testTopicRegistry.Downvote("1", "user3")
	result := testTopicRegistry.Get("1")
	if result == nil || result.Downvotes != 3 {
		t.Fatalf("topic downvote failed... got %v", result)
	}
}

func TestTopicShowTop(t *testing.T) {
	testMaxTopTopics := 3
	testTopicRegistry := &topicRegistry{m: make(map[string]*Topic), top: make([]string, testMaxTopTopics+1), topSize: 0, maxTopTopics: testMaxTopTopics}
	t1 := newTopic("1")
	t1Copy := newTopic("1")
	t1Copy.CreatedAt = t1.CreatedAt
	t2 := newTopic("2")
	t2Copy := newTopic("2")
	t2Copy.CreatedAt = t2.CreatedAt
	t3 := newTopic("3")
	t3Copy := newTopic("3")
	t3Copy.CreatedAt = t3.CreatedAt
	t4 := newTopic("4")
	t4Copy := newTopic("4")
	t4Copy.CreatedAt = t4.CreatedAt
	testTopicRegistry.Set(t1)
	testTopicRegistry.Set(t2)

	expected := []*Topic{t1Copy, t2Copy}
	result := testTopicRegistry.ShowTop()
	if result == nil || !isEqual(result, expected) {
		t.Fatalf("topic showtop (less than max) failed... Expected %v but got %v", expected, result)
	}

	testTopicRegistry.Upvote("2", "user99")
	testTopicRegistry.Upvote("1", "user99")
	testTopicRegistry.Upvote("2", "user99")
	t1Copy.Upvotes++
	t2Copy.Upvotes += 2
	expected = []*Topic{t2Copy, t1Copy}
	result = testTopicRegistry.ShowTop()
	if result == nil || !isEqual(result, expected) {
		t.Fatalf("topic showtop (less than max with upvotes) failed... Expected %v but got %v", expected, result)
	}

	testTopicRegistry.Set(t3)
	testTopicRegistry.Set(t4)
	expected = []*Topic{t2Copy, t1Copy, t3Copy}
	result = testTopicRegistry.ShowTop()
	if result == nil || !isEqual(result, expected) {
		t.Fatalf("topic showtop (more than max) failed... Expected %v but got %v", expected, result)
	}

	testTopicRegistry.Upvote("4", "user99")
	testTopicRegistry.Upvote("4", "user99")
	t4Copy.Upvotes += 2
	expected = []*Topic{t4Copy, t2Copy, t1Copy}
	result = testTopicRegistry.ShowTop()
	if result == nil || !isEqual(result, expected) {
		t.Fatalf("topic showtop (more than max with upvotes) failed... Expected %v but got %v", expected, result)
	}

}

func TestUserSetLogin(t *testing.T) {
	u := newUser("user1", "pwd1")
	UserRegistry.Set(u)
	u1, err := UserRegistry.Login("user1", "wrongpwd")
	if err == nil {
		t.Fatal("user login with wrong password should be failed...")
	}

	u1, err = UserRegistry.Login("user1", "pwd1")
	if err != nil || u1 == nil || u.Name != u1.Name || u.Country != u1.Country {
		t.Fatalf("user login failed... expected %v got %v", u, u1)
	}
}

func TestUserDelete(t *testing.T) {
	u := newUser("user1", "pwd1")
	UserRegistry.Set(u)
	u1 := UserRegistry.get("user1")
	if u1 == nil {
		t.Fatal("user set failed...")
	}
	UserRegistry.Delete("user1")
	u1 = UserRegistry.get("user1")
	if u1 != nil {
		t.Fatal("user delete failed...")
	}
}
