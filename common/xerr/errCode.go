package xerr

// 成功返回
const OK uint32 = 200

// 系统

// SystemSaveExcelFail 保存excel失败
const SystemSaveExcelFail uint32 = 1001

// SystemNoDataDownloadError 没有数据可以下载
const SystemNoDataDownloadError uint32 = 1002

// AntChainInitError 蚂蚁区块链初始化失败
const AntChainInitError uint32 = 1003

// SysCacheSetFail 系统缓存设置失败
const SysCacheSetFail uint32 = 1004

// SysCacheIncFail 系统缓存自增失败
const SysCacheIncFail uint32 = 1005

// SysCacheGetFail 系统缓存获取失败
const SysCacheGetFail uint32 = 1006

// CacheSetFail 缓存设置失败
const CacheSetFail uint32 = 1007

// ConfigError 配置参数错误
const ConfigError uint32 = 1008

// SystemConfigQueryError 系统配置查询错误
const SystemConfigQueryError uint32 = 1009

// ApiRequestRepeatError 接口请求重复
const ApiRequestRepeatError uint32 = 1010

// FileReadError 文件读取错误
const FileReadError = 1011

// FileSaveError 文件保存错误
const FileSaveError = 1012

// Base64DecodeError base64解码错误
const Base64DecodeError = 1013

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const DATA_NOEXIST_ERROR uint32 = 100003
const DATA_EXIST_ERROR uint32 = 100004
const DATA_CONVERSION_ERROR uint32 = 100005 // "数据转换错误"
const DATA_NOT_UNIQUE_ERROR uint32 = 100006 // "数据不唯一"
const DataQueryError uint32 = 100007        // "数据查询错误"

const TOKEN_EXPIRE_ERROR uint32 = 100007
const TOKEN_GENERATE_ERROR uint32 = 100008
const TOKEN_TIMEOUT_ERROR uint32 = 100009
const TOKEN_WRONG_ERROR uint32 = 100010
const TOKEN_TYPE_ERROR uint32 = 100011
const REFRESH_TOKEN_ERROR uint32 = 100012
const TOO_MANY_REQUESTS_ERROR uint32 = 100013

const DB_ERROR uint32 = 100016
const DB_INSERT_ERROR uint32 = 100017
const DB_UPDATE_ERROR uint32 = 100018
const DB_DELETE_ERROR uint32 = 100019
const DB_QUERY_ERROR uint32 = 100020

const COPY_DATA_ERROR uint32 = 100021
const JSON_PARSE_ERROR uint32 = 100022
const VERIFY_SIGN_ERROR uint32 = 100023

const CTX_GET_USERID_ERROR uint32 = 100024
const CTX_GET_STOREID_ERROR uint32 = 100025

const NO_STORE_ACCOUNT_ERROR uint32 = 100026

// DataError 数据错误
const DataError uint32 = 100027

// AlipayApiRequestError 支付宝接口请求错误
const AlipayApiRequestError uint32 = 100031

// AlipayApiResponseError 支付宝接口响应错误
const AlipayApiResponseError uint32 = 100032

// DownloadImageUrlError 下载图片错误
const DownloadImageUrlError uint32 = 100033

// AiSignApiResponseError 爱签接口响应错误
const AiSignApiResponseError uint32 = 100034

// AiSignApiRequestError 爱签接口请求错误
const AiSignApiRequestError uint32 = 100035

// HeePayApiRequestError 汇付宝接口请求错误
const HeePayApiRequestError uint32 = 100036

// HeePayApiResponseError 汇付宝接口响应错误
const HeePayApiResponseError uint32 = 100037

// AppIdParamsError AppId错误
const AppIdParamsError uint32 = 100038

// DomainError 域名错误
const DomainError uint32 = 100039

// AntChainApiRequestError 蚂蚁区块链接口请求错误
const AntChainApiRequestError uint32 = 100040

// AntChainApiResponseError 蚂蚁区块链接口响应错误
const AntChainApiResponseError uint32 = 100041

// DownloadPdfFileError 下载PDF文件错误
const DownloadPdfFileError uint32 = 100042

// DataParseError 数据解析错误
const DataParseError uint32 = 100043

// SPI
const SPI_POSTPONE_DAYS_ERROR uint32 = 100101
const SPI_POSTPONE_ONLY_ONE uint32 = 100102
const SPI_POSTPONE_OVERDUE_ERROR uint32 = 100103
const SPI_SURRENDER_STATUS_ERROR uint32 = 100201
const SPI_MERCHANT_SETTLE_RATE_ERROR uint32 = 100202

// 商品服务
const PMS_ATTR_NOEXIST_ERROR uint32 = 200001
const PMS_ATTR_VALUE_NOEXIST_ERROR uint32 = 200101
const PMS_SKU_NOEXIST_ERROR uint32 = 200201
const PMS_SPU_NOEXIST_ERROR uint32 = 200301
const PMS_SKU_INFO_ERROR uint32 = 200401
const PMS_AMOUNT_MAX_LIMIT uint32 = 200501
const PMS_SKU_STATUS_DOWN uint32 = 200402
const PMS_SKU_STATUS_NORMAL_NODELETED uint32 = 200002
const PMS_SKU_STATUS_NOT_SALE uint32 = 200003
const PMS_COMMERCE_IMAGE_SIZE_ERROR uint32 = 200502
const PMS_COMMERCE_IMAGE_NOT_RATIO_ERROR uint32 = 200503

// PmsSkuPeriodsMinAmount05Error 当前渠道商品的每期金额最低不能低于0.5元
const PmsSkuPeriodsMinAmount05Error uint32 = 200403

// PmsSkuPeriodsMinAmount001Error 当前渠道商品的每期金额最低不能低于0.01元
const PmsSkuPeriodsMinAmount001Error uint32 = 200404

// PmsSkuCustomQueryError 自定义商品查询错误
const PmsSkuCustomQueryError = 200405

// PmsSkuCustomNotExist 自定义商品不存在
const PmsSkuCustomNotExist = 200406

// PmsSkuSnNotMatch 自定义商品sn不匹配
const PmsSkuSnNotMatch = 200407

