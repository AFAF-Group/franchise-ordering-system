package common

import "github.com/labstack/echo/v4"

type Controller struct {
}

func (c Controller) BindAndValidate(ctx echo.Context, data interface{}) error {
	// Bind data
	if err := ctx.Bind(data); err != nil {
		return err
	}

	// Validate data
	// if err := ctx.Validate(data); err != nil {
	// 	return err
	// }

	return nil
}
