# go zap sample

## develop environment

### WSL2に直接Goをインストールする環境

- WSL2でUbuntuをインストールする。
- WSL2上でGOをインストールする。
- ホスト環境でVScodeをインストールする
  - Remote Developmentプラグインをインストールする
- コンテナ環境で以下のコマンドを実行する。  

  ```bash
  code .
  ```

  - ホスト環境側のVScodeが立ち上がり、Remote Developmentを経由してコンテナ環境側のVScode serverに接続される。
  - VScodeの左下のアイコンが ![img](./vscode.jpg) になっている。

> [how to install go](https://go.dev/doc/install)

### Docker on WSL2の環境
