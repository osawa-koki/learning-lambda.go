# learning-lambda.go

🪪🪪🪪 Go言語でLambdaを書くためのテンプレートです！  

## 開発環境の構築方法

最初にAWS CLIをインストールします。  
<https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/install-cliv2.html>  

以下のコマンドを実行して、AWS CLIのバージョンが表示されればOKです。  

```shell
aws --version
```

認証情報を設定します。  

```shell
aws configure
```

以下のように聞かれるので、適宜入力してください。

```shell
AWS Access Key ID [None]: アクセスキーID
AWS Secret Access Key [None]: シークレットアクセスキー
Default region name [None]: リージョン名
Default output format [None]: json
```

これらの情報は、AWSのコンソール画面から確認できます。  
IAMのページから指定のユーザを選択肢、アクセスキーを発行してください。  

続いて、AWS SAMをインストールします。  
こちらはサーバレスアプリケーションを構築するためのツールです。  
<https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html>  

以下のコマンドを実行して、AWS SAMのバージョンが表示されればOKです。  

```shell
sam --version
```

サーバサイドアプリケーションを開発用に実行するためには、以下のコマンドを実行します。  
ビルドにはDockerが必要です。  

```shell
sam build --use-container
sam local start-api
```

<http://localhost:3000/api/hello>にリクエストを投げて、`hello world`が返ってくればOKです。  

---

`main`ブランチにプッシュすると、GitHub Actionsが実行されて、Lambdaがデプロイされます。  
デプロイには以下のGitHubのSecretが必要です。  

| Name | Value |
| --- | --- |
| PROJECT_NAME | プロジェクト名 |
| AWS_ACCESS_KEY_ID | アクセスキーID |
| AWS_SECRET_ACCESS_KEY | シークレットアクセスキー |
| AWS_REGION | リージョン名 |
| LAMBDA_DOTENV | Lambdaで使用する`.env`ファイルの内容 |

---

手動でデプロイする場合は、以下のコマンドを実行します。  

```shell
sam build --use-container
sam deploy --stack-name <プロジェクト名>
```