// PmsSkuAntChainProductVersionEqual 区块链商品版本号相同
const PmsSkuAntChainProductVersionEqual = 200408

// PmsSkuMerchantNotMatch 商品与商家不匹配
const PmsSkuMerchantNotMatch = 200409

// PmsSkuAntChainProductQrCodeNotExist 商品二维码不存在,请先生成商品二维码
const PmsSkuAntChainProductQrCodeNotExist = 200410

// PmsSkuCommerceQueryError 安心付商品查询错误
const PmsSkuCommerceQueryError = 200411

// PmsSkuCommerceNotExist 安心付商品不存在
const PmsSkuCommerceNotExist = 200412
const PmsSkuCommerceVerifyError = 200413
const PmsSkuCommerceSyncAddButton = 200414
const PmsSkuCommerceSyncModifyButton = 200415
const PmsSkuCommerceUpdateError = 200416

// PmsSkuAntchainQueryError 区块链商品查询错误
const PmsSkuAntchainQueryError = 200417
const PmsSkuAntchainNotExist = 200418
const PmsSkuAntchainUpdateError = 200419

// PmsSkuCommercePriceOrPeriodNotMatch 安心付商品价格或期数不一致
const PmsSkuCommercePriceOrPeriodNotMatch = 200420

// PmsSkuCommercePriceListQueryFail 安心付商品价格列表查询失败
const PmsSkuCommercePriceListQueryFail = 200421

// PmsSkuAntchainEditTypeNotEdit 区块链商品为不可编辑类型
const PmsSkuAntchainEditTypeNotEdit uint32 = 200422

// PmsSkuAntchainTotalPriceNotMatch 区块链总价与期数总价不匹配
const PmsSkuAntchainTotalPriceNotMatch uint32 = 200423

// PmsSkuCommercePriceListTypeError 安心付一客一价类型错误
const PmsSkuCommercePriceListTypeError uint32 = 200424

// PmsSkuCommerceCycleTypeError 安心付商品周期类型错误
const PmsSkuCommerceCycleTypeError uint32 = 200425

// 用户服务
const USER_NOEXIST_ERROR uint32 = 300001
const USER_EXIST_ERROR uint32 = 300002
const USER_LOGIN_ERROR uint32 = 300003
const USER_INFO_ERROR uint32 = 300004
const USER_PASSWORD_ERROR uint32 = 300021
const USER_OLDPASSWORD_ERROR uint32 = 300022
const USER_PASSWORD_SAME_ERROR uint32 = 300023
const USER_PASSWORD_NOT_SAME_ERROR uint32 = 300024
const USER_PASSWORD_RESET_ERROR uint32 = 300025
const USER_PASSWORD_NOT_EMPTY uint32 = 300026
const USER_STATUS_ERROR uint32 = 300031
const USERNAME_USED_ERROR uint32 = 300032
const USER_NO_PERMISSION_ERROR uint32 = 300033
const USER_LOGIN_SYSTEM_OAUTH_TOKEN_ERROR uint32 = 300034
const USET_LOGIN_INFO_SHARE_ERROR uint32 = 300035
const GET_USET_INFO_ERROR uint32 = 300036
const USER_LOGIN_ILLEGAL uint32 = 300037
const UserRegisterError uint32 = 300038
const UserPlatformDataNotExist uint32 = 300039
const USER_REGISTER_TYPE_ERROR uint32 = 300040
const USER_OPERATE_ILLEGAL uint32 = 300041
const USER_REFERRALCODE_EXIST uint32 = 300042
const USER_REFERRALCODE_REPEAT_ERROR uint32 = 300043

// UserInfoQueryFail 用户信息查询失败
const UserInfoQueryFail uint32 = 301001

// UserInfoNotExist 用户信息不存在
const UserInfoNotExist uint32 = 301002

// 店铺服务 ams
const AMS_STORE_NEW_ERROR uint32 = 400000
const AMS_STORE_INFO_ERROR uint32 = 400001
const AMS_STORE_NOEXIST_ERROR uint32 = 400002
const AMS_STORE_STATUS_ERROR uint32 = 400003
const AMS_STORE_NO_PERMISSION_ERROR uint32 = 400004
const AMS_STORE_NOEXIST_BY_USERID_ERROR uint32 = 400005
const AMS_STORE_NOEXIST_BY_STOREID_ERROR uint32 = 400006
const AMS_STORE_NOEXIST_BY_STOREIDS_ERROR uint32 = 400007

// AmsHeepayMerchantBalanceQueryError 汇付宝商户余额变动信息查询失败
const AmsHeepayMerchantBalanceQueryError uint32 = 400008

// AmdHeepayMerchantWithdrawCountQueryError 汇付宝商户提现次数查询失败
const AmdHeepayMerchantWithdrawCountQueryError uint32 = 400009

// AmsHeepayMerchantWithdrawQueryError 汇付宝商户提现信息查询失败
const AmsHeepayMerchantWithdrawQueryError uint32 = 400010

// AmsHeepayMerchantWithdrawNotExistError 汇付宝商户提现信息不存在
const AmsHeepayMerchantWithdrawNotExistError uint32 = 400011

// AmsHeepayMerchantBalanceNotExistError 汇付宝商户余额变动信息不存在
const AmsHeepayMerchantBalanceNotExistError uint32 = 400012

// AmsStoreShopQueryFail 门店信息查询失败
const AmsStoreShopQueryFail = 400013

// AmsStoreShopNotExist 门店信息不存在
const AmsStoreShopNotExist = 400014

// AmsAntChainStoreSignTemplateQueryFail 区块链商户签约模板查询错误
const AmsAntChainStoreSignTemplateQueryFail uint32 = 400015

// AmsAntChainStoreSignTemplateNotExist 区块链商户签约模板不存在
const AmsAntChainStoreSignTemplateNotExist uint32 = 400016

// AmsAntChainStoreBalanceNotEnough 区块链商户余额不足
const AmsAntChainStoreBalanceNotEnough uint32 = 400017

// AmsAntChainStoreSignTemplateValidityError 区块链签约模板签约有效期错误
const AmsAntChainStoreSignTemplateValidityError uint32 = 400018

