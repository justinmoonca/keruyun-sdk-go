package kry_standard

type ShopStoreDetailResp struct {
	AddressProvince string `json:"addressProvince"` // 省
	Latitude        string `json:"latitude"`        // 纬度
	OrgMode         string `json:"orgMode"`         // 机构类型：单店SINGLE,连锁CHAIN
	ViceMeals       string `json:"viceMeals"`       // 附加业态
	OrgType         string `json:"orgType"`         // 机构类别
	AddressDetail   string `json:"addressDetail"`   // 地址详情
	ShopId          int64  `json:"shopId"`          // 客如云门店id
	CityName        string `json:"cityName"`        // 市
	AreaName        string `json:"areaName"`        // 区
	Contact         string `json:"contact"`         // 联系人
	Longitude       string `json:"longitude"`       // 经度
	AddressCity     string `json:"addressCity"`     // 城市
	MainMeal        string `json:"mainMeal"`        // 主营业态
	ContactMobile   string `json:"contactMobile"`   // 联系电话
	DetailPictures  string `json:"detailPictures"`  // 门店图片地址
	Name            string `json:"name"`            // 门店名称
	ServiceMobile   string `json:"serviceMobile"`   // 服务电话
	ProvinceName    string `json:"provinceName"`    // 省
	AddressArea     string `json:"addressArea"`     // 地区
	Status          string `json:"status"`          // 门店状态：正常NORMAL,禁用DISABLE
	ServerCode      int    `json:"serverCode"`      // 服务编码
	MessageId       string `json:"messageId"`       // 服务消息ID
}

type ShopTableListResp struct {
	PageNum   int `json:"pageNum"`   // 当前分页页数
	PageSize  int `json:"pageSize"`  // 每页条目数
	TotalNum  int `json:"totalNum"`  // 总条目数
	TotalPage int `json:"totalPage"` // 总页数
	TableList []struct {
		TableId          string `json:"tableId"`          // 桌台ID
		TableName        string `json:"tableName"`        // 桌台名称
		AreaId           string `json:"areaId"`           // 桌台所属区域ID
		AreaName         string `json:"areaName"`         // 桌台所属区域名称
		TablePersonCount int    `json:"tablePersonCount"` // 桌台用餐人数
	} `json:"tableList"` // 桌台列表
	MessageId  string `json:"messageId"`  // 服务消息ID
	ServerCode int    `json:"serverCode"` // 服务编码
}

type ShopServiceFeeListResp struct {
	MessageId       string `json:"messageId"` // 服务消息ID
	ExtraChargeList []struct {
		ExtraChargeId          string `json:"extraChargeId"`          // 服务费ID
		Name                   string `json:"name"`                   // 服务费名称
		ExtraChargeType        string `json:"extraChargeType"`        // 服务费类型
		CalcWay                string `json:"calcWay"`                // 计算方式,PERCENT按消费比例收取,PERSON按用餐人数收取,FIXED按固定金额收取
		CalcAmount             int    `json:"calcAmount"`             // 计算数额:消费比例/固定金额/每人金额
		AutoAddOrderFlag       string `json:"autoAddOrderFlag"`       // 是否自动加入订单，ON是，OFF否
		DiscountAfterCalcFlag  string `json:"discountAfterCalcFlag"`  // 是否优惠后计算，ON是，OFF否
		AllowDiscountShareFlag string `json:"allowDiscountShareFlag"` // 是否参与优惠分摊，ON是，OFF否
		AllowDiscountFlag      string `json:"allowDiscountFlag"`      // 是否参与折扣，ON是，OFF否
		EnabledFlag            string `json:"enabledFlag"`            // 启用状态，ON启用，OFF禁用
	} `json:"extraChargeList"` // 服务费列表
	ServerCode int `json:"serverCode"` // 服务编码
}

