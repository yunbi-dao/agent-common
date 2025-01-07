package baseConst

type SpaceCategory struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

var (
	SpaceCategoryList = []*SpaceCategory{
		{
			Code: "space cabin homestay",
			Name: "太空舱民宿",
			Desc: "未来感十足的胶囊式住宿，提供独特的入住体验",
		},
		{
			Code: "mobile dining cart",
			Name: "移动餐车",
			Desc: "灵活设计的餐车，可用于流动餐饮服务",
		},
		{
			Code: "mobile shop",
			Name: "移动商铺",
			Desc: "便携式零售设计，提供灵活的购物解决方案",
		},
		{
			Code: "business class",
			Name: "商业舱",
			Desc: "专为商务或专业用途设计的高端空间",
		},
		{
			Code: "container shops",
			Name: "集装箱商铺",
			Desc: "由改装的集装箱建成的商铺",
		},
		{
			Code: "yacht cabin",
			Name: "游艇舱",
			Desc: "游艇上的奢华舱室体验",
		},
		{
			Code: "boat house",
			Name: "船屋",
			Desc: "建在船上或类似船结构的房屋",
		},
		{
			Code: "triangle house",
			Name: "三角屋",
			Desc: "一种独特的三角形房屋，具有独特的设计感",
		},
		{
			Code: "thatched cottage",
			Name: "茅草屋",
			Desc: "一种以茅草为屋顶的乡村风格房屋，散发传统魅力",
		},
		{
			Code: "holiday villa",
			Name: "度假别墅",
			Desc: "为度假和放松设计的豪华别墅",
		},
		{
			Code: "camping car",
			Name: "露营车",
			Desc: "配备舒适设施的露营专用车辆",
		},
		{
			Code: "travel trailer",
			Name: "拖挂式房车",
			Desc: "可拖挂的移动居住空间",
		},
	}
)

type SpaceColor struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	SpaceDefaultColor = &SpaceColor{
		Code: "default",
		Name: "默认",
	}
)

var (
	SpaceColorList = []*SpaceColor{
		SpaceDefaultColor,
		{
			Code: "black",
			Name: "黑色",
		},
		{
			Code: "white",
			Name: "白色",
		},
		{
			Code: "brown",
			Name: "棕色",
		},
		{
			Code: "gray",
			Name: "灰色",
		},
		{
			Code: "blue",
			Name: "蓝色",
		},
		{
			Code: "red",
			Name: "红色",
		},
		{
			Code: "green",
			Name: "绿色",
		},
		{
			Code: "beige",
			Name: "米色",
		},
		{
			Code: "pink",
			Name: "粉色",
		},
		{
			Code: "purple",
			Name: "紫色",
		},
		{
			Code: "yellow",
			Name: "黄色",
		},
		{
			Code: "orange",
			Name: "橙色",
		},
	}
)
