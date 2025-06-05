
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
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0Iiwicm9sZSI6InVzZXIiLCJleHAiOjE3NDg3NjQ3MTd9.17ODbiF2VzBg1M8urFoXAg2e0TVsxbHHp-jwWuuWkPs" \
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

curl -X GET http://localhost:8080/api/memory \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InR0Iiwicm9sZSI6InVzZXIiLCJleHAiOjE3NDg3NjQ3MTd9.17ODbiF2VzBg1M8urFoXAg2e0TVsxbHHp-jwWuuWkPs"
