package baseConst

const (
	PersonGenderMale   = 1
	PersonGenderFemale = 2
)

const (
// PersonAgeGroupTeenager   = "teenager"
)

const (
	PersonAgeGroupYoungAdult = "youngAdult"
	PersonAgeGroupMiddleAged = "middleAged"
	PersonAgeGroupSenior     = "senior"
)

type PersonAgeGroup struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	PersonAgeGroupList = []*PersonAgeGroup{
		/*{
			Code: PersonAgeGroupTeenager,
			Name: "少年",
		},*/
		{
			Code: PersonAgeGroupYoungAdult,
			Name: "青年",
		},
		{
			Code: PersonAgeGroupMiddleAged,
			Name: "中年",
		},
		{
			Code: PersonAgeGroupSenior,
			Name: "老年",
		},
	}
)

const (
	PersonBodyTypeSlim       = "slim"
	PersonBodyTypeAverage    = "average"
	PersonBodyTypeOverweight = "overweight"
)

type PersonBodyType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	PersonBodyTypeList = []*PersonBodyType{
		{
			Code: PersonBodyTypeSlim,
			Name: "偏瘦",
		},
		{
			Code: PersonBodyTypeAverage,
			Name: "标准",
		},
		{
			Code: PersonBodyTypeOverweight,
			Name: "偏胖",
		},
	}
)

type PersonHairStyle struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	PersonHairStyleList = []*PersonHairStyle{
		{
			Code: "short",
			Name: "短发",
		},
		{
			Code: "medium",
			Name: "中发",
		},
		{
			Code: "shoulderLength",
			Name: "长发",
		},
	}
)

type PersonHairColor struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var (
	PersonHairColorList = []*PersonHairColor{
		{
			Code: "black",
			Name: "黑色",
		},
		{
			Code: "brown",
			Name: "棕色",
		},
		{
			Code: "blonde",
			Name: "金色",
		},
		{
			Code: "gray",
			Name: "灰色",
		},
	}
)