type ShopTableInfoResp struct {
	MessageId        string `json:"messageId"`        // 服务消息ID
	TableId          string `json:"tableId"`          // 桌台ID
	TableName        string `json:"tableName"`        // 桌台名称
	AreaId           string `json:"areaId"`           // 桌台所属区域ID
	AreaName         string `json:"areaName"`         // 桌台所属区域名称
	TablePersonCount int    `json:"tablePersonCount"` // 桌台用餐人数
	ServerCode       int    `json:"serverCode"`       // 服务编码
}

type DishPageResp struct {
	Value struct {
		Total    int `json:"total"` // 数据总数
		DataList []struct {
			DishId     string `json:"dishId"`     // 菜品ID
			DishName   string `json:"dishName"`   // 菜品名称
			DishDesc   string `json:"dishDesc"`   // 菜品描述
			CategoryId string `json:"categoryId"` // 分类ID
			Sort       int    `json:"sort"`       // 排序值
			HelpCode   string `json:"helpCode"`   // 助记码
			DishType   string `json:"dishType"`   // 菜品类型， SINGLE：单菜 ，COMBO：套餐， SIDE：配料
			State      string `json:"state"`      // 菜品状态
			WeighFlag  string `json:"weighFlag"`  // 称重菜标识, Y:是，N:否
		} `json:"dataList"` // 具体数据
	} `json:"value"` // 业务结果
	Success    string `json:"success"`    // 是否成功, true:成功，false:失败
	ErrorDesc  string `json:"errorDesc"`  // 错误信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type DishDetailResp struct {
	Value []struct {
		DishId       string `json:"dishId"`       //菜品ID
		DishName     string `json:"dishName"`     //菜品名称
		DishCode     string `json:"dishCode"`     //菜品编码
		DishType     string `json:"dishType"`     //菜品类型。SINGLE：单菜 ，COMBO：套餐， SIDE：配料
		CategoryId   string `json:"categoryId"`   //菜品分类ID
		CategoryName string `json:"categoryName"` //菜品分类名称
		DishSkuList  []struct {
			SpecName       string `json:"specName"`       //规格名称
			DefaultSkuFlag string `json:"defaultSkuFlag"` //是否为默认规格
			SellPrice      int    `json:"sellPrice"`      //售卖价，单位：分
			BarCode        string `json:"barCode"`        //条形码
			SkuId          string `json:"skuId"`          //菜品SKU ID
			DishSkuCode    string `json:"dishSkuCode"`    //菜品SKU Code
		} `json:"dishSkuList"` //菜品SKU列表
		WeighFlag      string `json:"weighFlag"` //是否为称重菜 Y:是；N:否
		Weight         int    `json:"weight"`    //重量，单位：毫克
		UnitName       string `json:"unitName"`  //单位名称
		HelpCode       string `json:"helpCode"`  //助记码
		ComboGroupList []struct {
			GroupName            string `json:"groupName"`  //分组名称
			MaxChoose            int    `json:"maxChoose"`  //套餐分组子菜的最大选择数
			MinChoose            int    `json:"minChoose"`  //套餐分组子菜的最小选择数
			Sort                 int    `json:"sort"`       //排序值
			Repeatable           string `json:"repeatable"` //是否可重复选 Y:是；N:否
			ComboGroupDetailList []struct {
				SingleDishId  string `json:"singleDishId"`  //子菜菜品ID
				MaxChoose     int    `json:"maxChoose"`     //最大可选数量
				MinChoose     int    `json:"minChoose"`     //最小可选数量
				FixChoose     int    `json:"fixChoose"`     //固定数量
				DishSkuId     string `json:"dishSkuId"`     //子菜SKU ID
				DishSkuPrice  int    `json:"dishSkuPrice"`  //套餐分组为可选分组时的子菜加价金额,单位：分
				OptType       string `json:"optType"`       //可选类型(OPTIONAL-可选/REQUIRED-必选)
				DefaultFlag   string `json:"defaultFlag"`   //是否为默认子菜
				DishName      string `json:"dishName"`      //子菜名称
				SpecName      string `json:"specName"`      //子菜规格名
				SellPrice     int    `json:"sellPrice"`     //子菜售卖价，单位：分
				MultiSpecFlag string `json:"multiSpecFlag"` //子菜是否为多规格 Y:是，N:否
				Sort          int    `json:"sort"`          //排序号
			} `json:"comboGroupDetailList"` //套餐分组子菜列表
			GroupType string `json:"groupType"` //套餐组类型（FIXED:固定，OPTIONAL:可选）
		} `json:"comboGroupList"` //菜品为套餐时的套餐分组信息
		ComboPriceIncludeChildDishSideDishPrice   string   `json:"comboPriceIncludeChildDishSideDishPrice"`   //套餐价格是否包含子菜加料价格， Y:是；N:否
		ComboPriceIncludeChildDishCookingWayPrice string   `json:"comboPriceIncludeChildDishCookingWayPrice"` //套餐价格是否包含子菜做法价格, Y:是；N:否
		StatsLabelNameList                        []string `json:"statsLabelNameList"`                        //统计标签名称列表
		SaleLabelNameList                         []string `json:"saleLabelNameList"`                         //售卖标签名称列表
		RemarkNameList                            []string `json:"remarkNameList"`                            //备注名称列表
		CookingWayGroupList                       []struct {
			CookingWayGroupName string `json:"cookingWayGroupName"` //做法组名
			OptType             string `json:"optType"`             //可选类型(OPTIONAL-可选/REQUIRED-必选)
			CookingWayList      []struct {
				CookingWayName string `json:"cookingWayName"` //做法名称
				AddPrice       int    `json:"addPrice"`       //加价，单位：分
				DefaultFlag    string `json:"defaultFlag"`    //是否为默认
				CookingWayId   string `json:"cookingWayId"`   //做法ID
			} `json:"cookingWayList"` //做法列表
			CookingWayGroupId string `json:"cookingWayGroupId"` //做法组ID
		} `json:"cookingWayGroupList"` //做法组信息
		SideDishGroupList []struct {
			GroupName               string `json:"groupName"` //分组名称
			SideDishGroupDetailList []struct {
				SideDishName string `json:"sideDishName"` //加料名称
				SideDishId   string `json:"sideDishId"`   //加料ID
				AddPrice     int    `json:"addPrice"`     //配料价格 单位： 分
				Sort         int    `json:"sort"`         //排序值
			} `json:"sideDishGroupDetailList"` //加料分组明细列表
		} `json:"sideDishGroupList"` //加料组信息
		State             string   `json:"state"`            //售卖状态(ONLINE-售卖, PAUSE-停售)
		StartNumber       int      `json:"startNumber"`      //起售数量
		StartInterval     int      `json:"startInterval"`    //增售数量
		ModifyPriceFlag   string   `json:"modifyPriceFlag"`  //是否可以手动改价
		DiscountFlag      string   `json:"discountFlag"`     //是否可以手动打折
		OrderSingleFlag   string   `json:"orderSingleFlag"`  //是否可以单点
		DishImageUrlList  []string `json:"dishImageUrlList"` //菜品图片列表
		SpicyLevel        int      `json:"spicyLevel"`       //辣度等级
		Sort              int      `json:"sort"`             //排序值
		DishStockInfoList []struct {
			DishId          string `json:"dishId"`          //菜品SPU ID
			SkuId           string `json:"skuId"`           //菜品SKU ID
			ResidualDecimal int64  `json:"residualDecimal"` //剩余售卖量
			SoldOutFlag     string `json:"soldOutFlag"`     //是否售罄（Y：售罄，N：未售罄）
		} `json:"dishStockInfoList"` //菜品库存信息列表
		UnitId string `json:"unitId"` //菜品单位ID
	} `json:"value"` //业务数据
	Success    bool   `json:"success"`    //是否成功,true:成功,false:失败
	ErrorDesc  string `json:"errorDesc"`  //错误描述
	ServerCode int    `json:"serverCode"` //服务编码
	MessageId  string `json:"messageId"`  //服务消息ID
}

