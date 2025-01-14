# ARPostcard_server

## 项目结构
```text
- main.go 项目入口
- conf 配置文件文件夹
- biz 业务逻辑文件夹
    - conf 配置，里面存储用于配置初始化的函数
    - consts 常量
    - dao 数据访问对象
    - model 数据模型
    - service 服务
    - handler 处理器，用于存储各个路由实际调用的函数
    - infra 基础设施，用于存储各种基础设施，包括数据库等配置
    - mw 中间件
    - utils 工具
- run_log 运行日志
```

## 项目启动
1. 安装项目所需要的包，命令行中使用 `go mod tidy`
2. 修改 biz/conf/enter.go 中的常量，让 ConfigFile 指向正确的配置文件
3. 运行项目，命令行中使用 `go run main.go`，或者使用IDE中的运行按钮

## 开发思路以及时间节点
### 关于系统接下来的设计（2025-01-13）
1. 增加 prefab 管理页面，prefab 提交之后可以直接在页面上查看（在前端采用three.js设置一个预览系统，然后可以查看）
2. 这种素材分为两种，一种是可以设置图片的3D模型，一种是角色模型，然后所有角色模型绑定一套骨骼动作，当用户使用APP的AR功能，可以让角色动起来，另外一种是可以设置图片的模型，这种模型可以设置图片，然后用户可以在APP中查看这个模型（这个还需要考虑）
3. 用户设置AR图片流程重新设计：（多个步骤进行定制）
```text
设计了新的Image类型
// Image 识别图
type Image struct {
	MODEL
	ImageID   string `gorm:"size:36" json:"image_id"`   // 图片ID，这个是EasyAR返回的图片ID
	ImageURL  string `gorm:"size:128" json:"image_url"` // 图片URL，存储到对象存储服务当中的URL
	ImageName string `gorm:"size:64" json:"image_name"` // 图片名称
	VisionType int    `gorm:"size:1" json:"vision_type"` // 图片类型，1代表视频，2代表角色模型，3代表其他模型
	// 需要添加一个字段的名称，这个字段用于存储用户选择的模型ID或者是视频ID，或者是其他模型ID
	ModelID string `gorm:"size:36" json:"model_id"` // 模型ID，这个是用户选择的模型ID或者视频ID或者其他模型ID，根据VisionType不同在不同的表中查找
}

1. 用户上传图片（此时数据库会存储这个图片，然后图片此时已经可以被识别了），图片大小和其他的进行设置
2. 用户选择显示的模型类型（视频/角色模型/其他模型），根据用户上一步不同的选择，进入不同的页面，（图片增加一个新的字段为visionType，1代表视频，2代表角色模型，3代表其他模型）视频则进行步骤3，角色模型则进行步骤4，其他模型（可以设置图片）则进行步骤5
3. 如果用户进入的是视频选择，此时用户可以选择一个视频上传（此时数据库也会存储这个video，并且更新图片的ModelID字段，video也需要对应起来UserID）
4. 如果用户进入的是角色模型选择，此时用户可以选择一个已经预设好的角色模型（角色模型从后端获取，可以自己选择，提交后更新数据库里面的ModelID字段）
5. 如果用户进入的是其他模型选择，此时用户可以选择一个已经预设好的模型（模型从后端获取，可以自己选择，提交后更新数据库里面的ModelID字段），然后可以提交一些图片（这里还需要设置一个新的表来存储用户提交的新的图片）（暂时搁置，没有那么重要）
```
4. 图片识别加载模型阶段流程重新设计
```text
1. EasyAR识别处理图片，得到ImageID，然后根据ImageID在数据库中查找对应的Image，然后根据Image的ModelID字段查找对应的模型信息，这里还是要数据结构化，数据设计如下
{
    "image_id": "xxxx",
    "image_url": "xxxx",
    "image_name": "xxxx",
    "vision_type": 1,
    "model_id": "xxxx",
    "model_info":{
        "model_id": "xxxx",
        "model_name": "xxxx", (如果是视频的话，这里就是默认值是video，然后下面url就是视频的URL，如果是角色模型的话，这里就是角色模型的名字，然后下面的url就是角色模型的URL)
        "model_url": "xxxx",
        "data": ["xxxx", "xxxx"], (如果是视频或者角色模型的话，这里就是空数组，如果是其他模型的话，这里就是图片的URL列表)
    }
}
2. 如果vision_type是1，那么unity端直接加载视频（根据url）进去，如果vision_type是2，那么unity端加载角色模型（根据url），然后绑定固定的骨骼动作，如果vision_type是3，那么unity端加载其他模型（根据name或者是url），然后绑定图片（根据data）
```