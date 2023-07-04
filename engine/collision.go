package engine

type Point struct {
	X, Y float64
}

type Object struct {
	Position Point
	Space    *Space
	Width    float64
	Height   float64
	Targets  []string
}

type Collision struct {
	checkingObject *Object
	dx, dy         float64
	Objects        []*Object
}

func (o *Object) Tags() []string {
	return o.Targets
}

func (o *Object) Check(x, y float64) *Collision {
	tmp := o
	tmp.Position.X = x
	tmp.Position.Y = y

	if o.Space != nil {
		for _, object := range o.Space.Objects {
			item := object
			if object == o { // 同一个元素忽略
				continue
			}
			// 判断是否满足碰撞需求
			if checkCollision(*tmp, *item) {
				return &Collision{
					checkingObject: tmp,
					dx:             x,
					dy:             y,
					Objects:        []*Object{object},
				}
			}
		}
	}

	return nil
}

type Space struct {
	Width, Height float64
	Objects       []*Object
}

func (s *Space) AddObject(objects ...*Object) {
	s.Objects = append(s.Objects, objects...)
	for ind := range objects {
		objects[ind].Space = s
	}
}

func NewSpace(w, h float64) *Space {
	return &Space{
		Width:   w,
		Height:  h,
		Objects: make([]*Object, 0),
	}
}

func NewObject(x, y, w, h float64, target ...string) *Object {
	o := &Object{
		Position: Point{x, y},
		Width:    w,
		Height:   h,
	}
	if len(target) != 0 {
		o.Targets = target
	}

	return o
}

func checkCollision(obj1, obj2 Object) bool {
	if obj1.Position.X < obj2.Position.X+obj2.Width &&
		obj1.Position.X+obj1.Width > obj2.Position.X &&
		obj1.Position.Y < obj2.Position.Y+obj2.Height &&
		obj1.Position.Y+obj1.Height > obj2.Position.Y {
		return true
	}
	return false
}

func SeparateRectangles(rect1, rect2 *Object) {
	dx := (rect1.Position.X + rect1.Width/2) - (rect2.Position.X + rect2.Width/2)
	dy := (rect1.Position.Y + rect1.Height/2) - (rect2.Position.Y + rect2.Height/2)

	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	if dx > dy {
		if rect1.Position.X < rect2.Position.X {
			rect1.Position.X -= dx
			// rect2.Position.X += dx
		} else {
			rect1.Position.X += dx
			// rect2.Position.X -= dx
		}
	} else {
		if rect1.Position.Y < rect2.Position.Y {
			rect1.Position.Y -= dy
			// rect2.Position.Y += dy
		} else {
			rect1.Position.Y += dy
			// rect2.Position.Y -= dy
		}
	}
}
