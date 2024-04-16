package baseConst

const (
	ClothingGenderMale    = 1
	ClothingGenderFemale  = 2
	ClothingGenderNeutral = 3
)

const (
	ClothingCategoryTop    = "top"
	ClothingCategoryBottom = "bottom"
	ClothingCategoryDress  = "dress"
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
			Code: ClothingCategoryDress,
			Name: "连衣裙",
		},
	}
)

type ClothingStyle struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ClothingTopStyleList = []*ClothingStyle{
		{
			Code: "tShirt",
			Name: "T恤",
		},
		{
			Code: "casualShirt",
			Name: "休闲衬衫",
		},
		{
			Code: "dressShirt",
			Name: "正装衬衫",
		},
		{
			Code: "poloShirt",
			Name: "polo衫",
		},
		{
			Code: "sweatshirt",
			Name: "卫衣",
		},
		{
			Code: "hoodie",
			Name: "连帽衫",
		},
		{
			Code: "sweater",
			Name: "毛衣",
		},
		{
			Code: "tankTop",
			Name: "背心",
		},
		{
			Code: "casualJacket",
			Name: "休闲夹克",
		},
		{
			Code: "suitJacket",
			Name: "西装外套",
		},
		{
			Code: "coat",
			Name: "外套",
		},
		{
			Code: "windbreaker",
			Name: "防风衣",
		},
		{
			Code: "cardigan",
			Name: "开衫",
		},
	}
)

var (
	ClothingBottomStyleList = []*ClothingStyle{
		{
			Code: "jeans",
			Name: "牛仔裤",
		},
		{
			Code: "casualPants",
			Name: "休闲裤",
		},
		{
			Code: "dressPants",
			Name: "正装裤",
		},
		{
			Code: "shorts",
			Name: "短裤",
		},
		{
			Code: "skirt",
			Name: "裙子",
		},
		{
			Code: "sweatpants",
			Name: "运动裤",
		},
		{
			Code: "leggings",
			Name: "打底裤",
		},
		{
			Code: "yogaPants",
			Name: "瑜伽裤",
		},
		{
			Code: "cargoPants",
			Name: "工装裤",
		},
		{
			Code: "culottes",
			Name: "阔腿裤",
		},
	}
)

var (
	ClothingDressStyleList = []*ClothingStyle{
		{
			Code: "miniDress",
			Name: "迷你裙",
		},
		{
			Code: "midiDress",
			Name: "迷笛裙",
		},
		{
			Code: "maxiDress",
			Name: "长裙",
		},
		{
			Code: "partyDress",
			Name: "派对礼服",
		},
		{
			Code: "casualDress",
			Name: "休闲连衣裙",
		},
	}
)

type ClothingMaterial struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ClothingMaterialList = []*ClothingMaterial{
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

type ClothingColor struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	ClothingColorList = []*ClothingColor{
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