// AmsAntChainStoreSignTemplateValidityUnitError 区块链签约模板签约有效期单位错误
const AmsAntChainStoreSignTemplateValidityUnitError uint32 = 400019

// AmsAntChainStoreFeeRateLetZero 区块链商户费率不能小于等于0
const AmsAntChainStoreFeeRateLetZero uint32 = 400020

// AmsAntChainMerchantBalanceQueryFail 区块链商户余额变动记录查询失败
const AmsAntChainMerchantBalanceQueryFail uint32 = 400021

// AmsStoreStatusInvalidError 该商户已被禁用，请联系商家或平台
const AmsStoreStatusInvalidError uint32 = 400022

// 订单服务 oms
const OMS_ORDER_INFO_ERROR uint32 = 500001
const OMS_ORDER_NOEXIST_ERROR uint32 = 500002
const OMS_ORDER_STATUS_ERROR uint32 = 500003
const OMS_ORDER_REFUND_ERROR uint32 = 500004
const OMS_ORDER_REFUND_STATUS_ERROR uint32 = 500005
const OMS_ORDER_REFUND_NOEXIST_ERROR uint32 = 500006

const OMS_CREATE_ORDER_ERROR uint32 = 500011
const OMS_UPDATE_ORDER_ERROR uint32 = 500012
const OMS_ORDER_COMMERCE_METHOD_ERROR uint32 = 500013

// 汇付宝订单
// OmsOrderDataQueryError 订单数据查询错误
const OmsOrderDataQueryError uint32 = 510001

// OmsNoOrderDataDownloadError 没有订单数据可以下载
const OmsNoOrderDataDownloadError uint32 = 510002

// OmsHeepayOrderQueryError 汇付宝订单查询失败
const OmsHeepayOrderQueryError uint32 = 510003

// OmsHeepayOrderNotExist 汇付宝订单不存在
const OmsHeepayOrderNotExist uint32 = 510004

// OmsHeepayOrderUpdateError 汇付宝订单更新失败
const OmsHeepayOrderUpdateError uint32 = 510005

// OmsHeepayOrderAllotNotAllowSelfError 汇付宝订单分润商户不能是自己
const OmsHeepayOrderAllotNotAllowSelfError uint32 = 510006

// OmsHeepayOrderAllotCountLimit 汇付宝订单商户分润数量限制 - 最多6个
const OmsHeepayOrderAllotCountLimit uint32 = 510007

// OmsHeepayOrderAllotAmountLimit 汇付宝订单商户分润比率超出限制
const OmsHeepayOrderAllotAmountLimit uint32 = 510008

// OmsHeepayOrderAllotRepeatError 汇付宝订单分润商户重复
const OmsHeepayOrderAllotRepeatError uint32 = 510009

// OmsHeepayOrderAllotQueryError 汇付宝订单分润信息查询失败
const OmsHeepayOrderAllotQueryError uint32 = 510010

// OmsHeepayOrderDeductedQueryError 汇付宝订单扣款信息查询失败
const OmsHeepayOrderDeductedQueryError uint32 = 510011

// OmsHeepayOrderDeductedNotExist 汇付宝订单扣款信息不存在
const OmsHeepayOrderDeductedNotExist uint32 = 510012

// OmsHeepayOrderPreviousDeductedNotSuccess 汇付宝订单扣款上一期未扣款成功
const OmsHeepayOrderPreviousDeductedNotSuccess uint32 = 510013

// OmsHeepayMerchantEntryNotSuccess 汇付宝商户未入驻成功
const OmsHeepayMerchantEntryNotSuccess uint32 = 510014

// OmsHeepayOrderAllotNotExist 汇付宝订单分润信息不存在
const OmsHeepayOrderAllotNotExist uint32 = 510015

// OmsHeepayDeductedLogCreateError 汇付宝扣款记录创建失败
const OmsHeepayDeductedLogCreateError uint32 = 510016

// OmsRequestDeductedError 汇付宝请求扣款失败
const OmsRequestDeductedError uint32 = 510017

// OmsHeepayOrderDeductedFailDayError 汇付宝订单扣款连续失败次数超限
const OmsHeepayOrderDeductedFailDayError uint32 = 510018

// OmsHeepayDeductFail 汇付宝扣款失败
const OmsHeepayDeductFail uint32 = 510019

// OmsHeepayOrderDeductedLogQueryError 汇付宝扣款记录查询失败
const OmsHeepayOrderDeductedLogQueryError uint32 = 510020

// OmsHeepayOrderDeductedLogNotExist 汇付宝扣款记录不存在
const OmsHeepayOrderDeductedLogNotExist uint32 = 510021

// OmsHeepayOrderDeductedInProgress 汇付宝订单扣款中
const OmsHeepayOrderDeductedInProgress uint32 = 510022

// OmsHeepayOrderAlreadyTerminated 订单已解约
const OmsHeepayOrderAlreadyTerminated uint32 = 510023

// OmsHeepayOrderPerformanceFinish 订单已履约完成
const OmsHeepayOrderPerformanceFinish uint32 = 510024

// OmsHeepayOrderAllotAddError 汇付宝订单添加失败
const OmsHeepayOrderAllotAddError uint32 = 510025

// OmsHeepayOrderAllotUpdateError 汇付宝订单分润信息更新失败
const OmsHeepayOrderAllotUpdateError uint32 = 510026

// OmsHeepayOrderFinishedTerminatedError 汇付宝履约完成的订单不可以解约
const OmsHeepayOrderFinishedTerminatedError uint32 = 510027

// OmsHeepayOrderDeductJobQueryError 汇付宝扣款任务查询失败
const OmsHeepayOrderDeductJobQueryError uint32 = 510028

// OmsHeepayOrderOverDelayCount 汇付宝订单超出延期次数
const OmsHeepayOrderOverDelayCount uint32 = 510029

// OmsHeepayOrderDeductedUpdateError 汇付宝订单扣款信息更新失败
const OmsHeepayOrderDeductedUpdateError uint32 = 510030

// OmsHeepayOrderStatusError 汇付宝订单状态异常
const OmsHeepayOrderStatusError uint32 = 510031

// OmsHeepayOrderDeductedSuccess 汇付宝订单扣款成功
const OmsHeepayOrderDeductedSuccess uint32 = 510032

