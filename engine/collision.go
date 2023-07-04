package engine

type Point struct {
	X, Y float64
}

type Object struct {
	position Point
	space    *Space
	Width    float64
	Height   float64
	Targets  []string
}

type Collision struct {
	checkingObject *Object
	dx, dy         float64
	Objects        []*Object
}

func (o *Object) Check(x, y float64) *Collision {
	tmp := o
	tmp.position.X = x
	tmp.position.Y = y

	if o.space != nil {
		for _, object := range o.space.Objects {
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
		objects[ind].space = s
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
		position: Point{x, y},
		Width:    w,
		Height:   h,
	}
	if len(target) != 0 {
		o.Targets = target
	}

	return o
}

func checkCollision(obj1, obj2 Object) bool {
	if obj1.position.X < obj2.position.X+obj2.Width &&
		obj1.position.X+obj1.Width > obj2.position.X &&
		obj1.position.Y < obj2.position.Y+obj2.Height &&
		obj1.position.Y+obj1.Height > obj2.position.Y {
		return true
	}
	return false
}
