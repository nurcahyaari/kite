package errcustom_test

import (
	"testing"

	"github.com/nurcahyaari/kite/internal/utils/errcustom"
	"github.com/stretchr/testify/assert"
)

func TestErrorCreation(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		errcstm := errcustom.NewErrorResp()
		errcstm.AddToErrList("error1")
		errcstm.AddToErrList("error2")
		errcstm.AddToErrList("error3")

		err := errcstm.ToError()

		assert.Error(t, err)
		assert.Equal(t, &errcustom.ErrorResp{
			ErrList: []string{
				"error1",
				"error2",
				"error3",
			},
		}, err)
	})

	t.Run("Test2", func(t *testing.T) {
		errcstm := errcustom.NewErrorResp()
		errcstm.AddToErrList("error1")
		errcstm.AddToErrList("error2")
		errcstm.AddToErrList("error3")

		err := errcstm.ToErrorAsString()

		assert.Error(t, err)
		assert.Equal(t, "err: error1,error2,error3", err.Error())
	})
}
