/**************
**
** Note: This methods is created as my initiative to propose a better approach 
**       to create the API. In these methods, the user id is added in the request header.
**		 Using this approach, we can create a cleaner and simpler request message and 
**       add better security.
**
**************/

package member

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/karneles/friend-management/errorcode"
	"github.com/karneles/friend-management/libs/apierror"
	//"../../errorcode"
	//"../../libs/apierror"
	"github.com/adam-hanna/arrayOperations"
	"regexp"
)

func (s *MemberService) AddFriend2(userId uuid.UUID, friendEmail string) (bool, error) {
	member, err := s.ResolveMemberByID(userId)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with ID: %v.", userId))
		return false, err
	}
	friend, err := s.ResolveMemberByEmail(friendEmail)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", friendEmail))
		return false, err
	}
	
	if member.Email == friendEmail {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return false, err
	}

	conn1Available, err := s.ConnectionRepository.ExistByMemberIDAndFriendID(userId, friend.ID)
	if err != nil {
		return false, err
	}
	if (!conn1Available) {
		connection1 := new(Connection)
		connection1.Create(userId, friend.ID)
		if err := s.ConnectionRepository.StoreConnection(*connection1); err != nil {
			return false, err
		}
	}
	
	conn2Available, err := s.ConnectionRepository.ExistByMemberIDAndFriendID(friend.ID, userId)
	if err != nil {
		return false, err
	}
	if (!conn2Available) {
		connection2 := new(Connection)
		connection2.Create(friend.ID, userId)
		if err := s.ConnectionRepository.StoreConnection(*connection2); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *MemberService) ResolveFriendList2(userId uuid.UUID) ([]Member, error) {
	member, err := s.ResolveMemberByID(userId)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with Id: %v.", userId))
		return nil, err
	}
	connections, err := s.ConnectionRepository.ResolveConnectionsByMemberID(member.ID)
	if err != nil {
		return nil, err
	}
	var ids = make([]uuid.UUID, 0)
	for _, connection := range connections {
		ids = append(ids, connection.FriendID)
	}
	members, err := s.ResolveMembersByIDs(ids)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	return members, nil
}

func (s *MemberService) ResolveCommonFriendList2(userId uuid.UUID, friendEmail string) ([]Member, error) {
	member, err := s.ResolveMemberByID(userId)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with ID: %v.", userId))
		return nil, err
	}
	friend, err := s.ResolveMemberByEmail(friendEmail)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", friendEmail))
		return nil, err
	}
	
	if member.Email == friendEmail {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return nil, err
	}

	connections1, err := s.ConnectionRepository.ResolveConnectionsByMemberID(userId)
	if err != nil {
		err := apierror.WithMessage(errorcode.FriendNotFound, fmt.Sprintf("Friend not found for %v.", member.Email))
		return nil, err
	}
	var ids1 = make([]uuid.UUID, 0)
	for _, connection := range connections1 {
		ids1 = append(ids1, connection.FriendID)
	}

	connections2, err := s.ConnectionRepository.ResolveConnectionsByMemberID(friend.ID)
	if err != nil {
		err := apierror.WithMessage(errorcode.FriendNotFound, fmt.Sprintf("Friend not found for %v.", friend.Email))
		return nil, err
	}
	var ids2 = make([]uuid.UUID, 0)
	for _, connection := range connections2 {
		ids2 = append(ids2, connection.FriendID)
	}

	var commonIds = make([]uuid.UUID, 0)
	if len(ids1) > 0 && len(ids2) > 0 {
		commonList, _ := arrayOperations.Intersect(ids1, ids2)
		commonIds, _ = commonList.Interface().([]uuid.UUID)
	}

	var members = make([]Member, 0)
	if (len(commonIds) > 0) {
		members, err = s.ResolveMembersByIDs(commonIds)
		if err != nil {
			fmt.Printf("%v", err)
			return nil, err
		}
	}
	return members, nil
}

