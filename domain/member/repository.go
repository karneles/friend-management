package member

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/karneles/friend-management/errorcode"
	"github.com/karneles/friend-management/libs/apierror"
	"github.com/karneles/friend-management/libs/logger"
	//"../../errorcode"
	//"../../libs/apierror"
	//"../../libs/logger"
	uuid "github.com/satori/go.uuid"
)

const (
	querySelectMember = `
		SELECT
			members.entity_id,
			members.email,
			members.name,
			members.password,
			members.created,
			members.updated
		FROM members `
	queryCountMember = `
		SELECT
			COUNT(members.entity_id)
		FROM members `
	queryExistMember = `
		SELECT
			COUNT(members.entity_id) > 0
		FROM members `

	querySelectConnection = `
		SELECT
			connections.member_entity_id,
			connections.friend_entity_id,
			connections.created,
			connections.updated
		FROM connections `
	queryCountConnection = `
		SELECT
			COUNT(connections.friend_entity_id)
		FROM connections `
	queryExistConnection = `
		SELECT
			COUNT(connections.friend_entity_id) > 0
		FROM connections `

	querySelectUpdates = `
		SELECT
			updates.member_entity_id,
			updates.target_entity_id,
			updates.is_blocked,
			updates.created,
			updates.updated
		FROM updates `
	queryCountUpdates = `
		SELECT
			COUNT(updates.target_entity_id)
		FROM updates `
	queryExistUpdates = `
		SELECT
			COUNT(updates.target_entity_id) > 0
		FROM updates `
)

type MemberRepository interface {
	ResolveMemberByID(id uuid.UUID) (Member, error)
	ResolveMemberByEmail(email string) (Member, error)
	ResolveMembersByIDs(ids []uuid.UUID) ([]Member, error)
	ExistMemberByID(id uuid.UUID) (bool, error)
	StoreMember(member Member) error
}

type MemberSQLRepository struct {
	DB *sqlx.DB `inject:""`
}

func (r *MemberSQLRepository) ResolveMemberByID(id uuid.UUID) (member Member, err error) {
	err = r.DB.Get(&member, querySelectMember+`
			WHERE members.entity_id = ?
		`, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			err = apierror.New(errorcode.MemberNotFound)
		} else {
			logger.Err("%v", err)
		}
	}
	return
}

func (r *MemberSQLRepository) ResolveMemberByEmail(email string) (member Member, err error) {
	err = r.DB.Get(&member, querySelectMember+`
			WHERE members.email = ?
		`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = apierror.New(errorcode.MemberNotFound)
		} else {
			logger.Err("%v", err)
		}
	}
	return
}

func (r *MemberSQLRepository) ResolveMembersByIDs(ids []uuid.UUID) (fms []Member, err error) {
	if len(ids) == 0 {
		return
	}
	idStrs := make([]string, 0)
	for _, id := range ids {
		idStrs = append(idStrs, "'" + id.String() + "'")
	}
	idsString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(idStrs), " "), ","), "[]")
	query := fmt.Sprintf(querySelectMember+`
			WHERE members.entity_id IN (%s)
		`, idsString)
	err = r.DB.Select(&fms, query)
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *MemberSQLRepository) ExistMemberByID(id uuid.UUID) (exist bool, err error) {
	err = r.DB.Get(&exist, queryExistMember + ` WHERE members.entity_id = ?
		  `, id.String())
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *MemberSQLRepository) StoreMember(member Member) (err error) {
	exist, err := r.ExistMemberByID(member.ID)
	if err != nil {
		return
	}
	if exist {
		statementUpdate, e := r.DB.Prepare(`
			UPDATE members
			SET
				members.name = ?,
				members.password = ?,
				members.updated = ?
			WHERE members.entity_id = ?`)
		if e != nil {
			err = e
			return
		}
		_, err = statementUpdate.Exec(
			member.Name,
			member.Password,
			member.Updated,
			member.ID)
		if e != nil {
			err = e
			return
		}
	} else {
		statementInsert, e := r.DB.Prepare(`
			INSERT INTO members (
				entity_id,
				email,
				name,
				password,
				created,
				updated)
			VALUES (?, ?, ?, ?, ?, ?)`)
		if e != nil {
			err = e
			return
		}
		_, e = statementInsert.Exec(
			member.ID.String(),
			member.Email,
			member.Name,
			member.Password,
			member.Created,
			member.Updated)
		if e != nil {
			err = e
			return
		}
	}
	return
}


