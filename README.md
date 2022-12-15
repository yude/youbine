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
        ADMIN_PASSWORD="ここにハッシュ化済みパスワードをペースト>"
        ````


## License
MIT License.