# Midnite Technincal Test

## Overview
This API provides an endpoint for providing alerts based on user actions within the system (deposits or withdrawals).  
It expects JSON input with user, action, amount, and timestamp fields.
All endpoints documented within are accessible via http://localhost:8080

## Event
### `POST /event`

Accepts an alert message payload and validates it.

### Headers
| Header | Value |
|---------|--------|
| `Content-Type` | `application/json` |

---

### Request Body
```json
{
  "userid": 123,
  "action": "withdrawal",
  "amount": 50.0,
  "time": 1729852800
}
```

### Response

#### 204 No Content

#### 405 Method Not Allowed 

#### 500 Internal Server Error