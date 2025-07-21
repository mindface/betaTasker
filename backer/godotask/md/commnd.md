
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"tt","email":"b@t.com","password":"word","role":"user"}'

curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"b@t.com","password":"word"}'

curl -X POST http://localhost:8080/api/memory \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "source_type": "book",
    "title": "サンプル本",
    "author": "著者名",
    "notes": "メモ内容",
    "tags": "tag1,tag2",
    "read_status": "unread",
    "read_date": "2025-06-01T00:00:00Z"
  }'

curl -X POST http://localhost:8080/api/memory \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0ciIsInJvbGUiOiJ1c2VyIiwiZXhwIjoxNzQ5Mzc2MjczfQ.JgOYNFIgBD-Xo50CfGhfTDMNYvC3CZRcBcyW19YDs30" \
  -d '{
    "user_id": 1,
    "source_type": "book",
    "title": "サンプル本",
    "author": "著者名",
    "notes": "メモ内容",
    "tags": "tag1,tag2",
    "read_status": "unread",
    "read_date": "2025-06-01T00:00:00Z"
  }'

curl -X PUT http://localhost:8080/api/memory/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0ciIsInJvbGUiOiJ1c2VyIiwiZXhwIjoxNzQ5Mzg1Nzg4fQ.zA39swpsusZIGdj0-ABdu7dLESJzm8OBLfbzem91I_g" \
  -d '{
    "user_id": 1,
    "source_type": "book",
    "title": "サンプル本2",
    "author": "著者名",
    "notes": "メモ内容",
    "tags": "tag1,tag2",
    "read_status": "unread",
    "read_date": "2025-06-01T00:00:00Z"
  }'

curl -X GET http://localhost:8080/api/memory \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0ciIsInJvbGUiOiJ1c2VyIiwiZXhwIjoxNzUxODQ5NzMzfQ.OqDvU98DDYNZdOl53D7Mzsgv9OrIDFSr1nDl2X5m3DY"

curl -X GET http://localhost:8080/api/memory/aid/MA-Q-02 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0ciIsInJvbGUiOiJ1c2VyIiwiZXhwIjoxNzUxODQ5NzMzfQ.OqDvU98DDYNZdOl53D7Mzsgv9OrIDFSr1nDl2X5m3DY"


curl -X GET http://localhost:8080/api/book

curl -X POST http://localhost:8080/api/book \
  -H "Content-Type: application/json" \
  -d '{"title":"サンプルタイトル","name":"著者名","text":"本文","disc":"説明","imgPath":"path/to/image.png","status":"active"}'

curl -X DELETE http://localhost:8080/api/deletebook/1

curl -X PUT http://localhost:8080/api/updatebook/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"新タイトル","name":"新著者名","text":"新本文","disc":"新説明","imgPath":"new/path.png","status":"inactive"}'


curl -X POST http://localhost:8080/api/assessment \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0Iiwicm9sZSI6InVzZXIiLCJleHAiOjE3NTMxNzA4Njh9._qS5LD_B7k6I5hOok95ujHE4dklB-3qsm2oe2P5Pvkk" \
  -d '{
    "task_id": 3,
    "user_id": 0,
    "effectiveness_score": 80,
    "effort_score": 60,
    "impact_score": 75,
    "qualitative_feedback": "新素材は有望だがコスト面の改善が必要"
  }'
