package team

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo interface {
	GetAllTeams() (err error, teams []Team)
	GetTeam(id int) (err error, team Team)
	CreateTeam(team Team) (err error)
	UpdateTeam(team Team) (err error)
	DeleteTeam(id int) (err error)
}

type repo struct {
	DB   *pgxpool.Pool
	psql sq.StatementBuilderType
}

func NewRepo(db *pgxpool.Pool) *repo {
	return &repo{
		DB:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repo) GetAllTeams() (err error, teams []Team) {
	sql, args, err := r.psql.Select("*").From("teams").ToSql()
	if err != nil {
		err = fmt.Errorf("could not generate query: %w", err)
		return
	}

	ctx := context.Background()

	err = pgxscan.Select(ctx, r.DB, &teams, sql, args...)
	if err != nil {
		err = fmt.Errorf("could not get teams: %w", err)
	}

	return
}

func (r *repo) GetTeam(id int) (err error, team Team) {
	sql, args, err := r.psql.Select("*").From("teams").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		err = fmt.Errorf("could not generate query: %w", err)
		return
	}

	ctx := context.Background()

	err = pgxscan.Get(ctx, r.DB, &team, sql, args...)
	if err != nil {
		err = fmt.Errorf("could not get team: %w", err)
	}

	return
}

func (r *repo) CreateTeam(team Team) (err error) {
	sql, args, err := r.psql.Insert("teams").Columns("name", "members", "image").Values(team.Name, team.Members, team.Image).ToSql()
	if err != nil {
		err = fmt.Errorf("could not generate query: %w", err)
		return
	}

	ctx := context.Background()

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		err = fmt.Errorf("could not begin transaction: %w", err)
		return
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		err = fmt.Errorf("could not create team: %w", err)
		return
	}

	if err = tx.Commit(ctx); err != nil {
		err = fmt.Errorf("could not commit transaction: %w", err)
	}
	return
}

func (r *repo) UpdateTeam(team Team) (err error) {
	sql, args, err := r.psql.Update("teams").Set("name", team.Name).Where(sq.Eq{"id": team.ID}).ToSql()
	if err != nil {
		err = fmt.Errorf("could not generate query: %w", err)
		return
	}

	ctx := context.Background()

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		err = fmt.Errorf("could not begin transaction: %w", err)
		return
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		err = fmt.Errorf("could not update team: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("could not commit transaction: %w", err)
	}
	return
}

func (r *repo) DeleteTeam(id int) (err error) {
	sql, args, err := r.psql.Delete("teams").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		err = fmt.Errorf("could not generate query: %w", err)
		return
	}

	ctx := context.Background()

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		err = fmt.Errorf("could not begin transaction: %w", err)
		return
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, sql, args...); err != nil {
		err = fmt.Errorf("could not delete team: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("could not commit transaction: %w", err)
	}
	return
}
