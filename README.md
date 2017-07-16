## Degoichi

Degoichiはgolangで書かれた、構成管理ツールになる予定のものです。

## 仕様予定
- config、recipe共にtomlで記述する
- providerはVagrant、Docker、AWS、VPSサーバー、ローカルサーバー
- プロビジョニングを行う

## 目標

- configを読んでVagrantfile、Dockerfileの自動生成
- recipe作成コマンド