// 安心付订单

// OmsCommerceOrderSubQueryError 安心付订购单查询失败
const OmsCommerceOrderSubQueryError uint32 = 510033

// OmsCommerceOrderSubNotExist 安心付订购单不存在
const OmsCommerceOrderSubNotExist uint32 = 510034

// OmsCommerceOrderSubAddError 安心付订购单添加失败
const OmsCommerceOrderSubAddError uint32 = 510035

// OmsCommerceOrderSubUpdateError 安心付订购单更新失败
const OmsCommerceOrderSubUpdateError uint32 = 510036

// OmsDeductionOrderSyncError 安心付核销/扣款订单同步失败
const OmsDeductionOrderSyncError uint32 = 510037

// OmsDeductionOrderQueryError 安心付核销/扣款订单查询失败
const OmsDeductionOrderQueryError uint32 = 510038

// OmsCommerceDeductedSyncError 安心付扣款计划同步失败
const OmsCommerceDeductedSyncError uint32 = 510039

// OmsCommerceDeductedQueryError 安心付扣款计划查询失败
const OmsCommerceDeductedQueryError uint32 = 510040

// OmsCommerceDeductedNotExistError 安心付扣款计划数据不存在
const OmsCommerceDeductedNotExistError uint32 = 510041

// OmsOrderStatusNeedNormal 订单状态需要是正常/履约中
const OmsOrderStatusNeedNormal uint32 = 510042

// OmsPeriodsError 订单总期数错误
const OmsPeriodsError uint32 = 510043

// OmsCommerceCardIdEmptyError 安心付卡号不能为空
const OmsCommerceCardIdEmptyError uint32 = 510044

// OmsOrderCommerceQueryFail 安心付售卖订单查询失败
const OmsOrderCommerceQueryFail uint32 = 510045

// OmsOrderCommerceNotExist 安心付售卖订单数据不存在
const OmsOrderCommerceNotExist uint32 = 510046

// OmsOrderCommerceDeductionQueryFail 安心付核销单查询失败
const OmsOrderCommerceDeductionQueryFail = 510047

// OmsOrderCommerceDeductionNotExist 安心付核销单数据不存在
const OmsOrderCommerceDeductionNotExist = 510048

// OmsCommerceDeductedUpdateFail 安心付扣款计划更新失败
const OmsCommerceDeductedUpdateFail = 510049

// ( start )区块链 OmsAntChain  1级 5,2级 2

// OmsAntChainOrderQueryError 区块链订单查询失败
const OmsAntChainOrderQueryError uint32 = 520001

// OmsAntChainOrderNotExist 区块链订单不存在
const OmsAntChainOrderNotExist uint32 = 520002

// OmsAntChainUserSignMaxCountOver 用户在商户签约次数超过限制,请联系商户解锁
const OmsAntChainUserSignMaxCountOver uint32 = 520003

// OmsAntChainOrderUpdateError 区块链订单更新失败
const OmsAntChainOrderUpdateError uint32 = 520004

// OmsAntChainOrderContractExpireError 区块链订单合同已过期请重新下单
const OmsAntChainOrderContractExpireError uint32 = 520005

// OmsAntChainOrderSignStatusError 区块链订单签约状态错误
const OmsAntChainOrderSignStatusError uint32 = 520006

// OmsAntChainOrderDeductedQueryError 区块链订单扣款信息查询失败
const OmsAntChainOrderDeductedQueryError uint32 = 520007

// OmsAntChainOrderDeductedNotExist 区块链订单扣款信息不存在
const OmsAntChainOrderDeductedNotExist uint32 = 520008

// OmsAntChainOrderStatusError - 区块链订单状态错误
const OmsAntChainOrderStatusError uint32 = 520009

// OmsAntChainTermIndexNotMatch 区块链期数不匹配
const OmsAntChainTermIndexNotMatch uint32 = 520010

// OmsAntChainPreTermNotEnd 当前期数之前的期数未结束
const OmsAntChainPreTermNotEnd uint32 = 520011

// OmsAntChainOrderDeductedUpdateError 区块链订单代扣计划信息更新失败
const OmsAntChainOrderDeductedUpdateError uint32 = 520012

// OmsAntChainOrderSignStatusAbnormal 区块链订单签约状态异常，请重新下单
const OmsAntChainOrderSignStatusAbnormal = 520013

// OmsAntChainOrderUnSignQueryError 区块链解约申请查询失败
const OmsAntChainOrderUnSignQueryError = 520014

// OmsAntChainContractSignNotComplete 区块链合同签约未完成
const OmsAntChainContractSignNotComplete = 520015

// OmsAntChainOrderUnSignNotExist 区块链解约申请不存在
const OmsAntChainOrderUnSignNotExist = 520016

// OmsAntChainOrderUnSignUpdateError 区块链解约申请更新失败
const OmsAntChainOrderUnSignUpdateError = 520017

// OmsAntChainOrderUnSignWaitHandle 区块链用户解约申请未处理，请先处理解约申请
const OmsAntChainOrderUnSignWaitHandle = 520018

// OmsAntChainOrderDeductedNeedReturnStateUnpaid 本期扣款状态需要是未支付的状态
const OmsAntChainOrderDeductedNeedReturnStateUnpaid = 520019

// OmsAntChainOrderDeductedNotYetRepayment 指定期数未到还款日
const OmsAntChainOrderDeductedNotYetRepayment = 520020

// OmsAntChainOrderDeductedRetryTodayIsSend 今日已发起，不可重复发起
const OmsAntChainOrderDeductedRetryTodayIsSend = 520021

// OmsAntChainPlanModeNotSupport 代扣模式不支持
const OmsAntChainPlanModeNotSupport = 520022

// OmsAntChainOrderDelayNotSupport 该订单模式不支持延期
const OmsAntChainOrderDelayNotSupport = 520023

// OmsCommerceDeductedPayOffAmountError 获取订购单线下还款金额统计错误
const OmsCommerceDeductedPayOffAmountError = 520024

// OmsOrderNotBelongToStore 此订单不属于该商户
const OmsOrderNotBelongToStore = 520025

