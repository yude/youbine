# youbine
郵便やさん

## Build
```shell
$ cd app; go build .
```

## Setup
* 起動前に、管理用パスワードを Bcrypt によってハッシュ化し `ADMIN_PASSWORD` 環境変数として読み込まれるよう設定してください。
    * Docker を利用する場合: `docker-compose.yaml` と同じディレクトリに `.env` を作成し、以下の形式でハッシュ化済みパスワードを保存します。
        ```
        ADMIN_PASSWORD="ここにハッシュ化済みパスワードをペースト"
        ````
* 以下の環境変数で動作を変更できます。
    | 値                | 説明                                                              | 既定の値                               | 
    | ----------------- | ----------------------------------------------------------------- | -------------------------------------- | 
    | POSTGRES_HOST     | PostgreSQL が稼働するホストの宛先                                 | localhost                              | 
    | POSTGRES_PORT     | PostgreSQL が稼働するポート                                       | 5432                                   | 
    | POSTGRES_DB       | PostgreSQL のデータベース名                                       | app                                    | 
    | POSTGRES_USER     | PostgreSQL のユーザー名                                           | app                                    | 
    | POSTGRES_PASSWORD | PostgreSQL のパスワード                                           | app                                    | 
    | ADMIN_PASSWORD    | 管理ページのパスワード、Bcrypt により事前にハッシュ化してください | <random>; to protect unintended access | 

## License
MIT License.