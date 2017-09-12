package member

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/karneles/friend-management/errorcode"
	"github.com/karneles/friend-management/libs/apierror"
	"github.com/adam-hanna/arrayOperations"
	"regexp"
)

type MemberService struct {
	MemberRepository *MemberSQLRepository `inject:""`
	ConnectionRepository *ConnectionSQLRepository `inject:""`
}

func (s *MemberService) ResolveMemberByID(id uuid.UUID) (Member, error) {
	return s.MemberRepository.ResolveMemberByID(id)
}

func (s *MemberService) ResolveMemberByEmail(email string) (Member, error) {
	return s.MemberRepository.ResolveMemberByEmail(email)
}

func (s *MemberService) ResolveMembersByIDs(ids []uuid.UUID) ([]Member, error) {
	return s.MemberRepository.ResolveMembersByIDs(ids)
}

func (s *MemberService) StoreMember(input MemberInput) (bool, error) {
	var member Member
	if input.ID != uuid.Nil {
		oldfm, err := s.ResolveMemberByID(input.ID)
		if err != nil {
			return false, err
		}
		member = oldfm
		member.Update(input)
	} else {
		member = input.ToMember()
	}
	if err := s.MemberRepository.StoreMember(member); err != nil {
		return false, err
	}
	return true, nil
}

func (s *MemberService) AddFriend(input FriendsInput) (bool, error) {
	var memberIds = make([]uuid.UUID, 0)
	if len(input.Friends) != 2 {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails in request are not two.")
		return false, err
	}
	if input.Friends[0] == input.Friends[1] {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return false, err
	}
	for _, friend := range input.Friends {
		member, err := s.ResolveMemberByEmail(friend)
		if err != nil {
			err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", friend))
			return false, err
		}
		memberIds = append(memberIds, member.ID)
	}
	
	conn1Available, err := s.ConnectionRepository.ExistByMemberIDAndFriendID(memberIds[0], memberIds[1])
	if err != nil {
		return false, err
	}
	if (!conn1Available) {
		connection1 := new(Connection)
		connection1.Create(memberIds[0], memberIds[1])
		if err := s.ConnectionRepository.StoreConnection(*connection1); err != nil {
			return false, err
		}
	}
	
	conn2Available, err := s.ConnectionRepository.ExistByMemberIDAndFriendID(memberIds[1], memberIds[0])
	if err != nil {
		return false, err
	}
	if (!conn2Available) {
		connection2 := new(Connection)
		connection2.Create(memberIds[1], memberIds[0])
		if err := s.ConnectionRepository.StoreConnection(*connection2); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *MemberService) ResolveFriendList(input EmailInput) ([]Member, error) {
	member, err := s.ResolveMemberByEmail(input.Email)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", input.Email))
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

func (s *MemberService) ResolveCommonFriendList(input FriendsInput) ([]Member, error) {
	var memberIds = make([]uuid.UUID, 0)
	
	if len(input.Friends) != 2 {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails in request are not two.")
		return nil, err
	}
	if input.Friends[0] == input.Friends[1] {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return nil, err
	}
	for _, friend := range input.Friends {
		member, err := s.ResolveMemberByEmail(friend)
		if err != nil {
			err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", friend))
			return nil, err
		}
		memberIds = append(memberIds, member.ID)
	}
	connections1, err := s.ConnectionRepository.ResolveConnectionsByMemberID(memberIds[0])
	if err != nil {
		err := apierror.WithMessage(errorcode.FriendNotFound, fmt.Sprintf("Friend not found for %v.", input.Friends[0]))
		return nil, err
	}
	var ids1 = make([]uuid.UUID, 0)
	for _, connection := range connections1 {
		ids1 = append(ids1, connection.FriendID)
	}

	connections2, err := s.ConnectionRepository.ResolveConnectionsByMemberID(memberIds[1])
	if err != nil {
		err := apierror.WithMessage(errorcode.FriendNotFound, fmt.Sprintf("Friend not found for %v.", input.Friends[1]))
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

func (s *MemberService) SubscribeUpdates(input UpdatesInput) (bool, error) {
	if input.Requestor == input.Target {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return false, err
	}

	requestor, err := s.ResolveMemberByEmail(input.Requestor)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", input.Requestor))
		return false, err
	}

	target, err := s.ResolveMemberByEmail(input.Target)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", input.Target))
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

func (s *MemberService) BlockUpdates(input UpdatesInput) (bool, error) {
	if input.Requestor == input.Target {
		err := apierror.WithMessage(errorcode.InvalidRequestData, "Emails cannot be same.")
		return false, err
	}

	requestor, err := s.ResolveMemberByEmail(input.Requestor)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", input.Requestor))
		return false, err
	}

	target, err := s.ResolveMemberByEmail(input.Target)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", input.Target))
		return false, err
	}

	updatesAvailable, err := s.ConnectionRepository.ExistByMemberIDAndTargetID(requestor.ID, target.ID)
	if err != nil {
		return false, err
	}
	if (updatesAvailable) {
		updates, err := s.ConnectionRepository.ResolveUpdatesByMemberIDAndTargetID(requestor.ID, target.ID)
		if err != nil {
			err = apierror.WithMessage(errorcode.InvalidSubscription, fmt.Sprintf("No subscriptioxn for: %v to: %v.", input.Requestor, input.Target))	
			return false, err
		}
		updates.Block(true)
		if err := s.ConnectionRepository.StoreUpdates(updates); err != nil {
			return false, err
		}
	} else {
		err = apierror.WithMessage(errorcode.InvalidSubscription, fmt.Sprintf("No subscription for: %v to: %v.", input.Requestor, input.Target))
		return false, err
	}
	return true, nil
}

func (s *MemberService) ResolveUpdatedMember(input SendUpdateInput) ([]Member, error) {
	sender, err := s.ResolveMemberByEmail(input.Sender)
	if err != nil {
		err = apierror.WithMessage(errorcode.MemberNotFound, fmt.Sprintf("No member with email: %v.", input.Sender))
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
	emails := re.FindAllString(input.Text, -1)
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