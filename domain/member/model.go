package member

import (
	"time"

	"github.com/karneles/friend-management/libs/common"
	//"../../libs/common"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

type Member struct {
	ID        	uuid.UUID   `json:"id" db:"entity_id"`
	Email     	string      `json:"email" db:"email"`
	Name	  	string	  	`json:"name" db:"name"`
	Password	string		`json:"password" db:"password"`
	Created   	time.Time   `json:"created" db:"created"`
	Updated   	null.Time   `json:"updated" db:"updated"`
}

type NewMemberInput struct {
	ID   		uuid.UUID 	`json:"id"`
	Email 		string   	`json:"email" validate:"required,min=3,max=45,email"`
	Name  		string   	`json:"name" validate:"required,max=45"`
	Password	string		`json:"password" validate:"required,min=8"`
	Password2	string		`json:"password2" validate:"required,eqfield=Password"`
}

type UpdateMemberInput struct {
	ID   		uuid.UUID 	`json:"id"`
	Name  		string   	`json:"name" validate:"required,max=45"`
	Password	string		`json:"password" validate:"omitempty,required"`
	Password2	string		`json:"password2" validate:"omitempty,required,eqfield=Password"`
}

func (member *Member) Update(input UpdateMemberInput) {
	member.Name = input.Name
	if len(input.Password) > 0 {
		member.Password = common.Hash(input.Password)
	}
	member.Updated = null.TimeFrom(time.Now())
}

func (input *NewMemberInput) ToMember() Member {
	return Member{
		ID:        	uuid.NewV4(),
		Email:     	input.Email,
		Name:		input.Name,
		Password:	common.Hash(input.Password),
		Created:   	time.Now(),
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
	Email 	string 	`json:"email" validate:"min=3,max=45,email"`
}

type UpdatesInput struct {
	Requestor 	string `json:"requestor" validate:"min=3,max=45,email"`
	Target		string `json:"target" validate:"min=3,max=45,email"`
}

type SendUpdateInput struct {
	Sender		string `json:"sender" validate:"min=3,max=45,email"`
	Text		string `json:"text"`
}

type TextInput struct {
	Text		string `json:"text"`
}