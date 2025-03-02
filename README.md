# README

## **🛠 開発環境セットアップ**

### **1. 事前準備**

1. **Docker Desktop のインストール**（[公式サイト](https://docs.docker.com/get-docker)）
2. **Docker ネットワークの作成**
   ```sh
   make network
   ```

### **2. 開発環境の起動**

```sh
make up
```

- 起動後、[http://localhost:3000](http://localhost:3000) にアクセス

### **3. データベースのセットアップ**

```sh
make migrate    # データベースマイグレーション
make seed       # サンプルデータの投入
```

### **4. 開発環境の停止**

```sh
make down
```

---

## **🏗 本番環境セットアップ**

### **1. 本番環境の起動**

```sh
make up-prod
```

### **2. 本番環境の停止**

```sh
make down-prod
```

### **3. クリーンアップ**

```sh
make clean
```

---

## **📌 よく使うコマンド**

### **1. 開発関連**

```sh
make rebuild   # 開発環境をクリーンアップして再構築
```

### **2. Docker ネットワーク**

```sh
make network   # Docker ネットワークを作成
```

### **3. コンテナの削除・クリーンアップ**

```sh
make clean     # 開発環境のクリーンアップ
```

### **4. MySQL で実行された SQL を確認**

```sql
SET GLOBAL general_log = 'ON';
SET GLOBAL log_output = 'TABLE';  -- ログをテーブルに保存

SELECT event_time, CONVERT(argument USING utf8) AS query_text
FROM mysql.general_log
ORDER BY event_time DESC
LIMIT 10;

SET GLOBAL general_log = 'OFF';
```

---

## **📡 API コマンド**

### **User API 操作**

```sh
# GET User by ID
curl -X GET "http://localhost:8080/api/user/v1/users/1"

# GET All Users
curl -X GET "http://localhost:8080/api/user/v1/users"

# Create User
curl -X POST "http://localhost:8080/api/user/v1/users" -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@example.com","password":"password","role":"user"}'

# Update User
curl -X PUT "http://localhost:8080/api/user/v1/users/1" -H "Content-Type: application/json" -d '{"name":"Alice Updated","email":"alice@example.com","role":"admin"}'

# Delete User
curl -X DELETE "http://localhost:8080/api/user/v1/users/1"
```

---

## **📖 Makefile ターゲット一覧**

```sh
make help       # 使用可能なターゲット一覧を表示
make up         # 開発環境を起動
make down       # 開発環境を停止
make migrate    # データベースマイグレーション
make seed       # サンプルデータの投入
make up-prod    # 本番環境を起動
make down-prod  # 本番環境を停止
make clean      # 不要なコンテナやボリュームの削除
make rebuild    # クリーンアップ後に開発環境を再構築
make network    # Docker ネットワークの作成
```
