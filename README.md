# Next.js + Go学習用のリポジトリ
## 導入するもの
- logger.goにsentry追加
- leftfook


## 環境構築にあたってインストールするもの
開発を進めるにあたって以下のツール、パッケージをインストールする必要があります。
```shell
golang 1.22.3
nodejs 22.1.0
task 3.37.2
lefthook 1.6.12
```

brew等でinstallしても問題ありませんが、[asdf](https://asdf-vm.com/)を使用している場合は`.tool-versions`に各ツール、パッケージが指定バージョンで設定済みなので簡単かと思います。
asdfで開発を進める場合、本PJのみ指定のバージョンで開発を進めることが可能です。

asdfで進める場合は事前に以下のpluginをインストールする必要があります。
```shell
asdf plugin add golang
asdf plugin add nodejs https://github.com/asdf-vm/asdf-nodejs.git
asdf plugin add task
asdf plugin add lefthook https://github.com/jtzero/asdf-lefthook.git
```
pluginを追加するのみで`.tool-versions`記載のバージョンで自動的にインストールされます。
