# JAWS-morning-34
JAWS-UG朝会#34で使用したソースコードです。

## 内容物
invokeするとGo Proverbsをランダムに一つ出力してくれるLambda関数を一つデプロイします。

## コマンド一覧
|コマンド|概要|
|---|---|
|`make bootstrap`|(初回デプロイ時のみ必須)CDK toolkitのデプロイ|
|`make deploy`|リソースのデプロイ|
|`make destroy`|CDK toolkitを含む全リソースの削除|
