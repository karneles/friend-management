package router

import (
	"github.com/gorilla/mux"

	"github.com/karneles/friend-management/handler"
	//"../handler"
)

func CreateRouter(rh handler.RootHandler) *mux.Router {
	router := mux.NewRouter()

	// Member
	router.HandleFunc("/member", rh.StoreNewMemberHandler).Methods("POST")
	router.HandleFunc("/member", rh.StoreUpdateMemberHandler).Methods("PUT")

	// Connection
	router.HandleFunc("/friend/add", rh.CreateFriendConnectionHandler).Methods("POST")
	router.HandleFunc("/friend/common", rh.ResolveCommonFriendsHandler).Methods("POST")
	router.HandleFunc("/friend/retrieve", rh.ResolveFriendsHandler).Methods("POST")

	// Updates
	router.HandleFunc("/update/subscribe", rh.SubscribeUpdatesHandler).Methods("POST")
	router.HandleFunc("/update/subscribe", rh.BlockUpdatesHandler).Methods("DELETE")
	router.HandleFunc("/update/send", rh.ResolveUpdatedMemberHandler).Methods("POST")

/**************
**
** Note: This methods is created as my initiative to propose a better approach 
**       to create the API. In these methods, the user id is added in the request header.
**		 Using this approach, we can create a cleaner and simpler request message and 
**       add better security.
**
**************/

	// Login
	router.HandleFunc("/login", rh.LoginHandler).Methods("POST")

	// Connections
	router.HandleFunc("/friends", rh.CreateFriendConnectionHandler2).Methods("POST")
	router.HandleFunc("/friends/common", rh.ResolveCommonFriendsHandler2).Methods("POST")
	router.HandleFunc("/friends", rh.ResolveFriendsHandler2).Methods("GET")

	// Updates
	router.HandleFunc("/updates", rh.SubscribeUpdatesHandler2).Methods("POST")
	router.HandleFunc("/updates", rh.BlockUpdatesHandler2).Methods("DELETE")
	router.HandleFunc("/updates/send", rh.ResolveUpdatedMemberHandler2).Methods("POST")
	
	return router
}
