# SAFECITY APP (REST) API DOCS

## Error Code list

* 200 - OK 
* 400 - Bad Request 
* 401 - service not found 
* 404 - URL Not Found 
* 453 - customer not found 
* 500 - Internal Server Error 
* 503 - Service Unavailable


**Response error:**
```
{
    "code": 404,
    "msg": "url not found"
}
```

**KEY = sha256( sha256(LOGIN) + sha256(PASSWORD) + TEXT )**
_________________
**LOGIN, PASSWORD, TEXT** - Генерируется и передаётся со стороны Сервиса.

 ****

## Check

**Request:**

**ACCOUNT = PHONE_NUMBER || USER_ID** \
Format **PHONE_NUMBER** = 931441244 


```text
GET method
url: https://bekhatar.tj/api/v1/payment/check-account?login=LOGIN&key=KEY&account=ACCOUNT
```
**Response success:**
```json
{
    "code": 200,
    "customer": {
        "phoneNo": "+992931441244",
        "balance": 50
    },
    "msg": "OK"
}
```

## Pay

**Request**
```text

POST method
url: https://bekhatar.tj/api/v1/payment/pay?login=LOGIN&key=KEY
Request Body: raw
Content/Type: application/json
{
    "transaction_id": 567890342,
    "phone": "931441244",
    "amount": 76.5
}
```

**Response success:**

```json
{
    "TRANSACTION_ID": 11,
    "code": 200,
    "msg": "OK"
}

```
## StatusCheck

**Request**
```text
POST method
url: https://bekhatar.tj/api/v1/payment/pay-check?login=LOGIN&key=KEY
Request Body: raw
Content/Type: application/json
{
    "transaction_id": 567890342,
    "phone": "931441244",
    "amount": 76.5
}
```

**Response success:**
```json
{
    "code": 200,
    "msg": "THE TRANSACTION IS SUCCESSFUL !!!"
}
```

