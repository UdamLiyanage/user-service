# User Microservice
Build Status: ![](https://github.com/UdamLiyanage/user-service/workflows/Go/badge.svg)
***
User Microservice for IoT Platform. This service is responsible to the following features:
*Create users
*Read user data
*Update user data
*Delete users

***
## Document Structure for User - Revision 1
Below is the structure of the JSON document that holds user data
```
{
	"first_name": "First Name",
	"last_name": "Last Name",
	"email": "email@email.com",
	"password": "password",
	"contact_number": "0000000000",
	"address": {
		"line_1": "Address Line 1",
		"line_2": "Address Line 2",
		"line_3": "Address Line 3",
		"postcode": "Postcode",
		"country": "Country"
	},
	"devices": [{
			"device_id": 1,
			"status": true,
			"device_token": "device_token"
		},
		{
			"device_id": 2,
			"status": false,
			"device_token": "device_token"
		}
	],
	"email_verified_at": "Verified Time",
	"created_at": "Created Time",
	"updated_at": "Updated Time"
}
```
