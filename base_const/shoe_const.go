package baseConst

type ShoeModel struct {
	ModelDid  string `json:"modelDid"`
	ModelName string `json:"modelName"`
}

const (
	ShoeGenderMale    = 1
	ShoeGenderFemale  = 2
	ShoeGenderNeutral = 3
)

const (
	ShoeCategoryLowTop  = "lowTop"
	ShoeCategoryMidTop  = "midTop"
	ShoeCategoryHighTop = "highTop"
)

type ShoeCategory struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ShoeCategoryList = []*ShoeCategory{
		{
			Code: ShoeCategoryLowTop,
			Name: "低帮",
		},
		{
			Code: ShoeCategoryMidTop,
			Name: "中帮",
		},
		{
			Code: ShoeCategoryHighTop,
			Name: "高帮",
		},
	}
)

type ShoeSole struct {
	Code     string `json:"code"`
	ResPath  string `json:"resPath"`
	MaskPath string `json:"maskPath"`
}

type ShoeStyle struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ShoeDefaultStyle = &ShoeStyle{
		Code: "default",
		Name: "默认",
	}
)

var (
	LowTopStyleList = []*ShoeStyle{
		ShoeDefaultStyle,
		{
			Code: "basketballSneaker",
			Name: "篮球鞋",
		},
		{
			Code: "runningShoe",
			Name: "跑鞋",
		},
		{
			Code: "casualShoe",
			Name: "休闲鞋",
		},
		{
			Code: "tennisShoe",
			Name: "网球鞋",
		},
		{
			Code: "trainingShoe",
			Name: "训练鞋",
		},
		{
			Code: "skateShoe",
			Name: "滑板鞋",
		},
		{
			Code: "sandal",
			Name: "凉鞋",
		},
		{
			Code: "slipper",
			Name: "拖鞋",
		},
		{
			Code: "boot",
			Name: "靴子",
		},
	}

	MidTopStyleList = []*ShoeStyle{
		ShoeDefaultStyle,
		{
			Code: "basketballSneaker",
			Name: "篮球鞋",
		},
		{
			Code: "runningShoe",
			Name: "跑鞋",
		},
		{
			Code: "casualShoe",
			Name: "休闲鞋",
		},
		{
			Code: "tennisShoe",
			Name: "网球鞋",
		},
		{
			Code: "trainingShoe",
			Name: "训练鞋",
		},
		{
			Code: "skateShoe",
			Name: "滑板鞋",
		},
		{
			Code: "hikingBoot",
			Name: "登山靴",
		},
		{
			Code: "boot",
			Name: "靴子",
		},
	}

	HighTopStyleList = []*ShoeStyle{
		ShoeDefaultStyle,
		{
			Code: "basketballSneaker",
			Name: "篮球鞋",
		},
		{
			Code: "runningShoe",
			Name: "跑鞋",
		},
		{
			Code: "casualShoe",
			Name: "休闲鞋",
		},
		{
			Code: "tennisShoe",
			Name: "网球鞋",
		},
		{
			Code: "trainingShoe",
			Name: "训练鞋",
		},
		{
			Code: "skateShoe",
			Name: "滑板鞋",
		},
		{
			Code: "hikingBoot",
			Name: "登山靴",
		},
		{
			Code: "boot",
			Name: "靴子",
		},
	}
)

func RemoveShoeDupStyleList(list []*ShoeStyle) []*ShoeStyle {
	result := make([]*ShoeStyle, 0)
	codeMap := make(map[string]bool)

	for _, style := range list {
		if _, ok := codeMap[style.Code]; !ok {
			codeMap[style.Code] = true
			result = append(result, style)
		}
	}

	return result
}

type ShoeMaterial struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ShoeDefaultMaterial = &ShoeMaterial{
		Code: "default",
		Name: "默认",
	}
)
var (
	ShoeMaterialList = []*ShoeMaterial{
		ShoeDefaultMaterial,
		{
			Code: "leather",
			Name: "皮革",
		},
		{
			Code: "genuineLeather",
			Name: "真皮",
		},
		{
			Code: "syntheticLeather",
			Name: "合成皮革",
		},
		{
			Code: "canvas",
			Name: "帆布",
		},
		{
			Code: "rubber",
			Name: "橡胶",
		},
		{
			Code: "mesh",
			Name: "网布",
		},
		{
			Code: "suede",
			Name: "绒面革",
		},
		{
			Code: "nylon",
			Name: "尼龙",
		},
		{
			Code: "denim",
			Name: "牛仔布",
		},
		{
			Code: "other",
			Name: "其他",
		},
	}
)

func RemoveShoeDupMaterialList(list []*ShoeMaterial) []*ShoeMaterial {
	result := make([]*ShoeMaterial, 0)
	codeMap := make(map[string]bool)

	for _, style := range list {
		if _, ok := codeMap[style.Code]; !ok {
			codeMap[style.Code] = true
			result = append(result, style)
		}
	}

	return result
}

type ShoeColor struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ShoeDefaultColor = &ShoeColor{
		Code: "default",
		Name: "默认",
	}
)

var (
	ShoeColorList = []*ShoeColor{
		ShoeDefaultColor,
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
		{
			Code: "multi",
			Name: "多色",
		},
		{
			Code: "other",
			Name: "其他",
		},
	}
)

func RemoveShoeDupColorList(list []*ShoeColor) []*ShoeColor {
	result := make([]*ShoeColor, 0)
	codeMap := make(map[string]bool)

	for _, style := range list {
		if _, ok := codeMap[style.Code]; !ok {
			codeMap[style.Code] = true
			result = append(result, style)
		}
	}

	return result
}