// OmsAntChainWithholdStatusNotComplete 区块链代扣签署未完成
const OmsAntChainWithholdStatusNotComplete = 520026

// OmsOperatorNotBelongOverdueMemo 操作人不属于逾期记录创建者
const OmsOperatorNotBelongOverdueMemo = 520027

// OmsAntPlanRetryQueryError 代扣重试次数查询失败
const OmsAntPlanRetryQueryError = 520028

// OmsAntChainOrderAbnormal 该区块链订单异常请联系服务商处理
const OmsAntChainOrderAbnormal = 520029

// OmsAntChainOrderPauseCountOver 订单暂停次数已超限
const OmsAntChainOrderPauseCountOver = 520030

// OmsAntChainOrderNeedPause 该订单需要已暂停状态
const OmsAntChainOrderNeedPause = 520031

// OmsAntChainOrderDelayCountOver 订单延期次数已超限
const OmsAntChainOrderDelayCountOver = 520032

// ( end )区块链 OmsAntChain  1级 5,2级 2

// 支付服务

// 商户小程序

// MiniNotExist 小程序不存在
const MiniNotExist = 600000

// MiniCreateError 小程序创建错误
const MiniCreateError = 600001

// MiniUpdateError 小程序更新失败
const MiniUpdateError = 600002

// MiniStoreInfoQueryError 商户基础数据查询失败
const MiniStoreInfoQueryError = 600003

// MiniStoreAppletQueryError 小程序查询失败
const MiniStoreAppletQueryError = 600004

// MiniStoreAppletExistError 商户已创建小程序
const MiniStoreAppletExistError = 600005

// MiniStoreAppletOrderIdExistError 小程序的order_no不存在
const MiniStoreAppletOrderIdExistError = 600006

// MiniStoreAppletNotExistError 小程序信息不存在
const MiniStoreAppletNotExistError = 600007

// MiniStoreAppletConfirmError 商家未确认小程序
const MiniStoreAppletConfirmError = 600008

// MiniStoreInfoNotExistError 商户基础信息不存在
const MiniStoreInfoNotExistError = 600009

// MiniUploadBusinessLicenseError 上传营业执照失败
const MiniUploadBusinessLicenseError = 600010

// MiniStoreImageQueryError 备案图片数据查询失败
const MiniStoreImageQueryError = 600011

// MiniStoreImageUploadError 备案图片上传错误
const MiniStoreImageUploadError = 600012

// MiniStoreImageUploadFailError 备案图片上传失败
const MiniStoreImageUploadFailError = 600013

// MiniStoreImageSaveError 备案图片保存错误
const MiniStoreImageSaveError = 600014

// MiniStoreQueryError 商户数据查询失败
const MiniStoreQueryError = 600015

// MiniStoreNotExistError 商户信息不存在
const MiniStoreNotExistError = 600016

// MiniStoreAppAuthTokenNotExistError 商户app_auth_token不存在
const MiniStoreAppAuthTokenNotExistError = 600017

// MiniUploadIdCardFrontError 上传身份证正面错误
const MiniUploadIdCardFrontError = 600018

// MiniUploadIdCardReverseError 上传身份证正背面错误
const MiniUploadIdCardReverseError = 600019

// MiniStoreAppletFaceAuthNotExistError 人脸识别凭证不存在
const MiniStoreAppletFaceAuthNotExistError = 600020

// MiniCertificateTypeNotExistError 主体证件类型无法找到
const MiniCertificateTypeNotExistError = 600021

// MiniStoreImageNotExistError 备案图片不存在
const MiniStoreImageNotExistError = 600022

// MiniMainCategoryIdNotExistError 指定的服务内容标识不存在
const MiniMainCategoryIdNotExistError = 600023

// MiniCreateStatusNotExistError 小程序创建状态不存在
const MiniCreateStatusNotExistError = 600024

// MiniTemplateNotExistError 小程序模板不存在
const MiniTemplateNotExistError = 600025

// MiniVersionQueryError 小程序版本号查询错误
const MiniVersionQueryError = 600026

// MiniStoreAppVersionNotExistError 小程序版本号不存在
const MiniStoreAppVersionNotExistError = 600027

// MiniBuildVersionError 小程序构建版本号错误
const MiniBuildVersionError = 600028

// MiniAttachmentImageTypeNotExistError 附件图片类型不存在
const MiniAttachmentImageTypeNotExistError = 600029

// MiniAuditInfoListParseError 小程序前置审批项信息解析失败
const MiniAuditInfoListParseError = 600030

// MiniSubjectAttachmentMaterialsParseError 小程序主体附件字段解析失败
const MiniSubjectAttachmentMaterialsParseError = 600031

// MiniAppletAttachmentMaterialsParseError 小程序附件字段解析失败
const MiniAppletAttachmentMaterialsParseError = 600032

// MiniAttachmentTypeNotRequireApprovalError 该行业需要附件类型-不涉及前置审批的承诺书
const MiniAttachmentTypeNotRequireApprovalError = 600033

// MiniUploadAudioInfoFailError 上传前置审批附件失败
const MiniUploadAudioInfoFailError = 600034

// MiniUploadSubjectAttachmentMaterialsFailError 上传主体附件失败
const MiniUploadSubjectAttachmentMaterialsFailError = 600035

// MiniAppletAttachmentMaterialsFailError 上传小程序附件失败
const MiniAppletAttachmentMaterialsFailError = 600036

// MiniBuildStatusTranslateError 小程序版本构建状态转换错误
const MiniBuildStatusTranslateError = 600037

// MiniIcpStatusNoNeedSubmit ICP审核状态不需要提交
const MiniIcpStatusNoNeedSubmit = 600038

// MiniIcpStatusNotExistError ICP备案状态不存在
const MiniIcpStatusNotExistError = 600039

// MiniFaceAuthInitiatedError 人脸识别已发起
const MiniFaceAuthInitiatedError = 600040

// MiniGuangDongMaterialsParamsError 广东省必须要填写: 互联网承诺书文件材料凭证和负责人授权书
const MiniGuangDongMaterialsParamsError = 600041

// MiniBuildStatusNotAudit 小程序版本不在审核中状态
const MiniBuildStatusNotAudit = 600042

