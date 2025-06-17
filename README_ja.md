<div align="center">
  <a href="https://alist-archive.pages.dev/">
    <img width="100px" alt="logo" src="https://cdn.jsdelivr.net/gh/alist-archive/logo@main/logo.svg"/>
  </a>
  <p><em>🗂️Gin と Solidjs による、複数のストレージをサポートするファイルリストプログラム。</em></p>
</div>

---

[English](./README.md) | [中文](./README_cn.md) | 日本語 | [Contributing](./CONTRIBUTING.md) | [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md)

## 特徴

- [x] マルチストレージ
    - [x] ローカルストレージ
    - [x] [Aliyundrive](https://www.alipan.com/)
    - [x] OneDrive / Sharepoint ([グローバル](https://www.office.com/), [cn](https://portal.partner.microsoftonline.cn),de,us)
    - [x] [189cloud](https://cloud.189.cn) (Personal, Family)
    - [x] [GoogleDrive](https://drive.google.com/)
    - [x] [123pan](https://www.123pan.com/)
    - [x] FTP / SFTP
    - [x] [PikPak](https://www.mypikpak.com/)
    - [x] [S3](https://aws.amazon.com/s3/)
    - [x] [Seafile](https://seafile.com/)
    - [x] [UPYUN Storage Service](https://www.upyun.com/products/file-storage)
    - [x] WebDav(Support OneDrive/SharePoint without API)
    - [x] Teambition([China](https://www.teambition.com/ ),[International](https://us.teambition.com/ ))
    - [x] [Mediatrack](https://www.mediatrack.cn/)
    - [x] [139yun](https://yun.139.com/) (Personal, Family, Group)
    - [x] [YandexDisk](https://disk.yandex.com/)
    - [x] [BaiduNetdisk](http://pan.baidu.com/)
    - [x] [Terabox](https://www.terabox.com/main)
    - [x] [UC](https://drive.uc.cn)
    - [x] [Quark](https://pan.quark.cn)
    - [x] [Thunder](https://pan.xunlei.com)
    - [x] [Lanzou](https://www.lanzou.com/)
    - [x] [ILanzou](https://www.ilanzou.com/)
    - [x] [Aliyundrive share](https://www.alipan.com/)
    - [x] [Google photo](https://photos.google.com/)
    - [x] [Mega.nz](https://mega.nz)
    - [x] [Baidu photo](https://photo.baidu.com/)
    - [x] SMB
    - [x] [115](https://115.com/)
    - [X] Cloudreve
    - [x] [Dropbox](https://www.dropbox.com/)
    - [x] [FeijiPan](https://www.feijipan.com/)
    - [x] [dogecloud](https://www.dogecloud.com/product/oss)
- [x] デプロイが簡単で、すぐに使える
- [x] ファイルプレビュー (PDF, マークダウン, コード, プレーンテキスト, ...)
- [x] ギャラリーモードでの画像プレビュー
- [x] ビデオとオーディオのプレビュー、歌詞と字幕のサポート
- [x] Office ドキュメントのプレビュー (docx, pptx, xlsx, ...)
- [x] `README.md` のプレビューレンダリング
- [x] ファイルのパーマリンクコピーと直接ダウンロード
- [x] ダークモード
- [x] 国際化
- [x] 保護されたルート (パスワード保護と認証)
- [x] WebDav (詳細は https://alist-archive-docs.pages.dev/guide/webdav.html を参照)
- [x] [Docker デプロイ](https://ghcr.io/alist-archive/alist)
- [x] Cloudflare ワーカープロキシ
- [x] ファイル/フォルダパッケージのダウンロード
- [x] ウェブアップロード(訪問者にアップロードを許可できる), 削除, mkdir, 名前変更, 移動, コピー
- [x] オフラインダウンロード
- [x] 二つのストレージ間でファイルをコピー
- [x] シングルスレッドのダウンロード/ストリーム向けのマルチスレッド ダウンロード アクセラレーション

## ドキュメント

<https://alist-archive.pages.dev/>

## デモ

<https://alist-archiv-demo.pages.dev/>

## コントリビューター

これらの素晴らしい人々に感謝します:

https://codeberg.org/alist/alist/activity/contributors

## ライセンス

`AList` は AGPL-3.0 ライセンスの下でライセンスされたオープンソースソフトウェアです。

## 免責事項
- このプログラムはフリーでオープンソースのプロジェクトです。ネットワークディスク上でファイルを共有するように設計されており、golang のダウンロードや学習に便利です。利用にあたっては関連法規を遵守し、悪用しないようお願いします;
- このプログラムは、公式インターフェースの動作を破壊することなく、公式 sdk/インターフェースを呼び出すことで実装されています;
- このプログラムは、302リダイレクト/トラフィック転送のみを行い、いかなるユーザーデータも傍受、保存、改ざんしません;
- このプログラムを使用する前に、アカウントの禁止、ダウンロード速度の制限など、対応するリスクを理解し、負担する必要があります;
