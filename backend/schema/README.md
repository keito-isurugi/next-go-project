# next_go

## テーブル一覧

| 名前 | カラム一覧 | コメント | タイプ |
| ---- | ------- | ------- | ---- |
| [public.todos](public.todos.md) | 8 | Todosテーブル | BASE TABLE |
| [public.users](public.users.md) | 7 | ユーザーテーブル | BASE TABLE |
| [public.categories](public.categories.md) | 5 | カテゴリテーブル | BASE TABLE |
| [public.todo_categories](public.todo_categories.md) | 6 | Todo, カテゴリ中間テーブル | BASE TABLE |

## ER図

```mermaid
erDiagram


"public.todos" {
  integer id
  integer user_id
  varchar_255_ title
  boolean done_flag
  timestamp_without_time_zone created_at
  timestamp_without_time_zone updated_at
  timestamp_without_time_zone deleted_at
  text attachment_file
}
"public.users" {
  integer id
  varchar_255_ name
  varchar_255_ email
  varchar_255_ password
  timestamp_without_time_zone created_at
  timestamp_without_time_zone updated_at
  timestamp_without_time_zone deleted_at
}
"public.categories" {
  integer id
  varchar_255_ name
  timestamp_without_time_zone created_at
  timestamp_without_time_zone updated_at
  timestamp_without_time_zone deleted_at
}
"public.todo_categories" {
  integer id
  integer todo_id
  integer category_id
  timestamp_without_time_zone created_at
  timestamp_without_time_zone updated_at
  timestamp_without_time_zone deleted_at
}
```

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
