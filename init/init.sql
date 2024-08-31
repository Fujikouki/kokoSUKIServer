CREATE TABLE chat_messages
(
    id      SERIAL PRIMARY KEY,                   -- 自動インクリメントされる一意のID
    name    VARCHAR(50) NOT NULL,                 -- メッセージを送信したユーザーの名前
    message TEXT        NOT NULL,                 -- チャットメッセージの内容
    time    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP -- メッセージの送信時間、デフォルトで現在のタイムスタンプ
);