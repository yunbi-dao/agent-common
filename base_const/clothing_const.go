package baseConst

type ClothingModel struct {
	ModelDid  string `json:"modelDid"`
	ModelName string `json:"modelName"`
}

const (
	ClothingGenderMale    = 1
	ClothingGenderFemale  = 2
	ClothingGenderNeutral = 3
)

const (
	ClothingCategoryTop      = "top"
	ClothingCategoryBottom   = "bottom"
	ClothingCategoryEnsemble = "ensemble"
	ClothingCategoryDress    = "dress"
)

type ClothingCategory struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ClothingCategoryList = []*ClothingCategory{
		{
			Code: ClothingCategoryTop,
			Name: "上装",
		},
		{
			Code: ClothingCategoryBottom,
			Name: "下装",
		},
		{
			Code: ClothingCategoryEnsemble,
			Name: "套装",
		},
		{
			Code: ClothingCategoryDress,
			Name: "连衣裙",
		},
	}
)

var (
	ClothingViewFront = "front"
	ClothingViewBack  = "back"
	ClothingViewSide  = "side"
)

var (
	ClothingViewList = []*ClothingView{
		{
			Code: ClothingViewFront,
			Name: "正面",
		},
		{
			Code: ClothingViewBack,
			Name: "背面",
		},
		{
			Code: ClothingViewSide,
			Name: "侧面",
		},
	}
)

type ClothingView struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ClothingMaterial struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ClothingDefaultMaterial = &ClothingMaterial{
		Code: "default",
		Name: "默认",
	}
)

var (
	ClothingMaterialList = []*ClothingMaterial{
		ClothingDefaultMaterial,
		{
			Code: "cotton",
			Name: "棉",
		},
		{
			Code: "silk",
			Name: "丝绸",
		},
		{
			Code: "wool",
			Name: "羊毛",
		},
		{
			Code: "linen",
			Name: "亚麻",
		},
		{
			Code: "polyester",
			Name: "涤纶",
		},
		{
			Code: "nylon",
			Name: "尼龙",
		},
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
			Code: "denim",
			Name: "牛仔布",
		},
		{
			Code: "lace",
			Name: "蕾丝",
		},
		{
			Code: "other",
			Name: "其他",
		},
	}
)

func RemoveClothingDupMaterialList(list []*ClothingMaterial) []*ClothingMaterial {
	result := make([]*ClothingMaterial, 0)
	codeMap := make(map[string]bool)

	for _, style := range list {
		if _, ok := codeMap[style.Code]; !ok {
			codeMap[style.Code] = true
			result = append(result, style)
		}
	}

	return result
}

type ClothingColor struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ClothingDefaultColor = &ClothingColor{
		Code: "default",
		Name: "默认",
	}
)

var (
	ClothingColorList = []*ClothingColor{
		ClothingDefaultColor,
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

func RemoveClothingDupColorList(list []*ClothingColor) []*ClothingColor {
	result := make([]*ClothingColor, 0)
	codeMap := make(map[string]bool)

	for _, style := range list {
		if _, ok := codeMap[style.Code]; !ok {
			codeMap[style.Code] = true
			result = append(result, style)
		}
	}

	return result
}