// MiniBuildVersionNotExistError 小程序构建版本号不存在
const MiniBuildVersionNotExistError = 600043

// MiniCategoryParseError 小程序类目解析出错
const MiniCategoryParseError = 600044

// MiniCategoryError 小程序类目错误
const MiniCategoryError = 600045

// MiniCategoryExistError 小程序类目已存在
const MiniCategoryExistError = 600046

// MiniStoreAppletVersionNotExistError 小程序版本号不存在
const MiniStoreAppletVersionNotExistError = 600047

// MiniBuildVersionExistError 小程序构建版本号已存在
const MiniBuildVersionExistError = 600048

// MiniBuildStatusNotVersionCreated 小程序构建版本号未创建
const MiniBuildStatusNotVersionCreated = 600049

// MiniAuditStatusAuditing 该版本已在审核中
const MiniAuditStatusAuditing = 600050

// MiniAuditVersionNotExistError 小程序目前没有审核的版本
const MiniAuditVersionNotExistError = 600051

// MiniVersionStatusError 小程序版本状态错误
const MiniVersionStatusError = 600052

// MiniStatusInitError 小程序为开发状态
const MiniStatusInitError = 600053

// MiniCancelOnlineVersionError 不能退回上架中的版本
const MiniCancelOnlineVersionError = 600054

// MiniIsNotOnlineStatusRelease 当前小程序版本不是上架状态
const MiniIsNotOnlineStatusRelease = 600055

// MiniIsNotOfflineStatusOffline 当前小程序版本不是下架状态
const MiniIsNotOfflineStatusOffline = 600056

// MiniIsNotAuditStatusPass 当前小程序版本未通过审核
const MiniIsNotAuditStatusPass = 600057

// MiniVersionIsOver 小程序版本生命周期已结束
const MiniVersionIsOver = 600058

// MiniIsOnlineRollbackYes 当前小程序版本已是回滚状态
const MiniIsOnlineRollbackYes = 600059

// HeepayMerchantInfoQueryError 汇付宝商户数据查询失败
const HeepayMerchantInfoQueryError = 600060

// HeepayMerchantExistError 汇付宝商户已存在
const HeepayMerchantExistError = 600061

// HeepayMerchantStoreIdExistError 汇付宝商户已存在store_id重复
const HeepayMerchantStoreIdExistError = 600062

// HeepayMerchantContactEmailExistError 汇付宝商户已存在联系人邮箱重复
const HeepayMerchantContactEmailExistError = 600063

// ShopSceneImageError 店内场景图片缺失
const ShopSceneImageError = 600064

// HeepayMerchantAddError 添加汇付宝商户失败
const HeepayMerchantAddError = 600065

// HeepayMerchantNotExistError 汇付宝商户不存在
const HeepayMerchantNotExistError = 600066

// HeepayMerchantEmailNotEqual 汇付宝商户邮箱不一致
const HeepayMerchantEmailNotEqual = 600067

// ChannelNotExistError 渠道不存在
const ChannelNotExistError = 600068

// MiniAppIdRepeatError 小程序appId重复
const MiniAppIdRepeatError = 600069

// HeepayMerchantEmailNotAllowModify 入驻成功后不允许修改邮箱
const HeepayMerchantEmailNotAllowModify = 600070

// StoreLogoNotAllowEmpty 商户logo不能为空
const StoreLogoNotAllowEmpty = 600071

// StoreLogoSizeError 商户logo大小错误:必须是160*160
const StoreLogoSizeError = 600072

// HeepayMerchantWithdrawTypeError 汇付宝商户提现类型错误
const HeepayMerchantWithdrawTypeError = 600073

// StoreAppIdAndMiniAppIdNotEqual 商户appId和小程序appId不一致
const StoreAppIdAndMiniAppIdNotEqual = 600074

// HeepayMerchantUpdateError 汇付宝商户信息更新失败
const HeepayMerchantUpdateError = 600075

// AntChainMerchantInfoQueryError 区块链商户数据查询失败
const AntChainMerchantInfoQueryError = 600076

// AntChainMerchantExistError 区块链商户已存在
const AntChainMerchantExistError = 600077

// AntChainMerchantIdRepeatError 区块链商户社会信用代码重复
const AntChainMerchantIdRepeatError = 600078

// AntChainSitUrlFormatError 区块链商户网站域名格式错误
const AntChainSitUrlFormatError = 600079

// AntChainMerchantAddError 区块链商户添加失败
const AntChainMerchantAddError = 600080

// AntChainMerchantNotExistError 区块链商户不存在
const AntChainMerchantNotExistError = 600081

// AntChainMerchantStatusOkNotAllowModify 区块链商户进件状态为入驻成功，不允许修改!
const AntChainMerchantStatusOkNotAllowModify = 600082

// AntChainMerchantUpdateError 区块链商户信息更新失败
const AntChainMerchantUpdateError = 600083

// AntChainMerchantAllotSetFail 区块链商户分润比率查询失败
const AntChainMerchantAllotSetFail = 600084

// AntChainMerchantAllotGetFail 区块链商户分润比率获取失败
const AntChainMerchantAllotGetFail = 600085

// AppletIsJuTypeNotMatch 小程序聚合类型不匹配
const AppletIsJuTypeNotMatch = 600086

// AntChainPrepaidRechargeError 区块链预付费充值失败
const AntChainPrepaidRechargeError = 600087

// AntChainDeductPrepaidAmountError 区块链扣除预付费金额失败
const AntChainDeductPrepaidAmountError = 600088

// AntChainContractSaveFail 区块链合同保存失败
const AntChainContractSaveFail = 600089

// AntChainContractQueryFail 区块链合同查询失败
const AntChainContractQueryFail = 600090

// AntChainContractNotExistError 区块链合同不存在
const AntChainContractNotExistError = 600091

// AntChainContractSignStatusNeedFinish 合同签署状态需要满足签署完成状态
const AntChainContractSignStatusNeedFinish = 600092

// AntCacheGetContractSignDownloadUrlError 缓存获取用户区块链合同下载地址失败
const AntCacheGetContractSignDownloadUrlError = 600093

