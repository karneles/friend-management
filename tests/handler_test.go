package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "bytes"
    "strconv"
    "time"
)

func TestStoreMemberHandlerSuccess(t *testing.T) {
    email := strconv.FormatInt(time.Now().Unix(), 10) + "@test.com"

    jsonStr := []byte(`{"email":"`+ email + `"}`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member/register", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreMemberHandlerFailedDuplicate(t *testing.T) {
    email := "test@test.com"

    jsonStr := []byte(`{"email":"`+ email + `"}`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member/register", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InternalServerError","data":null,"message":"Error 1062: Duplicate entry '` + email + `' for key 'email_UNIQUE'"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreMemberHandlerFailedWrongJson(t *testing.T) {
    email := "test@test.com"

    jsonStr := []byte(`{"emailx":"`+ email + `"}`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member/register", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'MemberInput.Email' Error:Field validation for 'Email' failed on the 'min' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

/*func TestCreateFriendConnectionHandlerSuccess(t *testing.T) {
    str := `{
        "friends": [
            "test4@test.com",
            "test3@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/add", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}*/

func TestCreateFriendConnectionHandlerFailedSameEmail(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com",
            "test@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/add", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails cannot be same."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestCreateFriendConnectionHandlerFailedLessEmails(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/add", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails in request are not two."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestCreateFriendConnectionHandlerFailedTooManyEmails(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com",
            "test2@test.com",
            "test3@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/add", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails in request are not two."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestCreateFriendConnectionHandlerFailedMemberNotFound1(t *testing.T) {
    str := `{
        "friends": [
            "test@testx.com",
            "test2@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/add", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestCreateFriendConnectionHandlerFailedMemberNotFound2(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com",
            "test2@testx.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/add", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test2@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveFriendsHandlerSuccess(t *testing.T) {
    str := `{
        "email": "test@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/retrieve", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"count":2,"friends":["test3@test.com","test2@test.com"],"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveFriendsHandlerFailedMemberNotFound(t *testing.T) {
    str := `{
        "email": "test@testx.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/retrieve", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandlerSuccessWithCommonFriend(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com",
            "test2@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"count":1,"friends":["test3@test.com"],"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandlerSuccessNoCommonFriend(t *testing.T) {
    str := `{
        "friends": [
            "test3@test.com",
            "test4@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"count":0,"friends":[],"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandlerFailedSameEmail(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com",
            "test@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails cannot be same."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandlerFailedLessEmail(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails in request are not two."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandlerFailedTooManyEmails(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com",
            "test2@test.com",
            "test3@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails in request are not two."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandlerFailedMemberNotFound1(t *testing.T) {
    str := `{
        "friends": [
            "test@testx.com",
            "test2@test.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandlerFailedMemberNotFound2(t *testing.T) {
    str := `{
        "friends": [
            "test@test.com",
            "test2@testx.com"
        ]
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friend/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test2@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

/*func TestSubscribeUpdatesHandlerSuccess(t *testing.T) {
    str := `{
        "requestor": "test3@test.com",
        "target": "test2@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}*/

func TestSubscribeUpdatesHandlerFailedSameEmail(t *testing.T) {
    str := `{
        "requestor": "test3@test.com",
        "target": "test3@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails cannot be same."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSubscribeUpdatesHandlerFailedRequestorEmailNotFound(t *testing.T) {
    str := `{
        "requestor": "test@testx.com",
        "target": "test2@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSubscribeUpdatesHandlerFailedTargetEmailNotFound(t *testing.T) {
    str := `{
        "requestor": "test@test.com",
        "target": "test2@testx.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test2@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestBlockUpdatesHandlerSuccess(t *testing.T) {
    str := `{
        "requestor": "test@test.com",
        "target": "test2@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestBlockUpdatesHandlerFailedRequestorSameEmail(t *testing.T) {
    str := `{
        "requestor": "test@test.com",
        "target": "test@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Emails cannot be same."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestBlockUpdatesHandlerFailedRequestorEmailNotFound(t *testing.T) {
    str := `{
        "requestor": "test@testx.com",
        "target": "test2@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestBlockUpdatesHandlerFailedTargetEmailNotFound(t *testing.T) {
    str := `{
        "requestor": "test@test.com",
        "target": "test2@testx.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test2@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestBlockUpdatesHandlerFailedInvalidSubscription(t *testing.T) {
    str := `{
        "requestor": "test@test.com",
        "target": "1504858412@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates/subscribe", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidSubscription","data":null,"message":"No subscription for: test@test.com to: 1504858412@test.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveUpdatedMemberHandlerSuccess(t *testing.T) {
    str := `{
        "sender": "test@test.com",
        "text": "Hello World! 1504848613@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/updates/send", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveUpdatedMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"recipients":["1504848613@test.com","test3@test.com"],"success":true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveUpdatedMemberHandlerFailedMemberNotFound1(t *testing.T) {
    str := `{
        "sender": "test@testx.com",
        "text": "Hello World! 1504848613@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/updates/send", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveUpdatedMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: test@testx.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveUpdatedMemberHandlerFailedMemberNotFound2(t *testing.T) {
    str := `{
        "sender": "test@test.com",
        "text": "Hello World! kate@example.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/updates/send", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveUpdatedMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with email: kate@example.com."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}