## Overview

Friend Management API

## API Endpoints:

- Add new member
```
POST https://friend-management.herokuapp.com/member
{
	"email": "test@test.com",
    	"name": "test",
    	"password": "test12345",
    	"password2": "test12345"
}
```

- Edit member
```
PUT https://friend-management.herokuapp.com/member
{
	"id": "873e2168-17b9-4302-9a32-4c8a9fe7da35",
    	"name": "test",
    	"password": "test12345",
    	"password2": "test12345"
}
```

- Create friend connection
```
POST https://friend-management.herokuapp.com/friend/add
{
	"friends": [
		"test@test.com",
		"test2@test.com"
	]
}
```

- Retrieve friend list
```
POST https://friend-management.herokuapp.com/friend/retrieve
{
	"email": "test@test.com"
}
```

- Get common friend list
```
POST https://friend-management.herokuapp.com/friend/common
{
	"friends": [
		"test3@test.com",
		"test@test.com"
	]
}
```

- Subscribe updates
```
POST https://friend-management.herokuapp.com/update/subscribe
{
	"requestor": "test@test.com",
	"target": "test2@test.com"
}
```

- Block updates
```
DELETE https://friend-management.herokuapp.com/update/subscribe
{
	"requestor": "test@test.com",
	"target": "test2@test.com"
}
```

- Send updates
```
POST https://friend-management.herokuapp.com/update/send
{
	"sender": "test3@test.com",
	"text": "Hello world! test2@test.com"
}
```

## New Endpoints
Here I propose new APIs for more secure and simpler request data. In this new API, request header will be used to send the requester ID in form of token. The token needs to be generated first via login endpoint.

- Login (to get token)
```
POST https://friend-management.herokuapp.com/login
{
    "email": "1505443741@test.com",
    "password": "1505443741"
}
```

- Create friend connection (new endpoint)
```
POST https://friend-management.herokuapp.com/friends
Token: (generated token from login process)
{
    "email": "test4@test.com"
}
```

- Retrieve friend list (new endpoint)
```
GET https://friend-management.herokuapp.com/friends
Token: (generated token from login process)
```

- Get common friend list (new endpoint)
```
POST https://friend-management.herokuapp.com/friends/common
Token: (generated token from login process)
{
    "email": "test4@test.com"
}
```

- Subscribe updates (new endpoint)
```
POST https://friend-management.herokuapp.com/updates
Token: (generated token from login process)
{
    "email": "test4@test.com"
}
```

- Block updates (new endpoint)
```
DELETE https://friend-management.herokuapp.com/updates
Token: (generated token from login process)
{
    "email": "test2@test.com"
}
```

- Send updates (new endpoint)
```
POST https://friend-management.herokuapp.com/updates/send
Request-Header:
Token: (generated token from login process)
{
    "text": "Hello World! 1504848613@test.com"
}
```