// PrepaidChannelNotSupport 预付费渠道不支持
const PrepaidChannelNotSupport = 600094

// PrepaidChangeTypeError 预付费变更类型错误
const PrepaidChangeTypeError = 600095

// ChannelNotSupport 渠道不支持
const ChannelNotSupport = 600096

// CacheGetSurrenderCodeError 缓存获取解约短信验证码失败
const CacheGetSurrenderCodeError = 600097

// CacheClearSurrenderCodeError 缓存清除解约短信验证码失败
const CacheClearSurrenderCodeError = 600098

// AppletConfigNotExistError 聚合小程序配置不存在
const AppletConfigNotExistError = 700001

// AppQrcodeCreatCreateError 聚合小程序推广二维码生成错误
const AppQrcodeCreatCreateError = 700002

// AppQrcodeCreatCreateFailError 聚合小程序推广二维码生成失败
const AppQrcodeCreatCreateFailError = 700003

// AppQrcodeSaveFailError 聚合小程序推广二维码保存失败
const AppQrcodeSaveFailError = 700004

// DuLiAppQrcodeCreatCreateError 独立小程序推广二维码生成错误
const DuLiAppQrcodeCreatCreateError = 700005

// DuLiAppQrcodeCreatCreateFailError 独立小程序推广二维码生成失败
const DuLiAppQrcodeCreatCreateFailError = 700006

// StoreImageSaveError 保存商户图片失败
const StoreImageSaveError = 700007

// StoreImageQueryError 查询商户图片失败
const StoreImageQueryError = 700008

// StoreImageUpdateError 商户图片更新失败
const StoreImageUpdateError = 700009

// HeepayMerchantContactEmailError 汇付宝商户联系人邮箱错误
const HeepayMerchantContactEmailError = 700010

// SkuQrcodeCreatCreateFailError sku商品二维码生成错误
const SkuQrcodeCreatCreateFailError = 800001

// SkuNonDivisibleInitialPeriodsError SKU - 分期金额有余数/除不尽的情况下前段期数不可以大于1
const SkuNonDivisibleInitialPeriodsError = 800002

// AlipayOpenMiniQrcodePatternCreateError 创建关联普通二维码模式失败
const AlipayOpenMiniQrcodePatternCreateError = 800003

// MiniQrcodePatternNotExistError 关联普通二维码模式数据不存在
const MiniQrcodePatternNotExistError = 800004

// AlipayOpenMiniQrcodeUnbindError 解绑关联普通二维码模式失败
const AlipayOpenMiniQrcodeUnbindError = 800005

// NeedSignMustTemplateIdError 需要签约的商户必须指定签约合同模板id
const NeedSignMustTemplateIdError = 800006

// NeedSignTemplateIdError 签约商品需要配置签约合同模板
const NeedSignTemplateIdError = 800007

// AddSignUserFail 添加签约用户失败
const AddSignUserFail = 800008

// QuerySignUserFail 查询签约用户失败
const QuerySignUserFail = 800009

// QuerySignTemplateFail 查询签约模板失败
const QuerySignTemplateFail = 800010

// SignTemplateNotExistError 签约模板不存在
const SignTemplateNotExistError = 800011

// QueryStoreAppletFail 查询商户的小程序失败
const QueryStoreAppletFail = 800012

// QuerySignMerchantFail 查询签约商户失败
const QuerySignMerchantFail = 800013

// SignMerchantNotExistError 签约商户不存在
const SignMerchantNotExistError = 800014

// AiSignUserExistError 爱签签约商户信息已存在
const AiSignUserExistError = 800015

// AiSignMerchantAuthFail 爱签签约商户认证失败
const AiSignMerchantAuthFail = 800016

// NotNeedSignError 非签约商户不需要签约
const NotNeedSignError = 800017

// AiSignGetSerialNoFail 获取签约认证流水号失败
const AiSignGetSerialNoFail = 800018

// AiSignCacheSerialNoFail 缓存签约认证流水号失败
const AiSignCacheSerialNoFail = 800019

// AiSignCacheStoreSealFail 缓存商户印章信息失败
const AiSignCacheStoreSealFail = 800020

// AiSignCacheGetStoreSealFail 缓存获取商户印章信息失败
const AiSignCacheGetStoreSealFail = 800021

// SkuPayTypeNotExistError 商品付款类型不存在
const SkuPayTypeNotExistError = 810000

// SkuPayMethodNotExistError 商品付款方式不存在
const SkuPayMethodNotExistError = 810001

// StoreUserTypeError 商户用户类型错误，请使用商户类型登录
const StoreUserTypeError = 810002

// SignTemplateFillDataParseError 签约模板合同参数解析错误
const SignTemplateFillDataParseError = 810003

// SignCreateContractError 创建上传待签署文件错误
const SignCreateContractError = 810004

// SignCreateContractFail 创建上传待签署文件失败
const SignCreateContractFail = 810005

// SignAddContractDataFail 添加合同数据失败
const SignAddContractDataFail = 810006

// SignAddContractFilesNotExist 合同文件数据不存在
const SignAddContractFilesNotExist = 810007

// SignTemplateParamsParseError 模板签约参数解析错误
const SignTemplateParamsParseError = 810008 // 签约TemplateParamsParseError

// SignStartFail 签约发起失败
const SignStartFail = 810009

// SignMerchantExistError 小程序签约商户信息已存在
const SignMerchantExistError = 810010

// SignSignUserError 签约用户错误
const SignSignUserError = 810011

// SignPersonFaceFail 签约人脸识别失败
const SignPersonFaceFail = 810012

// QueryAppletSignMerchantError 查询小程序签约商户信息失败
const QueryAppletSignMerchantError = 810013

// QueryAppletSignMerchantNotExist 小程序签约商户信息不存在
const QueryAppletSignMerchantNotExist = 810014

// IdCardNameNotMatch 身份证和姓名不匹配
const IdCardNameNotMatch = 810015

// IdCardFormatError 身份证号码格式不正确
const IdCardFormatError = 810016

// UpdateSignUserFail 签约用户信息更新失败
const UpdateSignUserFail = 810017

// SignContractNotExist 签约合同不存在
const SignContractNotExist = 810018

