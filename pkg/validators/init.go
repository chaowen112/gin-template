package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func Init() (*validator.Validate, error) {
	engine := binding.Validator.Engine()
	var err error

	if v, ok := engine.(*validator.Validate); ok {
		validators := make(map[string]validator.Func)

		for key, _validator := range validators {
			err = v.RegisterValidation(key, _validator)
			if err != nil {
				return nil, errors.Wrapf(err, "register validator '%v' failed", key)
			}
		}

		return v, nil
	}

	err = errors.Errorf("unknown validator '%v'", engine)
	return nil, err
}
