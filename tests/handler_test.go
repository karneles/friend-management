package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "bytes"
    "strconv"
    "time"
    uuid "github.com/satori/go.uuid"
    "fmt"
    "encoding/json"
)

// Member management

func TestStoreNewMemberHandlerSuccess(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerSuccess")
    name := strconv.FormatInt(time.Now().Unix(), 10)
    password := name
    email := name + "@test.com"

    jsonStr := []byte(`
        {
            "email": "`+ email + `", 
            "name": "` + name + `",
            "password": "` + password + `",
            "password2": "` + password + `"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
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

func TestStoreNewMemberHandlerFailedNoName(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerFailedNoName")
    name := strconv.FormatInt(time.Now().Unix(), 10)
    email := name + "@test.com"
    password := name

    jsonStr := []byte(`
        {
            "email" : "` + email + `",
            "password" : "` + password + `",
            "password2" : "` + password + `"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'NewMemberInput.Name' Error:Field validation for 'Name' failed on the 'required' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreNewMemberHandlerFailedNoPassword(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerFailedNoPassword")
    name := strconv.FormatInt(time.Now().Unix(), 10)
    email := name + "@test.com"
    
    jsonStr := []byte(`
        {
            "email" : "` + email + `",
            "name" : "` + name + `"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'NewMemberInput.Password' Error:Field validation for 'Password' failed on the 'required' tag\nKey: 'NewMemberInput.Password2' Error:Field validation for 'Password2' failed on the 'required' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreNewMemberHandlerFailedPasswordNoPasswordConfirmation(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerFailedPasswordNoPasswordConfirmation")
    name := strconv.FormatInt(time.Now().Unix(), 10)
    email := name + "@test.com"
    password := name
    
    jsonStr := []byte(`
        {
            "email" : "` + email + `",
            "name" : "` + name + `",
            "password" : "` + password + `"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'NewMemberInput.Password2' Error:Field validation for 'Password2' failed on the 'required' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreNewMemberHandlerFailedPasswordConfirmationNotEqual(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerFailedPasswordConfirmationNotEqual")
    name := strconv.FormatInt(time.Now().Unix(), 10)
    email := name + "@test.com"
    password := name
    
    jsonStr := []byte(`
        {
            "email" : "` + email + `",
            "name" : "` + name + `",
            "password" : "` + password + `",
            "password2" : "` + password + `x"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'NewMemberInput.Password2' Error:Field validation for 'Password2' failed on the 'eqfield' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreNewMemberHandlerFailedPasswordTooShort(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerFailedPasswordTooShort")
    name := strconv.FormatInt(time.Now().Unix(), 10)
    email := name + "@test.com"
    password := "name"
    
    jsonStr := []byte(`
        {
            "email" : "` + email + `",
            "name" : "` + name + `",
            "password" : "` + password + `",
            "password2" : "` + password + `"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'NewMemberInput.Password' Error:Field validation for 'Password' failed on the 'min' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreNewMemberHandlerFailedDuplicateEmail(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerFailedDuplicateEmail")
    email := "test@test.com"
    name := "test12345"
    password := name

    jsonStr := []byte(`
        {
            "email":"`+ email + `", 
            "name": "` + name + `",
            "password": "` + password + `",
            "password2": "` + password + `"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InternalServerError","data":null,"message":"Error 1062: Duplicate entry '` + email + `' for key 'email_UNIQUE'"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreNewMemberHandlerFailedWrongJson(t *testing.T) {
    fmt.Println("Running TestStoreNewMemberHandlerFailedWrongJson")
    email := "test@test.com"
    name := "test12345"
    password := name

    jsonStr := []byte(`
        {
            "emailx":"`+ email + `", 
            "name": "` + name + `",
            "password": "` + password + `",
            "password2": "` + password + `"
        }`)
    req := httptest.NewRequest("POST", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreNewMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'NewMemberInput.Email' Error:Field validation for 'Email' failed on the 'required' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreUpdateMemberHandlerSuccess(t *testing.T) {
    fmt.Println("Running TestStoreUpdateMemberHandlerSuccess")
    id := "873e2168-17b9-4302-9a32-4c8a9fe7da35"
    name := "test2"
    password := name

    jsonStr := []byte(`
        {
            "id": "` + id + `", 
            "password": "` + password + `",
            "password2": "` + password + `",
            "name": "` + name + `"
        }`)
    req := httptest.NewRequest("PUT", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreUpdateMemberHandler)
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

func TestStoreUpdateMemberHandlerSuccessNoPassword(t *testing.T) {
    fmt.Println("Running TestStoreUpdateMemberHandlerSuccessNoPassword")
    id := "873e2168-17b9-4302-9a32-4c8a9fe7da35"
    name := "test2"

    jsonStr := []byte(`
        {
            "id": "` + id + `", 
            "name": "` + name + `"
        }`)
    req := httptest.NewRequest("PUT", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreUpdateMemberHandler)
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

func TestStoreUpdateMemberHandlerFailedNoName(t *testing.T) {
    fmt.Println("Running TestStoreUpdateMemberHandlerFailedNoName")
    id := "873e2168-17b9-4302-9a32-4c8a9fe7da35"
    password := "test2"

    jsonStr := []byte(`
        {
            "id": "` + id + `", 
            "password": "` + password + `",
            "password2": "` + password + `"
        }`)
    req := httptest.NewRequest("PUT", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreUpdateMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'UpdateMemberInput.Name' Error:Field validation for 'Name' failed on the 'required' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestStoreUpdateMemberHandlerFailedMemberNotFound(t *testing.T) {
    fmt.Println("Running TestStoreUpdateMemberHandlerFailedMemberNotFound")
    id := uuid.NewV4().String()
    name := strconv.FormatInt(time.Now().Unix(), 10)
    email := name + "@test.com"

    jsonStr := []byte(`
        {
            "id": "` + id + `", 
            "email":"`+ email + `", 
            "name": "` + name + `"
        }`)
    req := httptest.NewRequest("PUT", "http://localhost:9111/member", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.StoreUpdateMemberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with ID: ` + id + `."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

// Create Connections

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

// Resolve Friends

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

// Resolve Common Friends

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

// Subscribe Updates

/*func TestSubscribeUpdatesHandlerSuccess(t *testing.T) {
    str := `{
        "requestor": "test3@test.com",
        "target": "test2@test.com"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/update/subscribe", bytes.NewBuffer(jsonStr))
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
    req := httptest.NewRequest("POST", "http://localhost:9111/update/subscribe", bytes.NewBuffer(jsonStr))
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
    req := httptest.NewRequest("POST", "http://localhost:9111/update/subscribe", bytes.NewBuffer(jsonStr))
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
    req := httptest.NewRequest("POST", "http://localhost:9111/update/subscribe", bytes.NewBuffer(jsonStr))
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

// Block Updates

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

// Resolve Updated Member

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


// Login
func TestLoginHandlerSuccess(t *testing.T) {
	fmt.Println("Running TestLoginHandlerSuccess")
    str := `{
        "email": "1505443741@test.com",
        "password": "1505443741"
    }`
    jsonStr := []byte(str)

    token := doLogin("1505443741@test.com", "1505443741")

    req := httptest.NewRequest("POST", "http://localhost:9111/login", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.LoginHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"success":true,"token":"` + token + `"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestLoginHandlerFailedWrongUsername(t *testing.T) {
	fmt.Println("Running TestLoginHandlerFailedWrongUsername")
    str := `{
        "email": "1505443741@test.comx",
        "password": "1505443741"
    }`
    jsonStr := []byte(str)

    req := httptest.NewRequest("POST", "http://localhost:9111/login", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.LoginHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Invalid username or password."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestLoginHandlerFailedWrongPassword(t *testing.T) {
	fmt.Println("Running TestLoginHandlerFailedWrongPassword")
    str := `{
        "email": "1505443741@test.com",
        "password": "xxx"
    }`
    jsonStr := []byte(str)

    req := httptest.NewRequest("POST", "http://localhost:9111/login", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.LoginHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Invalid username or password."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func doLogin(email string, password string) string {
    type LoginResult struct {
        Success     bool         `json:"success"`
        Token       string       `json:"token"`
    }

    var result LoginResult

    str := `{
        "email": "` + email + `",
        "password": "` + password + `"
    }`
    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/login", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.LoginHandler)
	handler.ServeHTTP(rr, req)

    if err := json.Unmarshal([]byte(rr.Body.String()), &result); err != nil {
        panic(err)
    }
    
    return result.Token
}

// Create Connections 2

func TestCreateFriendConnectionHandler2Success(t *testing.T) {
	fmt.Println("Running TestCreateFriendConnectionHandler2Success")
    str := `{
        "email": "test4@test.com"
    }`
    jsonStr := []byte(str)
    
    token := doLogin("test5@test.com", "test5")
    
    req := httptest.NewRequest("POST", "http://localhost:9111/friends", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler2)
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

func TestCreateFriendConnectionHandlers2FailedSameEmail(t *testing.T) {
	fmt.Println("Running TestCreateFriendConnectionHandlers2FailedSameEmail")
    str := `{
        "email": "test5@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test5@test.com", "test5")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler2)
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

func TestCreateFriendConnectionHandler2FailedWrongJson1(t *testing.T) {
	fmt.Println("Running TestCreateFriendConnectionHandler2FailedWrongJson1")
    str := `{
        "emailx": "test4@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test5@test.com", "test5")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'EmailInput.Email' Error:Field validation for 'Email' failed on the 'min' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestCreateFriendConnectionHandler2FailedWrongJson2(t *testing.T) {
	fmt.Println("Running TestCreateFriendConnectionHandler2FailedWrongJson2")
    str := `{
        "email": [
            "test@test.com",
            "test2@test.com"
        ]
    }`
    jsonStr := []byte(str)

    token := doLogin("test5@test.com", "test5")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"json: cannot unmarshal array into Go struct field EmailInput.email of type string"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestCreateFriendConnectionHandler2FailedMemberNotFound1(t *testing.T) {
	fmt.Println("Running TestCreateFriendConnectionHandler2FailedMemberNotFound1")
    str := `{
        "email": "test@testx.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test5@test.com", "test5")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler2)
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

func TestCreateFriendConnectionHandler2FailedMemberNotFound2(t *testing.T) {
	fmt.Println("Running TestCreateFriendConnectionHandler2FailedMemberNotFound2")
    str := `{
        "email": "test4@test.com"
    }`
    jsonStr := []byte(str)

    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIwM2ViYmE5Yi0yOGU2LTQzZTEtOTkyZi1jYWJjYzhmMzRhZDEiLCJpc3MiOiJmbSJ9.CtGWBwdY0t-RlW6sx5gbUq-U5uyQVNEues5-ymtuo8o"

    req := httptest.NewRequest("POST", "http://localhost:9111/friends", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.CreateFriendConnectionHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with ID: 03ebba9b-28e6-43e1-992f-cabcc8f34ad1."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

// Resolve Friends 2

func TestResolveFriendsHandler2Success(t *testing.T) {
    fmt.Println("Running TestResolveFriendsHandler2Success")

    token := doLogin("test@test.com", "test")    

    req := httptest.NewRequest("GET", "http://localhost:9111/friends", nil)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveFriendsHandler2)
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

func TestResolveFriendsHandler2FailedMemberNotFound(t *testing.T) {
    fmt.Println("Running TestResolveFriendsHandler2FailedMemberNotFound")

    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJhMTNlYmI1Mi05MzQ1LTRmYmItYjc2Yi02NDFiNWY4NzE3OWYiLCJpc3MiOiJmbSJ9.Oe2NmB2JEfsq5cxPuxreilOTlMAlt_2O9O0W0u-FzAs" 

    req := httptest.NewRequest("GET", "http://localhost:9111/friend/retrieve", nil)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveFriendsHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with Id: a13ebb52-9345-4fbb-b76b-641b5f87179f."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

// Resolve Common Friends 2

func TestResolveCommonFriendsHandler2SuccessWithCommonFriend(t *testing.T) {
    fmt.Println("Running TestResolveCommonFriendsHandler2SuccessWithCommonFriend")
    str := `{
        "email": "test2@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler2)
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

func TestResolveCommonFriendsHandler2SuccessNoCommonFriend(t *testing.T) {
    fmt.Println("Running TestResolveCommonFriendsHandler2SuccessNoCommonFriend")
    str := `{
        "email": "test4@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test3@test.com", "test3")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler2)
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

func TestResolveCommonFriendsHandler2FailedSameEmail(t *testing.T) {
    fmt.Println("Running TestResolveCommonFriendsHandler2FailedSameEmail")
    str := `{
        "email": "test@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler2)
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

func TestResolveCommonFriendsHandler2FailedWrongJson1(t *testing.T) {
    fmt.Println("Running TestResolveCommonFriendsHandler2FailedWrongJson1")
    str := `{
        "emailx": "test@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"Key: 'EmailInput.Email' Error:Field validation for 'Email' failed on the 'min' tag"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandler2FailedWrongJson2(t *testing.T) {
    fmt.Println("Running TestResolveCommonFriendsHandler2FailedWrongJson2")
    str := `{
        "email": [
            "test@test.com",
            "test2@test.com"
        ]
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"InvalidRequestData","data":null,"message":"json: cannot unmarshal array into Go struct field EmailInput.email of type string"}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandler2FailedMemberNotFound1(t *testing.T) {
    fmt.Println("Running TestResolveCommonFriendsHandler2FailedMemberNotFound1")
    str := `{
        "email": "test2@test.com"
    }`

    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJhMTNlYmI1Mi05MzQ1LTRmYmItYjc2Yi02NDFiNWY4NzE3MGYiLCJpc3MiOiJmbSJ9.c8fAi23jI1xdusybwgbj3RemKrRSxRnQ04c5_C3OXCY"

    jsonStr := []byte(str)
    req := httptest.NewRequest("POST", "http://localhost:9111/friends/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with ID: a13ebb52-9345-4fbb-b76b-641b5f87170f."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveCommonFriendsHandler2FailedMemberNotFound2(t *testing.T) {
    fmt.Println("Running TestResolveCommonFriendsHandler2FailedMemberNotFound2")
    str := `{
        "email": "test2@testx.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("POST", "http://localhost:9111/friends/common", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveCommonFriendsHandler2)
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

// Subscribe Updates 2

func TestSubscribeUpdatesHandler2Success(t *testing.T) {
    fmt.Println("Running TestSubscribeUpdatesHandler2Success")
    str := `{
        "email": "test2@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test5@test.com", "test5")

    req := httptest.NewRequest("POST", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler2)
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

func TestSubscribeUpdatesHandler2FailedSameEmail(t *testing.T) {
    fmt.Println("Running TestSubscribeUpdatesHandler2FailedSameEmail")
    str := `{
        "email": "test3@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test3@test.com", "test3")

    req := httptest.NewRequest("POST", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler2)
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

func TestSubscribeUpdatesHandler2FailedRequestorIDNotFound(t *testing.T) {
    fmt.Println("Running TestSubscribeUpdatesHandler2FailedRequestorIDNotFound")
    str := `{
        "email": "test2@test.com"
    }`
    jsonStr := []byte(str)

    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIwNmJmYmRmOC04NGYzLTQ1MjctYjM2ZC05MTBiODU1OTBlN2UiLCJpc3MiOiJmbSJ9.5f-aPyuRO49y3n2o6m7h2Ob_JJjIS0sqRN3S2iu1Yc4"

    req := httptest.NewRequest("POST", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with ID: 06bfbdf8-84f3-4527-b36d-910b85590e7e."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestSubscribeUpdatesHandler2FailedTargetEmailNotFound(t *testing.T) {
    fmt.Println("Running TestSubscribeUpdatesHandler2FailedTargetEmailNotFound")
    str := `{
        "email": "test2@testx.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test3@test.com", "test3")

    req := httptest.NewRequest("POST", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.SubscribeUpdatesHandler2)
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

// Block Updates 2

func TestBlockUpdatesHandler2Success(t *testing.T) {
    fmt.Println("Running TestBlockUpdatesHandler2Success")
    str := `{
        "email": "test2@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler2)
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

func TestBlockUpdatesHandler2FailedRequestorSameEmail(t *testing.T) {
    fmt.Println("Running TestBlockUpdatesHandler2FailedRequestorSameEmail")
    str := `{
        "email": "test@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler2)
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

func TestBlockUpdatesHandler2FailedRequestorIDNotFound(t *testing.T) {
    fmt.Println("Running TestBlockUpdatesHandler2FailedRequestorIDNotFound")
    str := `{
        "email": "test2@test.com"
    }`
    jsonStr := []byte(str)

    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJhMTNlYmI1Mi05MzQ1LTRmYmItYjc2Yi02NDFiNWY4NzE3OWYiLCJpc3MiOiJmbSJ9.Oe2NmB2JEfsq5cxPuxreilOTlMAlt_2O9O0W0u-FzAs"

    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with ID: a13ebb52-9345-4fbb-b76b-641b5f87179f."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestBlockUpdatesHandler2FailedTargetEmailNotFound(t *testing.T) {
    fmt.Println("Running TestBlockUpdatesHandler2FailedTargetEmailNotFound")
    str := `{
        "email": "test2@testx.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler2)
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

func TestBlockUpdatesHandler2FailedInvalidSubscription(t *testing.T) {
    fmt.Println("Running TestBlockUpdatesHandler2FailedInvalidSubscription")
    str := `{
        "email": "1504858412@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("DELETE", "http://localhost:9111/updates", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.BlockUpdatesHandler2)
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

// Resolve Updated Member 2

func TestResolveUpdatedMemberHandler2Success(t *testing.T) {
    fmt.Println("Running TestResolveUpdatedMemberHandler2Success")
    str := `{
        "text": "Hello World! 1504848613@test.com"
    }`
    jsonStr := []byte(str)

    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("POST", "http://localhost:9111/updates/send", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-ID", "a13ebb52-9345-4fbb-b76b-641b5f87179e")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveUpdatedMemberHandler2)
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

func TestResolveUpdatedMemberHandler2FailedMemberNotFound1(t *testing.T) {
    fmt.Println("Running TestResolveUpdatedMemberHandler2FailedMemberNotFound1")
    str := `{
        "text": "Hello World! 1504848613@test.com"
    }`
    jsonStr := []byte(str)
    
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJhMTNlYmI1Mi05MzQ1LTRmYmItYjc2Yi02NDFiNWY4NzE3OWYiLCJpc3MiOiJmbSJ9.Oe2NmB2JEfsq5cxPuxreilOTlMAlt_2O9O0W0u-FzAs"    

    req := httptest.NewRequest("POST", "http://localhost:9111/updates/send", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveUpdatedMemberHandler2)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusInternalServerError)
    }

    // Check the response body is what we expect.
    expected := `{"data":null,"error":{"code":"MemberNotFound","data":null,"message":"No member with ID: a13ebb52-9345-4fbb-b76b-641b5f87179f."}}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestResolveUpdatedMemberHandler2FailedMemberNotFound2(t *testing.T) {
    fmt.Println("Running TestResolveUpdatedMemberHandler2FailedMemberNotFound2")
    str := `{
        "text": "Hello World! kate@example.com"
    }`
    jsonStr := []byte(str)
    
    token := doLogin("test@test.com", "test")

    req := httptest.NewRequest("POST", "http://localhost:9111/updates/send", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Token", token)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(rh.ResolveUpdatedMemberHandler2)
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