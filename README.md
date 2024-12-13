# webappframeworks

[TechEmpower Framework Benchmarks](https://github.com/TechEmpower/FrameworkBenchmarks) をもとに、おおよそ同様なアプリケーションコードを使って、足元環境でベンチマークして実際のところを確認する。

TechEmpower の結果で興味深い部分を実際に自分でも試して、再現することを見る。再現しなかった場合はできれば、なぜなのか、どちらの結果の方が適切なのかを調べる。

1. TechEmpower のテストケースである JSON シリアライゼーション、単一 DB クエリ、複数 DB クエリ、Fortunes、DB 更新、プレーンテキストに対応したアプリケーションを用意する。
2. DB は足元で MySQL を動かす。
3. 足元の Intel Mac 環境で（本当は Linux 環境の方がいいだろうが）、Docker などを介さずにアプリケーションを実行する。
4. [hey](https://github.com/rakyll/hey) コマンドを使って、localhost の上のアプリケーションのポートにリクエストして、毎秒リクエスト数を計測する。

## Go

Gin と Echo とでは、TechEmpower では Gin の方が少し上位にあった。ただ TechEmpower のベンチマークコードを見ると、Gin は `runtime.GOMAXPROCS(runtime.NumCPU())` を呼んでいるが、Echo は呼んでいない。自分の環境で Echo でも `runtime.GOMAXPROCS(runtime.NumCPU())` 付きで試すと、差はなくなった。

DB 操作は、Go 標準の `database/sql` を使っているが、ほかはもう少しリッチな ORM なり SQL ビルダなりを使っているため、ずるい。

TODO: Ent でも試す。

フレームワークを使わず Go 標準の `net/http` ライブラリを直接使っても、Gin / Echo と有意な差は見られない。それどころか `net/http` のルーティング機能 `ServeMux` はパフォーマンスが悪く、それを使うと Gin / Echo より悪い結果になる。パフォーマンスが悪くなるといっても、そもそもこのベンチマークテストは CPU-bound にはならないように思うのだが、どういう実装なのか有意に悪くなる。手書きで switch 文でシンプルなルーティングを実装したら Gin / Echo と差がなくなる。

[fasthttp](https://github.com/valyala/fasthttp) ライブラリで試すと確かに `net/http` より速いが、あまり有意でなかった。

Go 標準の JSON エンコーダー・デコーダーである `encoding/json` は速くないらしいが、高速とされる [Sonnet](https://github.com/sugawarayuuta/sonnet) に置き換えてもあまり有意な差は見られなかった。大量・複雑な JSON 処理では恐らく差が出るだろうけれど、さすがに TechEcmpower の少量・単純なパターンではボトルネックでないんだろう。

fasthttp を使っている [fiber](https://github.com/gofiber/fiber) はさすがに Gin / Echo より速いが、そこまで有意な差でもなかった。ベンチマークの仕方の問題もあるかもしれないが、これくらいなら、フレームワークとしての信頼性やミドルウェアライブラリの充実性の方を優先して選んだ方がよさそう。

## Node.js

Express アプリケーションは、TechEmpower では cluster を使っていないのでないかと邪推したが、ベンチマークコードを見ると、ちゃんと cluster 化されていた。それでも確かにスコアは Gin や Echo には劣る。ただ、自分の環境では cluster を外して実施しても cluster 化したときと総合的な差はなくて、テストによってはより速くなったりより遅くなったりした。あと、Node.js 標準の cluster パッケージを使ったが、pm2 を使ったら違ったりするのかな。

逆を言えば、libuv がスレッドプールを持ちマルチスレッドであるのは無論として、JavaScript プログラマ目線では実質シングルスレッドのアプリケーションであるのにここまでのパフォーマンスが出せるとも言える。

cluster 関係なく、Go と比べて遅いのはどこがボトルネックなんだろう。もちろん V8 エンジンがいかに強力でも、さすがに Go と JavaScript のオブジェクトモデルでは前者の方が高性能になるだろうが、それが原因だという証拠があるわけでもない。あと SQL ライブラリも意味があるかも。

Fastify, Hono は Express より明確に速かった。それでも Gin / Echo よりは明確に遅かった。Fastify と Hono の間に有意な差はないので、今ならモダンな Hono が一番よい気もする。周辺ライブラリ、いわゆるミドルウェアの数は少ないが、必要性があるかどうか次第。

## Deno

Deno で、[node-mysql2](https://github.com/sidorares/node-mysql2) と Preact を使って Web アプリケーションフレームワークなどは使わず `Deno.serve` そのままで試した。Express よりも Fastify よりも速く、Go の Echo などに近づいていた。

Deno は cluster ライブラリを提供していないのでシングルプロセスでのみ試した。

## Bun

Bun で、Deno とほぼ同じように、[node-mysql2](https://github.com/sidorares/node-mysql2) と [Handlebars](https://handlebarsjs.com/) だけを使って `Bun.serve` そのままで試した。すると、Deno よりも Go の Echo よりも若干速かった。Deno も Bun も Intel Mac のときより Apple Silicon に変えて如実に高速化したので、Apple Silicon では Deno, Bun が強い模様。

TechEmpower のコードによれば、Bun は spawn によるサブプロセス起動と、`Bun.serve` の `reusePort` オプションによって cluster っぽいことができるらしいが、これは Linux でしか有効でない。[issue](https://github.com/oven-sh/bun/issues/2428#issuecomment-1750544989) 参照。 Node.js の cluster でも有意な差は見られなかったので、一応確認したが、Node.js cluster はちゃんと足元環境でも複数 worker プロセスに負荷分散してくれてはいた。

## ASP.NET Core

TechEmpower ではほぼ最速の結果を叩き出した ASP.NET Core だが、自分の Intel Mac / .NET 7 環境で試すと Gin / Echo と同レベルだった。Mac の dotnet のパフォーマンスがよくないというのはありえそうだが、分からない。何かチューニングがあるのかもしれないし、MySQL ドライバーや Entity フレームワークがよくないとかかもしれない。TechEmpower では最新の .NET 8 が使われているのでその差もありそうだけれど、.NET 8 はまだ正式リリースされていないから 7 にしたんだよね。正直、ASP.NET Core は Gin / Echo と同レベルの方が納得ができる。

## Actix-Web

Rust の Actix-Web を足元で試したが、実装がどこかまずいのか Intel Mac では最適化されないのか、Gin / Echo と同レベルのパフォーマンスだった。ただ触ってみて感じたのは、コンパイルにまあまあ時間がかかり、こんなにかかるなら開発の快適さにも影響を与えるように思った。

## Jooby (Kotlin Jooby)

Jooby も、チューニングの問題か Intel Mac では Java が最適化されていないのか、Gin / Echo より悪いくらいのパフォーマンスだった。実際、TechEmpower ソースコードの Dockerfile をもとにチューニング用の java オプションを付与して実行したのに、チューニング用のオプションなしと比べて全然よくならなかったので、なんかよくないっぽい。まあともかく事実として残ったのは、少なくとも Intel Mac では Gin / Echo より遅い。

## FastAPI

TODO

## まとめ

自分が足元で試す分には、Gin / Echo が普通に最速クラスだったからそれでいいじゃんとなった。メモリ使用量までは確認しなかったが、Go はメモリを食わないのを仕事で触っていて知っている。メモリの使い方と GC の設定次第で、昔の Google Chrome 的なリッチな食い方をすることはあるらしいが、会ったことないから分からないし、それはチューニングすればよさそう。

↑上は Intel Mac のときの話で、Apple Silicion だと Deno, Bun が強かった。