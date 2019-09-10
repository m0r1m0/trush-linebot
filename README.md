川崎市幸区小倉の明日のゴミ通知してくれるやつ
`secret.yaml` に LINE のシークレットとトークン書いて、GAE にデプロイすれば動く
`cron.yaml` に定期実行の時間設定しておく
（需要はない）

```
$ gcloud app deploy
$ gcloud app deploy cron.yaml // 忘れずに
```
