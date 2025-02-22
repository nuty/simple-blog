
# Simple Blog API 文档



## API URL
```
http://localhost:3001
```

## URI

### 1. **GET /comments/{article_id}**

#### 描述:
评论列表接口

#### Query 参数:
- `sort` (可选): 按照创建时间的排序方式
  - `asc`: 正序
  - `desc`: 倒序（默认）

#### 响应状态码:
- **200 OK**: 返回评论列表.
- **400 Bad Request**: 参数错误

#### 请求:
```
GET /comments/1?sort=desc
```

#### 响应:
```json
{
  "comments": [
    {
      "id": 1,
      "article_id": 1,
      "parent_comment_id": null,
      "content": "This is a root comment",
      "created_at": "2025-02-22T19:01:19.247816+08:00",
      "children": [
        {
          "id": 2,
          "article_id": 1,
          "parent_comment_id": 1,
          "content": "This is a reply to the root comment",
          "created_at": "2025-02-22T19:02:19.247816+08:00",
          "children": []
        }
      ]
    }
  ]
}
```

### 2. **POST /comments**

#### 描述:
创建一个新评论

#### 请求体:
```json
{
  "article_id": 1,
  "parent_comment_id": null,
  "content": "This is a new comment"
}
```

#### 响应:
- **200 OK**: 创建成功
- **400 Bad Request**: 参数错误.
- **500 Internal Server Error**: 创建失败，服务器错误

#### 请求:
```
POST /comments
```

#### 响应:
```json
{
  "message": "Comment posted successfully",
  "comment": {
    "id": 3,
    "article_id": 1,
    "parent_comment_id": null,
    "content": "This is a new comment",
    "created_at": "2025-02-22T19:05:00.247816+08:00"
  }
}
```

### 3. **DELETE /comments/{comment_id}**

#### 描述:
删除一个评论

#### Response:
- **200 OK**: 删除成功
- **404 Not Found**: 评论找不到
- **500 Internal Server Error**: 删除失败，服务器错误

#### 请求:
```
DELETE /comments/3
```

#### 响应:
```json
{
  "message": "deleted successfully"
}
```

## Redis 缓存规则

### 获取评论列表:
根据文章id获取评论，优先从redis中获取，如不能命中，则在数据库中获取，并存入redis。

### 缓存清理规则:
最简单的实现，当新增或者删除后，清空该文章下的所有缓存

## Redis 连接代码:
```go
func InitRedis() {
    Rdb = redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "",
        DB:       0,
    })

    pong, err := Rdb.Ping(context.Background()).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }
    fmt.Println("Connected to Redis! Response:", pong)
}
```


## 数据库配置

`config/config.toml` 

```toml
[database]
host = "db"
port = 5432
user = "postgres"
password = "helloworld"
dbname = "postgres"
sslmode = "disable"
```

## `docker-compose.yml` 

```yaml
version: '3.8'

services:
  app:
    build: .
    container_name: simple-blog-app
    ports:
      - "3001:3001"
    depends_on:
      - db
      - redis
    environment:
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=helloworld
      - DB_NAME=postgres
      - REDIS_HOST=redis
      - REDIS_PORT=6379
    volumes:
      - .:/app
    networks:
      - app-network

  db:
    image: postgres:15-alpine
    container_name: simple-blog-db
    environment:
      POSTGRES_PASSWORD: helloworld
      POSTGRES_DB: postgres
    ports:
      - "5432:5432"
    networks:
      - app-network

  redis:
    image: redis:latest
    container_name: simple-blog-redis
    ports:
      - "6379:6379"
    networks:
      - app-network

networks:
  app-network:
    driver: bridge
```

## 运行应用

1. 使用 Docker Compose 或 Podman Compose 运行.

    ```
    docker-compose up --build
    ```
3.  `http://localhost:3001`.
4. 执行测试用例 （需要手动修改 config.toml 和 providers/redis.go 中数据库和redis的连接参数）
    ```
    go test ./tests -v
    ```
也可直接执行python的测试脚本（需要修改测试api的地址）
  ```
    python tests/test_api.py
    ```