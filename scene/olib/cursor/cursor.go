package cursor

import "github.com/sparkoo/sparkengine/scene"

type Cursor struct {
	*scene.Base
}

func (c *Cursor) GetPixels() []scene.Pixel {
	panic("implement me")
}
