package baseConst

const (
	MqHandleShoeModelTagTopic    = "com.vt.service.mq.handle.shoe.model.tag.topic"
	MqHandleShoeModelTagGroup    = "com.vt.service.mq.handle.shoe.model.tag.group"
	MqHandleShoeModelTagConsumer = "com.vt.service.mq.handle.shoe.model.tag.consumer"
)

const (
	MqHandleShoeModelTrainTopic    = "com.vt.service.mq.handle.shoe.model.train.topic"
	MqHandleShoeModelTrainGroup    = "com.vt.service.mq.handle.shoe.model.train.group"
	MqHandleShoeModelTrainConsumer = "com.vt.service.mq.handle.shoe.model.train.consumer"
)

const (
	MqHandleShoeModelPublishTopic    = "com.vt.service.mq.handle.shoe.model.publish.topic"
	MqHandleShoeModelPublishGroup    = "com.vt.service.mq.handle.shoe.model.publish.group"
	MqHandleShoeModelPublishConsumer = "com.vt.service.mq.handle.shoe.model.publish.consumer"
)

const (
	MqHandleClothingModelTagTopic    = "com.vt.service.mq.handle.clothing.model.tag.topic"
	MqHandleClothingModelTagGroup    = "com.vt.service.mq.handle.clothing.model.tag.group"
	MqHandleClothingModelTagConsumer = "com.vt.service.mq.handle.clothing.model.tag.consumer"
)

const (
	MqHandleClothingModelTrainTopic    = "com.vt.service.mq.handle.clothing.model.train.topic"
	MqHandleClothingModelTrainGroup    = "com.vt.service.mq.handle.clothing.model.train.group"
	MqHandleClothingModelTrainConsumer = "com.vt.service.mq.handle.clothing.model.train.consumer"
)

const (
	MqHandleClothingModelPublishTopic    = "com.vt.service.mq.handle.clothing.model.publish.topic"
	MqHandleClothingModelPublishGroup    = "com.vt.service.mq.handle.clothing.model.publish.group"
	MqHandleClothingModelPublishConsumer = "com.vt.service.mq.handle.clothing.model.publish.consumer"
)

// 处理状态: 20.清洗待处理 21.清洗中 22.清洗失败 23.清洗成功 30.打标待处理 31.打标中 32.打标失败 33.打标成功 40.训练待处理 41.训练中 42.训练失败 43.训练成功 50.发布待处理 51.发布中 52.发布失败 53.发布成功
var (
	ModelHandleCleanPending   = &ModelHandleStatus{Status: 20, Msg: "清洗待处理"}
	ModelHandleCleanProcess   = &ModelHandleStatus{Status: 21, Msg: "清洗中"}
	ModelHandleCleanFailed    = &ModelHandleStatus{Status: 22, Msg: "清洗失败"}
	ModelHandleCleanSuccess   = &ModelHandleStatus{Status: 23, Msg: "清洗成功"}
	ModelHandleTagPending     = &ModelHandleStatus{Status: 30, Msg: "打标待处理"}
	ModelHandleTagProcess     = &ModelHandleStatus{Status: 31, Msg: "打标中"}
	ModelHandleTagFailed      = &ModelHandleStatus{Status: 32, Msg: "打标失败"}
	ModelHandleTagSuccess     = &ModelHandleStatus{Status: 33, Msg: "打标成功"}
	ModelHandleTrainPending   = &ModelHandleStatus{Status: 40, Msg: "训练待处理"}
	ModelHandleTrainProcess   = &ModelHandleStatus{Status: 41, Msg: "训练中"}
	ModelHandleTrainFailed    = &ModelHandleStatus{Status: 42, Msg: "训练失败"}
	ModelHandleTrainSuccess   = &ModelHandleStatus{Status: 43, Msg: "训练成功"}
	ModelHandlePublishPending = &ModelHandleStatus{Status: 50, Msg: "发布待处理"}
	ModelHandlePublishProcess = &ModelHandleStatus{Status: 51, Msg: "发布中"}
	ModelHandlePublishFailed  = &ModelHandleStatus{Status: 52, Msg: "发布失败"}
	ModelHandlePublishSuccess = &ModelHandleStatus{Status: 53, Msg: "发布成功"}
)

type ModelHandleStatus struct {
	Status int
	Msg    string
}

func GetModelHandleStatus(handleStatus int) *ModelHandleStatus {
	switch handleStatus {
	case ModelHandleCleanPending.Status:
		return ModelHandleCleanPending
	case ModelHandleCleanProcess.Status:
		return ModelHandleCleanProcess
	case ModelHandleCleanFailed.Status:
		return ModelHandleCleanFailed
	case ModelHandleCleanSuccess.Status:
		return ModelHandleCleanSuccess
	case ModelHandleTagPending.Status:
		return ModelHandleTagPending
	case ModelHandleTagProcess.Status:
		return ModelHandleTagProcess
	case ModelHandleTagFailed.Status:
		return ModelHandleTagFailed
	case ModelHandleTagSuccess.Status:
		return ModelHandleTagSuccess
	case ModelHandleTrainPending.Status:
		return ModelHandleTrainPending
	case ModelHandleTrainProcess.Status:
		return ModelHandleTrainProcess
	case ModelHandleTrainFailed.Status:
		return ModelHandleTrainFailed
	case ModelHandleTrainSuccess.Status:
		return ModelHandleTrainSuccess
	case ModelHandlePublishPending.Status:
		return ModelHandlePublishPending
	case ModelHandlePublishProcess.Status:
		return ModelHandlePublishProcess
	case ModelHandlePublishFailed.Status:
		return ModelHandlePublishFailed
	case ModelHandlePublishSuccess.Status:
		return ModelHandlePublishSuccess
	default:
		return nil
	}
}