type ConnectionRepository interface {
	ResolveConnectionsByMemberID(id uuid.UUID) ([]Connection, error)
	ExistByMemberIDAndFriendID(memberId uuid.UUID, friendId uuid.UUID) (bool, error)
	StoreConnection(connection Connection) error

	ResolveUpdatesByMemberIDAndTargetID(memberId uuid.UUID, targetId uuid.UUID)(Updates, error)
	ResolveUnblockedUpdatesByMemberId(id uuid.UUID)([]Updates, error)
	ExistByMemberIDAndTargetID(memberId uuid.UUID, targetId uuid.UUID) (bool, error)
	StoreUpdates(updates Updates) error
}

type ConnectionSQLRepository struct {
	DB *sqlx.DB `inject:""`
}

func (r *ConnectionSQLRepository) ResolveConnectionsByMemberID(id uuid.UUID) (connections []Connection, err error) {
	idString := fmt.Sprint(id)
	query := fmt.Sprintf(querySelectConnection+`
			WHERE connections.member_entity_id = '%s'
		`, idString)
	err = r.DB.Select(&connections, query)
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *ConnectionSQLRepository) ExistByMemberIDAndFriendID(memberId uuid.UUID, friendId uuid.UUID) (exist bool, err error) {
	err = r.DB.Get(&exist, queryExistConnection + ` WHERE member_entity_id = ? 
			AND friend_entity_id = ?`, memberId.String(), friendId.String())
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *ConnectionSQLRepository) StoreConnection(connection Connection) (err error) {
	exist, err := r.ExistByMemberIDAndFriendID(connection.MemberID, connection.FriendID)
	if err != nil {
		return
	}
	if exist {
		return
	} else {
		statementInsert, e := r.DB.Prepare(`
			INSERT INTO connections (
				member_entity_id,
				friend_entity_id,
				created,
				updated)
			VALUES (?, ?, ?, ?)`)
		if e != nil {
			err = e
			return
		}
		_, e = statementInsert.Exec(
			connection.MemberID.String(),
			connection.FriendID.String(),
			connection.Created,
			connection.Updated)
		if e != nil {
			err = e
			return
		}
	}
	return
} 

func (r *ConnectionSQLRepository) ResolveUpdatesByMemberIDAndTargetID(memberId uuid.UUID, targetId uuid.UUID) (updates Updates, err error) {
	err = r.DB.Get(&updates, querySelectUpdates+`
		WHERE updates.member_entity_id = ? AND updates.target_entity_id = ?
	`, memberId.String(), targetId.String())
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *ConnectionSQLRepository) ResolveUnblockedUpdatesByMemberID(id uuid.UUID) (updates []Updates, err error) {
	idString := fmt.Sprint(id)
	query := fmt.Sprintf(querySelectUpdates+`
			WHERE updates.member_entity_id = '%s' AND updates.is_blocked = 0
		`, idString)
	err = r.DB.Select(&updates, query)
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *ConnectionSQLRepository) ResolveBlockedUpdatesByMemberID(id uuid.UUID) (updates []Updates, err error) {
	idString := fmt.Sprint(id)
	query := fmt.Sprintf(querySelectUpdates+`
			WHERE updates.member_entity_id = '%s' AND updates.is_blocked = 1
		`, idString)
	err = r.DB.Select(&updates, query)
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *ConnectionSQLRepository) ExistByMemberIDAndTargetID(memberId uuid.UUID, friendId uuid.UUID) (exist bool, err error) {
	err = r.DB.Get(&exist, queryExistUpdates + ` WHERE member_entity_id = ? 
			AND target_entity_id = ?`, memberId.String(), friendId.String())
	if err != nil {
		logger.Err("%v", err)
	}
	return
}

func (r *ConnectionSQLRepository) StoreUpdates(updates Updates) (err error) {
	exist, err := r.ExistByMemberIDAndTargetID(updates.MemberID, updates.TargetID)
	if err != nil {
		return
	}
	if exist {
		statementUpdate, e := r.DB.Prepare(`
			UPDATE updates
			SET
				updates.is_blocked = ?,
				updates.updated = ?
			WHERE updates.member_entity_id = ?
				AND updates.target_entity_id = ?`)
		if e != nil {
			err = e
			return
		}
		_, err = statementUpdate.Exec(
			updates.IsBlocked,
			updates.Updated,
			updates.MemberID,
			updates.TargetID)
		if e != nil {
			err = e
			return
		}
	} else {
		statementInsert, e := r.DB.Prepare(`
			INSERT INTO updates (
				member_entity_id,
				target_entity_id,
				is_blocked,
				created,
				updated)
			VALUES (?, ?, ?, ?, ?)`)
		if e != nil {
			err = e
			return
		}
		_, e = statementInsert.Exec(
			updates.MemberID.String(),
			updates.TargetID.String(),
			updates.IsBlocked,
			updates.Created,
			updates.Updated)
		if e != nil {
			err = e
			return
		}
	}
	return
} 