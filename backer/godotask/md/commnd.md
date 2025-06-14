
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
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0ciIsInJvbGUiOiJ1c2VyIiwiZXhwIjoxNzQ5MzEwNDM5fQ.EyfPtEBqLCUgyYTXZO5mKSKV6mv1zG3TsMX1wjt8nGI"

