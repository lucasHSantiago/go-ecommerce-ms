package util

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StringToText(s *string) pgtype.Text {
	text := pgtype.Text{
		Valid: s != nil,
	}

	if text.Valid {
		text.String = *s
	}

	return text
}

func TimeToTimestamptz(t *time.Time) pgtype.Timestamptz {
	time := pgtype.Timestamptz{
		Valid: t != nil,
	}

	if time.Valid {
		time.Time = *t
	}

	return time
}

func BoolToBool(b *bool) pgtype.Bool {
	bool := pgtype.Bool{
		Valid: b != nil,
	}

	if bool.Valid {
		bool.Bool = *b
	}

	return bool
}
