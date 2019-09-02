# 商品登録API

## 環境
### 使用言語
・golang:1.12.6-alpine<br>
・mysql:5.7<br>
#### ORマッパー
・gorm

## 起動手順
### Qiitaアプリケーション登録(認証で使います)
#### リダイレクトURL設定
Qiitaアプリケーション登録画面のリダイレクトURLに下記を指定
```localhost:8080/login/redirect```

#### 環境変数設定
```export QIITA_CLIANT="Qiitaアプリケーション登録で指定されたクライアントID"```<br>
```export QIITA_SECRET="Qiitaアプリケーション登録で指定されたシークレットキー"```

#### dockerコンテナ起動<br>
```docker-compose up```<br>
#### ネットワーク作成<br>
```docker network create --driver bridge go-product```

## ディレクトリ構成
.<br>
├── Dockerfile<br>
├── README.md<br>
├── api<br>
│   ├── handlers<br>
│   │   ├── authorizationHandler.go<br>
│   │   ├── customeErrorHandler.go<br>
│   │   ├── productHandler.go<br>
│   │   ├── timeHandler.go<br>
│   │   └── validationErrorHandler.go<br>
│   ├── middlewares<br>
│   │   ├── authorizationMiddleware.go<br>
│   │   └── mainMiddleware.go<br>
│   └── models<br>
│       ├── authorization.go<br>
│       ├── error.go<br>
│       ├── product.go<br>
│       └── user.go<br>
├── db<br>
│   ├── Dockerfile<br>
│   ├── connection.go<br>
│   └── mysql_data<br>
├── docker-compose.yml<br>
├── go.mod<br>
├── go.sum<br>
├── router<br>
│   └── router.go<br>
├── server.go<br>
├── static<br>
    └── images

## API認証
```localhost:8080/login``` でQiitaログイン<br>
アクセストークンが返却される<br>
Authorizationヘッダーに ```Bearer: "返却されたアクセストークン"``` を指定してリクエスト

## APIリクエストURL
共通: ```localhost:8080```
<br>
概要|httpメソッド|URL
:-:|:-:|:-:
ログイン|GET|login
全件取得|GET|products
1件取得|GET|products/{id}
登録|POST|products
編集|PATCH|products/{id}
削除|DELETE|products/{id}
検索|GET|products/search
画像取得|GET|products/images

## APIリクエスト例
Content-Type: application/json
```
{
    title: "タイトル",
    description: "説明",
    price: 100,
    image: "base64エンコードした画像データ"
}
```

