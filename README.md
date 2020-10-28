# SweetPet
A simple tracking solution (back-end) for your pets' blood sugar levels.
당뇨에 걸린 애완동물 혈당을 기록하는 간단한 Solution

목표 : Backend를 최대한 가볍게 만들기(Client와 Overhead 분담)

## How To Use
*There is no front at the moment!!*

### Create User
    URI : /api/v1/user/create
    
    Example Request 1
    curl -X POST localhost:8080/api/v1/user/create -H 'Content-Type: application/json' -d '{"name":"test"}'
    
    Example Response
    {"name":"test","uuid":"531bb6dc-1919-11eb-8549-98e7437c2803"}
    
    Example Request 2
    curl -X POST localhost:8080/api/v1/user/create -H 'Content-Type: application/json' -d '{"name":"test2", "low":45, "high": 190}'
    
    Example Response
    {"name":"test2","uuid":"de76ad3f-191e-11eb-af22-98e7437c2803","low":45,"high":190}

### Modify Threshold
    URI : /api/v1/user/modify
    Example Request 1
    curl -X PUT localhost:8080/api/v1/user/modify -H 'Content-Type: application/json' -d '{"id": 6, "low": 55, "high": 165}'
    
    Example Response
    {"id":6,"name":"","low":55,"high":165}
    
    Example Request 2
    curl -X PUT localhost:8080/api/v1/user/modify -H 'Content-Type: application/json' -d '{"id": 7, "low": 80}'
    
    Example Response
    {"id":7,"name":"","low":80}

### Generate Link
    URI : /api/v1/user/new_link
    Example Request 1
    curl -X GET localhost:8080/api/v1/user/new_link -d '{"id": 5}'

    Example Response
    {"name":"","link":"OQRKqsGsRR"}
    
    Example Request 2
    curl -X GET localhost:8080/api/v1/user/new_link -d '{"id": 5}'
    
    Example Response
    {"name":"","link":"ktQqhpxLNB"}

### Add Entry
    URI : /api/v1/health/:link
    Example Request
    **In Progress**

### Modify Entry
    URI : /api/v1/health/:link/:ts
    Example Request
    **In Progress**

### Delete Entry
    URI : /api/v1/health/:link/:ts
    Example Request
    **In Progress**

### Retrieve Entries
    URI : /api/v1/health/:link
    Example Request 1
    curl -X GET localhost:8080/api/v1/health/PLRlrIotKP -H 'Content-Type: application/json' -d '{"id": 1}'
    
    Example Response
    {"user":{"id":1,"name":"Test","low":50,"high":120},"entries":[{"blood_sugar":150,"ts":"2020-10-28 07:28:23"},{"blood_sugar":120,"ts":"2020-10-28 07:28:26"},{"blood_sugar":110,"ts":"2020-10-28 07:28:29"},{"blood_sugar":130,"ts":"2020-10-28 07:28:31"},{"blood_sugar":200,"ts":"2020-10-28 07:28:35"},{"blood_sugar":160,"ts":"2020-10-28 07:28:41"},{"blood_sugar":140,"ts":"2020-10-28 07:28:44"},{"blood_sugar":110,"ts":"2020-10-28 07:28:47"},{"blood_sugar":80,"ts":"2020-10-28 07:28:50"}]}

## To Add
