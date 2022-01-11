/*
Package main は qiitan インタプリタの本体 Golang パッケージです。

インタプリタ（`qiitan` コマンド）自身の機能のみを提供するため、qiitan スクリプトで使える標準モジュールなどは、このパッケージに実装しないでください。

標準モジュール（同梱モジュール）の追加による Qiitan スクリプトの拡張は `../modules/` ディレクトリに設置してください。

*/
package main