func (s *MemberService) SubscribeUpdates2(userId uuid.UUID, friendEmail string) (bool, error) {
	requestor, err := s.ResolveMemberByID(userId)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with ID: %v.", userId))
		return false, err
	}

	if requestor.Email == friendEmail {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return false, err
	}

	target, err := s.ResolveMemberByEmail(friendEmail)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", friendEmail))
		return false, err
	}

	updatesAvailable, err := s.ConnectionRepository.ExistByMemberIDAndTargetID(requestor.ID, target.ID)
	if err != nil {
		return false, err
	}
	if (!updatesAvailable) {
		updates := new(Updates)
		updates.Create(requestor.ID, target.ID)
		if err := s.ConnectionRepository.StoreUpdates(*updates); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (s *MemberService) BlockUpdates2(userId uuid.UUID, friendEmail string) (bool, error) {
	requestor, err := s.ResolveMemberByID(userId)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with ID: %v.", userId))
		return false, err
	}

	if requestor.Email == friendEmail {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return false, err
	}

	target, err := s.ResolveMemberByEmail(friendEmail)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", friendEmail))
		return false, err
	}

	updatesAvailable, err := s.ConnectionRepository.ExistByMemberIDAndTargetID(requestor.ID, target.ID)
	if err != nil {
		return false, err
	}
	if (updatesAvailable) {
		updates, err := s.ConnectionRepository.ResolveUpdatesByMemberIDAndTargetID(requestor.ID, target.ID)
		if err != nil {
			err = apierror.WithMessage(errorcode.InvalidSubscription, fmt.Sprintf("No subscription for: %v to: %v.", requestor.Email, friendEmail))	
			return false, err
		}
		updates.Block(true)
		if err := s.ConnectionRepository.StoreUpdates(updates); err != nil {
			return false, err
		}
	} else {
		err = apierror.WithMessage(errorcode.InvalidSubscription, fmt.Sprintf("No subscription for: %v to: %v.", requestor.Email, friendEmail))
		return false, err
	}
	return true, nil
}

func (s *MemberService) ResolveUpdatedMember2(userId uuid.UUID, text string) ([]Member, error) {
	sender, err := s.ResolveMemberByID(userId)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with ID: %v.", userId))
		return nil, err
	}

	var friendIds = make([]uuid.UUID, 0)
	friendList, err := s.ConnectionRepository.ResolveConnectionsByMemberID(sender.ID)
	for _, friend := range friendList {
		friendIds = append(friendIds, friend.FriendID)
	}
	
	var updateIds = make([]uuid.UUID, 0)
	updateList, err := s.ConnectionRepository.ResolveUnblockedUpdatesByMemberID(sender.ID)
	for _, target := range updateList {
		updateIds = append(updateIds, target.TargetID)
	}
	
	var blockedIds = make([]uuid.UUID, 0)
	blockedList, err := s.ConnectionRepository.ResolveBlockedUpdatesByMemberID(sender.ID)
	for _, target := range blockedList {
		blockedIds = append(blockedIds, target.TargetID)
	}

	var combineIds = make([]uuid.UUID, 0)
	if len(friendIds) > 0 && len(updateIds) == 0 {
		combineIds = friendIds
	} else if len(friendIds) == 0 && len(updateIds) > 0 {
		combineIds = updateIds
	} else if len(friendIds) > 0 && len(updateIds) > 0 {
		combineList, _ := arrayOperations.Union(friendIds, updateIds)
		combineIds, _ = combineList.Interface().([]uuid.UUID)
	}

	var receiverIds = make([]uuid.UUID, 0)
	if len(combineIds) > 0 && len(blockedIds) == 0 {
		receiverIds = combineIds
	} else if len(combineIds) > 0 && len(blockedIds) > 0 {
		receiverList, _ := arrayOperations.Difference(combineIds, blockedIds)
		receiverIds, _ = receiverList.Interface().([]uuid.UUID)
	}

	re := regexp.MustCompile("[a-zA-Z0-9-_.]+@[a-zA-Z0-9-_.]+")
	emails := re.FindAllString(text, -1)
	for _, email := range emails {
		member, err := s.MemberRepository.ResolveMemberByEmail(email)
		if err != nil {
			err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", email))
			return nil, err			
		}
		receiverIds = append(receiverIds, member.ID)
	}

	members, err := s.ResolveMembersByIDs(receiverIds)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	return members, nil 
}