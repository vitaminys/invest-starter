package QIWIRequests

import (
    "time"
)

var curRequest_ID uint  = 0 //Уникальный номер последнего запроса

var curUnmadeRequest uint = 0 //Номер самой старой непроизведенной транзакции

type QIWIRequest struct {
  recipient string //QIWI_ID получателя
  sender string //QIWI_ID отправителя
  sum uint //Сумма, которая переводится
  date time.Time //Время транзакции
  request_ID uint //Уникальный номер запроса
  transactionStatus bool //Была ли осуществлена транзакция
}

var m = make(map[uint] QIWIRequest)

func addRequest(_sender, _recipient string, _sum uint) {
  var a QIWIRequest
  a.sender = _sender
  a.recipient = _recipient
  a.sum = _sum
  a.date = time.Now()
  a.request_ID = curRequest_ID
  a.transactionStatus = false
  m[curRequest_ID] = a
  curRequest_ID = curRequest_ID + 1
}
