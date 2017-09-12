package member

import (
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

type Member struct {
	ID        uuid.UUID   `json:"id" db:"entity_id"`
	Email      string      `json:"email" db:"email"`
	Created   time.Time   `json:"created" db:"created"`
	Updated   null.Time   `json:"updated" db:"updated"`
}

type MemberInput struct {
	ID   uuid.UUID `json:"id"`
	Email string    `json:"email" validate:"min=3,max=100"`
}

func (member *Member) Update(input MemberInput) {
	member.Email = input.Email
	member.Updated = null.TimeFrom(time.Now())
}

func (input *MemberInput) ToMember() Member {
	return Member{
		ID:        uuid.NewV4(),
		Email:      input.Email,
		Created:   time.Now(),
	}
}

type Connection struct {
	MemberID 	uuid.UUID 	`json:"memberId" db:"member_entity_id"`
	FriendID 	uuid.UUID 	`json:"friendId" db:"friend_entity_id"`
	Created   	time.Time   `json:"created" db:"created"`
	Updated   	null.Time   `json:"updated" db:"updated"`
}

func (connection *Connection) Create(memberId uuid.UUID, friendId uuid.UUID) {
	connection.MemberID = memberId
	connection.FriendID = friendId
	connection.Created = time.Now()
}

type Updates struct {
	MemberID 	uuid.UUID 	`json:"memberId" db:"member_entity_id"`
	TargetID 	uuid.UUID 	`json:"targetId" db:"target_entity_id"`
	IsBlocked 	bool 		`json:"isBlocked" db:"is_blocked"`
	Created   	time.Time   `json:"created" db:"created"`
	Updated   	null.Time   `json:"updated" db:"updated"`
}

func (updates *Updates) Create(memberId uuid.UUID, targetId uuid.UUID) {
	updates.MemberID = memberId
	updates.TargetID = targetId
	updates.IsBlocked = false
	updates.Created = time.Now()
}

func (updates *Updates) Block(block bool) {
	updates.IsBlocked = block
	updates.Updated = null.TimeFrom(time.Now())
}

type FriendsInput struct {
	Friends []string `json:"friends"`
}

type EmailInput struct {
	Email 	string 	`json:"email" validate:"min=3,max=100"`
}

type UpdatesInput struct {
	Requestor 	string `json:"requestor"`
	Target		string `json:"target"`
}

type SendUpdateInput struct {
	Sender		string `json:"sender"`
	Text		string `json:"text"`
}

/*type CommonOutput struct {
	Success 	bool 	`json:"success"`
}

type FriendsOutput struct {
	Success 	bool 		`json:"success"`
	Friends 	[]string	`json:"friends"`
	Count		int			`json:"count"`
}

type UpdatesOutput struct {
	Success		bool		`json:"success"`
	Recipients	[]string	`json:"recipients"`
}*/