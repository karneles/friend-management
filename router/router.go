package router

import (
	"github.com/gorilla/mux"

	"./../handler"
)

func CreateRouter(rh handler.RootHandler) *mux.Router {
	router := mux.NewRouter()

	// Member
	router.HandleFunc("/member", rh.StoreMemberHandler).Methods("POST")

	// Connection
	router.HandleFunc("/friend/add", rh.CreateFriendConnectionHandler).Methods("POST")
	router.HandleFunc("/friend/common", rh.ResolveCommonFriendsHandler).Methods("POST")
	router.HandleFunc("/friend/retrieve", rh.ResolveFriendsHandler).Methods("POST")

	// Updates
	router.HandleFunc("/update/subscribe", rh.SubscribeUpdatesHandler).Methods("POST")
	router.HandleFunc("/update/subscribe", rh.BlockUpdatesHandler).Methods("DELETE")
	router.HandleFunc("/update/send", rh.ResolveUpdatedMemberHandler).Methods("POST")

	return router
}
