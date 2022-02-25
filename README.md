# sap-api-integrations-quality-info-record-reads-rmq-kube
sap-api-integrations-quality-info-record-reads-rmq-kube は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 品質情報レコードデータを取得するマイクロサービスです。    
sap-api-integrations-quality-info-record-reads-rmq-kube には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-quality-info-record-reads-rmq-kube は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_QUALITYINFORECORD_SRV_0001/overview

## 動作環境  
sap-api-integrations-quality-info-record-reads-rmq-kube は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）
・ RabbitMQ on Kubernetes  
・ RabbitMQ Client　　

## クラウド環境での利用
sap-api-integrations-quality-info-record-reads-rmq-kube は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## RabbitMQ からの JSON Input

sap-api-integrations-quality-info-record-reads-rmq-kube は、Inputとして、RabbitMQ からのメッセージをJSON形式で受け取ります。 
Input の サンプルJSON は、Inputs フォルダ内にあります。  

## RabbitMQ からのメッセージ受信による イベントドリヴン の ランタイム実行

sap-api-integrations-quality-info-record-reads-rmq-kube は、RabbitMQ からのメッセージを受け取ると、イベントドリヴンでランタイムを実行します。  
AION の仕様では、Kubernetes 上 の 当該マイクロサービスPod は 立ち上がったまま待機状態で当該メッセージを受け取り、（コンテナ起動などの段取時間をカットして）即座にランタイムを実行します。　

## RabbitMQ への JSON Output

sap-api-integrations-quality-info-record-reads-rmq-kube は、Outputとして、RabbitMQ へのメッセージをJSON形式で出力します。  
Output の サンプルJSON は、Outputs フォルダ内にあります。  

## RabbitMQ の マスタサーバ環境

sap-api-integrations-quality-info-record-reads-rmq-kube が利用する RabbitMQ のマスタサーバ環境は、[rabbitmq-on-kubernetes](https://github.com/latonaio/rabbitmq-on-kubernetes) です。  
当該マスタサーバ環境は、同じエッジコンピューティングデバイスに配置されても、別の物理(仮想)サーバ内に配置されても、どちらでも構いません。

## RabbitMQ の Golang Runtime ライブラリ
sap-api-integrations-quality-info-record-reads-rmq-kube は、RabbitMQ の Golang Runtime ライブラリ として、[rabbitmq-golang-client](https://github.com/latonaio/rabbitmq-golang-client)を利用しています。

## デプロイ・稼働
sap-api-integrations-quality-info-record-reads-rmq-kube の デプロイ・稼働 を行うためには、aion-service-definitions の services.yml に、本レポジトリの services.yml を設定する必要があります。

kubectl apply - f 等で Deployment作成後、以下のコマンドで Pod が正しく生成されていることを確認してください。
```
$ kubectl get pods
```

## 本レポジトリ が 対応する API サービス
sap-api-integrations-quality-info-record-reads-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_QUALITYINFORECORD_SRV_0001/overview  
* APIサービス名(=baseURL): API_QUALITYINFORECORD_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-quality-info-record-reads-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* QualityInProcurement（SAP 品質情報レコード - ヘッダ）

## API への 値入力条件 の 初期値
sap-api-integrations-quality-info-record-reads-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.QualityInfoRecord.Material（品目）
* inoutSDC.QualityInfoRecord.Supplier（仕入先）
* inoutSDC.QualityInfoRecord.Plant（プラント）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "sap.s4.beh.qualityInfo.v1.QualityInfo.Created.v1",
	"accepter": ["Header"],
	"quality_info_record_no": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.qualityInfo.v1.QualityInfo.Created.v1",
	"accepter": ["All"],
	"quality_info_record_no": "",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetQualityInfoRecord(material, supplier, plant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(material, supplier, plant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 品質情報レコード の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"Material" ～ "QltyInProcmtLongText" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-quality-info-record-reads-rmq-kube/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-quality-info-record-reads-rmq-kube/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"Material": "QM002",
			"QltyInProcmtIntID": "1",
			"Supplier": "11300006",
			"Plant": "1110",
			"MaterialRevisionLevel": "",
			"QltyInProcmtReleaseValidTo": "9999-12-31T09:00:00+09:00",
			"BaseUnit": "PC",
			"ReleasedQuantity": "15",
			"BlockReason": "Quality Limitation (Q-Inf)",
			"ProcurementBlock": "",
			"QltyInProcmtRelQtyIsActive": "Y",
			"QltyInProcmtInspControl": "",
			"ProdnPieceApprovalIsRequired": "N",
			"ProductionPieceApproval": "",
			"ProductionPieceApprovalLevel": "",
			"ProductionPieceApprovalStatus": "",
			"IsDeleted": "",
			"QualityManagementSystem": "",
			"QltyInProcmtCertfnValidTo": "",
			"InspectionLotType": "",
			"QltyInProcmtLotCrtnLeadTime": 0,
			"QltyInProcmtCertificateCtrl": "",
			"QltyInProcmtLongText": ""
		}
	],
	"time": "2022-01-28T16:46:53+09:00"
}
```　　