// SignContractStatusNotDownload 签约完成的合同才可以下载
const SignContractStatusNotDownload = 810019

// SignContractStatusNotInvalid 签约完成的合同才能发起作废
const SignContractStatusNotInvalid = 810020

// NoSignUserError 没有签约人信息
const NoSignUserError = 810021

// QueryOrderSignContractFail 查询订单签约合同失败
const QueryOrderSignContractFail = 810022

// SignOrderRepeatContractError 同一个订单不可以发起2次签约
const SignOrderRepeatContractError = 810023

// SignContractStatusInvalid 该合同已在作废流程状态中
const SignContractStatusInvalid = 810024

// QueryUserCertVerifyInfoFail 查询用户实名信息失败
const QueryUserCertVerifyInfoFail = 810025

// UserCertverifyNameNotMatch 实名的姓名错误，与身份证信息不一致
const UserCertverifyNameNotMatch = 810026

// MiniCategoryAppletQueryFail 指定类目的小程序查询失败
const MiniCategoryAppletQueryFail = 810027

// MiniAppletVersionQueryFail 小程序版本信息查询失败
const MiniAppletVersionQueryFail = 810028

// MiniAppletVersionNotExist 小程序版本信息不存在
const MiniAppletVersionNotExist = 810029

// NoTestAppletNeedPassword 没有测试小程序就批量上架需要密码确认
const NoTestAppletNeedPassword = 810030

// NoTestAppletMustSetApplet 不指定测试小程序的情况下，必须明确指定一个要发布的小程序
const NoTestAppletMustSetApplet = 810031

// MiniCategoryAppletBatchQueryFail 批量查询指定类目的小程序错误
const MiniCategoryAppletBatchQueryFail = 810032

// MiniCategoryAppletNotExist 指定类目的小程序信息不存在
const MiniCategoryAppletNotExist = 810033

// TestVersionMismatch 测试小程序上架版本号与目标版本号不一致
const TestVersionMismatch = 810034

// QueryAutoAppletTaskStatusError 查询自动上架任务状态
const QueryAutoAppletTaskStatusError = 810035

// AutoAppletTaskIsRunning 该小程序自动上架任务正在进行中
const AutoAppletTaskIsRunning = 810036

// JobTypeStatusConflict jobType与Status冲突,只能二选一
const JobTypeStatusConflict = 810037

// MiniCategoryNotMatch 小程序类目不一致
const MiniCategoryNotMatch = 810038

// UserCertverifyMobileNotMatch 用户实名信息手机号不匹配
const UserCertverifyMobileNotMatch = 810039

// SkuDataNotExist 商品数据不存在
const SkuDataNotExist = 810040

// SkuChannelError 商品渠道错误
const SkuChannelError = 810041

// SkuStoreError 此商品不属于该商户
const SkuStoreError = 810042

// SkuChannelIsNotAntChain 商品渠道不属于蚂蚁区块链渠道
const SkuChannelIsNotAntChain = 810043

// IdCardCheckError 身份证校验失败
const IdCardCheckError = 810044

// CheckAgeError 年龄校验失败
const CheckAgeError = 810045

// DuLiSignMerchantStoreIdNotMatch 独立小程序的情况下-store_id不匹配
const DuLiSignMerchantStoreIdNotMatch = 810046

// SignMerchantAccountExistError 签约商户账号已存在
const SignMerchantAccountExistError = 810047

// SignMerchantAuthInfoNotExist 签约商户认证信息不存在或已过期
const SignMerchantAuthInfoNotExist = 810048

// SignContractUpdateFail 更新签约合同失败
const SignContractUpdateFail = 810049

// PhoneFormatError 手机号格式错误
const PhoneFormatError = 810050

// ======= 汇付宝/卡代扣 ======

// HeepayQuerySignBankCardFail  - 查询签约银行卡信息失败
const HeepayQuerySignBankCardFail = 820001

// HeepaySignDataDecryptFail 签约数据解密失败
const HeepaySignDataDecryptFail = 820002

// HeepaySignBankCardNotExist 签约银行卡信息不存在
const HeepaySignBankCardNotExist = 820003

// HeepayUserBindCardNotExist 用户绑卡信息不存在
const HeepayUserBindCardNotExist = 820004

// HeepayUserBindCardFail 用户绑卡失败
const HeepayUserBindCardFail = 820005

// HeepaySignStatusNotSuccess 签约状态需要满足签约成功状态
const HeepaySignStatusNotSuccess = 820006

// HeepayChannelNotSwitchAli 原始渠道是卡代扣 - 不允许切换到芝麻先享
const HeepayChannelNotSwitchAli = 820007

// HeepayNoSignBankCard 没有签约成功的银行卡信息
const HeepayNoSignBankCard = 820008

// HeepayPayDataDecryptFail 支付数据解密失败
const HeepayPayDataDecryptFail = 820009

// HeepayChannelOnlyTestCanSwitchChannel 渠道切换功能只允许测试小程序使用
const HeepayChannelOnlyTestCanSwitchChannel = 820010

// HeepayUserBindCardQueryFail 汇付宝用户绑卡信息查询失败
const HeepayUserBindCardQueryFail = 820011

// HeepayOrderBindCardError 汇付宝订单下单绑卡失败错误提示
const HeepayOrderBindCardError = 820012

// ======= 短信 ======

// SmsSendFail 短信发送失败
const SmsSendFail = 830001

// SmsSendRepeat 短信已发送，请勿重复发送
const SmsSendRepeat = 830002

// SmsCodeNotExist 短信验证码不存在，请重新发起
const SmsCodeNotExist = 830003

// SmsCodeError 短信验证码错误
const SmsCodeError = 830004

// ======= 图片验证码 ======

// ImageCaptchaCreateError 图片验证码生成错误
const ImageCaptchaCreateError = 840001

// ImageCaptchaNotExist 图片验证码不存在
const ImageCaptchaNotExist = 840002

// ImageCaptchaChannelError 图片验证码渠道类型错误
const ImageCaptchaChannelError = 840003

// ImageCaptchaSceneError 图片验证码场景错误
const ImageCaptchaSceneError = 840004

// ImageCaptchaError 图片验证码错误
const ImageCaptchaError = 840005
