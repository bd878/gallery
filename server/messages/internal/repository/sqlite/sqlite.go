package repository

import (
  "context"
  "errors"
  "fmt"
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
  "github.com/bd878/gallery/server/messages/pkg/model"
  "github.com/bd878/gallery/server/messages/internal/repository"
  usermodel "github.com/bd878/gallery/server/users/pkg/model"
)

type Repository struct {
  db *sql.DB

  insertSt *sql.Stmt
}

func New(dbfilepath string) (*Repository, error) {
  db, err := sql.Open("sqlite3", "file:" + dbfilepath)
  if err != nil {
    return nil, err
  }

  insertSt, err := db.Prepare(
    "INSERT INTO messages(" +
      "user_id, " +
      "createtime, " +
      "message, " +
      "file, " +
      "file_id, " +
      "log_index, " +
      "log_term" +
    ") VALUES (?,?,?,?,?,?,?)",
  )
  if err != nil {
    return nil, err
  }

  return &Repository{
    db: db,

    insertSt: insertSt,
  }, nil
}

func (r *Repository) Put(ctx context.Context, msg *model.Message) error {
  res, err := r.insertSt.ExecContext(ctx,
    msg.UserId,
    msg.CreateTime,
    msg.Value,
    msg.FileName,
    msg.FileId,
    msg.LogIndex,
    msg.LogTerm,
  )
  id, _ := res.LastInsertId()
  fmt.Println("added msg with id:", id)
  return err
}

func (r *Repository) Truncate(ctx context.Context) error {
  fmt.Println("truncate")
  _, err := r.db.ExecContext(ctx,
    "DELETE FROM messages",
  )
  if err != nil {
    return err
  }
  return nil
}

func (r *Repository) HasByLog(ctx context.Context, logIndex, logTerm uint64) (bool, error) {
  row := r.db.QueryRowContext(ctx,
    "SELECT id FROM messages WHERE log_index = ? AND log_term = ?",
    logIndex, logTerm,
  )

  var id int
  if err := row.Scan(&id); err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return false, nil
    }
    return false, err
  }
  return true, nil
}

func (r *Repository) Get(ctx context.Context, userId usermodel.UserId) ([]model.Message, error) {
  rows, err := r.db.QueryContext(ctx,
    "SELECT id, user_id, createtime, message, file, file_id, log_index, log_term " +
    "FROM messages WHERE user_id = ?",
    int(userId),
  )
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var res []model.Message
  for rows.Next() {
    var id int
    var userId int
    var createtime string
    var value string
    var fileCol sql.NullString
    var fileIdCol sql.NullString
    var logIndex uint64
    var logTerm uint64
    if err := rows.Scan(
      &id,
      &userId,
      &createtime,
      &value,
      &fileCol,
      &fileIdCol,
      &logIndex,
      &logTerm,
    ); err != nil {
      return nil, err
    }

    var fileName string
    if fileCol.Valid {
      fileName = fileCol.String
    }
    var fileId string
    if fileIdCol.Valid {
      fileId = fileIdCol.String
    }
    res = append(res, model.Message{
      Id: model.MessageId(id),
      UserId: userId,
      CreateTime: createtime,
      Value: value,
      FileName: fileName,
      FileId: model.FileId(fileId),
      LogIndex: logIndex,
      LogTerm: logTerm,
    })
  }
  return res, nil
}

func (r *Repository) GetOne(ctx context.Context, userId usermodel.UserId, id model.MessageId) (model.Message, error) {
  row := r.db.QueryRowContext(ctx,
    "SELECT id, user_id, createtime, message, file, file_id, log_index, log_term " +
    "FROM messages WHERE user_id = ? AND id = ?",
    int(userId), int(id),
  )

  var msg model.Message
  var fileCol sql.NullString
  var fileIdCol sql.NullString
  if err := row.Scan(
    &msg.Id,
    &msg.UserId,
    &msg.CreateTime,
    &msg.Value,
    &fileCol,
    &fileIdCol,
    &msg.LogIndex,
    &msg.LogTerm,
  ); err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return msg, repository.ErrNotFound
    }
    return msg, err
  }
  if fileCol.Valid {
    msg.FileName = fileCol.String
  }
  if fileIdCol.Valid {
    msg.FileId = model.FileId(fileIdCol.String)
  }
  return msg, nil
}

func (r *Repository) PutBatch(ctx context.Context, batch []*model.Message) error {
  for _, msg := range batch {
    _, err := r.insertSt.ExecContext(ctx,
      msg.UserId,
      msg.CreateTime,
      msg.Value,
      msg.FileName,
      msg.FileId,
      msg.LogIndex,
      msg.LogTerm,
    )
    if err != nil {
      return err
    }
  }
  return nil
}

func (r *Repository) GetBatch(ctx context.Context) ([]model.Message, error) {
  rows, err := r.db.QueryContext(ctx,
    "SELECT id, user_id, createtime, message, file, file_id, log_index, log_term " +
    "FROM messages",
  )
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var res []model.Message
  for rows.Next() {
    var id int
    var userId int
    var value string
    var createtime string
    var fileCol sql.NullString
    var fileIdCol sql.NullString
    var logIndex uint64
    var logTerm uint64

    if err := rows.Scan(
      &id,
      &userId,
      &createtime,
      &value,
      &fileCol,
      &fileIdCol,
      &logIndex,
      &logTerm,
    ); err != nil {
      return nil, err
    }

    var fileName string
    if fileCol.Valid {
      fileName = fileCol.String
    }
    var fileId string
    if fileIdCol.Valid {
      fileId = fileIdCol.String
    }
    res = append(res, model.Message{
      Id: model.MessageId(id),
      UserId: userId,
      CreateTime: createtime,
      Value: value,
      FileName: fileName,
      FileId: model.FileId(fileId),
      LogIndex: logIndex,
      LogTerm: logTerm,
    })
  }

  return res, nil
}