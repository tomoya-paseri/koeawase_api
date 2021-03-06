# パワースペクトルを用いて音声の類似度を計算してみる

## 注意

これは別に学術的なものではなく, そういう計算で特徴量が出せるならこんなことできるんじゃね? というお遊びです

## やっていること

人によって, 声が高い/低いと特徴があります

なので, フーリエ変換を行い, 各周波数ごとのパワースペクトルを算出し, それを各周波数帯を1次元とするN次元のベクトルとします

そして, 作成したベクトルの類似度を計算すれば声がどれくらい似ているかを算出できるのではないかと考えました

...というお遊びです

1. 入力された音声をそのままフーリエ変換(だいたいなんで時間区切りとかはしていません)
2. 絶対値をだして周波数帯ごとのパワースペクトルに変換
3. 人間の音声はだいたい100 ~ 2000hzなのでその範囲でスライス
4. コサイン類似度で類似度を計算する

## 使い方

- [CLI](./src/Task/Task.md)

## goについて

このプロジェクトはGOPATH以下に配置されていなくても正しく動作するようにしてあります. なので各ローカルパッケージごとにgo.modを設定する必要があります

## 参考ドキュメント

- [fftとパワースペクトルについて](https://jp.mathworks.com/help/signal/examples/measuring-signal-similarities.html)
- [DIコンテナ](https://github.com/google/wire)