type DishCategoryResp struct {
	Value []struct {
		CategoryId   string `json:"categoryId"`   // 分类ID
		ParentId     string `json:"parentId"`     // 父级分类ID
		CategoryName string `json:"categoryName"` // 分类名称
		Sort         int    `json:"sort"`         // 排序值
		CategoryType string `json:"categoryType"` // 分类类型，DISH：菜品分类 SIDE_DISH_GROUP:加料分组
	} `json:"value"` // 业务结果
	Success    bool   `json:"success"`    // 是否成功, true:成功，false:失败
	ErrorDesc  string `json:"errorDesc"`  // 错误信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmCreateResp struct {
	Data struct {
		CustomerId string `json:"customerId"` // 用户id
	} `json:"data"`
	Success    bool   `json:"success"`    // 是否成功, true成功, false失败
	MsgCode    string `json:"msgCode"`    // success为false时的异常码
	MsgInfo    string `json:"msgInfo"`    // success为false时的异常信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmCustomerInfoResp struct {
	Data struct {
		Birthday   string `json:"birthday"`   // 生日日期,yyyy-MM-dd格式
		CustomerId string `json:"customerId"` // 用户id
		Gender     int    `json:"gender"`     // 用户性别,0代表女,1代表男,2代表其他
		LevelDTO   struct {
			LevelName string `json:"levelName"` // 等级名称
			LevelNo   int    `json:"levelNo"`   // 等级级别
		} `json:"levelDTO"` // 等级对象
		Mobile string `json:"mobile"` // 手机号
		Name   string `json:"name"`   // 用户名称
		State  int    `json:"state"`  // 用户状态,1代表启用,0代表停用
	} `json:"data"` // 数据体对象
	Success    bool   `json:"success"`    // 请求请求结果, true代表成功, false代表失败
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmCustomerPropertyResp struct {
	Data struct {
		CustomerDTO struct {
			CustomerId string `json:"customerId"` // 用户id
			State      int    `json:"state"`      // 用户状态
			Mobile     string `json:"mobile"`     // 手机号
		} `json:"customerDTO"` // 会员基础信息
		PosCardDTOList []struct {
			PosRechargeAccountList []struct {
				RemainAvailableValue struct {
					TotalValue int `json:"totalValue"` // 当前剩余可用储值总金额 = 当前剩余可用实储总金额 + 当前剩余可用赠储总金额，单位：分
					RealValue  int `json:"realValue"`  // 当前剩余可用实储总金额，单位：分
					GiftValue  int `json:"giftValue"`  // 当前剩余可用赠储总金额，单位：分
				} `json:"remainAvailableValue"` // 当前剩余可用储值
			} `json:"posRechargeAccountList"` // 储值账户
			Status   string `json:"status"`   // 卡状态，SOLD：已出售；ACTIVED：已激活；STOP：已停用；INVALID：已作废；EXPIRED：已过期；REFUND：已退卡
			CardId   string `json:"cardId"`   // 卡号id
			CardType string `json:"cardType"` // 卡类型，MEMBER：会员卡；GIFT：礼品卡；PAY_MEMBER：付费会员卡
		} `json:"posCardDTOList"` // 卡列表
		NormalVoucherInstanceCount int `json:"normalVoucherInstanceCount"` // 可使用有效券张数
		PointAccountDTO            struct {
			RemainAvailableValue int `json:"remainAvailableValue"`
		} `json:"pointAccountDTO"` // 积分账户
	} `json:"data"` // 数据体
	Success    bool   `json:"success"`    // 是否成功,true:成功,false:失败
	ServerCode int    `json:"serverCode"` // 服务端返回码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmDirectChargeResp struct {
	Success bool   `json:"success"` // 是否成功,true:成功，false:失败
	MsgCode string `json:"msgCode"` // success为false时的异常码
	MsgInfo string `json:"msgInfo"` // success为false时的异常信息
	Data    struct {
		AccountId        string `json:"accountId"` // 储值账户id
		RemainTotalValue struct {
			RealValue  int64 `json:"realValue"`  // 实储余额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 赠储余额，单位：分
			PreValue   int64 `json:"preValue"`   // 预储余额，单位：分
			TotalValue int64 `json:"totalValue"` // 总余额=实储余额 + 赠储余额 + 预储余额，单位：分
		} `json:"remainTotalValue"` // 储值账户余额 = 储值账户可用余额 + 储值账户预扣金额
		RemainAvailableValue struct {
			RealValue  int64 `json:"realValue"`  // 可用实储余额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 可用赠储余额，单位：分
			PreValue   int64 `json:"preValue"`   // 可用预储余额，单位：分
			TotalValue int64 `json:"totalValue"` // 可用总余额=可用实储余额 + 可用赠储余额 +可用预储余额，单位：分
		} `json:"remainAvailableValue"` // 储值账户可用余额
		PreDeductValue struct {
			RealValue  int64 `json:"realValue"`  // 预扣实储金额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 预扣赠储金额，单位：分
			PreValue   int64 `json:"preValue"`   // 预扣预储金额，单位：分
			TotalValue int64 `json:"totalValue"` // 预扣总金额=预扣实储金额 + 预扣赠储金额 + 预扣预储金额，单位：分
		} `json:"preDeductValue"` // 储值账户预扣金额
		TotalValue struct {
			RealValue  int64 `json:"realValue"`  // 累计实储金额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 累计赠储金额，单位：分
			PreValue   int64 `json:"preValue"`   // 累计预储金额，单位：分
			TotalValue int64 `json:"totalValue"` // 累计储值总金额=累计实储金额 +累计赠储金额 + 累计预储金额，单位：分
		} `json:"totalValue"` // 储值累计总额
	} `json:"data"` // 业务数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type OrderDetailResp struct {
	Success  bool   `json:"success"`  // 是否成功,true:成功，false:失败
	MsgCode  string `json:"msgCode"`  // success为false时的异常码
	MsgInfo  string `json:"msgInfo"`  // success为false时的异常信息
	CanRetry bool   `json:"canRetry"` // 可以重试
	Data     struct {
		OrderBaseVO struct {
			ShopId             string      `json:"shopId"`             // 门店ID
			ShopName           string      `json:"shopName"`           // 门店名称
			BrandId            string      `json:"brandId"`            // 品牌ID
			BrandName          string      `json:"brandName"`          // 品牌名称
			OrderId            string      `json:"orderId"`            // 订单ID
			BusiOrderNo        string      `json:"busiOrderNo"`        // 业务订单号
			ThirdOrderNo       string      `json:"thirdOrderNo"`       // 三方订单号
			OpenTime           string      `json:"openTime"`           // 下单时间
			SettleTime         string      `json:"settleTime"`         // 结账时间
			FinishBusiDate     string      `json:"finishBusiDate"`     // 营业日
			OrderSource        string      `json:"orderSource"`        // 订单来源
			OrderType          string      `json:"orderType"`          // 订单类型
			OrderStatus        string      `json:"orderStatus"`        // 订单状态
			OrderAmt           interface{} `json:"orderAmt"`           // 订单金额
			PromoAmt           string      `json:"promoAmt"`           // 优惠金额
			OrderReceivedAmt   string      `json:"orderReceivedAmt"`   // 订单收入
			OpenOperatorName   string      `json:"openOperatorName"`   // 开单人
			SettleOperatorName string      `json:"settleOperatorName"` // 结账人
			MemberId           string      `json:"memberId"`           // 会员ID
			MemberPhone        string      `json:"memberPhone"`        // 会员手机号
			MemberName         string      `json:"memberName"`         // 会员姓名
			OrderPeopleCnt     string      `json:"orderPeopleCnt"`     // 就餐人数
			SerialNo           string      `json:"serialNo"`           // 流水号
			Subject            string      `json:"subject"`            // 整单备注
			RelatedOrderId     string      `json:"relatedOrderId"`     // 关联订单ID，如整单退或反结时的原单订单ID
			ThirdSerialNo      string      `json:"thirdSerialNo"`      // 第三方订单流水号
		} `json:"orderBaseVO"` // 主单信息
		OrderTableVoList []struct {
			TableName string `json:"tableName"` // 桌台名称
			TableId   string `json:"tableId"`   // 桌台ID
		} `json:"orderTableVoList"` // 桌台信息
		OrderItemVoList []struct {
			ItemType           string `json:"itemType"`           // 商品类型
			GiftFlag           bool   `json:"giftFlag"`           // 是否赠送
			WeighFlag          bool   `json:"weighFlag"`          // 是否是称重商品
			TempFlag           bool   `json:"tempFlag"`           // 是否是临时菜
			PromoFlag          bool   `json:"promoFlag"`          // 是否是优惠菜
			BigTypeName        string `json:"bigTypeName"`        // 商品大类名称
			MidTypeName        string `json:"midTypeName"`        // 商品中类名称
			ItemCode           string `json:"itemCode"`           // 商品编码
			ItemName           string `json:"itemName"`           // 商品名称
			SaleStatusType     string `json:"saleStatusType"`     // 商品售卖状态类型
			SaleStatusTypeCode string `json:"saleStatusTypeCode"` // 商品售卖状态类型编码
			Id                 string `json:"id"`                 // 商品id
			ParentId           string `json:"parentId"`           // 父商品id
			ProductionDeptId   string `json:"productionDeptId"`   // 出品部门ID
			UnitName           string `json:"unitName"`           // 单位名称
			SpecName           string `json:"specName"`           // 规格
			SpecNameConcat     string `json:"specNameConcat"`     // 规格名称全称
			ItemPrice          string `json:"itemPrice"`          // 商品原始单价
			SalePrice          string `json:"salePrice"`          // 商品售价
			PracticeVoList     []struct {
				PracticeName string `json:"practiceName"` // 做法名称
				PracticeAmt  int    `json:"practiceAmt"`  // 做法金额
			} `json:"practiceVoList"` // 做法
			ItemSubject           string         `json:"itemSubject"`           // 商品备注
			Quantity              string         `json:"quantity"`              // 销售数量、退菜数量或赠菜数量 1. 是否是赠菜通过giftFlag判断 2. 是否是退菜通过saleStatusTypeCode判断，DISCARD表示退菜，NORMAL表示销售
			ItemSaleAmt           string         `json:"itemSaleAmt"`           // 商品销售金额
			ExtraFeeApportionAmt  string         `json:"extraFeeApportionAmt"`  // 服务费分摊金额
			ItemPromoApportionAmt string         `json:"itemPromoApportionAmt"` // 商品优惠分摊
			ItemReceivedAmt       string         `json:"itemReceivedAmt"`       // 商品收入
			Children              map[string]any `json:"children"`              // 子节点
		} `json:"orderItemVoList"` // 菜品明细
		OpenOrderPromoVoList []struct {
			PromoName    string `json:"promoName"`    // 优惠名称
			PromoType    string `json:"promoType"`    // 优惠类型
			PromoTime    string `json:"promoTime"`    // 优惠时间
			OperatorName string `json:"operatorName"` // 操作人姓名
			PromoAmt     int    `json:"promoAmt"`     // 优惠金额
		} `json:"openOrderPromoVoList"` // 订单优惠信息
		OpenExtraFeeVoList []struct {
			ExtraFeeName              string `json:"extraFeeName"`              // 服务费名称
			ExtraFeeType              string `json:"extraFeeType"`              // 服务费类型
			ExtraFeeAmt               int    `json:"extraFeeAmt"`               // 服务费金额
			ExtraFeePromoApportionAmt int    `json:"extraFeePromoApportionAmt"` // 服务费优惠分摊金额
			ExtraFeeReceivedAmt       int    `json:"extraFeeReceivedAmt"`       // 服务费收入
			ApportionedAmt            int    `json:"apportionedAmt"`            // 被分摊金额
			PromoTotalAmt             int    `json:"promoTotalAmt"`             // 服务费优惠总金额
		} `json:"openExtraFeeVoList"` // 服务费信息
		OpenPaymentDetailVoList []struct {
			PayDetailNo         string `json:"payDetailNo"`         // 支付单号
			FaceAmt             string `json:"faceAmt"`             // 面额
			PayAmt              string `json:"payAmt"`              // 支付金额
			ShopPromoAmt        string `json:"shopPromoAmt"`        // 商户优惠金额
			PlatformServiceAmt  string `json:"platformServiceAmt"`  // 平台抽佣/服务费
			ActualReceiveAmt    string `json:"actualReceiveAmt"`    // 商户实收
			PayMethodName       string `json:"payMethodName"`       // 支付方式名称
			PayDetailStatus     string `json:"payDetailStatus"`     // 支付状态
			PayDetailStartTime  string `json:"payDetailStartTime"`  // 创建时间/支出开始时间
			PayDetailEndTime    string `json:"payDetailEndTime"`    // 支付/退款完成时间
			PayPromoAmt         string `json:"payPromoAmt"`         // 支付优惠
			CouponCnt           string `json:"couponCnt"`           // 券数量
			CouponName          string `json:"couponName"`          // 券名称
			OperatorName        string `json:"operatorName"`        // 操作人姓名/收银员
			ActualPayAmt        string `json:"actualPayAmt"`        // 实际支付金额
			PlatformPromoAmt    string `json:"platformPromoAmt"`    // 平台优惠金额
			PayMethodId         int    `json:"payMethodId"`         // 支付方式ID，-3:现金，-4:银行卡，-129:扫码支付，-130:收银码，-5:微信，-6:支付宝，-37:云闪付，-1:会员卡，-15:实体卡，-20:匿名卡，-127:储值补录，-128:挂账，-24:美团团购券，-36:口碑团购券，0:抵用券
			CouponReconcileFlag bool   `json:"couponReconcileFlag"` // 团购券是否已对账 true:已对账 false:未对账
		} `json:"openPaymentDetailVoList"` // 支付信息
	} `json:"data"` // 数据
	ExtInfo struct {
		TraceId                string `json:"traceId"`                // 跟踪标识
		ServerCurrentTimestamp string `json:"serverCurrentTimestamp"` // 服务器当前时间戳
		ServerCurrentTime      string `json:"serverCurrentTime"`      // 服务器当前时间
	} `json:"extInfo"` // 附加信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type OrderListResp struct {
	Success    bool `json:"success"`    // 是否成功,true:成功,false:失败
	ServerCode int  `json:"serverCode"` // 服务编码
	Data       struct {
		TotalCount string `json:"totalCount"` // 总数
		PageNo     string `json:"pageNo"`     // 页号
		PageSize   string `json:"pageSize"`   // 页大小
		TotalPage  string `json:"totalPage"`  // 总页数
		List       []struct {
			BrandId          string `json:"brandId"`          // 品牌ID
			ShopId           string `json:"shopId"`           // 门店ID
			OrderId          string `json:"orderId"`          // 订单ID
			BusiOrderNo      string `json:"busiOrderNo"`      // 业务订单号
			OrderStatus      string `json:"orderStatus"`      // 订单状态{1:待处理(WAIT_PROCESSED) 2:已接单(ORDER_RECEIVED) 3:已完成(SUCCESS) 4:待结账(WAIT_SETTLED) 5:已结账(SETTLED) 6:已退单(REFUND) 7:已关闭(CLOSED) 8:已作废(INVALID) 9:已取消(CANCELLED) 10:已拒绝(REJECTED) 11:已反结账(ANTI_SETTLED)}
			OrderType        string `json:"orderType"`        // 订单类型{1:堂食(FOR_HERE) 2:平台外卖(PLATFORM_TAKE_OUT) 3:自营外卖(SELF_TAKE_OUT) 4:自提(SELF_TAKE) 5:无单收银(NO_ORDER_CASHIER) 6:会员充值(MEMBER_STORE) 7:会员补录(MEMBER_MANUAL_STORE) 8:销账订单(REPAYMENT_ORDER)}
			OrderAmt         string `json:"orderAmt"`         // 订单金额(单位/分)
			PromoAmt         string `json:"promoAmt"`         // 优惠金额(单位/分)
			OrderReceivedAmt string `json:"orderReceivedAmt"` // 订单收入/订单实收(单位/分)
			OpenTime         string `json:"openTime"`         // 下单时间
			FinishTime       string `json:"finishTime"`       // 完结时间
			ThirdOrderNo     string `json:"thirdOrderNo"`     // 第三方订单号
			SerialNo         string `json:"serialNo"`         // 订单流水号
			ThirdSerialNo    string `json:"thirdSerialNo"`    // 第三方流水号
		} `json:"list"` // 列表
		PrevPage     string `json:"prevPage"`     // 上一页
		NextPage     string `json:"nextPage"`     // 下一页
		EmptyForList bool   `json:"emptyForList"` // 当前list是否为null
	} `json:"data"` // 	数据
	MessageId string `json:"messageId"` // 服务消息ID
}

type StockInOutListResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    struct {
		List []struct {
			Id             string `json:"id"`             // 单据id
			OrgId          string `json:"orgId"`          // 机构id
			OrgName        string `json:"orgName"`        // 机构名字
			DepotId        string `json:"depotId"`        // 仓库id
			DepotName      string `json:"depotName"`      // 仓库名字
			BussDate       string `json:"bussDate"`       // 业务日期
			BillNo         string `json:"billNo"`         // 单号
			Status         int    `json:"status"`         // 状态 961：暂存，963：提交，962：已审核，964：已驳回，965：已反审，966：已作废
			TotalAmt       int    `json:"totalAmt"`       // 成本总金额
			TotalNum       int    `json:"totalNum"`       // 总数量
			Remarks        string `json:"remarks"`        // 备注
			CreateUserName string `json:"createUserName"` // 创建人
			UpdateUserName string `json:"updateUserName"` // 修改人
			CreateTime     string `json:"createTime"`     // 创建时间
			UpdateTime     string `json:"updateTime"`     // 修改时间
		} `json:"list"` // 单据列表
		Total int `json:"total"` // 总条数
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockInOutDetailResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    []struct {
		Id         string `json:"id"` // 单据id
		DetailList []struct {
			GoodsId       string `json:"goodsId"`       // 物品id
			GoodsName     string `json:"goodsName"`     // 物品名字
			GoodsCode     string `json:"goodsCode"`     // 物品编码
			GoodsSpec     string `json:"goodsSpec"`     // 物品规格
			UnitId        string `json:"unitId"`        // 标准单位id
			UnitName      string `json:"unitName"`      // 标准单位名字
			DualUnitName  string `json:"dualUnitName"`  // 辅助单位名字
			DualUnitId    string `json:"dualUnitId"`    // 辅助单位id
			GoodsQty      int    `json:"goodsQty"`      // 出库数量
			DualGoodsQty  int    `json:"dualGoodsQty"`  // 辅助单位出库数量
			UnitPrice     int    `json:"unitPrice"`     // 成本单价
			GoodsAmt      int    `json:"goodsAmt"`      // 成本金额
			ProDate       string `json:"proDate"`       // 生产日期
			ExpDate       string `json:"expDate"`       // 有效期至
			QualityPeriod int    `json:"qualityPeriod"` // 保质期天数
			GoodsBatchNo  string `json:"goodsBatchNo"`  // 批号
			Remarks       string `json:"remarks"`       // 备注
		} `json:"detailList"` // 单据明细列表
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}
