# Ozon Code Platform Question API
by Sergey Gerasimov

## Entity

**Question** - информация о вопросе от обучающегося

```
Question {
    Id uint64
    UserId uint64
    Text string
}
```

## Usage
The service accepts gRPC connections at port 82 and HTTP at 8081.

Supports methods:
* Get question list: [GET] /v1/questions
* Create new question: [POST] /v1/questions
* Get question description: [GET] /v1/questions/{questionId}
* Remove question: [DELETE] /v1/questions/{questionId}

## Installation

```bash
git clone https://github.com/ozoncp/ocp-question-api.git
cd ocp-question-api
make deps && make build
```
