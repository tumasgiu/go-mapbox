package styles_test

import (
	"github.com/ryankurte/go-mapbox/lib/base"
	. "github.com/ryankurte/go-mapbox/lib/styles"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestStyles(t *testing.T) {
	b, err := base.NewBase(os.Getenv("MAPBOX_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}

	mapboxStyles := NewStyles(b)

	//t.Run("Create", func(t *testing.T) {
	//	style := Style{}
	//	mapboxStyles.Create(username, style)
	//})

	t.Run("List", func(t *testing.T) {
		styles, err := mapboxStyles.List("majidshabir", nil)
		if err != nil {
			t.Error(err)
		}

		logrus.Debug(styles)
	})


	t.Run("Retrieve", func(t *testing.T) {
		style, err := mapboxStyles.Retrieve("majidshabir", "ck3pnytqh4h441ckyvch3pfo7")
		if err != nil {
			t.Error(err)
		}

		logrus.Debug(style)
	})

}
