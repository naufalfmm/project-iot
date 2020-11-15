package service

import "github.com/labstack/echo/v4"

func (s *service) Count(ctx echo.Context) (int64, error) {
	return s.repository.Count(ctx)
}
