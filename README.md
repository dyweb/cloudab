# Project cloudab

云原生的 A/B 测试系统

## 介绍

A/B 测试起源于农业工程。人们将土地划分为不同的地块，通过种植不同的农作物来确定在这些土地上更适合种植何种作物。随后 A/B 测试被广泛地应用于医学、工业等不同领域。随着谷歌论文 [Overlapping Experiment Infrastructure: More, Better, Faster Experimentation](https://storage.googleapis.com/pub-tools-public-publication-data/pdf/36500.pdf) 的发表，将 A/B 测试引入了互联网领域。

A/B实验通过在线上流量中取出一部分，完全随机地分给不同的策略，再结合一定的统计方法，得到对于两种策略相对效果的准确估计。这一估计的结果可以一定程度上反映出不同的策略的优劣好坏。

随着泛互联网行业竞争的加剧，过去产品经理决策驱动的策略不再能够在充分竞争的市场中取得优势。通过 A/B 测试，不同的经营策略，不同的推荐算法之间的优劣都可以通过精确估计的方式予以量化。数据比产品经理主观的意愿更具有说服力。通过数据驱动的方式，能够以更加科学的方式进行决策。

## 编译

```bash
make
```

## 运行

### 本地运行

```bash
$ mongod &
$ ./bin/cloudab --logger-debug
```

### Docker 运行（TODO）

## 测试

可以通过导入 Postman collection 文件或者使用 cURL 进行测试。

### Postman

可以导入 [cloudab.postman_collection.json](docs/dev/cloudab.postman_collection.json) 进行测试。

### cURL

创建一个新的实验：

```bash
curl --location --request POST 'http://localhost:9999/apis/v1/experiments' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test",
    "versions": [
        {
            "name": "version-1",
            "traffic": 100,
            "features": [
                {
                    "name": "feature-1",
                    "value": "value-1"
                }
            ]
        }
    ]
}'
```

返回如下：

```json
{
    "id": "602ca4df9f4e8f5088966f6c",
    "name": "test",
    "versions": [
        {
            "id": "602ca4df9f4e8f5088966f6b",
            "name": "version-1",
            "traffic": 100,
            "features": [
                {
                    "name": "feature-1",
                    "value": "value-1"
                }
            ]
        }
    ]
}
```

进行分流（将 experimentID 替换成创建实验返回结构体的 id，上述例子中是 `602ca4df9f4e8f5088966f6c`，userID 可取任意值）：

```bash
curl --location --request GET 'http://localhost:9999/apis/v1/experiments/{experimentID}/abconfig?userID={userID}'
```

返回如下：

```json
{
    "features": [
        {
            "name": "feature-1",
            "value": "value-1"
        }
    ],
    "experiment_name": "test",
    "experiment_id": "602ca4df9f4e8f5088966f6c",
    "versions": [
        "602ca4df9f4e8f5088966f6b"
    ]
}
```

上报指标 （TODO）

## SDK（TODO）

## License

Apache 2.0
